package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Token    TokenConfig
	SMTP     SMTPConfig
	Log      LogConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type TokenConfig struct {
	Secret       string
	ExpiryHours  int
	ExpiryTime   time.Duration
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type LogConfig struct {
	Level string
	File  string
}

// LoadConfig membaca konfigurasi dari file .env
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	expiryHours := viper.GetInt("TOKEN_EXPIRY_HOURS")
	if expiryHours == 0 {
		expiryHours = 24 // default 24 jam
	}

	config := &Config{
		App: AppConfig{
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetString("APP_PORT"),
			Env:  viper.GetString("APP_ENV"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
		},
		Token: TokenConfig{
			Secret:      viper.GetString("TOKEN_SECRET"),
			ExpiryHours: expiryHours,
			ExpiryTime:  time.Duration(expiryHours) * time.Hour,
		},
		SMTP: SMTPConfig{
			Host:     viper.GetString("SMTP_HOST"),
			Port:     viper.GetInt("SMTP_PORT"),
			Username: viper.GetString("SMTP_USERNAME"),
			Password: viper.GetString("SMTP_PASSWORD"),
			From:     viper.GetString("SMTP_FROM"),
		},
		Log: LogConfig{
			Level: viper.GetString("LOG_LEVEL"),
			File:  viper.GetString("LOG_FILE"),
		},
	}

	return config, nil
}

// GetDatabaseDSN mengembalikan connection string untuk PostgreSQL
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}
