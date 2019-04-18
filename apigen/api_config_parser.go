package apigen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"os"
)

const (
	configFile = "./apigen/config/api-campus.json"
)

// ParseAPIConfig - This function parse the api config file
func ParseAPIConfig() {
	file, e := ioutil.ReadFile(configFile)
	if e != nil {
		fmt.Printf("Config File error: %v\n", e)
		os.Exit(1)
	}

	var apimodel API
	json.Unmarshal(file, &apimodel)
	fmt.Printf("Information provided in config files: %v\n", apimodel)
	GenerateAPI(apimodel)
}
