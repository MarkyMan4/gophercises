package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Dog struct {
	Breed      string `json:"breed"`
	Size       string `json:"size"`
	Temperment string `json:"temperment"`
}

// returning HTML
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>hello</h1>"))
}

func GetDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dog := Dog{
		Breed:      "Shepadoodle",
		Size:       "Medium",
		Temperment: "Friendly",
	}

	json.NewEncoder(w).Encode(dog)
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/dog", GetDog)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
