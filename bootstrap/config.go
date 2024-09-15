package bootstrap

import (
	"github.com/spf13/viper"
	"time"
)

type Security struct {
	JWTAccessSecret  string `mapstructure:"jwt_access_secret"`
	JWTRefreshSecret string `mapstructure:"jwt_refresh_secret"`
	JWTAccessExpire  time.Duration
	JWTRefreshExpire time.Duration
}

type Server struct {
	Mode           string   `mapstructure:"gin_mode"`
	TrustedProxies []string `mapstructure:"trusted_proxies"`
	Address        string   `mapstructure:"server_address"`
}

type Database struct {
	Name     string `mapstructure:"db_name"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
}

type Config struct {
	Security
	Server
	DB Database
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.LoadEnv("./.env"); err != nil {
		return nil, err
	}
	cfg.loadJWTExpire()
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

	if err := viper.Unmarshal(&cfg.Server); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.Security); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) LoadEnv(Path string) error {
	return cfg.loadFromFile(Path, "env")
}

func (cfg *Config) loadJWTExpire() {
	// TODO: Move to .yml config file.
	cfg.JWTAccessExpire = time.Hour * 1
	cfg.JWTRefreshExpire = (time.Hour * 24) * 7
}
