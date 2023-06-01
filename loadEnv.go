package main

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBHost         string `mapstructure:"POSTGRES_HOST"`
    DBUserName     string `mapstructure:"POSTGRES_USERNAME"`
    DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
    DBName         string `mapstructure:"POSTGRES_DBNAME"`
    DBPort         string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigType("env")
    viper.SetConfigName("database")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}
