package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(lg *logrus.Logger) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName("config.yml")
	config.AddConfigPath(".")
	config.SetConfigType("yml")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		lg.WithFields(logrus.Fields{"class": "main", "event": "config.file.read"}).Warn(err)
		return nil, err
	}
	return config, nil
}
