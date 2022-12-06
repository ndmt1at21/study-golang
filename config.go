package main

import "github.com/spf13/viper"

type Config struct {
	MySQLHost     string `mapstructure:"MYSQL_HOST"`
	MySQLPort     string `mapstructure:"MYSQL_PORT"`
	MySQLUser     string `mapstructure:"MYSQL_USER"`
	MySQLPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySQLDatabase string `mapstructure:"MYSQL_DATABASE"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
