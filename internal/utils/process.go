package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

var (
	ollamaCmd *exec.Cmd
	cmdMutex  sync.Mutex
	lastStop  time.Time
)

func StartOllamaProcess() error {
	cmdMutex.Lock()
	defer cmdMutex.Unlock()

	// If we recently stopped, give it a moment
	if time.Since(lastStop) < time.Second {
		time.Sleep(time.Second)
	}

	// Reset state
	ollamaCmd = nil

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", "/b", "ollama", "serve")
	} else {
		cmd = exec.Command("ollama", "serve")
	}

	// Suppress output
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start Ollama: %v", err)
	}

	ollamaCmd = cmd
	return nil
}

func StopOllamaProcess() {
	cmdMutex.Lock()
	defer cmdMutex.Unlock()

	if ollamaCmd != nil && ollamaCmd.Process != nil {
		if runtime.GOOS == "windows" {
			exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprint(ollamaCmd.Process.Pid)).Run()
		} else {
			ollamaCmd.Process.Kill()
		}
		ollamaCmd.Wait()
		ollamaCmd = nil
	}
	
	lastStop = time.Now()
}
