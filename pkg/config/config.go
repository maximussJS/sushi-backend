package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sushi-backend/pkg/logger"
	"time"
)

type Config struct {
	logger                         logger.ILogger
	httpPort                       string
	ipRateLimitRate                int
	ipRateLimitBurst               int
	ipRateLimitExpirationInMs      int
	ipRateLimitCleanupIntervalInMs int
}

func NewConfig(deps ConfigDependencies) *Config {
	_logger := deps.Logger

	if err := godotenv.Load(); err != nil {
		_logger.Error("No .env file found")
	}

	config := &Config{
		logger: _logger,
	}

	config.httpPort = config.getOptionalString("HTTP_PORT", ":8080")
	config.ipRateLimitRate = config.getOptionalInt("IP_RATE_LIMIT_RATE", 60)
	config.ipRateLimitBurst = config.getOptionalInt("IP_RATE_LIMIT_BURST", 20)
	config.ipRateLimitExpirationInMs = config.getOptionalInt("IP_RATE_LIMIT_EXPIRATION_IN_MS", 20000)
	config.ipRateLimitCleanupIntervalInMs = config.getOptionalInt("IP_RATE_LIMIT_CLEANUP_INTERVAL_IN_MS", 10000)

	return config
}

func (c *Config) GetHttpPort() string {
	return c.httpPort
}

func (c *Config) GetIpRateLimitRate() int {
	return c.ipRateLimitRate
}

func (c *Config) GetIpRateLimitBurst() int {
	return c.ipRateLimitBurst
}

func (c *Config) GetIpRateLimitExpiration() time.Duration {
	return time.Duration(c.ipRateLimitExpirationInMs) * time.Millisecond
}

func (c *Config) GetIpRateLimitCleanupInterval() time.Duration {
	return time.Duration(c.ipRateLimitCleanupIntervalInMs) * time.Millisecond
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
