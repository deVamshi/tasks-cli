package task

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/devamshi/tasks-cli/internal/model"
	"github.com/devamshi/tasks-cli/internal/store"
)

func AddTask(desc string) error {

	// validations
	if len(desc) == 0 {
		return errors.New("Description can't be empty")
	}

	tasks, err := store.LoadTasksFromDB()
	if err != nil {
		return fmt.Errorf("error while loading tasks %w", err)
	}

	newTask := model.Task{
		Id:          rand.Intn(100000),
		Description: desc,
		Status:      STATUS_TODO,
	}

	tasks = append(tasks, newTask)

	store.SaveTasksToDB(tasks)

	return nil
}

func ListTasks(status string) error {

	tasks, err := store.LoadTasksFromDB()
	if err != nil {
		return fmt.Errorf("error while loading tasks %w", err)
	}

	fmt.Println(tasks)

	return nil
}

func Delete(id int) {}

func UpdateTask(id int) {}
