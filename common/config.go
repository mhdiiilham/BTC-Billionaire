package common

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port       int    `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func ReadConfig() *Configuration {
	var config Configuration

	log.Println("reading config from env...")
	viper.AddConfigPath("./common")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed viper.ReadInConfig: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("failed unmarshal config: %v", err)
	}

	return &config
}
