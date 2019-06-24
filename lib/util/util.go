package util

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/metaslim/challenge/lib/config"
	"github.com/spf13/viper"
)

//ReadInput will read user keyboard input and remove the trailing newline
func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	return input
}

//LoadConfiguration will load config file
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

//CaptureOutput will capture stdout
func CaptureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
