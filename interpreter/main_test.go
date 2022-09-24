package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	m := map[string]string{}

	if res, ok := m["test"]; !ok {
		fmt.Println("nil")
	} else {
		fmt.Println(res)
	}
}
