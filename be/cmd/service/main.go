package main

import (
	"app/config"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "Path to the config file")
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			zap.S().Errorf("Recover when start project error: %v", err)
			os.Exit(0)
		}
	}()

	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("Load config error: %v", err)
		panic(err)
	}

	zap.S().Infof("Start project with config: %v", cfg)
	fmt.Println("Start project with config:", cfg)
}
