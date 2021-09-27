package stream

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/go-rabbitmq/internal/app/consumers"
)

type configConn struct {
	host     string
	user     string
	password string
	port     int
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Start stream messages
func Start() {
	var connection *amqp.Connection = GetConnection()
	consumers.Init(connection)
}

func getConfigConnection() configConn {
	port, _ := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))

	return configConn{
		host:     os.Getenv("RABBITMQ_HOST"),
		user:     os.Getenv("RABBITMQ_USER"),
		password: os.Getenv("RABBITMQ_PASSWORD"),
		port:     port,
	}
}

// GetConnection return connection database
func GetConnection() *amqp.Connection {
	var configConn configConn = getConfigConnection()
	strConn := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		configConn.user, configConn.password, configConn.host, configConn.port)

	conn, err := amqp.Dial(strConn)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
