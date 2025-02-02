package utils

import (
    "os/exec"
    "runtime"
    "strings"
)

func CleanupOllama() {
    if runtime.GOOS == "windows" {
        // Get all Ollama-related processes
        cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq ollama.exe", "/FO", "CSV")
        output, _ := cmd.Output()
        
        // Parse output and kill processes
        for _, line := range strings.Split(string(output), "\n") {
            if strings.Contains(line, "ollama.exe") {
                fields := strings.Split(strings.Trim(line, "\""), "\",\"")
                if len(fields) > 1 {
                    pid := strings.TrimRight(fields[1], "\"")
                    // Force kill the process and its children
                    exec.Command("taskkill", "/F", "/T", "/PID", pid).Run()
                }
            }
        }

        // Additional cleanup for any remaining instances
        exec.Command("taskkill", "/F", "/IM", "ollama.exe").Run()
        exec.Command("taskkill", "/F", "/IM", "ollama_llama_server.exe").Run()
    } else {
        // Unix-like systems - kill all Ollama processes
        exec.Command("pkill", "-9", "-f", "ollama").Run()
    }

    // No sleep delay - immediate termination
}
