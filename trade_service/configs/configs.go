package configs

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func setUpConfig() {
	configName := flag.String("configname", "debug", "config file name")
	configPath := flag.String("configpath", "./configs/yaml", "config path")

	viper.AddConfigPath(*configPath)
	viper.SetConfigName(*configName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
