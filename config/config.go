package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type BrokerConfig struct {
	URL          string `yaml:"url"`
	QoS          byte   `yaml:"qos"`
	EphemeralTTL int    `yaml:"ephemeral_ttl"`
}

type AppConfigStruct struct {
	Broker        BrokerConfig `yaml:"broker"`
	DefaultTopics []string     `yaml:"default_topics"`
	BannedWords   []string     `yaml:"banned_words"`
}

var AppConfig AppConfigStruct

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
