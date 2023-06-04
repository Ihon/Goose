package system

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBHost         string `mapstructure:"MARIADB_HOST"`
    DBUserName     string `mapstructure:"MARIADB_USERNAME"`
    DBUserPassword string `mapstructure:"MARIADB_PASSWORD"`
    DBName         string `mapstructure:"MARIADB_DBNAME"`
    DBPort         string `mapstructure:"MARIADB_PORT"`
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
