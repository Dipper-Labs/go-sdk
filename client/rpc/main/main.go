package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tendermint/tendermint/libs/pubsub/query"
	"github.com/tendermint/tendermint/rpc/client"
)

func main() {
	tmRpcClient := client.NewHTTP("tcp://127.0.0.1:26657", "/websocket")
	err := tmRpcClient.Start()
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100000)
	defer cancel()

	//query := query.MustParse("tm.event='Tx' AND message.action='contract_call'")
	query := query.MustParse("tm.event='Tx' AND contract_called.address='dip13uldak6495vup0kjqenwhpa5etyxqmel96r0e5'")
	fmt.Print(query.String())

	var x bool
	fmt.Print(x)

	eventChannel, err := tmRpcClient.Subscribe(ctx, "1", query.String())
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	go func() {
		for event := range eventChannel {
			fmt.Println("event:", event.Events)
		}
	}()

	for true {
		time.Sleep(time.Second)
	}
}
