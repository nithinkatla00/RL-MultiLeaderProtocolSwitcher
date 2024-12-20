package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "net/http/pprof"

	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/cmd/master/masterpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/pkg/logger"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dispel"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dqpbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/mirbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
)

var (
	help   = flag.Bool("help", false, "")
	maddr  = flag.StringP("master", "m", "master-ip:5060", "address of the master node to bootstrap with")
	msleep = flag.DurationP("msleep", "s", 0, "Time (in s) to wait before connecting to master")
)

func startMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}

func NewPeer(cfg *peer.LocalConfig) *peer.Peer {
	t := transport.NewGRPCTransport()

	var p peer.Protocol
	switch cfg.Algorithm {
	case peerpb.Algorithm_MirBFT:
		p = mirbft.NewMirBFT(cfg)
	case peerpb.Algorithm_Dispel:
		p = dispel.NewDispel(cfg)
	case peerpb.Algorithm_DQPBFT:
		p = dqpbft.NewDQPBFT(cfg)
	default:
		cfg.Logger.Panicf("Unknown protocol specified %v", cfg.Algorithm)
	}

	return peer.New(cfg, t, p)
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	go startMetricsServer()

	// Start a Datadog tracer, optionally providing a set of options,
	// returning an opentracing.Tracer which wraps it.
	// t := opentracer.New(tracer.WithServiceName("go-consensus"))
	// defer tracer.Stop() // important for data integrity (flushes any leftovers)

	// Use it with the Opentracing API. The (already started) Datadog tracer
	// may be used in parallel with the Opentracing API if desired.
	// opentracing.SetGlobalTracer(t)

	flag.CommandLine.MarkHidden("help")
	flag.Parse()
	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprint(os.Stderr, flag.CommandLine.FlagUsagesWrapped(120))
		return
	}

	logger := logger.NewDefaultLogger()

	id, config, err := registerWithMaster(*maddr)
	if err != nil {
		logger.Fatalf("Unable to register with master: %v", err)
		return
	}

	if config.LogVerbose {
		logger.EnableDebug()
	}

	peers := make([]peerpb.PeerID, len(config.PeerDetails))
	peerAddrs := make(map[peerpb.PeerID]string)
	for i, p := range config.PeerDetails {
		peers[i] = p.PeerID
		peerAddrs[p.PeerID] = fmt.Sprintf("%s:%d", p.PodIP, config.ListenPort)
	}

	lConfig := &peer.LocalConfig{
		PeerConfig: config,
		Peers:      peers,
		PeerAddrs:  peerAddrs,
		ListenAddr: fmt.Sprintf(":%d", config.ListenPort),
		ID:         id,
		Logger:     logger,
		RandSeed:   time.Now().Unix(),
	}

	p := NewPeer(lConfig)

	// Make sure we clean up before exiting.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		p.Stop()
		os.Exit(0)
	}()

	logger.Infof("Started peer")
	p.Run()
}

func registerWithMaster(maddr string) (peerpb.PeerID, *peerpb.PeerConfig, error) {
	time.Sleep(*msleep)

	conn, err := grpc.Dial(maddr, grpc.WithInsecure())
	if err != nil {
		return 0, nil, err
	}

	cli := masterpb.NewServiceDiscoveryClient(conn)

	hostname := os.Getenv("HOSTNAME")
	region := hostname
	pieces := strings.Split(hostname, "-")
	if len(pieces) == 4 {
		region = pieces[1]
	}

	nid := &peerpb.BasicPeerInfo{
		PodName:     os.Getenv("PODNAME"),
		HostMachine: hostname,
		PodIP:       os.Getenv("PODIP"),
		Region:      region,
	}

	res, err := cli.Register(context.Background(), nid)
	if err != nil {
		return 0, nil, err
	}

	return res.PeerID, res.PeerConfig, nil
}
