package cmd

import (
    "fmt"
    "project-app-todo-list-cli/utils"
    "github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
    Use:   "done [judul]",
    Short: "Menandai tugas sebagai selesai",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        title := args[0]
        tasks, err := utils.LoadTasks()
        if err != nil {
            fmt.Println("Gagal memuat tugas:", err)
            return
        }

        found := false
        for i, task := range tasks {
            if task.Title == title {
                tasks[i].Done = true
                found = true
                break
            }
        }

        if !found {
            fmt.Println("Tugas tidak ditemukan.")
            return
        }

        utils.SaveTasks(tasks)
        fmt.Println("Tugas ditandai sebagai selesai.")
    },
}

func init() {
    rootCmd.AddCommand(doneCmd)
}