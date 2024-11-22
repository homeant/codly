package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type app struct {
	Name string
	Port int
}

type databaseConfig struct {
	DriverType string `json:"driver_type"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	DBName     string `yaml:"db_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type config struct {
	App      app            `json:"app"`
	Database databaseConfig `json:"database"`
}

var App *app

var Database *databaseConfig

func init() {
	fileBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	Config := config{}
	err = yaml.Unmarshal(fileBytes, &Config)
	if err != nil {
		return
	}
	App = &Config.App
	Database = &Config.Database
}
