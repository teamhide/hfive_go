package configs

import (
	"os"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

type Configuration struct {
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassword   string
	DbName       string
	JwtSecretKey string
}

func GetConfiguration() Configuration {
	config := Configuration{}
	config.DbHost = getEnv("DB_HOST", "localhost")
	config.DbPort = getEnv("DB_PORT", "5432")
	config.DbUser = getEnv("DB_USER", "hfive")
	config.DbPassword = getEnv("DB_PASSWORD", "hfive")
	config.DbName = getEnv("DB_NAME", "hfive")
	config.JwtSecretKey = getEnv("JWT_SECRET_KEY", "hfive")
	return config
}
