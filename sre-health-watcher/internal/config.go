package internal

import (
	"os"
	"gopkg.in/yaml.v2"
)

type Endpoint struct {
	Name              string `yaml:"name"`
	URL               string `yaml:"url"`
	LatencyThresholdMS int    `yaml:"latency_threshold_ms"`
	AlertOn           []int  `yaml:"alert_on"`
}

type EmailConfig struct {
	SMTPServer string `yaml:"smtp_server"`
	Port       int    `yaml:"port"`
	From       string `yaml:"from"`
	To         string `yaml:"to"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type AlertConfig struct {
	SlackWebhook string      `yaml:"slack_webhook"`
	Email        EmailConfig `yaml:"email"`
}

type Config struct {
	IntervalSeconds int         `yaml:"interval_seconds"`
	Endpoints       []Endpoint  `yaml:"endpoints"`
	Alert           AlertConfig `yaml:"alert"`
}

func LoadConfig() Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
