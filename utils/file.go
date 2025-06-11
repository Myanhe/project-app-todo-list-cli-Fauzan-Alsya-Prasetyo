package utils

import (
    "encoding/json"
    "os"
    "project-app-todo-list-cli/models"
)

const dataFile = "data/tasks.json"

func LoadTasks() ([]models.Task, error) {
    var tasks []models.Task
    file, err := os.ReadFile(dataFile)
    if err != nil {
        return tasks, err
    }
    json.Unmarshal(file, &tasks)
    return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(dataFile, data, 0644)
}