package main

import (
	"encoding/xml"
	"log"
	"os"
)

type Book struct {
	XMLName  xml.Name    `xml:"book"`
	Title    string      `xml:"title"`
	ChapList ChapterList `xml:"chapters"`
}

type ChapterList struct {
	Chapters []Chapter `xml:"chapter"`
}

type Chapter struct {
	Name  string `xml:",chardata"`
	Pages int    `xml:"pages,attr"`
}

func main() {
	chapters := []Chapter{
		{Name: "chapter 1", Pages: 32},
		{Name: "chapter 2", Pages: 53},
		{Name: "chapter 3", Pages: 46},
		{Name: "chapter 4", Pages: 35},
	}

	chapList := &ChapterList{Chapters: chapters}
	book := &Book{Title: "test", ChapList: *chapList}
	res, err := xml.MarshalIndent(book, "", "  ")

	if err != nil {
		log.Fatalln("Failed to marshal data")
	}

	// write the result to a file
	os.WriteFile("book.xml", res, 0644)
}
