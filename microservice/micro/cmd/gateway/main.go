package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	gw "github.com/GGGxie/dataStructure/microservice/micro/service/gateway"
)

func main() {
	existChan := make(chan bool)
	gateway := gw.NewGateway()
	if err := gateway.Run(existChan); err != nil {
		fmt.Printf("run service error %v\n", err)
		return
	}
	// listen system call
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	sig := <-sc
	existChan <- true
	switch sig {
	case syscall.SIGTERM:
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
