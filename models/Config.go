package models

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		Driver       string `json:"driver"`
		Host         string `json:"host"`
		User         string `json:"user"`
		Password     string `json:"password"`
		DatabaseName string `json:"databaseName"`
	} `json:"database"`
	CPU struct {
		Maxprocs int `json:"maxprocs"`
	} `json:"cpu"`
	Port string `json:"port"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
