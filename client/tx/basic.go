package tx

import (
	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	"github.com/Dipper-Labs/go-sdk/client/lcd"
	"github.com/Dipper-Labs/go-sdk/client/rpc"
	"github.com/Dipper-Labs/go-sdk/client/types"
	"github.com/Dipper-Labs/go-sdk/keys"
)

type Client interface {
	keys.KeyManager
	SendToken(toAddrBech32 string, amount []types.Coin, memo string, commit bool) (types.BroadcastTxResult, error)
	ContractCall(contractAddrBech32 string, payload []byte, amount sdk.Coin, commit bool) (types.BroadcastTxResult, error)
	QueryContractEvents(contractBech32Addr string, startBlockNum int64, endBlockNum int64) ([]string, error)
}

type client struct {
	keys.KeyManager
	chainId   string
	lcdClient lcd.LcdClient
	rpcClient rpc.RpcClient
}

func NewClient(chainId string, keyManager keys.KeyManager, liteClient lcd.LcdClient, rpcClient rpc.RpcClient) Client {
	return &client{
		chainId:    chainId,
		KeyManager: keyManager,
		lcdClient:  liteClient,
		rpcClient:  rpcClient,
	}
}
