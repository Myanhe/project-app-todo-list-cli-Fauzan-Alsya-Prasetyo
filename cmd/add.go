package cmd

import (
    "fmt"
    "project-app-todo-list-cli/models"
    "project-app-todo-list-cli/utils"
    "github.com/spf13/cobra"
)

var (
    title, description, dueDate, urgency string
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Menambahkan tugas baru",
    Run: func(cmd *cobra.Command, args []string) {
        if title == "" {
            fmt.Println("Judul tidak boleh kosong.")
            return
        }

        if urgency != "Low" && urgency != "Medium" && urgency != "High" {
            fmt.Println("Urgency harus Low, Medium, atau High.")
            return
        }

        tasks, _ := utils.LoadTasks()
        for _, t := range tasks {
            if t.Title == title {
                fmt.Println("Tugas dengan judul tersebut sudah ada.")
                return
            }
        }

        task := models.Task{
            Title:       title,
            Description: description,
            DueDate:     dueDate,
            Urgency:     urgency,
            Done:        false,
        }

        tasks = append(tasks, task)
        utils.SaveTasks(tasks)
        fmt.Println("Tugas berhasil ditambahkan.")
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
    addCmd.Flags().StringVarP(&title, "title", "t", "", "Judul tugas")
    addCmd.Flags().StringVarP(&description, "desc", "d", "", "Deskripsi")
    addCmd.Flags().StringVar(&dueDate, "due", "", "Batas waktu")
    addCmd.Flags().StringVarP(&urgency, "urgency", "u", "Low", "Urgensi (Low/Medium/High)")
}
