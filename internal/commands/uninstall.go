package commands

import (
	"fmt"

	"github.com/AriajSarkar/ariaj/internal/utils"
	"github.com/spf13/cobra"
)

func UninstallCmd() *cobra.Command {
	var forceUninstall bool

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall ariaj CLI",
		Long: `Uninstall ariaj CLI completely from your system including:
- Binary files
- Configuration directory
- PATH environment entries`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !forceUninstall {
				fmt.Print("Are you sure you want to uninstall ariaj? [y/N]: ")
				var response string
				fmt.Scanln(&response)
				if response != "y" && response != "Y" {
					fmt.Println("Uninstallation cancelled")
					return nil
				}
			}
			return utils.Uninstall()
		},
	}

	cmd.Flags().BoolVarP(&forceUninstall, "force", "f", false, "Force uninstall without confirmation")
	return cmd
}
