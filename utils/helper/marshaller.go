package helper

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/utils/signer"
)

func VerifyAndUnmarshall(raw, sign []byte, msg proto.Message, signer *signer.Signer) {
	if signer.Verify(raw, sign) != nil {
		panic(fmt.Sprintf("Unable to verify signature: %v", raw))
	}
	err := proto.Unmarshal(raw, msg)
	if err != nil {
		panic(fmt.Sprintf("Unable to unmarshal msg %v: %v", msg, err))
	}
}

func MarshallAndSign(msg proto.Message, signer *signer.Signer) ([]byte, []byte) {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		panic(fmt.Sprintf("Unable to marshal msg %v: %v", msg, err))
	}
	sign, err := signer.Sign(msgBytes)
	if err != nil {
		panic(fmt.Sprintf("Unable to sign msg: %v", err))
	}
	return msgBytes, sign
}
