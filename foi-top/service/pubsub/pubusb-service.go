package notific_service

import (
	"bytes"
	"cloud.google.com/go/pubsub"
	"encoding/gob"
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
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		log.Fatal("Erro ao conventer conteudo da msg ",err)
		return nil
	}
	return buf.Bytes()
}