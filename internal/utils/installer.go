package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func getInstallPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		// First try Program Files, fallback to AppData
		destPath := filepath.Join("C:\\", "Program Files", "ariaj")
		if err := os.MkdirAll(destPath, 0755); err != nil {
			destPath = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "ariaj")
		}
		return destPath, nil
	case "darwin":
		// macOS uses /usr/local/bin for user-installed programs
		return "/usr/local/bin", nil
	default:
		// Linux and others use /usr/local/bin or $HOME/.local/bin
		if os.Getuid() == 0 {
			return "/usr/local/bin", nil
		}
		return filepath.Join(os.Getenv("HOME"), ".local", "bin"), nil
	}
}

func updateUnixPath(binPath string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Check common shell config files
	shellConfigs := []string{
		filepath.Join(homeDir, ".bashrc"),
		filepath.Join(homeDir, ".zshrc"),
		filepath.Join(homeDir, ".profile"),
	}

	pathEntry := fmt.Sprintf("\nexport PATH=\"%s:$PATH\"\n", binPath)

	for _, config := range shellConfigs {
		if _, err := os.Stat(config); err == nil {
			content, err := os.ReadFile(config)
			if err != nil {
				continue
			}

			if !strings.Contains(string(content), binPath) {
				f, err := os.OpenFile(config, os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					continue
				}
				f.WriteString(pathEntry)
				f.Close()
			}
		}
	}
	return nil
}

func Install() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	destPath, err := getInstallPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(destPath, 0755); err != nil {
		return fmt.Errorf("failed to create install directory: %v", err)
	}

	// Determine binary name based on OS
	binName := "ariaj"
	if runtime.GOOS == "windows" {
		binName = "ariaj.exe"
	}

	destFile := filepath.Join(destPath, binName)

	// Copy the executable
	input, err := os.ReadFile(execPath)
	if err != nil {
		return err
	}

	// Set appropriate permissions for Unix systems
	var fileMode os.FileMode = 0755
	if runtime.GOOS == "windows" {
		fileMode = 0644
	}

	if err := os.WriteFile(destFile, input, fileMode); err != nil {
		return err
	}

	// Update PATH based on OS
	switch runtime.GOOS {
	case "windows":
		currentPath := os.Getenv("PATH")
		if !strings.Contains(currentPath, destPath) {
			cmd := exec.Command("powershell", "-Command",
				fmt.Sprintf(`[Environment]::SetEnvironmentVariable("PATH", "%s;$env:PATH", [EnvironmentVariableTarget]::User)`, destPath))
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to update PATH: %v", err)
			}
		}
	default:
		if err := updateUnixPath(destPath); err != nil {
			fmt.Printf("Warning: Could not automatically update PATH: %v\n", err)
			fmt.Printf("Please add the following line to your shell configuration:\nexport PATH=\"%s:$PATH\"\n", destPath)
		}
	}

	fmt.Printf("Ariaj installed successfully to %s\n", destPath)
	if runtime.GOOS != "windows" {
		fmt.Println("Please restart your terminal or source your shell configuration file")
	}
	return nil
}
