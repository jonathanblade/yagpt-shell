package config

import (
	"github.com/jonathanblade/yagpt-shell/internal/config"
	"github.com/jonathanblade/yagpt-shell/internal/help"
	"github.com/spf13/cobra"
)

func newShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show config",
		Long:  "Show yagpt config.",
		Args:  help.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				return
			}
			config := config.Read()
			config.Show()
		},
	}
	// Setup flags
	cmd.Flags().BoolP("help", "h", false, "Print help")
	return cmd
}
