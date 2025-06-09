package models

// Task merepresentasikan entri tugas pada To-Do List.
type Task struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Done        bool   `json:"done"`
}
