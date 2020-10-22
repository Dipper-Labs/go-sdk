package rpc

import (
	"context"
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
	Subscribe(ctx context.Context, subscriber, query string) (out <-chan tmrpctypes.ResultEvent, err error)
	UnSubscribe(ctx context.Context, subscriber, query string) error
	UnSubscribeAll(ctx context.Context, subscriber string) error
	Start() error
}

type client struct {
	rpc *tmrpcclient.HTTP
}

func NewClient(nodeUrl string) RpcClient {
	rpc := tmrpcclient.NewHTTP(nodeUrl, "/websocket")
	return &client{rpc: rpc}
}
