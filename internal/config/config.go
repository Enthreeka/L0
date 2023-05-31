package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		TypeServer string `json:"typeserver"`
		Port       string `json:"port"`
	} `json:"server"`

	Db struct {
		URL string `json:"url"`
	} `json:"db"`

	Nats struct {
	} `json:"nats"`
}

func New(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
