package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Handles getting data from secrets.json
func GetSecret(secretName string) (interface{}, error) {
	secretsContent, readErr := os.ReadFile("secrets.json")
	var secret interface{}
	var err error

	if readErr != nil {
		fmt.Println("error reading secrets.json")
	}

	var secrets map[string]interface{}
	json.Unmarshal(secretsContent, &secrets)

	if val, ok := secrets[secretName]; ok {
		secret = val
	} else {
		err = os.ErrNotExist
	}

	return secret, err
}
