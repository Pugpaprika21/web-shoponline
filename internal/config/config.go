package config

import (
	"fmt"
	"os"
)

// Config holds all application configuration
type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Port string
	Env  string
}

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret     string
	ExpireHour int
}

// DSN returns the PostgreSQL connection string
func (d DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode,
	)
}

// Load reads configuration from environment variables
func Load() *Config {
	return &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "3000"),
			Env:  getEnv("APP_ENV", "development"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "shoponline"),
			Password: getEnv("DB_PASSWORD", "shoponline_secret"),
			Name:     getEnv("DB_NAME", "shoponline_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "shoponline-jwt-secret-key-change-in-production"),
			ExpireHour: 24,
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
