package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env       string `env:"APP_ENV" envDefault:"dev"`
	Url       string `env:"APP_URL" envDefault:"http://localhost"`
	Port      string `env:"PORT" envDefault:":8080"`
	TimeZone  string `env:"TIME_ZONE" envDefault:"Asia/Tokyo"`
	AccessLog string `env:"ACCESS_LOG" envDefault:"./logs/access/"`
	Log       string `env:"LOG" envDefault:"./logs/log/log.log"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"DEBUG"`
}

var singleton = &Config{}

// singleton
func GetConfig() (*Config, error) {
	cnf := singleton
	if err := env.Parse(cnf); err != nil {
		return nil, err
	}

	return cnf, nil
}
