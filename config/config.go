package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TokenID          string `mapstructure:"TOKEN_ID"`
	ClientID         string `mapstructure:"CLIENT_ID"`
	ClientSecret     string `mapstructure:"CLIENT_SECRET"`
	RedirectURL      string `mapstructure:"REDIRECT_URL"`
	State            string `mapstructure:"STATE"`
	OwnerPhoneNumber string `mapstructure:"OWNER_PHONE_NUMBER"`
}

var Cfg Config

func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	viper.AutomaticEnv()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&Cfg)
	return
}
