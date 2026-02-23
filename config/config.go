package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type BrokerConfig struct {
	URL          string `yaml:"url"`
	DefaultRoom  string `yaml:"default_room"`
	QoS          byte   `yaml:"qos"`
	EphemeralTTL int    `yaml:"ephemeral_ttl"`
}

type Config struct {
	Broker BrokerConfig `yaml:"broker"`
	Rooms  []string     `yaml:"rooms"`
}

var AppConfig Config

func LoadConfig(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %v", err))
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse YAML: %v", err))
	}
}
