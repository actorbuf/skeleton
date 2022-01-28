package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var Cfg *ServerConf

type ServerConf struct {
	ConfigFile       string
	ServiceName      string `yaml:"service-name"`
	Environment      string `yaml:"environment"`
	HttpServerListen string `yaml:"http-server-listen"`
}

func NewDefaultConfig() *ServerConf {
	c := &ServerConf{}
	return c
}

func (c *ServerConf) SetPath(path string) *ServerConf {
	c.ConfigFile = path
	return c
}

func (c *ServerConf) LoadConfigFile() error {
	f, err := os.Open(c.ConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}
