package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

// TODO: вынести в конфиг
const configPath = "config/local.yaml"

type Config struct {
	Env                string             `yaml:"env" env-default:"local"`
	HTTPServer         HTTPServer         `yaml:"httpcontroller-server"`
	DataBase           DataBase           `yaml:"data-base"`
	ClientInfoEnricher ClientInfoEnricher `yaml:"client-info-enricher"`
	Kafka              Kafka              `yaml:"kafka"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

type ClientInfoEnricher struct {
	GetAgeURL     string `yaml:"get-age-url"`
	GetGenderURL  string `yaml:"get-gender-url"`
	GetCountryURL string `yaml:"get-country-url"`
}

type Kafka struct {
	Address         string `yaml:"address"`
	AddClientTopic  string `yaml:"add-client-topic"`
	DeadClientTopic string `yaml:"dead-client-topic"`
	GroupId         string `yaml:"group-id"`
	Partition       int    `yaml:"partition"`
	MaxBytes        int    `yaml:"max-bytes"`
	DeadlineSec     int    `yaml:"deadline-sec"`
}

type DataBase struct {
	URL string `yaml:"url" env:"DATABASE_URL"`
}

func MustLoad() Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println()
		log.Fatalf("Config file does not exists: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Unable to read config: %s", err)
	}
	return cfg
}
