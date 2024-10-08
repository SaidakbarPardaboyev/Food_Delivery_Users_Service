package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	UserServiceGrpcHost string
	UserServiceGrpcPort string

	OrdersServiceGrpcHost string
	OrdersServiceGrpcPort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresName     string
	PostgresPassword string

	ServiceName string
	LoggerLevel string
	LogPath     string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".env file not found: %s", err)
	}
	config := Config{}

	config.UserServiceGrpcHost = cast.ToString(coalesce("USER_SERVICE_GRPC_HOST", "localhost"))
	config.UserServiceGrpcPort = cast.ToString(coalesce("USER_SERVICE_GRPC_PORT", ":1111"))

	config.OrdersServiceGrpcHost = cast.ToString(coalesce("ORDERS_SERVICE_GRPC_HOST", "localhost"))
	config.OrdersServiceGrpcPort = cast.ToString(coalesce("ORDERS_SERVICE_GRPC_PORT", ":5555"))

	config.PostgresHost = cast.ToString(coalesce("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToString(coalesce("POSTGRES_PORT", "5432"))
	config.PostgresUser = cast.ToString(coalesce("POSTGRES_USER", "postgres"))
	config.PostgresName = cast.ToString(coalesce("POSTGRES_DBNAME", "food_delivery_user_service"))
	config.PostgresPassword = cast.ToString(coalesce("POSTGRES_PASSWORD", "root"))

	config.ServiceName = cast.ToString(coalesce("SERVICE_NAME", "auth_service"))
	config.LoggerLevel = cast.ToString(coalesce("LOGGER_LEVEL", "debug"))
	config.LogPath = cast.ToString(coalesce("LOG_PATH", "app.log"))
	fmt.Println(config.PostgresName)

	return &config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	if res, exists := os.LookupEnv(key); exists {
		return res
	}
	return defaultValue
}
