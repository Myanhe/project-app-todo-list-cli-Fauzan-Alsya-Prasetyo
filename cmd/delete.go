package cmd

import (
    "fmt"
    "project-app-todo-list-cli/utils"
    "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [judul]",
    Short: "Menghapus tugas berdasarkan judul",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        title := args[0]
        tasks, err := utils.LoadTasks()
        if err != nil {
            fmt.Println("Gagal memuat tugas:", err)
            return
        }

        newTasks := []models.Task{}
        deleted := false
        for _, task := range tasks {
            if task.Title == title {
                deleted = true
                continue
            }
            newTasks = append(newTasks, task)
        }

        if !deleted {
            fmt.Println("Tugas tidak ditemukan.")
            return
        }

        utils.SaveTasks(newTasks)
        fmt.Println("Tugas berhasil dihapus.")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
