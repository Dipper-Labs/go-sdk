package rpc

import (
	"encoding/hex"

	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/auth"
	"github.com/Dipper-Labs/go-sdk/types"
)

func (c *client) GetTx(txHash string) (res Tx, err error) {
	hash, err := hex.DecodeString(txHash)
	if err != nil {
		return
	}

	resultTx, err := c.rpc.Tx(hash, true)
	if err != nil {
		return
	}

	res.Hash = resultTx.Hash.String()
	res.Height = resultTx.Height
	res.Index = resultTx.Index
	res.TxResult = resultTx.TxResult
	res.Proof = resultTx.Proof

	var stdTx auth.StdTx
	if err := types.Cdc.UnmarshalBinaryLengthPrefixed(resultTx.Tx, &stdTx); err != nil {
		return res, err
	} else {
		res.StdTx = stdTx
	}

	return
}
