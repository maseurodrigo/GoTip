package DataManage

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadStruct(fileStruct string, config interface{}) bool {

	// Read the JSON file
	file, jsonStruct := ioutil.ReadFile(fileStruct)
	if jsonStruct != nil {
		log.Fatal(jsonStruct)
		return false
	}

	// Unmarshal the JSON data into the config interface
	jsonStruct = json.Unmarshal(file, config)
	if jsonStruct != nil {
		log.Fatal(jsonStruct)
		return false
	}
	return true
}
