package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type app struct {
	Name string
	Port int
}

type databaseConfig struct {
	DriverType string `yaml:"driver_type"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DBName     string `yaml:"db_name"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type config struct {
	App            app            `yaml:"app"`
	Database       databaseConfig `yaml:"database"`
	BcryptPassword string         `yaml:"bcrypt_password"`
	JwtSecretKey   string         `yaml:"jwt_secret_key"`
}

var Config *config

var App *app

var Database *databaseConfig

func init() {
	fmt.Printf("Loading config file")
	fileBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(fileBytes, &Config)
	if err != nil {
		return
	}
	App = &Config.App
	Database = &Config.Database
}
