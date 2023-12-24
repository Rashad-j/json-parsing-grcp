package server

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/xid"

	"github.com/rashad-j/go-grpc-json-svc/config"
	"github.com/rashad-j/go-grpc-json-svc/rpc/parser"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// The server is the gRPC server that implements the JsonParsingServiceServer interface

type jsonParsingServer struct {
	fileProvider FileProviderService
	parser.UnimplementedJsonParsingServiceServer
}

func NewJsonParsingServer(fileListingProvider FileProviderService) *jsonParsingServer {
	return &jsonParsingServer{
		fileProvider: fileListingProvider,
	}
}

func (s *jsonParsingServer) ParseJsonFiles(ctx context.Context, req *parser.EmptyRequest) (*parser.JsonResponse, error) {
	var personList []*parser.JsonResponse_Person

	for _, data := range s.fileProvider.GetAll(ctx) {
		// check if data of type parser.JsonResponse_Person
		var person parser.JsonResponse_Person
		if err := json.Unmarshal(data, &person); err != nil {
			log.Printf("failed to unmarshal data %v", err)
			continue
		}
		personList = append(personList, &person)
	}

	if len(personList) == 0 {
		return nil, errors.New("no data found")
	}

	return &parser.JsonResponse{PersonList: personList}, nil
}

func MakeJsonParsingServiceServer(fileProvider FileProviderService) error {
	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read config")
	}

	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	parser.RegisterJsonParsingServiceServer(s, NewJsonParsingServer(fileProvider))

	// Enable reflection
	reflection.Register(s)

	log.Info().Msgf("gRPC listening on %s", cfg.Addr)

	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

// loggingInterceptor is a unary interceptor that logs each incoming RPC.
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// set requestID in context
	requestID := xid.New().String()
	ctx = context.WithValue(ctx, "requestID", requestID)

	// metrics & logging
	log.Info().Str("requestID", requestID).Str("method", info.FullMethod).Msg("starting RPC request")
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Str("requestID", requestID).Dur("duration", duration).Msg("finished RPC request")
	}(time.Now())

	// calling the actual handler to process the RPC
	resp, err := handler(ctx, req)

	return resp, err
}
