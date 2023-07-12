package config

import (
	"github.com/spf13/viper"
)

var config *MaskingConfig

type MaskingConfig struct {
	DeniedKeyList []string `mapstructure:"denied_key_list"`
	UseRegex      bool     `mapstructure:"use_regex"`
	Format        bool     `mapstructure:"format"`
}

func Initialize(fileName string) error {
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	c := &MaskingConfig{}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	config = c
	return nil
}

func GetConfig() *MaskingConfig {
	if config == nil {
		panic("config is not initialized.")
	}
	return config
}
