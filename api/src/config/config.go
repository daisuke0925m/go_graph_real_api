package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App   App   `toml:"app"`
	Db    Db    `toml:"db"`
	Redis Redis `toml:"redis"`
}

type App struct {
	Port string `toml:"port"`
}

type Db struct {
	User     string `toml:"user"`
	Name     string `toml:"name"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

type Redis struct {
	URL      string `toml:"url"`
	Password string `toml:"password"`
}

var Conf *Config

func Load() {
	Conf = new(Config)

	fmt.Println("exec init!")

	GO_ENV := os.Getenv("GO_ENV")
	if GO_ENV == "" {
		GO_ENV = "development"
	}

	viper.SetConfigName(GO_ENV)
	viper.AddConfigPath("./src/config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("failed to unmarshal err: %s \n", err))
	}
}
