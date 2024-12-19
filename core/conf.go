package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

// InitCore InitCore读取yaml的配置
func InitCore(path string) {
	if path == "" {
		path = ConfigFile
	}
	c := &config.Config{}
	yamlConf, err := os.ReadFile(path)
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
func SetYaml() error {

	byteDate, err := yaml.Marshal(global.Config)
	// fmt.Println(string(byteDate))
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFile, byteDate, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Logger.Info("配置文件修改成功")
	return nil
}
