package peer

import (
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/pkg/command/commandpb"
)

type Protocol interface {
	Step(peerpb.Message)
	Tick()
	Request(*commandpb.Command)
	Callback(ExecCallback)

	MakeReady() Ready
	ClearMsgs()
	ClearExecutedCommands()

	AsyncCallback()
}
