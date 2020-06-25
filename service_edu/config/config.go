package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Yml *Yaml

type Yaml struct {
	SqlConf Mysql `yaml:"mysql"`
}

type Mysql struct{
	Username string`yaml:"username"`
	Password string`yaml:"password"`
	Addr string`yaml:"addr"`
	Name string`yaml:"name"`
}

func GetDataBaseConfig() *Mysql {
	if Yml != nil {
		return &Yml.SqlConf
	}
	yamlFile,err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.Errorf("读取配置文件失败，err: %v\n",err)
		return nil
	}
	d := &Yaml{}
	err = yaml.Unmarshal(yamlFile,d)
	if err != nil {
		logrus.Errorf("解析配置文件失败，err: %v\n",err)
		return nil
	}
	return &d.SqlConf
}