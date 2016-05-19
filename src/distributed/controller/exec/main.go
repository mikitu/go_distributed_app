package main

import (
	"distributed/controller"
	"fmt"
)

func main() {
	ql := controller.NewQueueListener()
	go ql.ListenForNewSources()
	var a string
	fmt.Scanln(&a)
}
