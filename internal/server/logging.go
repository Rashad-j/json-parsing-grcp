package server

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

// loggerService is a decorator for FileProviderService
// It logs the duration of each method call and an Info message when the method is called
// Metrics can be added to this decorator to measure the performance of each method

type loggerService struct {
	next FileProviderService
}

// add a compiler check to ensure that loggerService implements FileProviderService
var _ FileProviderService = (*loggerService)(nil)

func NewLoggerService(next FileProviderService) *loggerService {
	return &loggerService{
		next: next,
	}
}

func (s *loggerService) GetAll(ctx context.Context) map[string][]byte {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}

	log.Info().Msg("GetAll called")
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Str("requestID", requestID).Dur("duration", duration).Msg("finished GetAll")
	}(time.Now())

	return s.next.GetAll(ctx)
}

func (s *loggerService) Bootstrap() error {
	log.Info().Msg("Bootstrap called")
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Dur("duration", duration).Int("num of files", len(s.GetAll(context.Background()))).Msg("finished Bootstrap")
	}(time.Now())

	return s.next.Bootstrap()
}

func (s *loggerService) StartWatching() error {
	log.Info().Msg("StartWatching called")
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Dur("duration", duration).Msg("StartWatching is running in the background")
	}(time.Now())

	return s.next.StartWatching()
}

func (s *loggerService) StopWatching() {
	log.Info().Msg("StopWatching called")
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Dur("duration", duration).Msg("finished StopWatching")
	}(time.Now())

	s.next.StopWatching()
}
