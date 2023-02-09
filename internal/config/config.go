package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen   Listen       `json:"listen" yaml:"listen"`
	Dingtalk Dingtalk     `json:"dingtalk" yaml:"dingtalk"`
	Feishu   FeishuConfig `json:"feishu" yaml:"feishu"`
}

type Listen struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

type Dingtalk struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Token   string `json:"token" yaml:"token"`
	Secret  string `json:"secret" yaml:"secret"`
}

type FeishuConfig struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Secret  string `json:"secret" yaml:"secret"`
}

func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Listen.Host, c.Listen.Port)
}

var (
	Conf Config
)

func LoadConfig(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(data, &Conf); err != nil {
		log.Fatalln(err)
	}
}
