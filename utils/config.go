package utils

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Ssl      string `yaml:"ssl"`
	} `yaml:"database"`
	Jwt struct {
		Secret string `yaml:"secret"`
	}
}

func readConf(filename string) (*Configuration, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Configuration{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}

func Config() *Configuration {
	c, err := readConf("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	return c
}
