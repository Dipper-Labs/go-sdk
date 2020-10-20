package rpc

import (
	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *client) GetBlock(height int64) (*tmrpctypes.ResultBlock, error) {
	return c.rpc.Block(&height)
}
