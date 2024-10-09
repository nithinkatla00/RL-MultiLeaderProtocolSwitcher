package transport

import (
	"context"

	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/pkg/command/commandpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/transport/transportpb"
)

type Request struct {
	Command *commandpb.Command
	ReturnC chan<- commandpb.CommandResult
}

// Transport handles RPC messages for Raft coordination.
type Transport interface {
	Init(id peerpb.PeerID, addr string, peers map[peerpb.PeerID]string)
	Serve(MessageHandler)
	Send(msgs []peerpb.Message)
	Close()
}

// MessageHandler is an object capable of accepting incoming protocol messages.
type MessageHandler interface {
	HandleMessage(*transportpb.TransMsg)
	HandleCommand(context.Context,
		*transportpb.ClientPacket) (*transportpb.ClientPacket, error)
}
