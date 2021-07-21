package config

import (
	"github.com/spf13/viper"
)

func Initialize(configFile string) error {
	// Only log the warning severity or above.
	viper.SetConfigName(configFile) // no need to include file extension
	viper.AddConfigPath(".\\etc")   // set the path of your config file
	err := viper.ReadInConfig()
	return err
}
