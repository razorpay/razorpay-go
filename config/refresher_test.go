package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TODO: might need to change the time once we have remote URL for config

func TestRefresher_Start(t *testing.T) {
	loadCount := 0
	cfg := new()
	cb := func(cfg *Config) error {
		loadCount++
		return nil
	}

	NewConfigRefresher(&cfg, 10*time.Millisecond, cb).Start()

	time.Sleep(95 * time.Millisecond)
	assert.Equal(t, 10, loadCount)
}

func TestRefresher_Update(t *testing.T) {
	loadCount := 0
	cfg := new()
	cb := func(cfg *Config) error {
		loadCount++
		return nil
	}

	NewConfigRefresher(&cfg, 10*time.Millisecond, cb).Start()

	assert.Equal(t, 1, loadCount)
}
