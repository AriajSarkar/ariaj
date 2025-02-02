package llm

import (
	"ariaj/internal/config"
	"ariaj/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/tmc/langchaingo/llms/ollama"
)

func isOllamaRunning() bool {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get("http://127.0.0.1:11434/api/version")
	return err == nil && resp.StatusCode == http.StatusOK
}

func startOllamaServer() error {
    // Try multiple times to start server
    maxAttempts := 3
    for attempt := 0; attempt < maxAttempts; attempt++ {
        if isOllamaRunning() {
            return nil
        }

        if err := utils.StartOllamaProcess(); err != nil {
            continue // Try again
        }

        // Wait for server to be available
        for i := 0; i < 15; i++ {
            if isOllamaRunning() {
                return nil
            }
            time.Sleep(500 * time.Millisecond)
        }
    }

    return fmt.Errorf("timeout waiting for Ollama server to start")
}

func GetLLMResponse(prompt string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	if err := startOllamaServer(); err != nil {
		return "", fmt.Errorf("failed to start Ollama server: %v\nPlease ensure Ollama is installed and run 'ollama serve' manually", err)
	}

	client, err := ollama.New(ollama.WithModel(cfg.SelectedModel))
	if err != nil {
		return "", fmt.Errorf("failed to create Ollama client: %v", err)
	}

	ctx := context.Background()
	completion, err := client.Call(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("failed to get response: %v", err)
	}

	return completion, nil
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func GetLLMStreamingResponse(prompt string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if err := startOllamaServer(); err != nil {
		return fmt.Errorf("failed to start Ollama server: %v", err)
	}

	reqBody := OllamaRequest{
		Model:  cfg.SelectedModel,
		Prompt: prompt,
		Stream: true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var response OllamaResponse
		if err := decoder.Decode(&response); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode response: %v", err)
		}

		fmt.Print(response.Response)
		if response.Done {
			break
		}
	}

	return nil
}
