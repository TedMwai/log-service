package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

func GetEnv(key, fallback string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}

	return fallback
}

func MustGetEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}

	// Panic if not in debug mode
	// This allows us to run unit tests without the envs set
	if !AppIsDebug {
		log.Panic().Msgf("%s environment variable is required", key)
	}

	return ""
}

const AppEnvDebug = "DEBUG"
const AppEnvStaging = "STAGING"
const AppEnvProduction = "PRODUCTION"

var AppEnv = GetEnv("APP_ENVIRONMENT", AppEnvDebug)
var AppIsDebug = AppEnv == AppEnvDebug
var AppName = GetEnv("APP_NAME", "service.log-management")
var AppPort = GetEnv("APP_PORT", "8080")
var DB_URL = GetEnv("POSTGRESQL_URL", "postgres://postgres:password@localhost:5432/log?sslmode=disable")