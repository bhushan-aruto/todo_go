package todo

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
	CompletedAt *time.Time
}
type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}
func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return fmt.Errorf("invalid index %d: %v", index, err)

	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return fmt.Errorf("invalid index %d: %v", index, err)
	}
	isCompleted := t[index].Completed
	if !isCompleted {
		complitionTime := time.Now()
		t[index].CompletedAt = &complitionTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return fmt.Errorf("invalid index %d: %v", index, err)
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) Printt() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todos {
		completed := "❌😣"
		completedAt := ""

		if todo.Completed {
			completed = "✅🫡"
			if todo.CompletedAt != nil {
				completedAt = todo.CreatedAt.Format(time.RFC1123)

			}
		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
