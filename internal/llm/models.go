package llm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type Model struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Digest  string `json:"digest"`
}

type ModelsResponse struct {
	Models []Model `json:"models"`
}

func modelExists(modelName string) bool {
	cmd := exec.Command("ollama", "show", modelName)
	err := cmd.Run()
	return err == nil
}

func ListAvailableModels() ([]Model, error) {
	if err := startOllamaServer(); err != nil {
		return nil, err
	}

	resp, err := http.Get("http://localhost:11434/api/tags")
	if err != nil {
		return nil, fmt.Errorf("failed to get models: %v", err)
	}
	defer resp.Body.Close()

	var modelsResp ModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&modelsResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Filter out models that don't actually exist
	var validModels []Model
	for _, model := range modelsResp.Models {
		if modelExists(model.Name) {
			validModels = append(validModels, model)
		}
	}

	if len(validModels) == 0 {
		return nil, fmt.Errorf("no models found. Please pull a model using 'ollama pull <model-name>'")
	}

	return validModels, nil
}
