/*

Make an API request, unmarshal the data and store it in a database

*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getKeyAndCoords() (interface{}, interface{}, interface{}, error) {
	var err error
	apiKey, apiKeyErr := GetSecret("apiKey")
	lat, latErr := GetSecret("lat")
	lon, lonErr := GetSecret("lon")

	if apiKeyErr != nil || latErr != nil || lonErr != nil {
		err = os.ErrNotExist
	}

	return apiKey, lat, lon, err
}

func main() {
	apiKey, lat, lon, secretErr := getKeyAndCoords()

	if secretErr != nil {
		log.Fatalln("failed to get secrets")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, apiKey)
	res, reqErr := http.Get(url)

	if reqErr != nil {
		log.Fatalln("failed to make request")
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatalln("failed to read response body")
	}

	fmt.Println(string(body))
}
