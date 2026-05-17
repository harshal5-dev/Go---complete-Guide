package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Context   string    `json:"context"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title:     title,
		Context:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n%v\n", note.Title, note.Context)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
