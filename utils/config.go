package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	JwtRefreshExpiredIn      time.Duration `mapstructure:"JWT_REFRESH_TOKEN_EXPIRED_IN"`
	JwtRefreshMaxAge         int           `mapstructure:"JWT_REFRESH_TOKEN_MAXAGE"`
	JwtAccessMaxAge          int           `mapstructure:"JWT_ACCESS_TOKEN_MAXAGE"`
	JwtAccessExpiredIn       time.Duration `mapstructure:"JWT_ACCESS_TOKEN_EXPIRED_IN"`
	JwtAccessTokenPublicKey  string        `mapstructure:"JWT_ACCESS_TOKEN_PRIVATE_KEY"`
	JwtAccessTokenPrivateKey string        `mapstructure:"JWT_ACCESS_TOKEN_PUBLIC_KEY"`
	Environment              string        `mapstructure:"NODE_ENV"`
	Port                     string        `mapstructure:"PORT"`
	HttpServerAddress        string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	PostgresDriver           string        `mapstructure:"POSTGRES_DRIVER"`
	MigrationURL             string        `mapstructure:"MIGRATION_URL"`
	DbSSLMode                string        `mapstructure:"SSL_MODE"`
	DbRootCert               string        `mapstructure:"DB_ROOT_CERT"`
	DbName                   string
	DbUser                   string
	DbPassword               string
	DbHost                   string
	DbPort                   string
	PostgresSource           string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	switch config.Environment {
	case "dev":
		config.DbName = viper.GetString("DEV_DB_NAME")
		config.DbUser = viper.GetString("DEV_DB_USER")
		config.DbPassword = viper.GetString("DEV_DB_PASSWORD")
		config.DbHost = viper.GetString("DEV_DB_HOST")
		config.DbPort = viper.GetString("DEV_DB_PORT")
	case "prod":
		config.DbName = viper.GetString("PROD_DB_NAME")
		config.DbUser = viper.GetString("PROD_DB_USER")
		config.DbPassword = viper.GetString("PROD_DB_PASSWORD")
		config.DbHost = viper.GetString("PROD_DB_HOST")
		config.DbPort = viper.GetString("PROD_DB_PORT")
	}
	//config.PostgresSource = fmt.Sprintf(
	//	"%s://%s:%s@%s:%s/%s?sslmode=%s&sslrootcert=%s", config.PostgresDriver, config.DbUser, config.DbPassword,
	//	config.DbHost, config.DbPort, config.DbName, config.DbSSLMode, config.DbRootCert)
	config.PostgresSource = fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s", config.PostgresDriver, config.DbUser, config.DbPassword,
		config.DbHost, config.DbPort, config.DbName, config.DbSSLMode)

	if err != nil {
		return config, err
	}
	return config, nil
}
