package client

import (
	"errors"
	"fmt"

	"github.com/Dipper-Labs/go-sdk/client/basic"
	"github.com/Dipper-Labs/go-sdk/client/lcd"
	"github.com/Dipper-Labs/go-sdk/client/rpc"
	"github.com/Dipper-Labs/go-sdk/client/tx"
	"github.com/Dipper-Labs/go-sdk/config"
	"github.com/Dipper-Labs/go-sdk/keys"
)

type Client interface {
	basic.HttpClient
	lcd.LcdClient
	rpc.RpcClient
	tx.Client
}

type client struct {
	basic.HttpClient
	lcd.LcdClient
	rpc.RpcClient
	tx.Client
}

func NewClient(sdkConfigFileAbsPath string) (Client, error) {
	config.Init(sdkConfigFileAbsPath)

	var fake client
	km, err := keys.NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd)
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient(config.LiteClientRpcEndpoint)
	lcdClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(config.RPCEndpoint)

	status, err := rpcClient.GetNodeStatus()
	if err != nil {
		return fake, err
	}

	if config.ChainID != status.NodeInfo.Network {
		return fake, errors.New(fmt.Sprintf("chainID dismatch:expected chainID[%s], actual chainID[%s]", config.ChainID, status.NodeInfo.Network))
	}

	txClient := tx.NewClient(status.NodeInfo.Network, km, lcdClient, rpcClient)

	client := client{
		HttpClient: basicClient,
		LcdClient:  lcdClient,
		RpcClient:  rpcClient,
		Client:     txClient,
	}

	return client, nil
}

func NewTxClient(sdkConfigFileAbsPath string) (tx.Client, error) {
	config.Init(sdkConfigFileAbsPath)

	km, err := keys.NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd)
	if err != nil {
		panic(err)
	}

	basicClient := basic.NewClient(config.LiteClientRpcEndpoint)
	liteClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(config.RPCEndpoint)

	status, err := rpcClient.GetNodeStatus()
	if err != nil {
		return nil, err
	}

	if config.ChainID != status.NodeInfo.Network {
		return nil, errors.New(fmt.Sprintf("chainID dismatch:expected chainID[%s], actual chainID[%s]", config.ChainID, status.NodeInfo.Network))
	}

	client := tx.NewClient(status.NodeInfo.Network, km, liteClient, rpcClient)

	return client, nil
}

func NewLcdClient(sdkConfigFileAbsPath string) (lcd.LcdClient, error) {
	config.Init(sdkConfigFileAbsPath)

	basicClient := basic.NewClient(config.LiteClientRpcEndpoint)
	liteClient := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient(config.RPCEndpoint)

	status, err := rpcClient.GetNodeStatus()
	if err != nil {
		return nil, err
	}

	if config.ChainID != status.NodeInfo.Network {
		return nil, errors.New(fmt.Sprintf("chainID dismatch:expected chainID[%s], actual chainID[%s]", config.ChainID, status.NodeInfo.Network))
	}

	return liteClient, nil
}
