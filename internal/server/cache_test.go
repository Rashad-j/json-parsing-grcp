package server

// test cases for Cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	// arrange

	// act
	cache := NewCache()

	// assert
	assert.NotNil(t, cache)
}

func TestCache_Get(t *testing.T) {
	// arrange
	cache := NewCache()
	cache.Set("test", []byte("test"))

	// act
	val, ok := cache.Get("test")

	// assert
	assert.Equal(t, []byte("test"), val)
	assert.True(t, ok)
}

func TestCache_GetAll(t *testing.T) {
	// arrange
	cache := NewCache()
	cache.Set("test", []byte("test"))

	// act
	data := cache.GetAll()

	// assert
	assert.Equal(t, []byte("test"), data["test"])
}

func TestCache_Delete(t *testing.T) {
	// arrange
	cache := NewCache()
	cache.Set("test", []byte("test"))

	// act
	cache.Delete("test")

	// assert
	assert.Equal(t, 0, cache.Len())
}

func TestCache_Set(t *testing.T) {
	// arrange
	cache := NewCache()

	// act
	cache.Set("test", []byte("test"))

	// assert
	assert.Equal(t, []byte("test"), cache.data["test"])
}

func TestCache_Len(t *testing.T) {
	// arrange
	cache := NewCache()
	cache.Set("test", []byte("test"))

	// act
	length := cache.Len()

	// assert
	assert.Equal(t, 1, length)
}
