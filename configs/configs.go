package configs

import "os"

type ServerConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
}

func LoadDBConfig() DBConfig {
	config := DBConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
	return config
}

func LoadServerConfig() ServerConfig {
	config := ServerConfig{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}
	return config
}
