package viper

import (
	"errors"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigFile("C:\\Users\\Ricardo\\source\\repos\\Flo-AI_Config\\viperconfig.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return errors.New("VIPER - Error reading configuration file")
	}

	return nil
}

func GetString(key string) (string, error) {
	if viper.IsSet(key) {
		return viper.GetString(key), nil
	}

	return "", errors.New("VIPER - Cannot find key")
}
