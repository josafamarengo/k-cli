package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install-tools",
	Short: "Install kubectl, krew and stern if not present",
	Run: func(cmd *cobra.Command, args []string) {
		installKubectl()
		installStern()
	},
}

func installKubectl() {
	if !isCommandAvailable("kubectl") {
		fmt.Println("kubectl not found. Installing...")
		downloadKubectl()
		if err := moveKubectlToPath(); err != nil {
			fmt.Println("Failed to move kubectl:", err)
			return
		}
		fmt.Println("kubectl installed successfully.")
	} else {
		fmt.Println("kubectl is already installed.")
	}
}

func downloadKubectl() {
	// Fazendo o request para obter a versão estável
	resp, err := http.Get("https://dl.k8s.io/release/stable.txt")
	if err != nil {
		fmt.Println("Erro ao obter versão estável:", err)
	}
	defer resp.Body.Close()

	// Lendo o corpo da resposta
	versionBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
	}

	// Convertendo bytes para string e removendo quebras de linha
	version := strings.TrimSpace(string(versionBytes))

	// Construindo a URL final
	kubectlUrl := fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/amd64/kubectl", version)

	if kubectlUrl == "" {
		fmt.Println("URL do kubectl não encontrada")
		return
	}

	if err := execCommand("curl", "-LO", kubectlUrl); err != nil {
		fmt.Println("Failed to download kubectl:", err)
	}
}

func moveKubectlToPath() error {
	if err := os.Rename("kubectl", "/usr/local/bin/kubectl"); err != nil {
		return fmt.Errorf("failed to move kubectl to /usr/local/bin: %w", err)
	}
	execCommand("chmod", "+x", "/usr/local/bin/kubectl")
	if !isCommandAvailable("kubectl") {
		return fmt.Errorf("kubectl not found after move")
	}
	return nil
}

func installStern() {
	if !isCommandAvailable("stern") {
		fmt.Println("stern not found. Installing...")
		if err := downloadStern(); err != nil {
			fmt.Println("Failed to install stern:", err)
			return
		}
	} else {
		fmt.Println("stern is already installed.")
	}
}

func downloadStern() error {
	stern := "stern_1.30.0_linux_amd64"
	sternURL := "https://github.com/stern/stern/releases/download/v1.30.0/stern_1.30.0_linux_amd64.tar.gz"

	fmt.Println("Downloading Stern...")
	if err := execCommand("curl", "-L", "-o" , stern, sternURL); err != nil {
		return fmt.Errorf("failed to download stern: %w", err)
	}

	fmt.Println("Extracting Stern...")
	if err := execCommand("tar", "zxvf", stern); err != nil {
		return fmt.Errorf("failed to extract Stern: %w", err)
	}

	if err := execCommand("rm", stern); err != nil {
		return fmt.Errorf("failed to remove stern: %w", err)
    }

	fmt.Println("Installing Stern...")
	if err := execCommand("chmod", "+x", "stern"); err != nil {
		return fmt.Errorf("failed to make Stern executable: %w", err)
    }
	if err := execCommand("mv", "stern", "/usr/local/bin/"); err != nil {
		return fmt.Errorf("failed to install Stern: %w", err)
	}

	fmt.Println("Stern installed successfully.")

	return nil
}

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func addToPath(cmd string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("erro ao obter o diretório home do usuário: %w", err)
	}
	homeDir := usr.HomeDir

	files := []string{
		filepath.Join(homeDir, ".bashrc"),
		filepath.Join(homeDir, ".zshrc"),
	}

	lineToAdd := fmt.Sprintf(`export PATH="%s:$PATH"`, cmd)

	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("erro ao abrir o arquivo %s: %w", file, err)
			}

			content, err := os.ReadFile(file)
			if err != nil {
				f.Close()
				return fmt.Errorf("erro ao ler o arquivo %s: %w", file, err)
			}

			if !strings.Contains(string(content), lineToAdd) {
				if _, err := f.WriteString("\n" + lineToAdd + "\n"); err != nil {
					f.Close()
					return fmt.Errorf("erro ao escrever no arquivo %s: %w", file, err)
				}
				fmt.Printf("Linha adicionada ao arquivo %s\n", file)
			} else {
				fmt.Printf("A linha já está presente no arquivo %s\n", file)
			}
			f.Close()
		} else {
			fmt.Printf("Arquivo %s não encontrado\n", file)
		}
	}
	return nil
}
