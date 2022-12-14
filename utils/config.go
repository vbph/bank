package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver                string        `mapstructure:"DB_DRIVER"`
	DbSource                string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddr          string        `mapstructure:"HTTP_SERVER_ADDR"`
	GrpcServerAddr          string        `mapstructure:"GRPC_SERVER_ADDR"`
	ServerType              string        `mapstructure:"SERVER_TYPE"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenExpiredTime  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_TIME"`
	RefreshTokenExpiredTime time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_TIME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
