package jconf

import (
	"os"
	"encoding/json"
	"fmt"
)

type Interface interface {
	Load() error
	Get(key string) (interface{}, error)
}

type Config struct {
	// storage
	data map[string]interface{}
	FileName string
}

// load json file
func (c *Config) Load() error {
	file, err := os.Open(c.FileName)
	defer file.Close()

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	return decoder.Decode(&c.data)
}

// getting value by key
func (c Config) Get(key string) (interface{}, error) {
	if value, success := c.data[key]; success {
		return value, nil
	}

	return nil, fmt.Errorf("the '%s' is not found in config", key)
}

// constructor
func NewConfig(fileName string) (*Config, error) {
	var conf Config = Config{FileName: fileName}
	err := conf.Load()
	if err != nil {
		return nil, err
	}

	return &conf, nil
}