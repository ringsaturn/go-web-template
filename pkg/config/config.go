package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HTTPServer struct {
		Addr string `yaml:"Addr"`
	} `yaml:"HTTPServer"`
	PProfServer struct {
		Addr string `yaml:"Addr"`
	} `yaml:"PProfServer"`
}

// NewConfigFromYAMLPath will init a config class instance from YAML file Path
func NewConfigFromYAMLPath(yamlPath string) (*Config, error) {
	log.Println("loading config from", yamlPath)
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}
	log.Println("config file content:\n", string(yamlFile))

	config := &Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	log.Printf("config Unmarshal result %+v\n", config)
	return config, nil
}
