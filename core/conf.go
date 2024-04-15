package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"
)

// InitCore InitCore读取yaml的配置
func InitCore() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error : %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("conf init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load init success")
	global.Config = c
}
