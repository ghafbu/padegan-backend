package config

import "time"

type RateLimiterConfig struct {
	Max        int
	Expiration time.Duration
}
