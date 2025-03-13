package config

import (
	"github.com/caarlos0/env/v11"
)

// Config holds the application configuration
type Config struct {
	PGHost     string `env:"PG_HOST,required"`
	PGPort     string `env:"PG_PORT" envDefault:"5432"`
	PGUsername string `env:"PG_USERNAME,required"`
	PGPassword string `env:"PG_PASSWORD,required"`
	PGDBName   string `env:"PG_DBNAME,required"`

	Port  int  `env:"PORT" envDefault:"8080"`
	Debug bool `env:"DEBUG" envDefault:"false"`
}

// LoadConfig loads environment variables into the Config struct
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
