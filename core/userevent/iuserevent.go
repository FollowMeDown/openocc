package userevent

import (
	"bytes"
	"github.com/OpenOCC/OCC/core/types"
	"github.com/OpenOCC/OCC/crypto"
)

func ValidateTransaction(transaction Transaction) bool {
	pubKey, err := crypto.RecoverPubKey(transaction.Msg(), transaction.GetSign())
	if err != nil {
		return false
	}
	result := bytes.EqualFold(types.FromPubKeyToAddress(pubKey), transaction.GetFrom())
	return result
}

func SignTransaction(transaction *Transaction, privKey []byte) error {
	sign, err := crypto.Crypto(transaction.Msg(), privKey)
	if err != nil {
		return err
	}
	transaction.SetSign(sign)
	return nil
}
