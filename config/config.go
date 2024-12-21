package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sushi-backend/internal/logger"
	"time"
)

type Config struct {
	logger                         logger.ILogger
	httpPort                       string
	postgresDSN                    string
	cloudinaryUrl                  string
	cloudinaryFolder               string
	telegramBotToken               string
	telegramOrdersChatId           string
	telegramDeliveryChatId         string
	runMigration                   bool
	ipRateLimitRate                int
	ipRateLimitBurst               int
	ipRateLimitExpirationInMs      int
	ipRateLimitCleanupIntervalInMs int
	errorStackTraceSizeInKb        int
	maxFileSizeInMb                int
}

func NewConfig(deps ConfigDependencies) *Config {
	_logger := deps.Logger

	if err := godotenv.Load(); err != nil {
		_logger.Error("No .env file found")
	}

	config := &Config{
		logger: _logger,
	}

	config.postgresDSN = config.getRequiredString("POSTGRES_DSN")
	config.telegramBotToken = config.getRequiredString("TELEGRAM_BOT_TOKEN")
	config.telegramOrdersChatId = config.getRequiredString("TELEGRAM_ORDERS_CHAT_ID")
	config.telegramDeliveryChatId = config.getRequiredString("TELEGRAM_DELIVERY_CHAT_ID")
	config.cloudinaryUrl = config.getRequiredString("CLOUDINARY_URL")
	config.cloudinaryFolder = config.getOptionalString("CLOUDINARY_FOLDER", "sushi")
	config.runMigration = config.getOptionalBool("RUN_MIGRATION", true)
	config.httpPort = config.getOptionalString("HTTP_PORT", ":8080")
	config.ipRateLimitRate = config.getOptionalInt("IP_RATE_LIMIT_RATE", 60)
	config.ipRateLimitBurst = config.getOptionalInt("IP_RATE_LIMIT_BURST", 20)
	config.ipRateLimitExpirationInMs = config.getOptionalInt("IP_RATE_LIMIT_EXPIRATION_IN_MS", 360_000)
	config.ipRateLimitCleanupIntervalInMs = config.getOptionalInt("IP_RATE_LIMIT_CLEANUP_INTERVAL_IN_MS", 360_000)
	config.errorStackTraceSizeInKb = config.getOptionalInt("ERROR_STACK_TRACE_SIZE_IN_KB", 4)
	config.maxFileSizeInMb = config.getOptionalInt("MAX_FILE_SIZE_IN_MB", 200)

	return config
}

func (c *Config) PostgresDSN() string {
	return c.postgresDSN
}

func (c *Config) TelegramBotToken() string {
	return c.telegramBotToken
}

func (c *Config) TelegramOrdersChatId() string {
	return c.telegramOrdersChatId
}

func (c *Config) TelegramDeliveryChatId() string {
	return c.telegramDeliveryChatId
}

func (c *Config) CloudinaryUrl() string {
	return c.cloudinaryUrl
}

func (c *Config) CloudinaryFolder() string {
	return c.cloudinaryFolder
}

func (c *Config) RunMigration() bool {
	return c.runMigration
}

func (c *Config) HttpPort() string {
	return c.httpPort
}

func (c *Config) IpRateLimitRate() int {
	return c.ipRateLimitRate
}

func (c *Config) IpRateLimitBurst() int {
	return c.ipRateLimitBurst
}

func (c *Config) IpRateLimitExpiration() time.Duration {
	return time.Duration(c.ipRateLimitExpirationInMs) * time.Millisecond
}

func (c *Config) IpRateLimitCleanupInterval() time.Duration {
	return time.Duration(c.ipRateLimitCleanupIntervalInMs) * time.Millisecond
}

func (c *Config) ErrorStackTraceSizeInKb() int {
	return c.errorStackTraceSizeInKb
}

func (c *Config) MaxFileSizeInMb() int64 {
	return int64(c.maxFileSizeInMb)
}

func (c *Config) getRequiredString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		c.logger.Error(`Environment variable "` + key + `" not found`)
		panic(fmt.Sprintf(`Environment variable "%s" not found`, key))
	}

	return value
}

func (c *Config) getOptionalString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		c.logger.Warn(`Environment variable "` + key + `" not found, used default ` + defaultValue)
		value = defaultValue
	}

	return value
}

func (c *Config) getOptionalInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	if value == "" {
		c.logger.Warn(fmt.Sprintf(`Environment variable "%s" not found, used default %d`, key, defaultValue))
		return defaultValue
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return valueInt
}

func (c *Config) getOptionalBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)

	if value == "" {
		c.logger.Warn(fmt.Sprintf(`Environment variable "%s" not found, used default %t`, key, defaultValue))
		return defaultValue
	}

	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return valueBool
}
