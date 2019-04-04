package main

import (
	"log"
	"os"

	"github.com/ubozov/grpc-example/cli"
)

func main() {
	log.Println("start program")
	defer log.Println("end program")

	cli.Execute(os.Args[1:])
}
