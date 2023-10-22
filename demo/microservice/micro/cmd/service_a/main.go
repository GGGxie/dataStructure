package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/GGGxie/dataStructure/microservice/micro/service/service_a"
)

func main() {
	existChan := make(chan bool)
	svc := service_a.NewService()
	if err := svc.Run(existChan); err != nil {
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
