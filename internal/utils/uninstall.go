package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func getGoPaths() []string {
	paths := []string{}

	// Check GOPATH first
	if goPath := os.Getenv("GOPATH"); goPath != "" {
		paths = append(paths, filepath.Join(goPath, "bin"))
	}

	// Check default Go installation paths
	switch runtime.GOOS {
	case "windows":
		paths = append(paths, filepath.Join(os.Getenv("USERPROFILE"), "go", "bin"))
	case "darwin":
		paths = append(paths,
			filepath.Join(os.Getenv("HOME"), "go", "bin"),
			"/usr/local/go/bin",
		)
	default: // Linux and others
		paths = append(paths,
			filepath.Join(os.Getenv("HOME"), "go", "bin"),
			"/usr/local/go/bin",
		)
	}
	return paths
}

func cleanupWindowsPATH() error {
	cmd := exec.Command("powershell", "-Command", `
		$userPath = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
		$systemPath = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::Machine)
		
		# Clean from User PATH
		$userPath = ($userPath -split ';' | Where-Object { $_ -notmatch 'ariaj' }) -join ';'
		[Environment]::SetEnvironmentVariable("PATH", $userPath, [EnvironmentVariableTarget]::User)
		
		# Clean from System PATH
		$systemPath = ($systemPath -split ';' | Where-Object { $_ -notmatch 'ariaj' }) -join ';'
		[Environment]::SetEnvironmentVariable("PATH", $systemPath, [EnvironmentVariableTarget]::Machine)
	`)
	return cmd.Run()
}

func cleanupUnixPath() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Common shell config files
	configFiles := []string{
		filepath.Join(home, ".bashrc"),
		filepath.Join(home, ".zshrc"),
		filepath.Join(home, ".profile"),
		filepath.Join(home, ".bash_profile"),
	}

	for _, file := range configFiles {
		if _, err := os.Stat(file); err == nil {
			// Read file content
			content, err := os.ReadFile(file)
			if err != nil {
				continue
			}

			// Remove ariaj-related PATH entries
			lines := strings.Split(string(content), "\n")
			newLines := []string{}
			for _, line := range lines {
				if !strings.Contains(line, "ariaj") {
					newLines = append(newLines, line)
				}
			}

			// Write back if changed
			if len(lines) != len(newLines) {
				err = os.WriteFile(file, []byte(strings.Join(newLines, "\n")), 0644)
				if err == nil {
					fmt.Printf("Cleaned PATH from %s\n", file)
				}
			}
		}
	}
	return nil
}

func Uninstall() error {
	fmt.Println("Starting Ariaj uninstallation...")

	// Get all possible installation paths
	possiblePaths := []string{}

	switch runtime.GOOS {
	case "windows":
		possiblePaths = append(possiblePaths,
			filepath.Join("C:\\", "Program Files", "ariaj"),
			filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "ariaj"),
		)
	case "darwin":
		possiblePaths = append(possiblePaths,
			"/usr/local/bin",
			filepath.Join(os.Getenv("HOME"), ".local", "bin"),
		)
	default: // Linux and others
		possiblePaths = append(possiblePaths,
			"/usr/local/bin",
			filepath.Join(os.Getenv("HOME"), ".local", "bin"),
		)
	}

	// Add Go paths
	possiblePaths = append(possiblePaths, getGoPaths()...)

	// Remove binary from all possible locations
	binName := "ariaj"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	for _, path := range possiblePaths {
		binPath := filepath.Join(path, binName)
		if err := os.Remove(binPath); err == nil {
			fmt.Printf("Removed binary from: %s\n", path)
		}

		// Only remove directories we created (not Go or system dirs)
		if !strings.Contains(path, "go") && !strings.Contains(path, "bin") {
			if err := os.Remove(path); err == nil {
				fmt.Printf("Removed directory: %s\n", path)
			}
		}
	}

	// Remove config directory
	configDir := filepath.Join(func() string {
		if runtime.GOOS == "windows" {
			return os.Getenv("USERPROFILE")
		}
		return os.Getenv("HOME")
	}(), ".ariaj")

	if err := os.RemoveAll(configDir); err == nil {
		fmt.Printf("Removed config directory: %s\n", configDir)
	}

	// Cleanup PATH based on platform
	if runtime.GOOS == "windows" {
		if err := cleanupWindowsPATH(); err != nil {
			fmt.Printf("Warning: Failed to clean Windows PATH: %v\n", err)
		}
	} else {
		if err := cleanupUnixPath(); err != nil {
			fmt.Printf("Warning: Failed to clean shell configs: %v\n", err)
		}
	}

	fmt.Println("\nUninstallation complete!")
	fmt.Println("Note: Please restart your terminal for changes to take effect")
	return nil
}
