package main

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// GConfig 全局配置
var GConfig *Config

// Config 保存本程序的设置
type Config struct {
	Action          string `yaml:"action,omitempty"`
	AccessKeyID     string `yaml:"accessKeyID,omitempty"`
	AccessSecret    string `yaml:"accessSecret,omitempty"`
	RegionID        string `yaml:"regionID,omitempty"`
	InstanceID      string `yaml:"instanceID,omitempty"`
	SecurityGroupID string `yaml:"securityGroupID,omitempty"`
}

// LoadConfig 加载配置文件
func LoadConfig(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	_Gconfig := Config{}
	GConfig = &_Gconfig
	if err != nil {
		log.Infof("未找到配置文件: %v", path)
		var regionID, accessKeyID, accessSecret string
		fmt.Printf("请依次输入regionID，accessKeyID和accessSecret，以空格分隔，按回车键结束输入: ")
		fmt.Scanln(&regionID, &accessKeyID, &accessSecret)
		log.Infof("你输入的信息为，regionID：%s，accessKeyID，%s，accessSecret，%s，保存于文件：%s", regionID, accessKeyID, accessSecret, path)
		GConfig.RegionID = regionID
		GConfig.AccessKeyID = accessKeyID
		GConfig.AccessSecret = accessSecret
		GConfig.Action = "describeInstanceStatus"
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("生成配置文件失败: %v", err)
		}
		defer f.Close()
		out, err := yaml.Marshal(&GConfig)
		if err != nil {
			log.Fatalf("序列化失败: %v", err)
		}
		if _, err := f.Write(out); err != nil {
			log.Fatalf("写入文件失败: %v", err)
		}

		return
	}
	err = yaml.Unmarshal(yamlFile, GConfig)
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
}
