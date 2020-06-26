package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Yml *Yaml

type Yaml struct {
	SqlConf Mysql `yaml:"mysql"`
	ServerConf Server `yaml:"server"`
}

type Mysql struct{
	Username string`yaml:"username"`
	Password string`yaml:"password"`
	Addr string`yaml:"addr"`
	Name string`yaml:"name"`
}

type Server struct {
	Port int `yaml:"port"`
}

func init() {
	yamlFile,err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Errorf("读取配置文件失败，err: %v\n",err)
	}
	Yml = &Yaml{}
	err = yaml.Unmarshal(yamlFile,Yml)
	if err != nil {
		logrus.Errorf("解析配置文件失败，err: %v\n",err)
	}
}

func GetDataBaseConfig() *Mysql {
	return &Yml.SqlConf
}