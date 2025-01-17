package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"sushi-backend/constants"
	"sushi-backend/internal/logger"
	"sushi-backend/utils"
	"time"
)

type Config struct {
	logger                         logger.ILogger
	appEnv                         constants.AppEnv
	jwtSecretKey                   string
	jwtExpirationInMs              int
	sslCertPath                    string
	sslKeyPath                     string
	allowedOrigins                 []string
	allowedMethods                 []string
	allowedHeaders                 []string
	allowCredentials               bool
	requestTimeoutInS              int
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
	adminPassword                  string
}

func NewConfig(deps ConfigDependencies) *Config {
	_logger := deps.Logger

	godotenv.Load() // ignore error, because in deployment we pass all env variables via docker run command

	config := &Config{
		logger: _logger,
	}

	appEnv := config.getRequiredString("APP_ENV")

	switch appEnv {
	case string(constants.DevelopmentEnv):
		config.appEnv = constants.DevelopmentEnv
	case string(constants.ProductionEnv):
		config.appEnv = constants.ProductionEnv
	default:
		panic(fmt.Sprintf("Invalid APP_ENV value: %s. Supported values: %s, %s", appEnv, constants.DevelopmentEnv, constants.ProductionEnv))
	}

	config.postgresDSN = config.getRequiredString("POSTGRES_DSN")
	config.telegramBotToken = config.getRequiredString("TELEGRAM_BOT_TOKEN")
	config.telegramOrdersChatId = config.getRequiredString("TELEGRAM_ORDERS_CHAT_ID")
	config.telegramDeliveryChatId = config.getRequiredString("TELEGRAM_DELIVERY_CHAT_ID")
	config.cloudinaryUrl = config.getRequiredString("CLOUDINARY_URL")
	config.adminPassword = config.getRequiredString("ADMIN_PASSWORD")
	config.jwtSecretKey = config.getRequiredString("JWT_SECRET_KEY")
	config.sslCertPath = config.getOptionalString("SSL_CERT_PATH", "./certs/cert.pem")
	config.sslKeyPath = config.getOptionalString("SSL_KEY_PATH", "./certs/priv.pem")
	config.jwtExpirationInMs = config.getOptionalInt("JWT_EXPIRATION_IN_MS", 86_400_000)
	config.cloudinaryFolder = config.getOptionalString("CLOUDINARY_FOLDER", "sushi")
	config.runMigration = config.getOptionalBool("RUN_MIGRATION", true)
	config.httpPort = config.getOptionalString("HTTP_PORT", ":8080")
	config.requestTimeoutInS = config.getOptionalInt("REQUEST_TIMEOUT_IN_SECONDS", 10)
	config.ipRateLimitRate = config.getOptionalInt("IP_RATE_LIMIT_RATE", 60)
	config.ipRateLimitBurst = config.getOptionalInt("IP_RATE_LIMIT_BURST", 20)
	config.ipRateLimitExpirationInMs = config.getOptionalInt("IP_RATE_LIMIT_EXPIRATION_IN_MS", 360_000)
	config.ipRateLimitCleanupIntervalInMs = config.getOptionalInt("IP_RATE_LIMIT_CLEANUP_INTERVAL_IN_MS", 360_000)
	config.errorStackTraceSizeInKb = config.getOptionalInt("ERROR_STACK_TRACE_SIZE_IN_KB", 4)
	config.maxFileSizeInMb = config.getOptionalInt("MAX_FILE_SIZE_IN_MB", 200)

	for _, origin := range strings.Split(config.getOptionalString("ALLOWED_ORIGINS", "*"), ",") {
		config.allowedOrigins = append(config.allowedOrigins, origin)
	}

	for _, method := range strings.Split(config.getOptionalString("ALLOWED_METHODS", "GET,POST,PUT,DELETE"), ",") {
		config.allowedMethods = append(config.allowedMethods, method)
	}

	for _, header := range strings.Split(config.getOptionalString("ALLOWED_HEADERS", "Content-Type,Authorization"), ",") {
		config.allowedHeaders = append(config.allowedHeaders, header)
	}

	config.allowCredentials = config.getOptionalBool("ALLOW_CREDENTIALS", true)

	return config
}

func (c *Config) AppEnv() constants.AppEnv {
	return c.appEnv
}

func (c *Config) JWTSecretKey() []byte {
	return []byte(c.jwtSecretKey)
}

func (c *Config) JWTExpiration() time.Duration {
	return time.Duration(c.jwtExpirationInMs) * time.Millisecond
}

func (c *Config) SSLCertPath() string {
	return c.sslCertPath
}

func (c *Config) SSLKeyPath() string {
	return c.sslKeyPath
}

func (c *Config) AdminPassword() string {
	return c.adminPassword
}

func (c *Config) AllowedOrigins() []string {
	return c.allowedOrigins
}

func (c *Config) AllowedMethods() []string {
	return c.allowedMethods
}

func (c *Config) AllowedHeaders() []string {
	return c.allowedHeaders
}

func (c *Config) AllowCredentials() bool {
	return c.allowCredentials
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

func (c *Config) RequestTimeout() time.Duration {
	return time.Duration(c.requestTimeoutInS) * time.Second
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

	valueInt := utils.PanicIfErrorWithResultReturning(strconv.Atoi(value))

	return valueInt
}

func (c *Config) getOptionalBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)

	if value == "" {
		c.logger.Warn(fmt.Sprintf(`Environment variable "%s" not found, used default %t`, key, defaultValue))
		return defaultValue
	}

	valueBool := utils.PanicIfErrorWithResultReturning(strconv.ParseBool(value))

	return valueBool
}
