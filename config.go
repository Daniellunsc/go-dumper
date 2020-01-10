package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"

	"gopkg.in/yaml.v2"
)

// Config represents a default config structure
type Config struct {
	Database struct {
		User     string `yaml:"user" envconfig:"user"`
		Hostname string `yaml:"hostname" envconfig:"hostname"`
		Password string `yaml:"password" envconfig:"password"`
		Db       string `yaml:"db" envconfig:"db"`
		Port     string `yaml:"port" envconfig:"port"`
	} `yaml:"database"`

	Telegram struct {
		Botkey string `yaml:"botkey" envconfig:"botkey"`
		ChatID int64  `yaml:"chatid" envconfig:"chatid"`
	} `yaml:"telegram"`
}

func readConfigFile(config *Config) {
	reader, _ := os.Open("config.yml")
	yamlFile, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}

func readConfigEnv(config *Config) {
	err := envconfig.Process("", config)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%+v\n", config)
}
