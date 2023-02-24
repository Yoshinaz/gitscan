package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
	"path"
	"runtime"
)

type Config struct {
	DB  Database `json:"db"`
	App App      `json:"app"`
}

type Database struct {
	Host     string `env:"CONFIG__DB__HOST" required:"true" json:"host"`
	Port     int    `env:"CONFIG__DB__PORT" default:"3306" json:"port"`
	User     string `env:"CONFIG__DB__USER" required:"true" json:"user"`
	Password string `env:"CONFIG__DB__PASSWORD" required:"true" json:"password"`
	Name     string `env:"CONFIG__DB__NAME" required:"true" json:"name"`
}

type App struct {
	Name       string `env:"CONFIG__APP_NAME" default:"mst-payment-service" json:"name"`
	Version    string `env:"APP_VERSION" default:"local"`
	Port       int    `env:"CONFIG__APP_CONFIG_PORT" default:"8080" json:"port"`
	MaxProcess int    `env:"CONFIG__APP_MAX_PROCESS" default:"16" json:"max_process"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := configor.
		New(&configor.Config{AutoReload: false}).
		Load(&config, fmt.Sprintf("%s/config.%s.json", getConfigLocation(), getEnv()))

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigLocation() string {
	_, filename, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(filename), "../config")
}

func getEnv() string {
	val := os.Getenv("APP_ENV")
	// todo: check our stage names and align with them
	switch val {
	case "prod":
		return "prod"
	case "test":
		return "test"
	case "qa":
		return "qa"
	default:
		return "dev"
	}
}
