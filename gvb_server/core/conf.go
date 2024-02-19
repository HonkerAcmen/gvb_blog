package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const configFile = "settings.yaml"

// InitConfig : 读取yaml文件的配置
func InitConfig() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("获取yaml文件错误: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Config init unmarshal : %v", err)
	}

	log.Println("配置文件解析成功!")
	global.Config = c
}

func UpdateConfigForSite() (err error) {
	biteinfo, err := yaml.Marshal(global.Config)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(configFile, biteinfo, fs.ModePerm)
	if err != nil {
		return
	}
	global.Log.Info("配置文件修改成功!")
	return nil
}
