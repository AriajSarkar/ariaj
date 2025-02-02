package commands

import (
	"github.com/AriajSarkar/ariaj/internal/utils"
	"github.com/spf13/cobra"
)

func UninstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall ariaj CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.Uninstall()
		},
	}
}
