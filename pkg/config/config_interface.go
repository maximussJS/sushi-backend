package config

import (
	"time"
)

type IConfig interface {
	GetHttpPort() string

	GetIpRateLimitRate() int
	GetIpRateLimitBurst() int
	GetIpRateLimitExpiration() time.Duration
	GetIpRateLimitCleanupInterval() time.Duration
}
