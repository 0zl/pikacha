package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiKey          []string `json:"api_key"`
	Port            int      `json:"port"`
	MaximumFileSize int64    `json:"maximum_file_size"`
}

func defaultConfig() Config {
	return Config{
		ApiKey:          []string{},
		Port:            49100,
		MaximumFileSize: 25 * 1024 * 1024,
	}
}

func NewConfig() (Config, error) {
	defaultCfg := defaultConfig()
	cfgFile := "config.yml"

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		yamlCfg, err := yaml.Marshal(defaultCfg)
		if err != nil {
			return Config{}, err
		}
		err = os.WriteFile(cfgFile, yamlCfg, 0644)
		if err != nil {
			return Config{}, err
		}
	}

	yamlCfg, err := os.ReadFile(cfgFile)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(yamlCfg, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) GetAdress() string {
	return "0.0.0.0:" + fmt.Sprint(cfg.Port)
}
