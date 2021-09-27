package consumers

import "github.com/streadway/amqp"

// Init consumers
func Init(conn *amqp.Connection) {
	handlerHellorWord(conn)
}
