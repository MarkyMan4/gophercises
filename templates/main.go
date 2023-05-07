package main

import (
	"os"
	"text/template"
)

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func main() {
	dog := Dog{"Yetti", "lab", 2}
	t, err := template.New("dogs").Parse("You have a {{ .Breed }} named {{ .Name }} who is {{ .Age }} years old")

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, dog)

	if err != nil {
		panic(err)
	}
}
