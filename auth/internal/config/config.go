package config

import (
	"strings"

	"app/internal/broker"
	"app/internal/server"
	db "app/internal/sql"
	"app/internal/token"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	TTLToken      int64  `mapstructure:"ttl"`
	Salt          string `mapstructure:"salt"`
	AdminName     string `mapstructure:"admin_name"`
	AdminPassword string `mapstructure:"admin_password"`
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{}
}

type Config struct {
	GRPC       *server.Config `yaml:"grpc" mapstructure:"grpc"`
	Token      *token.Config  `yaml:"token" mapstructure:"token"`
	DB         *db.Config     `yaml:"db" mapstructure:"db"`
	Broker     *broker.Config `yaml:"broker" mapstructure:"broker"`
	AuthConfig *AuthConfig    `yaml:"auth" mapstructure:"auth"`
}

// New Default configuration.
func New() *Config {
	return &Config{
		GRPC:       server.NewConfig(),
		Token:      token.NewConfig(),
		DB:         db.NewConfig(),
		Broker:     broker.NewConfig(),
		AuthConfig: NewAuthConfig(),
	}
}

func (c *Config) Load() *Config {
	var v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("/app/config")
	v.AddConfigPath("../config")
	v.AddConfigPath("/etc/auth")

	v.SetEnvPrefix("Auth")                             //
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
