package config

import (
	"bytes"
	_ "embed"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all settings
//
//go:embed config.yaml
var defaultConfig []byte

type Config struct {
	Server Server            `json:"server" yaml:"server" mapstructure:"server"`
	Hooks  []map[string]Hook `json:"hooks" yaml:"hooks" mapstructure:"hooks"`
}

type Server struct {
	SericeName string     `json:"service_name" yaml:"service_name" mapstructure:"service_name"`
	HTTP       HTTPConfig `json:"http" yaml:"http" mapstructure:"http"`
}

type HTTPConfig struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port int32  `json:"port" yaml:"port" mapstructure:"port"`
}

type Hook struct {
	Backend       string         `json:"backend" yaml:"backend" mapstructure:"backend"`
	Notifications []Notification `json:"notifications" yaml:"notifications" mapstructure:"notifications"`
}

type Notification struct {
	Backend  string `json:"backend" yaml:"backend" mapstructure:"backend"`
	Endpoint string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatal("Failed to read viper config", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config", err)
	}

	log.Println("Config loaded")
	return cfg
}
