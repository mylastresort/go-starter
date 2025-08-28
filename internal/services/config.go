package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Conf AppConfig

type AppConfig struct {
	DB struct {
		NAME string `mapstructure:"NAME"`
		USER string `mapstructure:"USER"`
		PASS string `mapstructure:"PASS"`
		HOST string `mapstructure:"HOST"`
		PORT int    `mapstructure:"PORT"`
	} `mapstructure:"DB"`
	HTTP struct {
		PORT int
	} `mapstructure:"HTTP"`
}

func LoadConfig() {
	log.Default().Print("LOADING Config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Conf.DB.HOST)
}
