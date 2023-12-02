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

	v.SetDefault("database.sqlite.datasource", "mq-watch.db")
	v.SetDefault("server.port", "8000")

	for k, val := range envKeys {
		err := v.BindEnv(k, strings.ToUpper(val))
		if err != nil {
			fmt.Println("Error binding env var", val, err)
		}
	}

	if err := v.ReadInConfig(); err != nil {
		return Configuration{}, err
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
	"mqtt.broker":                    "MQTT_BROKER",
	"mqtt.client_id":                 "MQTT_CLIENT_ID",
	"mqtt.username":                  "MQTT_USERNAME",
	"mqtt.password":                  "MQTT_PASSWORD",
	"mqtt.tls_enabled":               "MQTT_TLS_ENABLED",
	"mqtt.tls_cert_file":             "MQTT_TLS_CERT_FILE",
}
