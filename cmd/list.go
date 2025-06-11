package cmd

import (
    "fmt"
    "os"
    "project-app-todo-list-cli/utils"
    "text/tabwriter"

    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "Menampilkan semua tugas",
    Run: func(cmd *cobra.Command, args []string) {
        tasks, err := utils.LoadTasks()
        if err != nil {
            fmt.Println("Gagal memuat tugas:", err)
            return
        }

        if len(tasks) == 0 {
            fmt.Println("Tidak ada tugas.")
            return
        }

        writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
        fmt.Fprintln(writer, "NO\tJUDUL\tDESKRIPSI\tBATAS WAKTU\tURGENSI\tSELESAI")
        for i, task := range tasks {
            status := "Belum"
            if task.Done {
                status = "Sudah"
            }
            fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\t%s\n", i+1, task.Title, task.Description, task.DueDate, task.Urgency, status)
        }
        writer.Flush()
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}