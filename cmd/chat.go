package cmd

import (
	"github.com/jonathanblade/yagpt-shell/internal/chat"
	"github.com/jonathanblade/yagpt-shell/internal/help"
	"github.com/jonathanblade/yagpt-shell/internal/logger"
	"github.com/spf13/cobra"
)

func newChatCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "Start new conversation",
		Long:  "Start new conversation.",
		Args:  help.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				return
			}
			debug, err := cmd.PersistentFlags().GetBool("debug")
			if err != nil {
				logger.FatalErr(err)
			}
			chat.Run(debug)
		},
	}
	// Setup flags
	cmd.PersistentFlags().BoolP("debug", "", false, "Start in debug mode")
	cmd.Flags().BoolP("help", "h", false, "Print help")
	return cmd
}
