package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Yml *Yaml

type Yaml struct {
	VodConf Vodconfig `yaml:"vod"`
	ServerConf Server `yaml:"server"`
}

type Vodconfig struct{
	EndPoint string`yaml:"end_point"`
	KeyId string`yaml:"key_id"`
	KeySecret string`yaml:"key_secret"`
	BucketName string`yaml:"bucket_name"`
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

func GetVodConfig() *Vodconfig {
	return &Yml.VodConf
}