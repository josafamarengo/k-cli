package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install-tools",
	Short: "Install kubectl, krew and stern if not present",
	Run: func(cmd *cobra.Command, args []string) {
		installKubectl()
		installKrew()
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
	if !isCommandAvailable("kubectl") {
		return fmt.Errorf("kubectl not found after move")
	}
	return nil
}

func installKrew() {
	if !isCommandAvailable("kubectl") {
		installKubectl()
	}
	if !isCommandAvailable("kubectl krew") {
		fmt.Println("krew not found. Installing krew...")
		if err := downloadKrew(); err != nil {
			fmt.Println("Failed to install krew:", err)
			return
		}
	}
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

func downloadKrew() error {
	// Criar diretório temporário
	tmpDir, err := os.MkdirTemp("", "krew-install")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Navegar até o diretório temporário
	if err := os.Chdir(tmpDir); err != nil {
		return fmt.Errorf("failed to change to temp directory: %w", err)
	}

	// Detectar o sistema operacional
	osType := strings.ToLower(runtime.GOOS)

	// Detectar a arquitetura
	arch, err := detectArch()
	if err != nil {
		return fmt.Errorf("failed to detect architecture: %w", err)
	}

	// Nome do arquivo krew
	krew := fmt.Sprintf("krew-%s_%s", osType, arch)

	// URL para baixar o krew
	krewURL := fmt.Sprintf("https://github.com/kubernetes-sigs/krew/releases/latest/download/%s.tar.gz", krew)

	// Baixar o arquivo krew.tar.gz
	fmt.Println("Downloading Krew...")
	if err := execCommand("curl", "-fsSLO", krewURL); err != nil {
		return fmt.Errorf("failed to download krew: %w", err)
	}

	// Extrair o arquivo tar.gz
	fmt.Println("Extracting Krew...")
	if err := execCommand("tar", "zxvf", fmt.Sprintf("%s.tar.gz", krew)); err != nil {
		return fmt.Errorf("failed to extract krew: %w", err)
	}

	// Instalar o Krew
	fmt.Println("Installing Krew...")
	if err := execCommand(fmt.Sprintf("./%s", krew), "install", "krew"); err != nil {
		return fmt.Errorf("failed to install krew: %w", err)
	}

	fmt.Println("Krew installed successfully.")
	return nil
}

func downloadStern() error {
	if err := execCommand("kubectl", "krew", "install", "stern"); err != nil {
		return fmt.Errorf("failed to install stern: %w", err)
	}
	fmt.Println("stern installed successfully.")
	return nil
}

func detectArch() (string, error) {
	arch := runtime.GOARCH

	switch arch {
	case "amd64":
		return "amd64", nil
	case "arm64":
		return "arm64", nil
	case "386":
		return "386", nil
	default:
		return "", fmt.Errorf("unsupported architecture: %s", arch)
	}
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
