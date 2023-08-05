package cmd

import (
	"github.com/jonathanblade/yagpt-shell/cmd/config"
	"github.com/jonathanblade/yagpt-shell/internal/help"
	"github.com/jonathanblade/yagpt-shell/internal/logger"
	"github.com/spf13/cobra"
)

func newRootCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "yagpt",
		Short:   "YaGPT Shell",
		Long:    "YaGPT Shell is a text-based user interface (TUI) that uses the Yandex GPT API.\nMade with Cobra, Viper, Bubble Tea and 💖.",
		Version: version,
		Args:    cobra.MaximumNArgs(1),
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// Setup flags
	cmd.Flags().BoolP("help", "h", false, "Print help")
	cmd.Flags().BoolP("version", "v", false, "Print version")
	// Setup funcs
	cmd.SetFlagErrorFunc(help.FlagErrorFunc)
	cmd.SetHelpFunc(help.HelpFunc)
	// Setup templates
	cmd.SetVersionTemplate(help.VersionTemplate())
	// Setup commands
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(newChatCommand())
	cmd.AddCommand(config.NewCommand())
	return cmd
}

func Execute(version string) {
	if err := newRootCommand(version).Execute(); err != nil {
		logger.FatalErr(err)
	}
}
