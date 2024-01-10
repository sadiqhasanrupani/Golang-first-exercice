package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

// ^ lets create a interface which will have a save method init
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	displayData(&note, "note")
	displayData(&todo, "todo")

	printSomething("bruh")
	printSomething(1 + 2)
	anotherAnyMethod(1.123)
}

func printSomething(value any) error {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return nil
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float: ", floatVal)
		return nil
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", stringVal)
		return nil
	}

	return nil
}

func anotherAnyMethod(value interface{}) {

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	}
}

func displayData(data outputtable, contentNm string) {
	data.Display()
	err := saveDate(data, contentNm)

	if err != nil {
		return
	}
}

func saveDate(data saver, contentName string) error {
	err := data.Save()

	if err != nil {
		fmt.Printf("\n%v was not able to save successfully.", contentName)
		return err
	}

	fmt.Printf("\n%v is saved successfully.", contentName)
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("\n%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
