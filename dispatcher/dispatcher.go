package dispatcher

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/OpenOCC/OCC/core/types"
	"github.com/OpenOCC/OCC/core/userevent"
	"github.com/OpenOCC/OCC/node"
)

func NewTransaction(transaction *userevent.Transaction) error {
	// 主币的tokenAddress为空
	if transaction.TokenAddress != "" {
		tokenAddress, err := hex.DecodeString(transaction.TokenAddress)
		if err != nil {
			return err
		}
		currentBlock := node.GetMainChain().LastHeader()
		var token types.Token
		err = currentBlock.TokenTree.GetInterfaceValue(tokenAddress, &token)
		if err != nil || token.Name == "" || token.Decimals <= 0 || token.Total <= 0 {
			return err
		}
	}
	if !userevent.ValidateTransaction(*transaction) {
		return errors.New("error signature")
	}
	if bytes.EqualFold(transaction.GetFrom(), transaction.GetTo()) {
		return errors.New("invalid address")
	}
	node.GetMainChain().NewTransaction(transaction)
	return nil
}
