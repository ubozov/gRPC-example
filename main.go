package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ubozov/grpc-example/cli"
)

func main() {
	log.Println("start program")
	defer log.Println("end program")

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("\ncaught sig: %+v", sig)
		log.Println("Wait for 2 second to finish processing")
		cli.Shutdown(context.Background())
		os.Exit(0)
	}()

	cli.Execute(os.Args[1:])
}
