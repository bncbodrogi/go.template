package config

import (
	"bytes"
	"io/ioutil"
	"util"

	"github.com/spf13/viper"
)

var (
	configPath = "./src/props/vars.yml"
)

//Config loads the configuration file for the project
func Config() map[string]interface{} {
	viper.SetConfigType("yml")
	configFile, err := ioutil.ReadFile(configPath)
	util.CheckErr(err)

	viper.ReadConfig(bytes.NewBuffer(configFile))
	return viper.AllSettings()
}
