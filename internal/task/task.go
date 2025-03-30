package task

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"

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

	for idx, task := range tasks {
		if idx == 0 {
			fmt.Printf("%-10s %-40s %-30s\n", "ID", "Description", "Status")
		}
		fmt.Printf("%-10d %-40s %-30s\n", task.Id, task.Description, task.Status)
	}

	return nil
}

func Delete(id int) {
	tasks, err := store.LoadTasksFromDB()
	if err != nil {
		fmt.Printf("error while loading tasks: " + err.Error())
		return
	}

	updatedTasks := slices.DeleteFunc(tasks, func(task model.Task) bool {
		return task.Id == id
	})

	store.SaveTasksToDB(updatedTasks)

	if len(updatedTasks) != len(tasks) {
		fmt.Printf("Deleted task with Id: %d\n", id)
	} else {
		fmt.Printf("No task found with id: %d\n", id)
	}

}

func UpdateTask(id int) {}
