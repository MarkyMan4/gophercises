package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// basic loading json into a map
	jsonData := []byte(`{"item1": 1, "item2": "test", "item3": 0.23}`)
	var data map[string]interface{}

	json.Unmarshal(jsonData, &data)

	fmt.Println(data)
	fmt.Println(data["item1"])

	// read json from file and load it into a map
	var fileData map[string]interface{}
	contents, err := os.ReadFile("data.json")

	if err != nil {
		log.Fatal("error reading file")
	}

	json.Unmarshal(contents, &fileData)

	// get keys of map and print each one
	for k := range fileData {
		fmt.Println(k, ":", fileData[k])
	}

	// Marshal json and print to standard out
	res, _ := json.MarshalIndent(fileData, "", "  ")

	fmt.Println(string(res))
}
