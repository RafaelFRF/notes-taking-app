package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Use “ for formating. On this case, json files will receive title instead of Title
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	// Marshal uses just public variables
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input.")
	}
	return Todo{
		Text: content,
	}, nil
}
