package client

import (
	"context"
	"fmt"
	lcdtype "github.com/Dipper-Labs/go-sdk/client/lcd"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/pubsub/query"
	"testing"
	"time"
)

func Test_Event(t *testing.T) {
	client, err := NewClient(yamlPath)
	require.Nil(t, err)

	err = client.Start()
	require.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100000)
	defer cancel()

	//query := query.MustParse("tm.event='Tx' AND message.action='contract_call'")
	//query := query.MustParse("tm.event='Tx' AND contract_called.address='dip13uldak6495vup0kjqenwhpa5etyxqmel96r0e5'")
	query := query.MustParse("tm.event='Tx'")
	fmt.Print(query.String())

	eventChannel, err := client.Subscribe(ctx, "1", query.String())
	fmt.Print(query.String())

	go func() {
		for event := range eventChannel {
			fmt.Println("event:", event.Events)
		}
	}()

	for true {
		time.Sleep(time.Second)
	}
}

func Test_SubVmEventWithTopic(t *testing.T) {
	client, err := NewClient(yamlPath)
	require.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100000)
	defer cancel()

	c, err := client.SubVmEventWithTopic(ctx, "1", "dip1y2mfz8ca284z8emzuzrta7r9vkvc84gplw3vq6", "xxx()")
	require.Nil(t, err)

	for e := range c {
		switch i := e.(type) {
		case lcdtype.VMLog:
			t.Log(fmt.Sprintf("%v", i))
		}
	}
}
