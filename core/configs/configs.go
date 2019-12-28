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

type Config struct {
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassword   string
	DbName       string
	JwtSecretKey string
}

type OAuthConfig struct {
	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectUrl  string
	KakaoClientId      string
	KakaoClientSecret  string
	KakaoRedirectUrl   string
}

func GetConfig() Config {
	config := Config{}
	config.DbHost = getEnv("DB_HOST", "localhost")
	config.DbPort = getEnv("DB_PORT", "5432")
	config.DbUser = getEnv("DB_USER", "hfive")
	config.DbPassword = getEnv("DB_PASSWORD", "hfive")
	config.DbName = getEnv("DB_NAME", "hfive")
	config.JwtSecretKey = getEnv("JWT_SECRET_KEY", "hfive")
	return config
}

func GetOAuthConfig() OAuthConfig {
	config := OAuthConfig{}
	config.GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
	config.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	config.GoogleRedirectUrl = os.Getenv("GOOGLE_REDIRECT_URL")
	config.KakaoClientId = os.Getenv("KAKAO_CLIENT_ID")
	config.KakaoClientSecret = os.Getenv("KAKAO_CLIENT_SECRET")
	config.KakaoRedirectUrl = os.Getenv("KAKAO_REDIRECT_URL")
	return config
}
