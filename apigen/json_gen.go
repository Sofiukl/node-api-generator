package apigen

import (
	"encoding/json"
	"fmt"
)

// WriteToJSONFile - This writes the api config into json file
func WriteToJSONFile(apimodel API) {
	lbConfig := apimodel.Methods.Detail.LbConfig
	b, err := json.Marshal(lbConfig)
	if err != nil {
		fmt.Println("Failt to  read loopback JSON file", err)
	}
	ReplaceFileContent(apimodel.Methods.Detail.FileName.JSONName, "#Replace#", string(b))
}
