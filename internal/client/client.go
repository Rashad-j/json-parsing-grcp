package client

import (
	"github.com/pkg/errors"
	"github.com/rashad-j/go-grpc-json-svc/config"
	"github.com/rashad-j/go-grpc-json-svc/rpc/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MakeJsonParsingServiceClient creates a new client to connect to the server
// and returns the client and the connection
func MakeJsonParsingServiceClient(cfg *config.Config) (parser.JsonParsingServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to dial server")
	}

	client := parser.NewJsonParsingServiceClient(conn)

	return client, conn, nil
}
