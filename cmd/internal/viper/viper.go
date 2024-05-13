package viper

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig(configFilePath string) error {
	viper.SetConfigFile("C:\\Users\\Ricardo\\source\\repos\\Flo-AI_Config\\viperconfig.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading configuration file:", err)
		return fmt.Errorf("error reading configuration file: %w", err)
	}

	return nil
}
