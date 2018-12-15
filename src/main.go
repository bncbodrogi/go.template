package main

import (
	"fmt"
	"io/ioutil"
	"util"

	"config"
)

var (
	conf   = config.Config()
	path   = conf["logo_path"].(string)
	dbName = conf["db_name"].(string)
)

func main() {
	printLogo(conf["logo_path"].(string))
}

func printLogo(path string) {
	logo, err := ioutil.ReadFile(path)
	util.CheckErr(err)
	fmt.Println(string(logo))
}
