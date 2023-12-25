package server

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/rashad-j/go-grpc-json-svc/internal/config"
	"github.com/rs/zerolog/log"
)

// FileProviderService is an interface that defines the methods that a file provider service should implement
// The file provider service is responsible for providing the files to the server
// The file provider service should be able to bootstrap itself and start watching for file changes

type FileProviderService interface {
	Bootstrap() error
	GetAll(context.Context) map[string][]byte
	StartWatching() error
	StopWatching()
}

type FileCacheProvider interface {
	Get(key string) ([]byte, bool)
	GetAll() map[string][]byte
	Set(key string, val []byte)
	Delete(key string)
}

type OSFileProviderService struct {
	directory string
	extension string
	watcher   *fsnotify.Watcher

	// fileCacheProvider is used to cache files in memory
	fileCacheProvider FileCacheProvider
}

// add a compiler check to ensure that OSFileProviderService implements FileProviderService
var _ FileProviderService = (*OSFileProviderService)(nil)

func NewOSFileProviderService(cfg *config.Config, fileCacheProvider FileCacheProvider) (FileProviderService, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create watcher")
	}

	return &OSFileProviderService{
		directory:         cfg.FilesDirectory,
		extension:         cfg.Extension,
		watcher:           watcher,
		fileCacheProvider: fileCacheProvider,
	}, nil
}

func (p *OSFileProviderService) watchFiles(ch chan<- []string) {
	for {
		select {
		case event, ok := <-p.watcher.Events:
			if !ok {
				return
			}

			fileName := filepath.Base(event.Name)
			if filepath.Ext(fileName)[1:] != p.extension {
				log.Info().Str("file", fileName).Msg("skipped non-JSON file")
				continue
			}
			// check if the file was removed
			if event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
				log.Printf("removed file: %s", fileName)
				p.fileCacheProvider.Delete(fileName)
				continue
			}

			// check if the file was created or modified
			if event.Op&fsnotify.Create != fsnotify.Create && event.Op&fsnotify.Write != fsnotify.Write {
				log.Info().Str("operation", event.Op.String()).Msg("skipped non-create, non-write, non-rename event")
				continue
			}

			fileContent, err := p.getFileContent(fileName)
			if err != nil {
				log.Error().Err(err).Str("file", fileName).Msg("failed to get file content")
				continue
			}

			// update the cache
			p.fileCacheProvider.Set(fileName, fileContent)

			log.Info().Str("file", fileName).Msg("successfully cached file")

		case err, ok := <-p.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("watcher error: %v", err)
		}
	}
}

func (p *OSFileProviderService) StartWatching() error {
	ch := make(chan []string)
	err := p.watcher.Add(p.directory)
	if err != nil {
		return errors.Wrap(err, "failed to add directory to watcher")
	}
	go p.watchFiles(ch)

	return nil
}

func (p *OSFileProviderService) StopWatching() {
	p.watcher.Close()
}

func (p *OSFileProviderService) listFiles() ([]fs.DirEntry, error) {
	return os.ReadDir(p.directory)
}

func (p *OSFileProviderService) Bootstrap() error {
	fileList, err := p.listFiles()
	if err != nil {
		return errors.Wrap(err, "failed to list files")
	}

	if len(fileList) == 0 {
		return errors.New("no JSON files found")
	}

	for _, fileInfo := range fileList {
		if fileInfo.IsDir() {
			// skip directories
			continue
		}

		// check if the file has the right extension
		if filepath.Ext(fileInfo.Name())[1:] != p.extension {
			log.Error().Str("file", fileInfo.Name()).Msg("skipped non-JSON file")
			continue
		}

		fileContent, err := p.getFileContent(fileInfo.Name())
		if err != nil {
			log.Error().Err(err).Str("file", fileInfo.Name()).Msg("failed to get file content")
			continue
		}
		// update the cache
		p.fileCacheProvider.Set(fileInfo.Name(), fileContent)
	}

	if len(fileList) == 0 {
		return errors.New("no valid JSON files found")
	}

	return nil
}

func (p *OSFileProviderService) getFileContent(fileName string) ([]byte, error) {
	filepath := filepath.Join(p.directory, fileName)
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file: %s", filepath)
	}

	return bytes, nil
}

func (p *OSFileProviderService) GetAll(_ context.Context) map[string][]byte {
	return p.fileCacheProvider.GetAll()
}
