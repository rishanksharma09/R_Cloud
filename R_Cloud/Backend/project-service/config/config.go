package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/tools/go/cfg"
)

type Config struct {
	Port string
	DatabaseURL string
	GitHubToken string
}
func LoadConfig() (*Config,error) {
	_=godotenv.Load()
	cfg:= &Config{
		Port: os.Getenv("PROJECT_SERVICE_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		GitHubToken: os.Getenv("GITHUB_TOKEN"),
	}
	if cfg.Port== ""{
		return nil,fmt.Errorf("Project_service_Port is not set")

	}
		if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	if cfg.GitHubToken == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN is not set")
	}

	return cfg, nil
	


}