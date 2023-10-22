package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string `mapstructure:"env" env-default:"development"`
	StoragePath string `mapstructure:"storage_path"`
	TelegramBot `mapstructure:"telegram_bot" env-required:"true"`
	HTTPServer  `mapstructure:"http_server"`
}

type TelegramBot struct {
	Token string `mapstructure:"token" env-required:"true"`
}

type HTTPServer struct {
	Address string        `mapstructure:"adress"`
	Timeout time.Duration `mapstructure:"timeout" env-default:"5s"`
}

func MustLoad() *Config {
	var cfg Config

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found")
		} else {
			log.Fatal("another error read config")
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
