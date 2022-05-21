package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/Shopify/sarama.v1"
)

// conexão com o kafka
func ConnectionProduce(brokerUrs []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokerUrs, config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

//Publicando mensagens na fila

func PushCommentToQueue(topic string, menssege []byte) error {
	brokerUrl := []string{"kafkahost1:9092", "kafkahost2:9092"}
	producer, err := ConnectionProduce(brokerUrl)

	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(menssege),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\\n", topic, partition, offset)
	return nil
}

type Comment struct {
	Text string `from:"text" json:"text"`
}

func createComment(c *fiber.Ctx) error {
	cmt := new(Comment)
	if err := c.BodyParser(cmt); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	// converte o corpo em bytes e envia para kafka
	cmtInBytes, err := json.Marshal(cmt)
	PushCommentToQueue("comments", cmtInBytes)
	// Retorna Comentário no formato JSON
	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comentário enviado com sucesso",
		"comment": cmt,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"sucesso": false,
			"message": "Erro ao criar produto",
		})
		return err
	}
	return err
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comment", createComment)
	app.Listen(":3001")
}
