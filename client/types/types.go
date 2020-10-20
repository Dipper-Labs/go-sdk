package types

import (
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
)

type (
	Coin struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	BroadcastTxResult struct {
		BroadcastResult ResultBroadcastTx       `json:"broadcast_result"`
		CommitResult    ResultBroadcastTxCommit `json:"commit_result"`
	}

	ResultBroadcastTx struct {
		Code uint32          `json:"code"`
		Data common.HexBytes `json:"data"`
		Log  string          `json:"log"`
		Hash common.HexBytes `json:"hash"`
	}

	ResultBroadcastTxCommit struct {
		CheckTx   types.ResponseCheckTx   `json:"check_tx"`
		DeliverTx types.ResponseDeliverTx `json:"deliver_tx"`
		Hash      common.HexBytes         `json:"hash"`
		Height    int64                   `json:"height"`
	}
)
