package common

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Version    string `mapstructure:"VERSION"`
	Port       int    `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSL_MODE"`
}

func (c Configuration) GetServerPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

func ReadConfig(version string) *Configuration {
	var config Configuration

	log.Printf("reading config from app.env")
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed viper.ReadInConfig: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("failed unmarshal config: %v", err)
	}

	config.Version = version
	log.Println("success reading config from app.env")
	return &config
}
