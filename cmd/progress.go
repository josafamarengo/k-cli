package cmd

import (
    "fmt"
    "github.com/schollz/progressbar/v3"
    "github.com/spf13/cobra"
    "time"
)

var progressCmd = &cobra.Command{
    Use:   "progress",
    Short: "Test progress bar",
    Run: func(cmd *cobra.Command, args []string) {
        bar := progressbar.Default(100)
        for i := 0; i < 100; i++ {
            time.Sleep(50 * time.Millisecond) // Simula trabalho
            bar.Add(1)
        }
        fmt.Println("\nProgress complete!")
    },
}

func init() {
    rootCmd.AddCommand(progressCmd)
}
