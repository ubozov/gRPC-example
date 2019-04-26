package cli

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

// Execute executes commands.
// db-create - command to create database
func Execute(args []string) {
	if len(args) == 0 {
		args = append(args, "help")
	}

	switch strings.ToLower(args[0]) {
	case "db-create":
		log.Println("create database")
		databaseCreateFlow(args)
	case "start-server":
		log.Println("start grpc server")
		grpcServerStartFlow(args)
	case "help":
		fmt.Println(helpMessage)
		return
	default:
		fmt.Println(helpMessage)
		return
	}

	log.Println("command was successfully executed")
}

// Shutdown graceful shutdowns program.
func Shutdown(ctx context.Context) {
	log.Println("The program is shutting down...")
	time.Sleep(time.Second)
	log.Println("Done")
}
