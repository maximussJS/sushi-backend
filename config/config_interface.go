package config

import (
	"time"
)

type IConfig interface {
	HttpPort() string

	CloudinaryUrl() string
	CloudinaryFolder() string

	PostgresDSN() string

	RunMigration() bool

	IpRateLimitRate() int
	IpRateLimitBurst() int
	IpRateLimitExpiration() time.Duration
	IpRateLimitCleanupInterval() time.Duration
}
