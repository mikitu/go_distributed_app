package exec

import (
    "bytes"
    "distributed/datamanager"
    "distributed/dto"
    "distributed/qutils"
    "encoding/gob"
    "log"
)

const url = "amqp://guest:guest@localhost:5672"

func main() {
    conn, ch := qutils.GetChannel(url)
    defer conn.Close()
    defer ch.Close()

    msgs, err := ch.Consume(
        qutils.PersistReadingsQueue, //queue string
        "", // consumer string
        false, // autoAck bool
        true, // exclusive bool
        false, //noLocal bool
        false, //noWait bool
        nil, //args amqp.Table
    )

    if err != nil {
        log.Fatal("Failed to get access to messages")
    }

    for msg := range msgs {

    }
}
