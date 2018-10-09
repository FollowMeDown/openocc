package param

import (
	"github.com/OpenOCC/OCC/conf"
	"github.com/OpenOCC/OCC/core/types"
)

var mapping = make(map[string][]types.Peer)
var MainChainDelegateNode []types.Peer

func InitBootNodes() {
	mapping["mainnet"] = MainNet
	mapping["testnet"] = TestNet
	mapping["localnet"] = LocalNet
	MainChainDelegateNode = mapping[conf.EKTConfig.Env]
}
