package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func Load(filePath string) (*AppConfig, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG_FILE")
	}

	zap.S().Infof("Load config file: %s", filePath)
	cfgBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		zap.S().Errorf("Read config file error: %v", err)
		return nil, err
	}

	cfgBytes = []byte(os.ExpandEnv(string(cfgBytes)))
	cfg := new(AppConfig)
	if err := yaml.Unmarshal(cfgBytes, cfg); err != nil {
		zap.S().Errorf("Unmarshal config file error: %v", err)
		return nil, err
	}

	zap.S().Infof("Load config file success")
	return cfg, nil
}

type AppConfig struct {
	Port        int    `yaml:"port"`
	Host        string `yaml:"host"`
	PostgresDSN string `yaml:"postgres_dsn"`
}
