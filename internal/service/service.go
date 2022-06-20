package service

import (
	"fmt"
)

type Config struct {
	Name string `json:"name" yaml:"name" toml:"name" mapstructure:"name"` //nolint
}

// No default config
func NewConfig() *Config {
	return &Config{}
}

// Service is a description of service
type Service struct {
	Name    string
	Methods map[string]bool
}

func New(cfg *Config) (*Service, error) {
	switch cfg.Name {
	case "access":
		return NewAccessService(), nil
	default:
		return nil, fmt.Errorf("no service with name %s", cfg.Name)
	}
}

type AllServices struct {
	Services []*Service
}

func All() *AllServices {

	var services []*Service
	services = append(services, NewAccessService())

	return &AllServices{
		Services: services,
	}
}
