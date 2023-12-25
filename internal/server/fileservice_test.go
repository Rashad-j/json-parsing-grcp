package server

// unit tests for OSFileProviderService

import (
	"context"
	"testing"

	"github.com/rashad-j/go-grpc-json-svc/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewOSFileProviderService(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "testdata",
		Extension:      ".json",
	}

	// act
	fileCacheProvider := NewCache()
	fileProvider, err := NewOSFileProviderService(cfg, fileCacheProvider)

	_, ok := fileProvider.(*OSFileProviderService)

	// assert
	assert.NotNil(t, fileProvider)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestOSFileProviderService_Bootstrap(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "testdata",
		Extension:      "json",
	}
	fileCacheProvider := NewCache()
	fileProvider, _ := NewOSFileProviderService(cfg, fileCacheProvider)

	// act
	err := fileProvider.Bootstrap()

	// assert
	assert.Nil(t, err)
}

func TestOSFileProviderService_StartWatching(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "testdata",
		Extension:      "json",
	}
	fileCacheProvider := NewCache()
	fileProvider, _ := NewOSFileProviderService(cfg, fileCacheProvider)

	// act
	err := fileProvider.StartWatching()

	// assert
	assert.Nil(t, err)
}

func TestOSFileProviderService_StopWatching(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "testdata",
		Extension:      "json",
	}
	fileCacheProvider := NewCache()
	fileProvider, _ := NewOSFileProviderService(cfg, fileCacheProvider)

	// act
	fileProvider.StopWatching()

	// assert
	assert.NotNil(t, fileProvider.(*OSFileProviderService).watcher)
}

func TestOSFileProviderService_GetAll(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "./testdata",
		Extension:      "json",
	}
	fileCacheProvider := NewCache()
	fileProvider, _ := NewOSFileProviderService(cfg, fileCacheProvider)
	fileProvider.Bootstrap()
	ctx := context.Background()
	// act
	files := fileProvider.GetAll(ctx)

	// assert
	assert.NotNil(t, files)
	assert.Equal(t, 2, len(files))
}

func TestOSFileProviderService_ListFiles(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "testdata",
		Extension:      "json",
	}
	fileCacheProvider := NewCache()
	fileProvider, _ := NewOSFileProviderService(cfg, fileCacheProvider)
	fileProvider.Bootstrap()

	// act
	files, err := fileProvider.(*OSFileProviderService).listFiles()

	// assert
	assert.NotNil(t, files)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(files))
}

func TestOSFileProviderService_GetFileContent(t *testing.T) {
	// arrange
	cfg := &config.Config{
		FilesDirectory: "testdata",
		Extension:      "json",
	}
	fileCacheProvider := NewCache()
	fileProvider, _ := NewOSFileProviderService(cfg, fileCacheProvider)
	fileProvider.Bootstrap()

	// act
	content, err := fileProvider.(*OSFileProviderService).getFileContent("person1.json")

	// assert
	assert.NotNil(t, content)
	assert.Nil(t, err)
}
