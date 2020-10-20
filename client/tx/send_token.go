package tx

import (
	"fmt"
	"strconv"

	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/auth"
	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/bank"
	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	clitypes "github.com/Dipper-Labs/go-sdk/client/types"
	"github.com/Dipper-Labs/go-sdk/config"
	"github.com/Dipper-Labs/go-sdk/constants"
	"github.com/Dipper-Labs/go-sdk/types"
)

func (c *client) SendToken(toAddrBech32 string, amount []clitypes.Coin, memo string, commit bool) (clitypes.BroadcastTxResult, error) {
	var result clitypes.BroadcastTxResult

	toAddr, err := sdk.AccAddressFromBech32(toAddrBech32)
	if err != nil {
		return result, err
	}

	coins, err := buildCoins(amount)
	if err != nil {
		return result, err
	}

	fromAddr := c.KeyManager.GetAddr()
	msg := bank.MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      coins,
	}

	toSpend := sdk.NewInt(config.TxDefaultFeeAmount)
	for _, coin := range coins {
		if coin.Denom == constants.TxDefaultDenom {
			toSpend = toSpend.Add(coin.Amount)
		}
	}

	accountBody, err := c.lcdClient.QueryAccount(fromAddr.String())
	if err != nil {
		return result, err
	}

	currentAmount := getCoin(accountBody.Result.Value.Coins, constants.TxDefaultDenom)
	if currentAmount.Amount.LT(toSpend) {
		return result, fmt.Errorf("account balance is not enough")
	}

	fee := sdk.Coins{
		{
			Denom:  constants.TxDefaultDenom,
			Amount: sdk.NewInt(config.TxDefaultFeeAmount),
		},
	}

	accountNumber, err := strconv.Atoi(accountBody.Result.Value.AccountNumber)
	if err != nil {
		return result, err
	}

	sequence, err := strconv.Atoi(accountBody.Result.Value.Sequence)
	if err != nil {
		return result, err
	}

	stdSignMsg := types.StdSignMsg{
		ChainID:       c.chainId,
		AccountNumber: uint64(accountNumber),
		Sequence:      uint64(sequence),
		Fee:           auth.NewStdFee(config.TxDefaultGas, fee),
		Msgs:          []sdk.Msg{msg},
		Memo:          memo,
	}

	for _, msg := range stdSignMsg.Msgs {
		if err := msg.ValidateBasic(); err != nil {
			return result, err
		}
	}

	txBytes, err := c.KeyManager.Sign(stdSignMsg)
	if err != nil {
		return result, err
	}

	var txBroadcastType string
	if commit {
		txBroadcastType = constants.TxBroadcastTypeCommit
	} else {
		txBroadcastType = constants.TxBroadcastTypeSync
	}

	return c.rpcClient.BroadcastTx(txBroadcastType, txBytes)
}

func buildCoins(inputCoins []clitypes.Coin) (sdk.Coins, error) {
	var coins []sdk.Coin

	if len(inputCoins) == 0 {
		return coins, nil
	}

	for _, coin := range inputCoins {
		if amount, ok := sdk.NewIntFromString(coin.Amount); ok {
			coins = append(coins, sdk.Coin{
				Denom:  coin.Denom,
				Amount: amount,
			})
		} else {
			return coins, fmt.Errorf("can't parse str to Int, coin is %+v", inputCoins)
		}
	}

	return coins, nil
}

func getCoin(inputCoins []clitypes.Coin, denom string) sdk.Coin {
	for _, coin := range inputCoins {
		if coin.Denom == denom {
			if amount, ok := sdk.NewIntFromString(coin.Amount); ok {
				return sdk.Coin{
					Denom:  coin.Denom,
					Amount: amount,
				}
			}

			break
		}
	}

	return sdk.Coin{}
}
