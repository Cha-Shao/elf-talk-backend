package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port     int
	Secret   string
	Database struct {
		Type     string
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `toml:"dbname"`
	}
	Server struct {
		Name     string
		Password string
		Max      int
	}
}

var Conf Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
