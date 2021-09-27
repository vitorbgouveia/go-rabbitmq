package usecases

import (
	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/go-rabbitmq/configs/stream"
	"github.com/vitorbgouveia/go-rabbitmq/internal/app/producers"
)

// HelloWord execute usecase and call producer
func HelloWord(messages []string) {
	var rabbitmqConn *amqp.Connection = stream.GetConnection()
	producers.HellorWord(rabbitmqConn, messages)
}
