package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	utils "github.com/devamshi/tasks-cli/internal"
	"github.com/devamshi/tasks-cli/internal/model"
)

func LoadTasksFromDB() ([]model.Task, error) {

	isFilePresent := utils.CheckIfFileExists("db.json")

	if !isFilePresent {
		return []model.Task{}, nil
	}

	data, err := os.ReadFile("db.json")
	if err != nil {
		log.Fatal(err)
	}

	taskList := []model.Task{}

	err = json.Unmarshal(data, &taskList)
	if err != nil {
		fmt.Println("ERROR: db.json contains corrupted data. Please backup and delete db.json file to continue.")
		log.Fatal(err)
	}

	return taskList, nil
}

func SaveTasksToDB(updated []model.Task) {

	file, err := os.Create("db.json")
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(updated)
	if err != nil {

		log.Fatal(err)
	}

}
