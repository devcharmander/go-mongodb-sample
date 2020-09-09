package main

import (
	"github.com/devcharmander/100-day-habits/grpc/client"
	"github.com/devcharmander/100-day-habits/grpc/server"
)

func main() {
	chb := make(chan bool)
	go server.Start()
	client.Init()
	<-chb
}
