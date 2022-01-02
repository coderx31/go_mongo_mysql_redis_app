package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mitchellh/mapstructure"
)

type App struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type Mongo struct {
	URI        string `mapstructure:"uri"`
	DB         string `mapstructure:"db"`
	Collection string `mapstructure:"collection"`
}

type Config struct {
	App
	Mongo
}

// reading configs

func Configs() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
	}

	// decoding into map
	var mapConfig map[string]interface{}
	err = json.Unmarshal(data, &mapConfig)

	if err != nil {
		log.Fatal(err)
	}

	// decoding into structs

	var config Config
	mapstructure.Decode(mapConfig, &config)

	return &config, nil

}
