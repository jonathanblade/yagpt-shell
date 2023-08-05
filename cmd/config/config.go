package config

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage config",
		Long:  "Manage yagpt config.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// Setup flags
	cmd.Flags().BoolP("help", "h", false, "Print help")
	// Setup commands
	cmd.AddCommand(newShowCommand())
	return cmd
}
