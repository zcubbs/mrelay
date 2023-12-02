package config

type ServiceType string

const (
	SmtpService   ServiceType = "smtp"
	AwsSesService ServiceType = "aws_ses"
)

type Configuration struct {
	Accounts map[string]string `mapstructure:"accounts"`
	Database DatabaseConfig    `mapstructure:"database"`

	HttpServer HttpServerConfig `mapstructure:"server"`
	Logging    LoggingConfig    `mapstructure:"logging"`

	Smtp   SmtpConfig   `mapstructure:"smtp"`
	AwsSes AwsSesConfig `mapstructure:"aws_ses"`
}

type HttpServerConfig struct {
	Port         string   `mapstructure:"port"`
	EnableCORS   bool     `mapstructure:"enable_cors"`
	AllowOrigins []string `mapstructure:"allow_origins"`
	TlsEnabled   bool     `mapstructure:"tls_enabled"`
	TlsCertFile  string   `mapstructure:"tls_cert_file"`
	TlsKeyFile   string   `mapstructure:"tls_key_file"`
}

type LoggingConfig struct {
	Level      string `mapstructure:"level"`
	TextFormat bool   `mapstructure:"text_format"`
}

type AccountRule struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

type DatabaseConfig struct {
	Sqlite   SQLiteConfig   `mapstructure:"sqlite"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type SQLiteConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	Datasource  string `mapstructure:"datasource"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

type PostgresConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname"`
	SSLMode     bool   `mapstructure:"ssl_mode"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

type SmtpConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	FromName    string `mapstructure:"from_name"`
	FromAddress string `mapstructure:"from_address"`
}

type AwsSesConfig struct {
	Region      string `mapstructure:"region"`
	AccessKeyID string `mapstructure:"access_key_id"`
}
