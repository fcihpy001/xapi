package service

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"xapi/model"
)

var (
	config *model.Config
)

func GetConfig() *model.Config {
	if config == nil {
		//yamlFile, err := os.ReadFile("./config.yaml")
		yamlFile, err := os.ReadFile("./config.yaml")
		if err != nil {
			log.Println(err)
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			fmt.Println("Error unmarshaling YAML:", err)
			return nil
		}
		log.Println("配置文件读取成功。。。")
	}
	return config
}
