package config

import (
	"log"
	"time"
)

const (
	refreshInterval = 24 * time.Hour
)

type callback func(cfg *Config) error

type refresher struct {
	cfg      *Config
	interval time.Duration
	cb       callback
}

func NewConfigRefresher(cfg *Config, interval time.Duration, cb callback) *refresher {
	return &refresher{
		cfg:      cfg,
		interval: interval,
		cb:       cb,
	}
}

func (r *refresher) Start() {
	r.update()

	go func() {
		ticker := time.NewTicker(r.interval)
		defer ticker.Stop()
		for range ticker.C {
			r.cfg.load()
			if err := r.cb(r.cfg); err != nil {
				log.Println("config update callback failed")
			}
		}
	}()
}

func (r *refresher) update() {
	r.cfg.load()
	if err := r.cb(r.cfg); err != nil {
		log.Println("config update callback failed")
	}
}
