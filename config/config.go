package config

import (
	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	ChatGpt ChatGptConfig `json:"chatgpt"`
	Line    LineConfig    `json:"line"`
}

type ChatGptConfig struct {
	Keyword  string `json:"keyword,omitempty"`
	Token    string `json:"token,omitempty" json:"token,omitempty"`
	Telegram string `json:"telegram"`
}

type LineConfig struct {
	ChannelId     string `json:"channel_id"`
	ChannelSecret string `json:"channel_secret"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	return config
}
