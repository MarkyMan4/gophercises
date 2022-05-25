/*

This is a program to load data from a JSON file and insert it
into a SQL database.

*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationYear"`
	Genre           string `json:"genre"`
}

// load the json data from a file
func loadData() []Book {
	data, err := os.ReadFile("data/books.json")

	if err != nil {
		log.Fatal("Failed to read file")
	}

	var jsonData []Book
	json.Unmarshal(data, &jsonData)

	return jsonData
}

// insert books into book table
func insertData(db *sql.DB, books []Book) int {
	numSuccess := 0

	for i := 0; i < len(books); i++ {
		stmt, prepErr := db.Prepare(`
			insert into book (title, author, publication_year, genre) 
			values (?, ?, ?, ?)
		`)

		if prepErr != nil {
			log.Fatal("Failed to prepare insert statement")
			continue
		}

		_, insertErr := stmt.Exec(books[i].Title, books[i].Author, books[i].PublicationYear, books[i].Genre)

		if insertErr != nil {
			fmt.Println("Failed to insert records")
			continue
		}

		numSuccess++
	}

	return numSuccess
}

func main() {
	data := loadData()
	db, err := sql.Open("sqlite3", "db.sqlite")

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	numSuccess := insertData(db, data)
	fmt.Printf("Successfully loaded %d records\n", numSuccess)
}
