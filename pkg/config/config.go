package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	BASE_URL   string `mapstructure:"BASE_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AUTHTOKEN  string `mapstructure:"TWILIO_AUTHTOKEN"`
	ACCOUNTSID string `mapstructure:"TWILIO_ACCOUNTSID"`
	SERVICESID string `mapstructure:"TWILIO_SERVIECSID"`

	KEY       string `mapstructure:"KEY"`
	KEY_ADMIN string `mapstructure:"KEY_ADMIN"`

	KEY_ID_FOR_PAY     string `mapstructure:"KEY_ID_FOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`
}

var envs = []string{
	"BASE_URL", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "TWILIO_AUTHTOKEN", "TWILIO_ACCOUNTSID", "TWILIO_SERVICESID", "KEY", "KEY_ADMIN", "KEY_ID_FOR_PAY", "SECRET_KEY_FOR_PAY",
}

func LoadConfig() (Config, error) {
	var confg Config
	viper.AddConfigPath("/")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return confg, err
		}
	}

	if err := viper.Unmarshal(&confg); err != nil {
		return confg, err
	}

	if err := validator.New().Struct(&confg); err != nil {
		return confg, err
	}
	return confg, nil
}
