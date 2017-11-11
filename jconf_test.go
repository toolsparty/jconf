package jconf

import (
	"testing"
)

var configFile = "./test/conf.json"
var errorFile = "./test/error.json"

func TestNewConfig(t *testing.T) {

	t.Run("No Errors", func(t *testing.T) {
		_, err := NewConfig(configFile)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Error file", func(t *testing.T) {
		_, err := NewConfig(errorFile)
		if err == nil {
			t.Error("no file errors")
		}
	})

	t.Run("Has Errors", func(t *testing.T) {
		_, err := NewConfig("./conf.json")
		if err == nil {
			t.Error("no errors")
		}
	})
}

func TestConfig_Load(t *testing.T) {
	conf := Config{}
	err := conf.Load()
	if err == nil {
		t.Error("no errors")
	}

	conf.FileName = "conf.json"
	err = conf.Load()
	if err == nil {
		t.Error("no errors")
	}

	conf.FileName = configFile
	err = conf.Load()
	if err != nil {
		t.Error(err)
	}
}

func TestConfig_Get(t *testing.T) {
	conf, _ := NewConfig(configFile)

	if res, err := conf.Get("name"); err != nil || res != "Name" {
		t.Error("Name Error", err)
	}

	if res, err := conf.Get("data"); err != nil || res != "test" {
		t.Error("Data Error", err)
	}

	if res, err := conf.Get("number"); err != nil || res != 10. {
		t.Error("number error", err)
	}

	if res, err := conf.Get("bool"); err != nil || res != false {
		t.Error("bool error", err)
	}
}