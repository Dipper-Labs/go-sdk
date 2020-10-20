package lcd

import (
	"github.com/Dipper-Labs/go-sdk/client/basic"
)

type LcdClient interface {
	QueryAccount(address string) (AccountBody, error)
	QueryContractLog(txId []byte) (ContractLog, error)
}

type client struct {
	httpClient basic.HttpClient
}

func NewClient(c basic.HttpClient) LcdClient {
	return &client{
		httpClient: c,
	}
}
