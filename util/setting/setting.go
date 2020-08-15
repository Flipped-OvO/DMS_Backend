package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type configModel struct {
	Runtime struct {
		Mode    string `yaml:"mode"`
		Port    int    `yaml:"port"`
		LogPath string `yaml:"logPath"`
	}
}

var Config configModel

func init() {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Println(err)
	}
}
