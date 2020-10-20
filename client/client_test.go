package client

import (
	"testing"

	"github.com/Dipper-Labs/go-sdk/client/types"
	"github.com/Dipper-Labs/go-sdk/util"
	"github.com/stretchr/testify/require"
)

const yamlPath = "/Users/sun/go/src/github.com/Dipper-Labs/go-sdk/config/sdk.yaml"

func TestNewClient(t *testing.T) {
	c, err := NewClient(yamlPath)
	if err != nil {
		t.Fatal(err)
	} else {
		if res, err := c.QueryAccount("dip1lzydk8mxwtm3gjszf94lq2jkfgqyd3y7yh3tc8"); err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToJsonIgnoreErr(res))
		}
	}
}

func TestClient_SendToken(t *testing.T) {
	c, err := NewTxClient(yamlPath)
	require.True(t, err == nil)

	coins := []types.Coin{
		{
			Denom:  "pdip",
			Amount: "100",
		},
	}

	if res, err := c.SendToken("dip1dcu73lw9uqkygpde4z4z22f079skta49vxs2r0", coins, "", false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
