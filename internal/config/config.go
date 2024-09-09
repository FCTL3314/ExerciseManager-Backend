package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Name     string `mapstructure:"db_name"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
}

type Config struct {
	DB Database
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.LoadEnv("./.env"); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) setupViper(Path string, ConfigType string) error {
	viper.SetConfigFile(Path)
	viper.SetConfigType(ConfigType)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) loadFromFile(Path string, ConfigType string) error {
	if err := cfg.setupViper(Path, ConfigType); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.DB); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) LoadEnv(Path string) error {
	return cfg.loadFromFile(Path, "env")
}
