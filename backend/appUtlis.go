package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigUtils struct {
	MongoConfig string `yaml:"MongoConfig"`
}

func get_config() (*ConfigUtils, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error occurred: %v", err)
	}
	defer file.Close()

	var config ConfigUtils
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error occurred during decoding: %v", err)
	}

	fmt.Printf("Decoded Config: %+v\n", config) // Debugging line
	return &config, nil
}
