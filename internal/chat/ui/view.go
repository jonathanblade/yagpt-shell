package ui

import (
	"strings"

	"github.com/jonathanblade/yagpt-shell/internal/style"
)

func (m model) View() string {
	var s string
	s += m.viewport.View() + "\n"
	s += style.MutedTextStyle.Render(strings.Repeat("─", m.viewport.Width)) + "\n"
	if !m.thinking {
		s += m.textinput.View() + "\n\n"
	} else {
		s += m.spinner.View() + "\n\n"
	}
	s += style.MutedTextStyle.Render("(press esc or ctrl+c to quit)")
	return s
}
