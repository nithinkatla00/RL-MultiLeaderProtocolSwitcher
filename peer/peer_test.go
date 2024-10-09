package peer

import (
	"testing"

	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/pkg/command/commandpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/pkg/logger"
	transport "github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/transport"
)

type mockProtocol struct {
	batches []*commandpb.Command
}

func (p *mockProtocol) Request(b *commandpb.Command) {
	p.batches = append(p.batches, b)
}

func (*mockProtocol) Tick() {

}

func (*mockProtocol) Step(peerpb.Message) {

}

func (*mockProtocol) Ready() {

}

func (*mockProtocol) AsyncCallback() {

}

func (*mockProtocol) Callback(ExecCallback) {

}

func (*mockProtocol) ClearExecutedCommands() {

}

func (*mockProtocol) ClearMsgs() {

}

func (*mockProtocol) MakeReady() Ready {
	return Ready{}
}

func TestPeerBatching(t *testing.T) {
	protocol := &mockProtocol{}
	cfg := &LocalConfig{
		PeerConfig: &peerpb.PeerConfig{},
		Logger:     logger.NewDefaultLogger(),
	}
	p := New(cfg, transport.NewGRPCTransport(), protocol)

	for i := 0; i < 38; i++ {
		ops := []commandpb.Operation{
			{
				Type: &commandpb.Operation_KVOp{},
			},
			{
				Type: &commandpb.Operation_KVOp{},
			},
			{
				Type: &commandpb.Operation_KVOp{},
			},
			{
				Type: &commandpb.Operation_KVOp{},
			},
		}
		p.rb.add(reqBufElem{cmd: &commandpb.Command{Timestamp: uint64(i + 1),
			Ops: ops}})
	}

	p.flushCmds(p.rb.b[:p.rb.i])

	for i := 0; i < 9; i++ {
		if len(protocol.batches[i].Ops) != 4 {
			t.Fatalf("Invalid batch size at %d: %d != 4", i, len(protocol.batches[i].Ops))
		}
	}
	if len(protocol.batches[9].Ops) != 4 {
		t.Fatalf("Invalid batch size at %d: %d != 4", 9, len(protocol.batches[9].Ops))
	}
}
