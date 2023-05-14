package config

import "github.com/spf13/viper"

type Config struct {
	Port     string `mapstructure:"PORT"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUrl    string `mapstructure:"DB_URL"`
}

func GetConfig() (*Config, error) {
	cfg := new(Config)

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}