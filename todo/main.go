package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

var items []Todo

const TodoFile string = "todos.json"

type Todo struct {
	Id         int
	Name       string
	IsComplete bool
}

func newTodo(name string) *Todo {
	todo := Todo{Id: getNextId(items), Name: name, IsComplete: false}

	return &todo
}

func getNextId(todos []Todo) int {
	maxId := 0

	for i := 0; i < len(todos); i++ {
		if todos[i].Id > maxId {
			maxId = todos[i].Id
		}
	}

	return maxId + 1
}

func getTodos() []Todo {
	contents, err := os.ReadFile(TodoFile)

	if err != nil {
		log.Fatal("failed to read todos file")
	}

	json.Unmarshal(contents, &items)

	return items
}

func writeItems() {
	// finally, write todos back to file
	data, err := json.Marshal(items)

	if err != nil {
		log.Fatal("cannot marshal todos")
	}

	os.WriteFile(TodoFile, data, 0644)
}

func addTodoItem() {
	var name string
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Todo name: ")
	name, err := in.ReadString('\n')

	if err != nil {
		log.Fatal("failed to read input")
	}

	name = name[:len(name)-1] // remove newline character

	// create a new todo and append it to the list
	todo := newTodo(name)
	items = append(items, *todo)

	writeItems()
}

func listTodoItems(showCompleted bool) {
	var itemsForDisplay []Todo

	if showCompleted {
		itemsForDisplay = items
	} else {
		for i := 0; i < len(items); i++ {
			if !items[i].IsComplete {
				itemsForDisplay = append(itemsForDisplay, items[i])
			}
		}
	}

	if len(itemsForDisplay) == 0 {
		fmt.Println("no todos added yet")
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(writer, "Id\tName\tIsComplete")

	for i := 0; i < len(itemsForDisplay); i++ {
		fmt.Fprintf(writer, "%d\t%s\t%t\n", itemsForDisplay[i].Id, itemsForDisplay[i].Name, itemsForDisplay[i].IsComplete)
	}

	writer.Flush()
}

func markTodoComplete() {
	if len(items) == 0 {
		fmt.Println("no todos added yet")
		return
	}

	// show todo items, then let the user select one to mark complete
	listTodoItems(false)

	var inp string
	fmt.Print("\nID of completed item: ")
	fmt.Scanln(&inp)
	id, err := strconv.Atoi(inp)

	if err != nil {
		log.Fatal("invalid ID")
	}

	for i := 0; i < len(items); i++ {
		if items[i].Id == id {
			items[i].IsComplete = true
			break
		}
	}

	writeItems()
}

func clearTodos() {
	items = make([]Todo, 0)
	writeItems()
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Specify a command")
		return
	}

	cmd := os.Args[1]
	items = getTodos()

	if cmd == "add" {
		addTodoItem()
	} else if cmd == "list" {
		showCompleted := false

		if len(os.Args) >= 3 && os.Args[2] == "-c" {
			showCompleted = true
		}

		listTodoItems(showCompleted)
	} else if cmd == "complete" {
		markTodoComplete()
	} else if cmd == "clear" {
		clearTodos()
	} else {
		fmt.Println("options are add, list and complete")
	}
}
