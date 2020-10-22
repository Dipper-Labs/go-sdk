package rpc

import (
	"context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *client) Subscribe(ctx context.Context, subscriber, query string) (out <-chan ctypes.ResultEvent, err error) {
	return c.rpc.Subscribe(ctx, subscriber, query)
}

func (c *client) UnSubscribe(ctx context.Context, subscriber, query string) error {
	return c.rpc.Unsubscribe(ctx, subscriber, query)
}

func (c *client) UnSubscribeAll(ctx context.Context, subscriber string) error {
	return c.rpc.UnsubscribeAll(ctx, subscriber)
}

func (c *client) Start() error {
	return c.rpc.Start()
}
