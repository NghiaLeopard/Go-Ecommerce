package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDrive             string        `mapstructure:"DB_DRIVE"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAction        string        `mapstructure:"SERVER_ACTION"`
	Symmetric           string        `mapstructure:"SYMMETRICKEY"`
	Access_token        time.Duration `mapstructure:"ACCESS_TOKEN"`
	Refresh_token       time.Duration `mapstructure:"REFRESH_TOKEN"`
	ForgotPasswordToken time.Duration `mapstructure:"FORGOT_PASSWORD_TOKEN"`
	Account_email       string        `mapstructure:"ACCOUNT_EMAIL"`
	Password_email      string        `mapstructure:"PASSWORD_EMAIL"`
	NameEmail           string        `mapstructure:"NAME_EMAIL"`
	AppUrlFE            string        `mapstructure:"URL_FE_APP"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// override env
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return
	}

	return
}
