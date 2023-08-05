package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/jonathanblade/yagpt-shell/internal/chat/api"
	"github.com/jonathanblade/yagpt-shell/internal/chat/domain"
	"github.com/jonathanblade/yagpt-shell/internal/config"
	"github.com/jonathanblade/yagpt-shell/internal/logger"
	"github.com/jonathanblade/yagpt-shell/internal/style"
)

type model struct {
	client       *api.Client
	logger       *logger.Logger
	thinking     bool
	conversation []domain.Message
	textinput    textinput.Model
	spinner      spinner.Model
	viewport     viewport.Model
}

func newModel(config *config.Config, logger *logger.Logger) model {
	ti := textinput.New()
	ti.Placeholder = "Type your request here"
	ti.Prompt = ""
	ti.Focus()

	vp := viewport.New(0, 0)

	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(style.Yellow)

	apiClient := api.NewClient(config, logger)

	return model{
		client:       apiClient,
		logger:       logger,
		thinking:     false,
		conversation: make([]domain.Message, 0),
		textinput:    ti,
		viewport:     vp,
		spinner:      s,
	}
}

func (m model) renderMessages() string {
	var renderedMessages []string

	user := style.UserMessageStyle.Render("You")
	assistant := style.AssistantMessageStyle.Render("YaGPT")

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(style.MarkdownStyle),
		glamour.WithWordWrap(m.viewport.Width),
	)
	if err != nil {
		logger.FatalErr(err)
	}

	for _, message := range m.conversation {
		var author, spaces string
		switch message.Role {
		case domain.UserRole:
			author = user
			// hh:mm:ss + You = 11 symbols
			spaces = strings.Repeat(" ", m.viewport.Width-11)
		case domain.AssistantRole:
			author = assistant
			// hh:mm:ss + YaGPT = 13 symbols
			spaces = strings.Repeat(" ", m.viewport.Width-13)
		default:
			continue
		}
		text, err := renderer.Render(message.Text)
		if err != nil {
			logger.FatalErr(err)
		}
		time := style.MutedTextStyle.Render(message.Time.Format("15:04:05"))
		renderedMessage := fmt.Sprintf("%s%s%s\n%s", author, spaces, time, text)
		renderedMessages = append(renderedMessages, renderedMessage)
	}
	return strings.Join(renderedMessages, "\n")
}
