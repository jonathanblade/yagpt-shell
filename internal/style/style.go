package style

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/lipgloss"
)

const (
	Black  = lipgloss.Color("#000000")
	Gray   = lipgloss.Color("#989898")
	Pink   = lipgloss.Color("#ff69b7")
	Purple = lipgloss.Color("#bd93f9")
	Red    = lipgloss.Color("#ff0000")
	Yellow = lipgloss.Color("#ffcc00")
	White  = lipgloss.Color("#ffffff")
)

var (
	BorderStyle           = lipgloss.NewStyle().Padding(1, 2, 0, 2).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(Gray)
	AccentTextStyle       = lipgloss.NewStyle().Foreground(Yellow)
	MutedTextStyle        = lipgloss.NewStyle().Foreground(Gray)
	UserMessageStyle      = lipgloss.NewStyle().Foreground(Pink).Bold(true)
	AssistantMessageStyle = lipgloss.NewStyle().Foreground(Yellow).Bold(true)
)

func markdownStyle() ansi.StyleConfig {
	margin := uint(0)
	styleConfig := glamour.NoTTYStyleConfig
	styleConfig.Document.Margin = &margin
	styleConfig.Document.BlockPrefix = ""
	styleConfig.Document.BlockSuffix = ""
	return styleConfig
}

var MarkdownStyle = markdownStyle()
