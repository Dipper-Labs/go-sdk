package rpc

import (
	tmrpcclient "github.com/tendermint/tendermint/rpc/client"
	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/Dipper-Labs/go-sdk/client/types"
)

type RpcClient interface {
	BroadcastTx(broadcastType string, tx tmtypes.Tx) (types.BroadcastTxResult, error)
	GetNodeStatus() (NodeStatus, error)
	GetTx(hash string) (Tx, error)
	GetBlock(height int64) (*tmrpctypes.ResultBlock, error)
}

type client struct {
	rpc *tmrpcclient.HTTP
}

func NewClient(nodeUrl string) RpcClient {
	rpc := tmrpcclient.NewHTTP(nodeUrl, "/websocket")
	return &client{rpc: rpc}
}
