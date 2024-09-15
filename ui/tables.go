package cmd

import (
    "github.com/olekukonko/tablewriter"
    "os"
    "github.com/spf13/cobra"
)

var tablesCmd = &cobra.Command{
    Use:   "tables",
    Short: "Test table output",
    Run: func(cmd *cobra.Command, args []string) {
        table := tablewriter.NewWriter(os.Stdout)
        table.SetHeader([]string{"Name", "Age"})
        table.Append([]string{"John Doe", "30"})
        table.Append([]string{"Jane Doe", "25"})
        table.Render()
    },
}

func init() {
    rootCmd.AddCommand(tablesCmd)
}
