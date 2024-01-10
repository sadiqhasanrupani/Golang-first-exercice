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
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note *Note) Display() {
	fmt.Printf("\nYour note titled %v has the following content:\n%v", note.Title, note.Content)
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

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("title is invalid")
	}

	if content == "" {
		return Note{}, errors.New("content is invalid")
	}

	note := Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	return note, nil
}
