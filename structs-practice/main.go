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
