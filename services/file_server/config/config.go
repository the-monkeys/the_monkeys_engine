package config

import "github.com/spf13/viper"

type Config struct {
	FileService string `mapstructure:"FILE_SERVICE"`
}

func LoadFileServerConfig() (config Config, err error) {
	viper.AddConfigPath("/the_monkeys/etc")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
