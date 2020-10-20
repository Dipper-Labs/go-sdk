package lcd

import "github.com/Dipper-Labs/go-sdk/client/types"

const (
	UriQueryAccount      = "/auth/accounts/%s"
	UriQueryContractLogs = "/vm/logs/%s"
)

type (
	AccountValue struct {
		Address       string       `json:"address"`
		Coins         []types.Coin `json:"coins"`
		AccountNumber string       `json:"account_number"`
		Sequence      string       `json:"sequence"`
	}

	AccountResult struct {
		Type  string       `json:"type"`
		Value AccountValue `json:"value"`
	}

	AccountBody struct {
		Height string        `json:"height"`
		Result AccountResult `json:"result"`
	}
)

type (
	ContractLog struct {
		Height string `json:"height"`
		Result VMLogs `json:"result"`
	}

	VMLogs struct {
		Logs []VMLog `json:"logs"`
	}

	VMLog struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
		BlockHash        string   `json:"blockHash"`
		LogIndex         string   `json:"logIndex"`
		Removed          bool     `json:"removed"`
	}
)
