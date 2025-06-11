package cmd

import (
    "fmt"
    "os"
    "strings"
    "project-app-todo-list-cli/utils"
    "text/tabwriter"

    "github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
    Use:   "search [keyword]",
    Short: "Mencari tugas berdasarkan judul atau deskripsi",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        keyword := strings.ToLower(args[0])
        tasks, err := utils.LoadTasks()
        if err != nil {
            fmt.Println("Gagal memuat tugas:", err)
            return
        }

        writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
        fmt.Fprintln(writer, "NO\tJUDUL\tDESKRIPSI\tBATAS WAKTU\tURGENSI\tSELESAI")

        count := 0
        for i, task := range tasks {
            if strings.Contains(strings.ToLower(task.Title), keyword) || strings.Contains(strings.ToLower(task.Description), keyword) {
                status := "Belum"
                if task.Done {
                    status = "Sudah"
                }
                fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\t%s\n", i+1, task.Title, task.Description, task.DueDate, task.Urgency, status)
                count++
            }
        }

        if count == 0 {
            fmt.Println("Tidak ada tugas yang cocok dengan keyword tersebut.")
        } else {
            writer.Flush()
        }
    },
}

func init() {
    rootCmd.AddCommand(searchCmd)
}
