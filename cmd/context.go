package cmd

import (
	"bufio"
	"fmt"
	"os"
    "time"
	"strings"
	"os/exec"

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
                return m, nil
            case 3:
                m.envType = m.input
                if m.name == "" || m.server == "" || m.token == "" || (m.envType != "openshift" && m.envType != "rancher") {
                    fmt.Println("Error: All fields are required and environment type must be 'openshift' or 'rancher'")
                    return m, tea.Quit
                }
                addContext(m.name, m.server, m.token, m.envType)
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
    case 3:
        return fmt.Sprintf("    Enter the environment type (openshift/rancher): %s", m.input)
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
			fmt.Printf("\n   %s   List available contexts", Cyan("list"))
			fmt.Printf("\n   %s    Get current context", Cyan("get"))
			fmt.Printf("\n   %s    Set a specific context", Cyan("set"))
			fmt.Printf("\n   %s    Add a new context", Cyan("add"))
			fmt.Printf("\n   %s    Update an existing context", Cyan("update"))
			fmt.Printf("\n   %s    Remove a context\n", Cyan("remove"))
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
			fmt.Printf("Invalid command: %s\n", InvalidArg(args[0]))
			fmt.Println(" ")
			fmt.Printf("Valid commands are: %s, %s, %s, %s, %s, %s\n", Yellow("list"), Yellow("get"), Yellow("set"), Yellow("add"), Yellow("update"), Yellow("remove"))
		}
	},
}

func listContexts() {
    
    file, err := os.Open(contextFile)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var contexts []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            continue // Ignora linhas em branco
        }
        contexts = append(contexts, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    if len(contexts) == 0 {
        fmt.Println("No contexts available. Please add a context first.")
        return
    }

    var maxKeyLength int
    contextMap := make(map[string]string)
    for _, line := range contexts {
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            key := parts[0]
            value := parts[1]
            if strings.HasSuffix(key, "-token") {
                continue
            }
            if _, exists := contextMap[key]; !exists {
                contextMap[key] = value
                if len(key) > maxKeyLength {
                    maxKeyLength = len(key)
                }
            }
        }
    }

    fmt.Println("Here is the list of available contexts to configure:")

    for key, value := range contextMap {
        fmt.Printf("   %-*s = %s\n", maxKeyLength, key, value)
    }
}


func getCurrentContext() {
	fmt.Println("Getting current context...")
	executeCommand("kubectl", "config", "current-context")
}

func addContext(name, server, token, envType string) {
    
    if envType == "openshift" {
        name = "os-" + name
    }

	file, err := os.OpenFile(contextFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	contextEntry := fmt.Sprintf("%s=%s\n", name, server)
	tokenEntry := fmt.Sprintf("%s-token=%s\n", name, token)
	
	_, err = file.WriteString(contextEntry)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	
	_, err = file.WriteString(tokenEntry)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Context '%s' with server '%s' added successfully.\n", name, server)
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
        if strings.HasPrefix(line, name+"=") || strings.HasPrefix(line, name+"-token=") {
            continue // Ignora linhas que correspondem ao contexto a ser removido
        }
        lines = append(lines, line)
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

    fmt.Printf("Context '%s' removed successfully.\n", name)
}

func setContext(context string) {
    contexts, err := loadContexts()
    if err != nil {
        fmt.Println("Error loading contexts:", err)
        return
    }

    value, ok := contexts[context]
    if !ok {
        fmt.Println("Invalid context. Use 'k context list' to see available contexts.")
        return
    }

    var server, token string
    var isOpenShift bool

    if strings.HasPrefix(context, "os-") {
        server, token = parseValue(value)
        isOpenShift = true
    } else {
        server, token = parseValue(value)
        isOpenShift = false
    }

    if isOpenShift {
        setOpenShiftContext(server, token)
    } else {
        setRancherContext(server, token)
    }
}

func parseValue(value string) (string, string) {
    parts := strings.SplitN(value, " ", 2)
    if len(parts) == 2 {
        return parts[0], parts[1]
    }
    return parts[0], "" // Return empty string if token is not present
}

func setOpenShiftContext(server, token string) {
    showSpinner("Setting the context to "+server+" [openshift]...", func() {
        executeCommand("oc", "login", "--token="+token, "--server="+server)
        time.Sleep(1 * time.Second)
    })

    fmt.Println("Current context:")
    executeCommand("kubectl", "config", "current-context")
}


func setRancherContext(server, token string) {
    showSpinner("Setting the context to "+server+" [rancher]...", func() {
        executeCommand("rancher", "login", "--token="+token, "--server="+server)
        time.Sleep(1 * time.Second)
    })
	fmt.Println("Current context:")
	executeCommand("kubectl", "config", "current-context")
}

func executeCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", Red(err))
	}
}

func loadContexts() (map[string]string, error) {
	contexts := make(map[string]string)
	file, err := os.Open(contextFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			contexts[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return contexts, nil
}

func init() {
	rootCmd.AddCommand(contextCmd)
}
