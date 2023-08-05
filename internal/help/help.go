package help

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jonathanblade/yagpt-shell/internal/style"
	"github.com/spf13/cobra"
)

const versionTemplate = `{{with .Name}}{{printf "%%s" .}}{{end}} {{printf "%%s" .Version}}
`

func VersionTemplate() string {
	return fmt.Sprintf(versionTemplate)
}

func HelpFunc(cmd *cobra.Command, s []string) {
	helpText := strings.Builder{}

	// Description
	longDesc := lipgloss.NewStyle().Align(lipgloss.Center).Render(cmd.Long)
	helpText.WriteString(longDesc + "\n\n")

	// Usage
	if cmd.Runnable() {
		usage := style.AccentTextStyle.Render("Usage: ")
		if cmd.HasAvailableSubCommands() {
			usage += fmt.Sprintf("%s [command] [flag]", cmd.CommandPath())
		} else {
			usage += fmt.Sprintf("%s [flag]", cmd.CommandPath())
		}
		helpText.WriteString(usage + "\n\n")
	}

	// Commands
	if cmd.HasAvailableSubCommands() {
		subCmds := cmd.Commands()
		subTitle := style.AccentTextStyle.Render("Commands:")
		subs := ""
		for i := range subCmds {
			if subCmds[i].IsAvailableCommand() {
				subs += lipgloss.NewStyle().PaddingLeft(2).Render(subCmds[i].Name()) +
					lipgloss.NewStyle().
						PaddingLeft(subCmds[i].NamePadding()-len(subCmds[i].Name())+1).
						Render(subCmds[i].Short) + "\n"
			}
		}
		helpText.WriteString(lipgloss.JoinVertical(lipgloss.Left, subTitle, subs) + "\n")
	}

	// Flags
	if cmd.HasAvailableLocalFlags() {
		flags := style.AccentTextStyle.Render("Flags:") + "\n" + cmd.LocalFlags().FlagUsages()
		helpText.WriteString(flags)
	}

	fmt.Println(style.BorderStyle.Render(helpText.String()))
}

func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return cmd.Help()
	}
	return nil
}

func FlagErrorFunc(cmd *cobra.Command, err error) error {
	return cmd.Help()
}
