package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App *AppConfig
	//DB  *DBConfig
}

type AppConfig struct {
	Port        int    `default:1203`
	VerifyToken string `required:"true"`
	AppID       string `required:"true"`
	AppSecret   string `required:"true"`
	AccessToken string `required:"true"`
	FBGroupID   string `required:"true"`
}

type DBConfig struct {
	Host string `default:"localhost"`
	Port string `default:"3306"`
	Name string `required:"true"`
	User string `required:"true"`
	Pass string `required:"true"`
}

func loadConfigFromEnv() (*Config, error) {
	var app AppConfig

	if err := envconfig.Process("message", &app); err != nil {
		return nil, err
	}

	return &Config{
		App: &app,
	}, nil
}
