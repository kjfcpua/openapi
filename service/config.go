package service

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Common struct {
		Version  string
		IsDebug  bool `yaml:"debug"`
		LogPath  string
		LogLevel string
	}

	Mysql struct {
		Addr     string
		Port     string
		Database string
		Acc      string
		Pw       string
	}
}

var Conf = &Config{}

func initConfig() {
	data, err := ioutil.ReadFile("openapi.yaml")
	if err != nil {
		log.Fatal("read config error :", err)
	}

	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal("yaml decode error :", err)
	}

	log.Println(Conf)
}