package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/exec"
	"k/ui"

    "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var contextFile = "assets/contexts.txt"

type model struct {
    name   string
    server string
    token  string
    envType string
    input  string
    stage  int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyMsg:
        keyMsg := msg.(tea.KeyMsg)
        switch keyMsg.String() {
        case "q":
            return m, tea.Quit
        case "enter":
            switch m.stage {
            case 0:
                m.name = m.input
                m.input = ""
                m.stage++
                return m, nil
            case 1:
                m.server = m.input
                m.input = ""
                m.stage++
                return m, nil
            case 2:
                m.token = m.input
                m.input = ""
                m.stage++
                addContext(m.name, m.server, m.token)
                return m, tea.Quit
            }
        case "backspace":
            if len(m.input) > 0 {
                m.input = m.input[:len(m.input)-1]
            }
        default:
            m.input += keyMsg.String()
        }
    }
    return m, nil
}

func (m model) View() string {
    switch m.stage {
    case 0:
        return fmt.Sprintf("    Enter the name of context: %s", m.input)
    case 1:
        return fmt.Sprintf("    Enter the server: %s", m.input)
    case 2:
        return fmt.Sprintf("    Enter the token: %s", m.input)
    default:
        return "Unknown stage"
    }
}

func addContextInteractive() {
    p := tea.NewProgram(model{})
    if err := p.Start(); err != nil {
        fmt.Fprintf(os.Stderr, "Error starting interactive prompt: %v", err)
        os.Exit(1)
    }
}

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage contexts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("\nusage: k context <command> [<arg>]")
			fmt.Println("\nThese are the available commands:")
			fmt.Printf("\n   %s   List available contexts", ui.Cyan("list"))
			fmt.Printf("\n   %s    Get current context", ui.Cyan("get"))
			fmt.Printf("\n   %s    Set a specific context", ui.Cyan("set"))
			fmt.Printf("\n   %s    Add a new context", ui.Cyan("add"))
			fmt.Printf("\n   %s    Update an existing context", ui.Cyan("update"))
			fmt.Printf("\n   %s    Remove a context\n", ui.Cyan("remove"))
			return
		}

		switch args[0] {
		case "list":
			listContexts()
		case "get":
			getCurrentContext()
		case "set":
			if len(args) < 2 {
				fmt.Println("Usage: k context set <context>")
				return
			}
			setContext(args[1])
		case "add":
			addContextInteractive() // Chama a função interativa
		case "update":
			if len(args) < 3 {
				fmt.Println("Usage: k context update <name> <new-value>")
				return
			}
			updateContext(args[1], args[2])
		case "remove":
			if len(args) < 2 {
				fmt.Println("Usage: k context remove <name>")
				return
			}
			removeContext(args[1])
		default:
			fmt.Println(" ")
			fmt.Printf("Invalid command: %s\n", ui.InvalidArg(args[0]))
			fmt.Println(" ")
			fmt.Printf("Valid commands are: %s, %s, %s, %s, %s, %s\n", ui.Yellow("list"), ui.Yellow("get"), ui.Yellow("set"), ui.Yellow("add"), ui.Yellow("update"), ui.Yellow("remove"))
		}
	},
}

func listContexts() {
    
    cmd := exec.Command("kubectl", "config", "get-contexts", "-o", "name")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Failed to list contexts:", err)
        return
    }

    fmt.Println("Available contexts:")
    fmt.Println(string(output))
}


func getCurrentContext() {
	fmt.Println("Getting current context...")
	executeCommand("kubectl", "config", "current-context")
}

func addContext(name, clusterURL, token string) {
	cmd := exec.Command("kubectl", "config", "set-cluster", name, "--server="+clusterURL)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to set cluster:", err)
		return
	}

	cmd = exec.Command("kubectl", "config", "set-credentials", name+"-user", "--token="+token)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to set credentials:", err)
		return
	}

	cmd = exec.Command("kubectl", "config", "set-context", name, "--cluster="+name, "--user="+name+"-user")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to set context:", err)
		return
	}

	cmd = exec.Command("kubectl", "config", "use-context", name)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to switch context:", err)
		return
	}

	fmt.Println("Successfully added and switched to context:", name)
}


func updateContext(name, newValue string) {
	file, err := os.Open(contextFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, name+"=") {
			lines = append(lines, fmt.Sprintf("%s=%s", name, newValue))
		} else {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	file, err = os.Create(contextFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Printf("Context '%s' updated successfully.\n", name)
}

func removeContext(name string) {
    // Remover o contexto
    cmd := exec.Command("kubectl", "config", "delete-context", name)
    err := cmd.Run()
    if err != nil {
        fmt.Println("Failed to remove context:", err)
        return
    }
    fmt.Printf("Successfully removed context: %s\n", name)

    // Remover o cluster associado
    cmd = exec.Command("kubectl", "config", "unset", "clusters."+name)
    err = cmd.Run()
    if err != nil {
        fmt.Println("Failed to remove cluster:", err)
        return
    }
    fmt.Printf("Successfully removed cluster: %s\n", name)

    // Remover o usuário associado
    cmd = exec.Command("kubectl", "config", "unset", "users."+name+"-user")
    err = cmd.Run()
    if err != nil {
        fmt.Println("Failed to remove user:", err)
        return
    }
    fmt.Printf("Successfully removed user: %s\n", name+"-user")
}

func setContext(context string) {
    cmd := exec.Command("kubectl", "config", "use-context", context)
    err := cmd.Run()
    if err != nil {
        fmt.Println("Failed to switch context:", err)
        return
    }

    fmt.Println("Successfully switched to context:", context)
}

func executeCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", ui.Red(err))
	}
}

func init() {
	rootCmd.AddCommand(contextCmd)
}
