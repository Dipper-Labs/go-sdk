package tx

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/vm/types"
	"github.com/Dipper-Labs/Dipper-Protocol/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/crypto"
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
				if msg.Type() == types.TypeMsgContractCreate || msg.Type() == types.TypeMsgContractCall {
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

func (c *client) SubVmEventWithTopic(ctx context.Context, subscriber, contractAddr, topic string) (out <-chan interface{}, err error) {
	err = c.rpcClient.Start()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("tm.event='Tx' AND contract_called.address='%s'", contractAddr)
	ch, err := c.rpcClient.Subscribe(ctx, subscriber, query)
	if err != nil {
		return nil, err
	}

	topicHash := fmt.Sprintf("%x", ethcrypto.Keccak256([]byte(topic)))

	c1 := make(chan interface{})

	go func() {
		for event := range ch {
			for _, txIdStr := range event.Events["tx.hash"] {
				txId, err := hexutil.Decode(txIdStr)
				if err != nil {
					continue
				}

				clog, err := c.lcdClient.QueryContractLog(txId)
				if err != nil {
					continue
				}

				for _, log := range clog.Result.Logs {
					for _, t := range log.Topics {
						if t == topicHash {
							c1 <- log
						}
					}
				}
			}
		}
	}()

	return c1, nil
}
