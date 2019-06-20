package util

import (
	"bufio"

	"github.com/metaslim/challenge/lib/config"
	"github.com/spf13/viper"
)

//ReadInput will read user keyboard input and remove the trialing newline
func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	return input
}

func LoadConfiguration() (config.Config, error) {
	appConfig := config.Config{}

	viper.AddConfigPath("config")
	viper.SetConfigName("default")

	err := viper.ReadInConfig()

	if err != nil {
		return config.Config{}, err
	}

	viper.Unmarshal(&appConfig)
	return appConfig, nil
}
