package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Postgres    PostgresConfig
		FileStorage FileStorageConfig
		Auth        AuthConfig
		HTTP        HTTPConfig
		SMTP        SMTPConfig
		Email       EmailConfig
	}
	PostgresConfig struct {
		Username string `mapstructure:"username"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
		Password string
	}
	FileStorageConfig struct {
		Endpoint  string
		Bucket    string `mapstructure:"bucket"`
		AccessKey string
		SecretKey string
		Website   string
	}
	EmailConfig struct {
		Templates EmailTemplates
	}
	EmailTemplates struct {
		Verification string `mapstructure:"verification_email"`
	}
	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}
	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}
	HTTPConfig struct {
		Host               string
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
	SMTPConfig struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		From string `mapstructure:"from"`
		Pass string
	}
)

func Init(configsDir string) (*Config, error) {

	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}
func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("fileStorage", &cfg.FileStorage); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("smtp", &cfg.SMTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("email.templates", &cfg.Email.Templates); err != nil {
		return err
	}
	return nil
}
func setFromEnv(cfg *Config) {
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")

	cfg.HTTP.Host = os.Getenv("HTTP_HOST")

	cfg.FileStorage.Endpoint = os.Getenv("STORAGE_ENDPOINT")
	cfg.FileStorage.AccessKey = os.Getenv("STORAGE_ACCESS_KEY")
	cfg.FileStorage.SecretKey = os.Getenv("STORAGE_SECRET_KEY")
	cfg.FileStorage.Bucket = os.Getenv("STORAGE_BUCKET")
	cfg.FileStorage.Website = os.Getenv("BUCKET_WEBSITE_URL")
	cfg.SMTP.Pass = os.Getenv("SMTP_PASSWORD")
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
