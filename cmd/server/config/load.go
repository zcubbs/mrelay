package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/zcubbs/x/pretty"
	"strings"
	"sync"
)

func LoadConfiguration(configFile string) (Configuration, error) {
	var configuration Configuration

	var onceEnv sync.Once
	onceEnv.Do(loadEnv)

	v := viper.New()
	v.SetConfigFile(configFile)

	for k, val := range defaults {
		v.SetDefault(k, val)
	}

	for k, val := range envKeys {
		err := v.BindEnv(k, strings.ToUpper(val))
		if err != nil {
			fmt.Println("Error binding env var", val, err)
		}
	}

	if configFile != "" {
		if err := v.ReadInConfig(); err != nil {
			return Configuration{}, err
		}
	}

	if err := v.Unmarshal(&configuration); err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("no .env file found")
	}
}

func PrintConfiguration(config Configuration) {
	// Print out the configuration
	pretty.PrintJson(config)
}

var envKeys = map[string]string{
	"database.sqlite.enabled":        "DATABASE_SQLITE_ENABLED",
	"database.sqlite.datasource":     "DATABASE_SQLITE_DATASOURCE",
	"database.postgres.enabled":      "DATABASE_POSTGRES_ENABLED",
	"database.postgres.host":         "DATABASE_POSTGRES_HOST",
	"database.postgres.port":         "DATABASE_POSTGRES_PORT",
	"database.postgres.user":         "DATABASE_POSTGRES_USER",
	"database.postgres.password":     "DATABASE_POSTGRES_PASSWORD",
	"database.postgres.dbname":       "DATABASE_POSTGRES_DBNAME",
	"database.postgres.sslmode":      "DATABASE_POSTGRES_SSLMODE",
	"database.postgres.auto_migrate": "DATABASE_POSTGRES_AUTOMIGRATE",
	"server.port":                    "SERVER_PORT",
	"server.tls_enabled":             "SERVER_TLS_ENABLED",
	"server.tls_cert_file":           "SERVER_TLS_CERT_FILE",
	"server.tls_key_file":            "SERVER_TLS_KEY_FILE",
	"smtp.host":                      "SMTP_HOST",
	"smtp.port":                      "SMTP_PORT",
	"smtp.username":                  "SMTP_USERNAME",
	"smtp.password":                  "SMTP_PASSWORD",
	"smtp.from_name":                 "SMTP_FROM_NAME",
	"smtp.from_address":              "SMTP_FROM_ADDRESS",
}

var defaults = map[string]string{
	"server.port":                    "8000",
	"server.enable_cors":             "true",
	"server.tls_enabled":             "false",
	"server.tls_cert_file":           "",
	"server.tls_key_file":            "",
	"database.postgres.sslmode":      "disable",
	"database.postgres.host":         "localhost",
	"database.postgres.user":         "postgres",
	"database.postgres.password":     "postgres",
	"database.postgres.dbname":       "postgres",
	"database.postgres.port":         "5432",
	"database.postgres.auto_migrate": "true",
	"database.postgres.enabled":      "false",
	"database.sqlite.enabled":        "false",
	"database.sqlite.datasource":     "mrelay.db?_loc=auto",
	"smtp.host":                      "localhost",
	"smtp.port":                      "25",
	"smtp.username":                  "",
	"smtp.password":                  "",
	"smtp.from_name":                 "mrelay",
	"smtp.from_address":              "mrelay@localhost",
}
