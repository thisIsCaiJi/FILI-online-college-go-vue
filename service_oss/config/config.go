package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Yml *Yaml

type Yaml struct {
	OssConf OSSconfig `yaml:"oss"`
}

type OSSconfig struct{
	EndPoint string`yaml:"end_point"`
	KeyId string`yaml:"key_id"`
	KeySecret string`yaml:"key_secret"`
	BucketName string`yaml:"bucket_name"`
}

func GetOssConfig() *OSSconfig {
	if Yml != nil {
		return &Yml.OssConf
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
	return &d.OssConf
}