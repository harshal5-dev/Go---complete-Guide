package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
	}

	flatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", flatVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// default:
	// 	fmt.Println(value)
	// }
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	// if err != nil {
	// 	return "", ""
	// }

	content := getUserInput("Note content:")

	// if err != nil {
	// 	return "", ""
	// }

	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo Text: ")
	return text
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func main() {
	printSomething("HI Any value")
	printSomething(21)
	printSomething(22.3)

	title, content := getNoteData()
	todoText := getTodoData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(note)

	if err != nil {
		return
	}
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func getUserInput(prompt string) string {
	// var value string

	fmt.Print(prompt)
	// fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	// if value == "" {
	// 	return "", errors.New("Invalid input.")
	// }

	return text
}
