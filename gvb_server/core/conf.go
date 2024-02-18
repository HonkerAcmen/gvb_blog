package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// InitConfig : 读取yaml文件的配置
func InitConfig() {
	const configFile = "settings.yaml"
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
