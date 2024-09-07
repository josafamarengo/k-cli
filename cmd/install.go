package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install-tools",
	Short: "Install kubectl, stern, k9s and Rancher CLI if not present",
	Run: func(cmd *cobra.Command, args []string) {
		installKubectl()
		installStern()
		installK9s()
		installRancherCLI()
	},
}

func installKubectl() {
	if !isCommandAvailable("kubectl") {
		fmt.Println("kubectl not found. Installing...")
		downloadKubectl()
	} else {
		fmt.Println("kubectl is already installed.")
	}
}

func installStern() {
	if !isCommandAvailable("stern") {
		fmt.Println("stern not found. Installing...")
		downloadStern()
	} else {
		fmt.Println("stern is already installed.")
	}
}

func installK9s() {
	if !isCommandAvailable("k9s") {
		fmt.Println("k9s not found. Installing...")
		downloadK9s()
	} else {
		fmt.Println("k9s is already installed.")
	}
}

func installRancherCLI() {
	if !isCommandAvailable("rancher") {
		fmt.Println("Rancher CLI not found. Installing...")
		downloadRancherCLI()
	} else {
		fmt.Println("Rancher CLI is already installed.")
	}
}

func isCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func downloadKubectl() {
	url := ""
	switch runtime.GOOS {
	case "linux":
		url = "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
	case "darwin":
		url = "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl"
	case "windows":
		url = "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/windows/amd64/kubectl.exe"
	}
	installTool(url, "kubectl")
}

func downloadStern() {
	url := ""
	switch runtime.GOOS {
	case "linux":
		url = "https://github.com/stern/stern/releases/latest/download/stern_linux_amd64"
	case "darwin":
		url = "https://github.com/stern/stern/releases/latest/download/stern_darwin_amd64"
	case "windows":
		url = "https://github.com/stern/stern/releases/latest/download/stern_windows_amd64.exe"
	}
	installTool(url, "stern")
}

func downloadK9s() {
	url := ""
	switch runtime.GOOS {
	case "linux":
		url = "https://github.com/derailed/k9s/releases/latest/download/k9s_Linux_x86_64.tar.gz"
	case "darwin":
		url = "https://github.com/derailed/k9s/releases/latest/download/k9s_Darwin_x86_64.tar.gz"
	case "windows":
		url = "https://github.com/derailed/k9s/releases/latest/download/k9s_Windows_x86_64.tar.gz"
	}
	installTool(url, "k9s")
}

func downloadRancherCLI() {
	url := ""
	switch runtime.GOOS {
	case "linux":
		url = "https://github.com/rancher/cli/releases/latest/download/rancher-linux-amd64.tar.gz"
	case "darwin":
		url = "https://github.com/rancher/cli/releases/latest/download/rancher-darwin-amd64.tar.gz"
	case "windows":
		url = "https://github.com/rancher/cli/releases/latest/download/rancher-windows-amd64.zip"
	}
	installTool(url, "rancher")
}

func installTool(url string, name string) {
	fmt.Printf("Downloading %s from %s...\n", name, url)

	cmd := exec.Command("curl", "-LO", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to download %s: %v\n", name, err)
		return
	}

	if runtime.GOOS != "windows" {
		cmd = exec.Command("chmod", "+x", name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Failed to make %s executable: %v\n", name, err)
		}
	}
	fmt.Printf("%s installed successfully.\n", name)
}

func init() {
	rootCmd.AddCommand(installCmd)
}
