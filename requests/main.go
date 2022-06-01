/*

Make a request to a url specified in the command line arguments and print it out

*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("provide a url as an argument")
		os.Exit(0)
	}

	url := args[1]
	res, reqErr := http.Get(url)

	if reqErr != nil {
		fmt.Printf("error making request to %s\n", url)
		os.Exit(0)
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		fmt.Printf("error parsing body of response")
		os.Exit(0)
	}

	fmt.Println(string(body))
}
