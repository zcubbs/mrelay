package config

type ServiceType string

const (
	SmtpService   ServiceType = "smtp"
	AwsSesService ServiceType = "aws_ses"
)

type Configuration struct {
	Accounts []Account      `mapstructure:"accounts"`
	Database DatabaseConfig `mapstructure:"database"`

	HttpServer HttpServerConfig `mapstructure:"server"`

	Smtp   SmtpConfig   `mapstructure:"smtp"`
	AwsSes AwsSesConfig `mapstructure:"aws_ses"`
}

type HttpServerConfig struct {
	Port        int    `mapstructure:"port"`
	TlsEnabled  bool   `mapstructure:"tls_enabled"`
	TlsCertFile string `mapstructure:"tls_cert_file"`
	TlsKeyFile  string `mapstructure:"tls_key_file"`
}

type Account struct {
	Name           string        `mapstructure:"name"`
	Password       string        `mapstructure:"password"`
	Enabled        bool          `mapstructure:"enabled"`
	Services       []ServiceType `mapstructure:"services"`
	AllowedDomains []string      `mapstructure:"allowed_domains"`
	AllowedIPs     []string      `mapstructure:"allowed_ips"`
	Rules          []AccountRule `mapstructure:"rules"`
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
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname"`
	SSLMode     bool   `mapstructure:"ssl_mode"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

type SmtpConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type AwsSesConfig struct {
	Region      string `mapstructure:"region"`
	AccessKeyID string `mapstructure:"access_key_id"`
}
