package controller

import (
	"bytes"
	"distributed/dto"
	"distributed/qutils"
	"encoding/gob"
	"fmt"

	"github.com/streadway/amqp"
)

const url = "amqp://guest:guest@localhost:5672"

type QueueListener struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	sources map[string]<-chan amqp.Delivery
	ea      *EventAgregator
}

func NewQueueListener() *QueueListener {
	ql := QueueListener{
		sources: make(map[string]<-chan amqp.Delivery),
		ea:      NewEventAgregator(),
	}
	ql.conn, ql.ch = qutils.GetChannel(url)
	return &ql
}

func (ql *QueueListener) ListenForNewSources() {
	q := qutils.GetQueue("", ql.ch)
	ql.ch.QueueBind(
		q.Name,
		"",
		"amq.fanout",
		false,
		nil,
	)

	msgs, _ := ql.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	for msg := range msgs {
		sourceChan, _ := ql.ch.Consume(
			string(msg.Body),
			"",
			true,
			false,
			false,
			false,
			nil)
		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(sd)
		ed := EventData{
			Name:      sd.Name,
			Timestamp: sd.Timestamp,
			Value:     sd.Value,
		}
		ql.ea.PublishEvent("Messagereceived_"+msg.RoutingKey, ed)
		fmt.Printf("Received message: %v\n", sd)
	}
}