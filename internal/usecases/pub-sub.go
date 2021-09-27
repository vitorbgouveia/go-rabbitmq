package usecases

import (
	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/go-rabbitmq/configs/stream"
	"github.com/vitorbgouveia/go-rabbitmq/internal/app/producers"
)

// PubSub execute usecase and call producer
func PubSub(messages []string) {
	var rabbitmqConn *amqp.Connection = stream.GetConnection()
	producers.PubSub(rabbitmqConn, messages)
}
