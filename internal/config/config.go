package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Listen struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

type Dingtalk struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Token   string `json:"token" yaml:"token"`
	Secret  string `json:"secret" yaml:"secret"`
}

type Config struct {
	Listen   Listen   `json:"listen" yaml:"listen"`
	Dingtalk Dingtalk `json:"dingtalk" yaml:"dingtalk"`
}

func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Listen.Host, c.Listen.Port)
}

func LoadConfig(configFile string) (*Config, error) {
	conf := Config{}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
