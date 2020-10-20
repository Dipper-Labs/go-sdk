package rpc

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/p2p"
	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"

	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/auth"
)

type NodeStatus struct {
	NodeInfo p2p.DefaultNodeInfo `json:"node_info"`
	SyncInfo tmrpctypes.SyncInfo `json:"sync_info"`
}

type Tx struct {
	Hash     string                 `json:"hash"`
	Height   int64                  `json:"height"`
	Index    uint32                 `json:"index"`
	TxResult abci.ResponseDeliverTx `json:"tx_result"`
	StdTx    auth.StdTx             `json:"std_tx"`
	Proof    types.TxProof          `json:"proof,omitempty"`
}
