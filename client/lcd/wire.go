package lcd

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/irisnet/irishub/codec"
)

var cdc = codec.New()

func init() {
	ctypes.RegisterAmino(cdc)
}
