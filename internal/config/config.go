package config

import (
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		Http Http
		DB   DB
	}

	Http struct {
		Host         string        `mapstructure:"host"`
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	}

	DB struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Username string `mapstructure:"username"`
		Database string `mapstructure:"database"`
	}
)

func Init(configFile string) (*Config, error) {
	var cfg Config
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	viper.AddConfigPath(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err

	} // Find and read the config file
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("server", &cfg.Http); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return err
	}
	return nil
}
