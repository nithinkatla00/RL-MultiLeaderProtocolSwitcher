package worker

import (
	"context"

	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/pkg/command/commandpb"
)

func (c *ClientWorker) Generic(
	req *commandpb.Command,
) (*commandpb.CommandResult, error) {
	req.Target = peerpb.PeerID(c.leaderID)

	ctx, cancel := context.WithCancel(context.Background())

	resultChan := make(chan *commandpb.CommandResult, c.peerCount)
	errChan := make(chan peerpb.PeerID, c.peerCount)

	c.logger.Debugf("sending request %v", req.Timestamp)
	c.conn.SendToAll(ctx, req, resultChan, errChan)

	firstVal := <-resultChan
	c.logger.Debugf("received 0th response for %v: %v", req.Timestamp, firstVal)
	for i := 1; i < c.failures+1; i++ {
		res := <-resultChan
		c.logger.Debugf("received %dth response for %v: %v", i, req.Timestamp, res)
		if !res.Equal(firstVal) {
			c.logger.Panicf("VALUES NOT EQUAL: %v != %v", firstVal, res)
		}
	}
	cancel()
	return firstVal, nil
}
