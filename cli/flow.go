package cli

import (
	"errors"
	"flag"
	"regexp"
	"strings"

	"github.com/ubozov/grpc-example/data"
	"github.com/ubozov/grpc-example/grpc/client"
	"github.com/ubozov/grpc-example/grpc/server"
	"github.com/ubozov/grpc-example/statik"
)

func databaseCreateFlow(args []string) {
	connStr := flag.String("conn", "", "Database connection string")

	flag.CommandLine.Parse(args[1:])

	if *connStr == "" {
		panic(errors.New("connection string is not detected"))
	}

	if err := createDatabase(*connStr); err != nil {
		panic("failed to create database: " + err.Error())
	}
}

func createDatabase(connStr string) error {
	db, err := data.NewDB(connStr)
	if err != nil {
		return err
	}
	defer data.CloseDB(db)

	file, err := statik.ReadFile("/scripts/create_db.sql")
	if err != nil {
		return err
	}

	s := string(file)
	re := regexp.MustCompile(`(?m)\\connect \w*`)
	separator := re.FindString(s)

	i := strings.LastIndex(s, separator)
	createStatements := s[:i]

	for _, query := range strings.Split(createStatements, ";") {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	// switch to grpc-example database and configurate
	re = regexp.MustCompile(`(?m)dbname=\w*`)
	connStr = re.ReplaceAllString(connStr, `dbname=`+separator[8:])
	conn, err := data.NewDB(connStr)
	if err != nil {
		return err
	}
	defer data.CloseDB(conn)

	configStatements := s[i+len(separator):]
	for _, query := range strings.Split(configStatements, ";") {
		if _, err := conn.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func grpcServerStartFlow(args []string) {
	if err := server.Start(); err != nil {
		panic("failed to run grpc server: " + err.Error())
	}
}

func grpcClientStartFlow(args []string) {
	if err := client.Start(); err != nil {
		panic("failed to run grpc client: " + err.Error())
	}
}
