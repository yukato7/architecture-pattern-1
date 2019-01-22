package config

import "github.com/BurntSushi/toml"

type Config struct {
	Server   Server `toml:"server"`
	DBMaster DB     `toml:"dbm"`
	DBSlave  DB     `toml:"dbs"`
}

type Server struct {
	Port string `toml:"port"`
}

type DB struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

func New(config *Config, configPath string) error {
	_, err := toml.DecodeFile(configPath, config)
	return err
}
