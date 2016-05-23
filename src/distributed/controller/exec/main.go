package main

import (
    "distributed/controller"
    "fmt"
)

var dc *controller.DatabaseConsumer

func main() {
    ea := controller.NewEventAgregator()
    dc := controller.NewDatabaseConsumer(ea)
    ql := controller.NewQueueListener(ea)
    go ql.ListenForNewSources()
    var a string
    fmt.Scanln(&a)
}
