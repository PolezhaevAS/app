package service

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

func New(cfg *Config) (s *Service, err error) {
	s = new(Service)
	return
}
