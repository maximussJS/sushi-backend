package config

import (
	"sushi-backend/constants"
	"time"
)

type IConfig interface {
	AppEnv() constants.AppEnv

	AllowedOrigins() []string
	AllowedMethods() []string
	AllowedHeaders() []string
	AllowCredentials() bool

	HttpPort() string

	TelegramBotToken() string
	TelegramOrdersChatId() string
	TelegramDeliveryChatId() string

	CloudinaryUrl() string
	CloudinaryFolder() string

	PostgresDSN() string

	RunMigration() bool

	IpRateLimitRate() int
	IpRateLimitBurst() int
	IpRateLimitExpiration() time.Duration
	IpRateLimitCleanupInterval() time.Duration

	ErrorStackTraceSizeInKb() int
	MaxFileSizeInMb() int64
}
