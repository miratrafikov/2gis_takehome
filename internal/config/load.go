package config

import "github.com/spf13/viper"

func Load() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
