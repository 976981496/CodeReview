package common

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)
type Log  struct {
	FileName string `yaml:"FileName"`
	LogLevel string `yaml:"LogLevel"`
}
var GlobalConfig Config

type Config struct {
	Log   Log    `yaml:"Log"`
}


func Init() error {
	f, openErr := os.Open("./conf/log.yaml")
	if openErr != nil {
		return openErr
	}
	data, readErr := ioutil.ReadAll(f)
	if readErr != nil {
		return readErr
	}
	umErr := yaml.Unmarshal(data, &GlobalConfig)
	if umErr != nil {
		return umErr
	}
	return nil
}
