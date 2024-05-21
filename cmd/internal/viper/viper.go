package viper

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigFile("C:\\Users\\Ricardo\\source\\repos\\Flo-AI_Config\\viperconfig.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading configuration file: ", err)
		return fmt.Errorf("error reading configuration file: %w", err)
	}

	return nil
}

func GetString(key string) (string, error) {
	if viper.IsSet(key) {
		return viper.GetString(key), nil
	}

	return "", errors.New("VIPER - Cannot find key")
}
