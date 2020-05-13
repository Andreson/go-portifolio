package notific_service

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/config"
	"log"
)

func SendMesage(topicName string, content interface{} ) error {
	client,ctx  := config.PubSubConfig()
	topic :=config.CreateTopicIfNotExists(client, topicName)
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(castContentMensage(content)),
	})
	_, err := res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}

func castContentMensage(key interface{}) ([]byte) {
	msg, error :=json.Marshal(key)
	if error != nil {
		log.Fatal("Erro ao conventer conteudo da msg ",error)
		return nil
	}
	return msg
}