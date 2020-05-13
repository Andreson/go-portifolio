package notify_service

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top-notify/config"
	"log"
	"sync"
	"context"
	"fmt"
)

func ReadMesage(subscriptionName string)([]byte, error){
	client, ctx :=config.PubSubConfig()
	var mu sync.Mutex
	var response  []byte
	received := 0
	sub := client.Subscription(subscriptionName)
	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()

		var teste Test

		err := json.Unmarshal(msg.Data , &teste)

		if err !=nil {
			log.Println("error log msg ", err)
		}
		fmt.Printf("Got message: %q\n", teste)
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 10 {
			cancel()
		}
		response = msg.Data
	})
	if err != nil {
		return nil, err
	}
	return response, nil

}

type Test struct {
	Nome string
	Idade int
}