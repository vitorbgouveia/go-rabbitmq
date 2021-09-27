package producers

import (
	"log"

	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/go-rabbitmq/tools"
)

// PubSub producer
func PubSub(conn *amqp.Connection, messages []string) {
	ch, err := conn.Channel()
	tools.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	tools.FailOnError(err, "Failed to declare an exchange")

	for _, m := range messages {
		body := m
		err = ch.Publish(
			"logs", // exchange
			"",     // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		tools.FailOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent %s", body)
	}
}
