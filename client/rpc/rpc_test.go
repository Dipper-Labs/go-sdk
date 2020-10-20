package rpc

import (
	"github.com/Dipper-Labs/go-sdk/types"
	"testing"

	"github.com/Dipper-Labs/go-sdk/util"
)

var (
	c RpcClient
)

func TestMain(m *testing.M) {
	c = NewClient("tcp://127.0.0.1:26657")
	m.Run()
}

func TestClient_GetStatus(t *testing.T) {
	if res, err := c.GetNodeStatus(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func TestClient_GetSyncStatus(t *testing.T) {
	if res, err := c.GetNodeStatus(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func TestClient_GetBlock(t *testing.T) {
	const height int64 = 25
	if res, err := c.GetBlock(height); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func TestClient_GetTx(t *testing.T) {
	hash := "7F8B5483D372D6FF8D69D2E543D2426FE643518894C9639E14615977F33209D1"
	if res, err := c.GetTx(hash); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}

func TestClient_GetTxMsgs(t *testing.T) {
	hash := "7F8B5483D372D6FF8D69D2E543D2426FE643518894C9639E14615977F33209D1"
	if res, err := c.GetTx(hash); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
		var msgSend types.MsgSend
		for _, msg := range res.StdTx.Msgs {
			err := types.Cdc.UnmarshalJSON(msg.GetSignBytes(), &msgSend)
			if err != nil {
				continue
			}

			t.Log(msgSend.FromAddress.String())
			t.Log(msgSend.ToAddress.String())
			t.Log(msgSend.Amount.String())
		}
	}
}
