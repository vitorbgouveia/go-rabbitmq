package producers

import (
	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/go-rabbitmq/tools"
)

// HellorWord producer
func HellorWord(conn *amqp.Connection, messages []string) {
	ch, err := conn.Channel()
	tools.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	tools.FailOnError(err, "Failed to declare a queue")

	for _, m := range messages {
		body := m
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		tools.FailOnError(err, "Failed to publish a message")
	}
}
