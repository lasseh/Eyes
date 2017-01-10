package main

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

var Config Conf

func init() {
	conf, err := getConfig("auth.json")
	if err != nil {
		log.Panicln(err)
	}
	Config = conf
}

func getConfig(filename string) (Conf, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Conf{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Conf{}

	err = decoder.Decode(&config)
	if err != nil {
		return Conf{}, err
	}
	return config, nil
}
