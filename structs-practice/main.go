package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

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

func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
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
