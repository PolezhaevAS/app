package config

import (
	"app/internal/broker"
	"app/internal/server"
	db "app/internal/sql"
	"app/internal/token"
	"strings"

	"github.com/spf13/viper"
)

type AccessConfig struct{}

func NewAccessConfig() *AccessConfig {
	return &AccessConfig{}
}

type Config struct {
	GRPC         *server.Config `yaml:"grpc" mapstructure:"grpc"`
	Token        *token.Config  `yaml:"token" mapstructure:"token"`
	DB           *db.Config     `yaml:"db" mapstructure:"db"`
	Broker       *broker.Config `yaml:"broker" mapstructure:"broker"`
	AccessConfig *AccessConfig  `yaml:"access" mapstructure:"access"`
}

// New Default configuration.
func New() *Config {
	return &Config{
		GRPC:         server.NewConfig(),
		Token:        token.NewConfig(),
		DB:           db.NewConfig(),
		Broker:       broker.NewConfig(),
		AccessConfig: NewAccessConfig(),
	}
}

func (c *Config) Load() *Config {
	var v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("/app/config")
	v.AddConfigPath("../config")
	v.AddConfigPath("/etc/access")

	v.SetEnvPrefix("Access")                           //
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) //

	var err error
	if err = v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.AutomaticEnv() // from environments

	if err = v.Unmarshal(c); err != nil {
		panic(err)
	}

	return c
}
