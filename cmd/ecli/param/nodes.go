package param

import (
	"github.com/OpenOCC/OCC/core/types"
	"github.com/OpenOCC/OCC/param"
)

var (
	Localnet bool = false
	Testnet  bool = false
	Mainnet  bool = false
)

func GetPeers() []types.Peer {
	if Localnet {
		return param.LocalNet
	} else if Testnet {
		return param.TestNet
	} else if Mainnet {
		return param.MainNet
	}
	return param.LocalNet
}
