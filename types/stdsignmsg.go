package types

import (
	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/auth"
	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
)

type StdSignMsg struct {
	ChainID       string      `json:"chain_id"`
	AccountNumber uint64      `json:"account_number"`
	Sequence      uint64      `json:"sequence"`
	Fee           auth.StdFee `json:"fee"`
	Msgs          []sdk.Msg   `json:"msgs"`
	Memo          string      `json:"memo"`
}

func (msg StdSignMsg) Bytes() []byte {
	return auth.StdSignBytes(msg.ChainID, msg.AccountNumber, msg.Sequence, msg.Fee, msg.Msgs, msg.Memo)
}
