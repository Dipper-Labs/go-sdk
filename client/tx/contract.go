package tx

import (
	"fmt"
	"strconv"

	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/auth"
	vmtypes "github.com/Dipper-Labs/Dipper-Protocol/app/v0/vm/types"
	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	clitypes "github.com/Dipper-Labs/go-sdk/client/types"
	"github.com/Dipper-Labs/go-sdk/config"
	"github.com/Dipper-Labs/go-sdk/constants"
	"github.com/Dipper-Labs/go-sdk/types"
)

func (c *client) ContractCall(contractBech32Addr string, payload []byte, amount sdk.Coin, commit bool) (r clitypes.BroadcastTxResult, err error) {
	var result clitypes.BroadcastTxResult

	if amount.Denom != constants.TxDefaultDenom {
		return result, err
	}

	var contractAddr sdk.AccAddress
	if contractBech32Addr != "" {
		contractAddr, err = sdk.AccAddressFromBech32(contractBech32Addr)
		if err != nil {
			return result, err
		}
	}

	from := c.KeyManager.GetAddr()
	accountBody, err := c.lcdClient.QueryAccount(from.String())
	if err != nil {
		return result, err
	}

	accountAmount := getCoin(accountBody.Result.Value.Coins, constants.TxDefaultDenom)

	toSpendAmount := sdk.NewInt(config.TxDefaultFeeAmount)
	if accountAmount.Amount.LT(toSpendAmount.Add(amount.Amount)) {
		return result, fmt.Errorf("account balance is not enough")
	}

	fee := sdk.Coins{
		{
			Denom:  constants.TxDefaultDenom,
			Amount: sdk.NewInt(config.TxDefaultFeeAmount),
		},
	}

	msg := vmtypes.NewMsgContract(from, contractAddr, payload, amount)

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
		Memo:          "",
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
