package cmd

import (
    "ariaj/internal/commands"
    "ariaj/internal/config"
    "ariaj/internal/llm"
    "ariaj/internal/utils"
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/spf13/cobra"
)

var (
    version string
    commit  string
    date    string
)

func SetVersion(v, c, d string) {
    version = v
    commit = c
    date = d
}

var rootCmd = &cobra.Command{
    Use:     "ariaj [prompt]",
    Short:   "Ariaj is a CLI tool for interacting with LLM",
    Long: `Ariaj is a command-line interface tool that allows you to interact 
with the Large Language Model directly from your terminal.

Available Commands:
  install     Install ariaj CLI globally
  model       Manage Ollama models
  help        Help about any command`,
    Version: version,
    Args:    cobra.MaximumNArgs(1),
    RunE:    func(cmd *cobra.Command, args []string) error {
        cfg, err := config.LoadConfig()
        if (err != nil) {
            return err
        }

        // Start server for interactive mode
        if len(args) == 0 {
            if err := utils.StartOllamaProcess(); err != nil {
                return err
            }
        }

        // If no model is selected, try to get the first available model
        if cfg.SelectedModel == "" {
            models, err := llm.ListAvailableModels()
            if err != nil {
                return fmt.Errorf("no models available: %v", err)
            }
            cfg.SelectedModel = models[0].Name
            if err := config.SaveConfig(cfg); err != nil {
                return err
            }
        }

        fmt.Printf("Using model: %s\n", cfg.SelectedModel)

        if len(args) == 0 {
            // Interactive mode
            for {
                fmt.Print("\nEnter your prompt (or 'exit' to quit): ")
                reader := bufio.NewReader(os.Stdin)
                input, err := reader.ReadString('\n')
                if err != nil {
                    utils.CleanupOllama()
                    return err
                }
                prompt := strings.TrimSpace(input)
                if prompt == "exit" {
                    utils.CleanupOllama()
                    return nil
                }
                if prompt == "" {
                    continue
                }
                
                if err := handlePrompt(prompt, false); err != nil {
                    fmt.Printf("Error: %v\n", err)
                }
            }
        } else {
            // Single prompt mode
            return handlePrompt(args[0], true)
        }
    },
}

func handlePrompt(prompt string, cleanup bool) error {
    fmt.Println()
    
    err := llm.GetLLMStreamingResponse(prompt)
    if err != nil {
        // Don't show error messages for server restarts
        if !strings.Contains(err.Error(), "timeout waiting for Ollama server") {
            fmt.Printf("\nError: %v\n", err)
            fmt.Println("\nTroubleshooting steps:")
            fmt.Println("1. Ensure Ollama is installed")
            fmt.Println("2. Try running 'ollama serve' manually")
            fmt.Println("3. Check if port 11434 is available")
        }
        return err
    }

    fmt.Println()
    
    // Only cleanup in single-prompt mode
    if cleanup {
        utils.CleanupOllama()
    }
    return nil
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(commands.InstallCmd())
    rootCmd.AddCommand(commands.UninstallCmd())
    rootCmd.AddCommand(ModelCmd())
}
