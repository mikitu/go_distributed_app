package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const SensorlistQueue = "SensorList"
const SensorDescoveryExchange = "SensorDescovery"
const PersistReadingsQueue = "PersistReadings"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to estabilish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")
	return conn, ch
}
func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name string
		false, //durable bool
		autoDelete, //autoDelete boolean
		false, //exclusive boolean,
		false, //noWait boolean,
		nil)   //args amqp.Table)
	failOnError(err, "Failed to declare queue")
	return &q
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
