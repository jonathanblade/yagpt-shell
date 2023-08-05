package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jonathanblade/yagpt-shell/internal/config"
	"github.com/jonathanblade/yagpt-shell/internal/logger"
)

type UI struct {
	program *tea.Program
}

func New(config *config.Config, logger *logger.Logger) *UI {
	model := newModel(config, logger)
	program := tea.NewProgram(model)
	return &UI{program: program}
}

func (ui *UI) Run(debug bool) {
	if debug {
		f, err := tea.LogToFile("yagpt.log", "")
		if err != nil {
			logger.FatalErr(err)
		}
		defer f.Close()
	}
	if _, err := ui.program.Run(); err != nil {
		logger.FatalErr(err)
	}
}
