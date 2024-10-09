package dispel

import (
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	pb "github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/protocols/dispel/dispelpb"
)

type epoch struct {
	epochNum pb.Epoch

	rbDoneCount   int
	consDoneCount int

	rbs  map[peerpb.PeerID]*rbroadcast
	cons map[peerpb.PeerID]*dbft

	d *Dispel
}

func makeEpoch(m *pb.RBMessage, d *Dispel) *epoch {
	e := &epoch{
		epochNum: m.EpochNum,

		rbs:  make(map[peerpb.PeerID]*rbroadcast),
		cons: make(map[peerpb.PeerID]*dbft),

		d: d,
	}

	return e
}
