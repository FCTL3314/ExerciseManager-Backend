package bootstrap

import (
	"github.com/spf13/viper"
	"time"
)

type JWT struct {
	AccessExpire  time.Duration `mapstructure:"access_expire"`
	RefreshExpire time.Duration `mapstructure:"refresh_expire"`
	AccessSecret  string        `mapstructure:"jwt_access_secret"`
	RefreshSecret string        `mapstructure:"jwt_refresh_secret"`
}

type Pagination struct {
	MaxUserLimit     int `mapstructure:"max_user_limit"`
	MaxWorkoutLimit  int `mapstructure:"max_workout_limit"`
	MaxExerciseLimit int `mapstructure:"max_exercise_limit"`
}

type Entities struct {
	Workout
}

type Workout struct {
	MaxExercisesCount int `mapstructure:"max_exercises_count"`
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
	Server
	Database
	JWT
	Pagination
	Entities
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.LoadEnv("./.env"); err != nil {
		return nil, err
	}
	if err := cfg.LoadYml("./config.yml"); err != nil {
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

	if err := viper.Unmarshal(&cfg.Server); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.Database); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.JWT); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.Pagination); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.Entities); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) LoadEnv(Path string) error {
	return cfg.loadFromFile(Path, "env")
}

func (cfg *Config) LoadYml(Path string) error {
	return cfg.loadFromFile(Path, "yml")
}
