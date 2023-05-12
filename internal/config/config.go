package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string
		Port string
		Host string
	}
	Web struct {
		StaticPath string
	}
	Connection struct {
		Postgresql struct {
			DbType   string
			Db       string
			Host     string
			User     string
			Password string
			MaxOpen  int
			MaxIdle  int
		}
	}
}

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cfg := &Config{}
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %s", err)
	}
	return cfg, nil
}
