package cmd

import (
    "ariaj/internal/config"
    "ariaj/internal/llm"
    "fmt"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
)

func ModelCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "model",
        Short: "Change or view the current model",
        RunE: func(cmd *cobra.Command, args []string) error {
            models, err := llm.ListAvailableModels()
            if err != nil {
                return fmt.Errorf("no models available - use 'ollama pull' to download a model first")
            }

            var items []string
            for _, model := range models {
                items = append(items, model.Name)
            }

            prompt := promptui.Select{
                Label: "Select Model",
                Items: items,
            }

            _, result, err := prompt.Run()
            if err != nil {
                return fmt.Errorf("prompt failed: %v", err)
            }

            cfg, err := config.LoadConfig()
            if (err != nil) {
                return err
            }

            cfg.SelectedModel = result
            if err := config.SaveConfig(cfg); err != nil {
                return err
            }

            fmt.Printf("Model changed to: %s\n", result)
            return nil
        },
    }
}
