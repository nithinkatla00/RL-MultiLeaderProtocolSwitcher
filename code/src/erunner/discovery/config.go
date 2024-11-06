package discovery

import (
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/cmd/master/masterpb"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
)

type Config struct {
	PeerCount       int32                     `json:"node_count"`
	ClientCount     int32                     `json:"client_count"`
	Algorithm       peerpb.Algorithm          `json:"algorithm"`
	MaxFailures     int32                     `json:"max_failures"`
	MaxFastFailures int32                     `json:"max_fast_failures"`
	LeaderRegion    string                    `json:"leader_region"`
	KeyFile         string                    `json:"key_file"`
	ServerConfig    masterpb.BaseServerConfig `json:"server"`
	ClientConfig    masterpb.BaseClientConfig `json:"client"`
}
