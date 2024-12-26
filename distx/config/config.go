package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// 从config.yaml读取配置，并且注入到Config结构体中

const yamlConfigPath = "config.yaml"

type Config struct {
	DataBase DataBase `yaml:"database"`
}

type DataBase struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	DB1DSN string `yaml:"db1_dsn"`
	DB2DSN string `yaml:"db2_dsn"`
}

func NewConfig() *Config {
	config := &Config{}
	yamlFile, err := os.ReadFile(yamlConfigPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML file: %s\n", err)
	}
	fmt.Printf("config: %+v\n", config)
	return config
}
