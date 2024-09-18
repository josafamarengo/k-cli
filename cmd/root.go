package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "k/ui"
)

var (
    appVersion = "1.0.0"
    appInfo    = "K CLI tool for Kubernetes"
    versionFlag bool
)

var rootCmd = &cobra.Command{
    Use:   "k",
    Short: "K is a ALL-IN-ONE CLI tool for Kubernetes",
    Run: func(cmd *cobra.Command, args []string) {

        if versionFlag {
            fmt.Printf("%s - Version: %s\n", appInfo, appVersion)
            return
        }

		if len(args) == 0 {
            fmt.Println("\nusage: k [command] [<arg>]")
            fmt.Println("\nAvailable commands:")
            fmt.Printf("   %s          Context command", ui.Cyan("context"))
            fmt.Printf("\n   %s    Tools instalation command", ui.Cyan("install-tools"))
            fmt.Println(" ")
            fmt.Println(" ")
            fmt.Println("\n" + appInfo + " - " + appVersion)
            return
        }
    },
}

// Execute é a função que executa o comando raiz.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    // Adiciona a flag --version na CLI
    rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Display the version of the tool")
}
