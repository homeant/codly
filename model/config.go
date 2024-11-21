package model

type App struct {
	Name string
	Port int
}

type DatabaseConfig struct {
	DriverType string `json:"driver_type"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	DBName     string `yaml:"db_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Config struct {
	App      App            `json:"app"`
	Database DatabaseConfig `json:"database"`
}
