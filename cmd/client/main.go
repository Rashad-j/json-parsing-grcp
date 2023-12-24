package main

import (
	"context"
	"log"

	"github.com/rashad-j/go-grpc-json-svc/config"
	"github.com/rashad-j/go-grpc-json-svc/internal/client"
	"github.com/rashad-j/go-grpc-json-svc/rpc/parser"
)

func main() {
	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	// establish connection to server
	client, conn, err := client.MakeJsonParsingServiceClient(cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// call ParseJsonFiles
	response, err := client.ParseJsonFiles(context.Background(), &parser.EmptyRequest{})
	if err != nil {
		log.Fatalf("Error calling ParseJsonFiles: %v", err)
	}

	for _, person := range response.PersonList {
		log.Printf("Person: %v", person)
	}
}
