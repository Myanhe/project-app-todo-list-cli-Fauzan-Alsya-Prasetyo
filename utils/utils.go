package utils

import (
    "encoding/json"
    "os"
    "fmt"
)

// SaveTasks menyimpan slice tugas ke file JSON (overwrite).
func SaveTasks(filename string, tasks []models.Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    // Tulis data JSON ke file (perizinan 0644)
    return os.WriteFile(filename, data, 0644)
}

// LoadTasks membaca file JSON dan mengembalikan slice tugas.
// Jika file tidak ada, kembalikan slice kosong.
func LoadTasks(filename string) ([]models.Task, error) {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        // File belum ada, kembalikan slice kosong
        return []models.Task{}, nil
    }
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    var tasks []models.Task
    if err := json.Unmarshal(data, &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}

