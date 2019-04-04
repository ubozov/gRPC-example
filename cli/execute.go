package cli

import (
	"fmt"
	"log"
	"strings"
)

// db-create - command to create database
func Execute(args []string) {
	if len(args) == 0 {
		args = append(args, "help")
	}

	switch strings.ToLower(args[0]) {
	case "db-create":
		log.Println("create database")
		databaseCreateFlow(args)
	case "help":
		fmt.Println(helpMessage)
		return
	default:
		fmt.Println(helpMessage)
		return
	}

	log.Println("command was successfully executed")
}
