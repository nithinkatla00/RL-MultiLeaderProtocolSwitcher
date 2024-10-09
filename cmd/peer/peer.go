package main

import (
	_ "net/http/pprof"

	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/chainhotstuff"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/destiny"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dispel"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dqpbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dqsbftslow"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/duobft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/hybster
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/hybsterx"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/linhybster"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/minbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/mirbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/multichainduobft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/multichainduobft/chainduobf
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/pbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/prime"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/rcc"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/sbft"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/sbftslow"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/sbftx"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/transport"
)

func newPeer(cfg *peer.LocalConfig) *peer.Peer {
	t := transport.NewGRPCTransport()

	var p peer.Protocol
	switch cfg.Algorithm {
	case peerpb.Algorithm_PBFT:
		p = pbft.NewPBFT(cfg)
	case peerpb.Algorithm_SBFT:
		p = sbft.NewSBFT(cfg)
	case peerpb.Algorithm_Hybster:
		p = hybster.NewHybster(cfg)
	case peerpb.Algorithm_Destiny:
		p = destiny.NewDestiny(cfg)
	case peerpb.Algorithm_LinHybster:
		p = linhybster.NewLinHybster(cfg)
	case peerpb.Algorithm_SBFTSlow:
		p = sbftslow.NewSBFTSlow(cfg)
	case peerpb.Algorithm_Prime:
		p = prime.NewPrime(cfg)
	case peerpb.Algorithm_MinBFT:
		p = minbft.NewMinBFT(cfg)
	case peerpb.Algorithm_DuoBFT:
		p = duobft.NewDuoBFT(cfg)
	case peerpb.Algorithm_Hybsterx:
		p = hybsterx.NewHybsterx(cfg)
	case peerpb.Algorithm_ChainHotstuff:
		p = chainhotstuff.NewChainHotstuff(cfg)
	case peerpb.Algorithm_RCC:
		p = rcc.NewRCC(cfg)
	case peerpb.Algorithm_MirBFT:
		p = mirbft.NewMirBFT(cfg)
	case peerpb.Algorithm_Dispel:
		p = dispel.NewDispel(cfg)
	case peerpb.Algorithm_SBFTx:
		p = sbftx.NewSBFTx(cfg)
	case peerpb.Algorithm_DQPBFT:
		p = dqpbft.NewDQPBFT(cfg)
	case peerpb.Algorithm_DQSBFTSlow:
		p = dqsbftslow.NewDQSBFTSlow(cfg)
	case peerpb.Algorithm_ChainDuoBFT:
		p = chainduobft.NewChainDuoBFT(cfg)
	case peerpb.Algorithm_MultiChainDuoBFT:
		p = multichainduobft.NewMultiChainDuoBFT(cfg)
	case peerpb.Algorithm_MultiChainDuoBFTRCC:
		cfg.MultiChainDuoBFT.RCCMode = true
		p = multichainduobft.NewMultiChainDuoBFT(cfg)
	default:
		cfg.Logger.Panicf("Unknown protocol specified %v", cfg.Algorithm)
	}

	return peer.New(cfg, t, p)
}
