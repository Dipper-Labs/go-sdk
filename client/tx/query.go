package tx

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/vm/types"
)

func (c *client) QueryContractEvents(contractBech32Addr string, startBlockNum int64, endBlockNum int64) (result []string, err error) {
	for i := startBlockNum; i < endBlockNum; i++ {
		blockResult, err := c.rpcClient.GetBlock(i)
		if err != nil {
			continue
		}
		txs := blockResult.Block.Data.Txs
		for _, tx := range txs {
			txhash := crypto.Sha256(tx)
			fmt.Println(fmt.Sprintf("block %d, txhash: %x", i, txhash))

			if res, err := c.rpcClient.GetTx(hex.EncodeToString(txhash)); err == nil {
				msg := res.StdTx.Msgs[0]
				if msg.Type() == types.TypeMsgContract {
					msgContract, _ := msg.(types.MsgContract)
					targetContractAddr := msgContract.To.String()
					if contractBech32Addr != targetContractAddr {
						continue
					}

					eventLog, _ := c.lcdClient.QueryContractLog(txhash)
					for _, e := range eventLog.Result.Logs {
						result = append(result, e.Data)
					}
				}
			}
		}
	}

	return result, nil
}
