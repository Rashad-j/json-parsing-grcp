package main

import (
	"os"

	"github.com/rashad-j/go-grpc-json-svc/config"
	"github.com/rashad-j/go-grpc-json-svc/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// zerolog basic config
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})

	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	fileCacheProvider := server.NewCache()
	fileProvider, err := server.NewOSFileProviderService(cfg, fileCacheProvider)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create file provider")
	}

	// add loggerService decorator to fileProvider
	fileProvider = server.NewLoggerService(fileProvider)

	// bootstrap file provider to load all files into memory
	if err := fileProvider.Bootstrap(); err != nil {
		log.Fatal().Err(err).Msg("failed to bootstrap file provider")
	}

	// start watching for file changes
	err = fileProvider.StartWatching()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start watching")
	}
	defer fileProvider.StopWatching()

	// start server
	server.MakeJsonParsingServiceServer(fileProvider)
}
