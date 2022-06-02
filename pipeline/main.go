/*

Make an API request, unmarshal the data and store it in a database

*/

package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	sec, err := GetSecret("lat")

	if err != nil {
		log.Fatalln("secret does not exist")
	}

	fmt.Println(reflect.TypeOf(sec))
}
