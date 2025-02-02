package commands

import (
    "ariaj/internal/utils"
    "github.com/spf13/cobra"
)

func InstallCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "install",
        Short: "Install ariaj CLI globally",
        Long:  `Install ariaj CLI globally on your system for easy access from anywhere.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return utils.Install()
        },
    }
}
