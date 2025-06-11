package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "todo",
    Short: "Aplikasi CLI untuk mengelola to-do list",
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}