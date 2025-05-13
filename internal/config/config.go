package config

import "github.com/BurntSushi/toml"

type Config struct {
	Postgres Postgres `toml:"postgres"`
}

type Postgres struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

func NewConfig(path string) (*Config, error) {
	if path == "" {
		path = "../config.toml"
	}
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
