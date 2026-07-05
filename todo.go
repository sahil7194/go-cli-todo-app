package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Print(err)

		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	isComplted := t[index].Completed
	if !isComplted {
		completionTime := time.Now()

		t[index].CompletedAt = completionTime
	}

	t[index].Completed = !isComplted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := todos.validateIndex(index); err != nil {

		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Titile", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❗️"
		completedAt := ""
		if t.Completed {
			completed = "✅"
			if !t.CompletedAt.IsZero() {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()

}
