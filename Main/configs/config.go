package config

import (
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

type TokenJWT struct {
	JWT string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     GetEnv("DB_HOST", "204.188.247.54"),
		Port:     GetEnv("DB_PORT", "5432"),
		User:     GetEnv("DB_USER", "postgres"),
		Password: GetEnv("DB_PASSWORD", "postgres"),
		DBName:   GetEnv("DB_NAME", "final_project"),
	}
}

func GetTokenJWTConfig() *TokenJWT {
	return &TokenJWT{
		JWT: GetEnv("JWT", "your-secret-key"),
	}
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
