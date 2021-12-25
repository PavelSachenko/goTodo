package config

import (
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		Http Http
	}

	Http struct {
		Host         string        `mapstructure:"host"`
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	}
)

func Init(configFile string) (*Config, error) {

	var cfg Config
	//viper.AddConfigPath(fmt.Sprintf("%s/%s", configPath, configFile))
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
	return nil
}
