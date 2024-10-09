package main

import (
	_ "net/http/pprof"

	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dispel"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dqpbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/mirbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/rcc"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/transport"
)

func newPeer(cfg *peer.LocalConfig) *peer.Peer {
    t := transport.NewGRPCTransport()

    var p peer.Protocol
    switch cfg.Algorithm {
    case peerpb.Algorithm_RCC:
        p = rcc.NewRCC(cfg) // Make sure NewRCC accepts *peer.LocalConfig
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
