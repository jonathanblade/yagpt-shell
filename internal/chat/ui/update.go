package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jonathanblade/yagpt-shell/internal/chat/api"
	"github.com/jonathanblade/yagpt-shell/internal/chat/domain"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var textinputCmd, spinnerCmd, viewportCmd tea.Cmd

	m.textinput, textinputCmd = m.textinput.Update(msg)
	m.spinner, spinnerCmd = m.spinner.Update(msg)
	m.viewport, viewportCmd = m.viewport.Update(msg)

	cmds := []tea.Cmd{textinputCmd, spinnerCmd, viewportCmd}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 4
		content := m.renderMessages()
		m.viewport.SetContent(content)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			requestText := strings.TrimSpace(m.textinput.Value())
			if strings.ToLower(requestText) == "clear" && !m.thinking {
				m.conversation = make([]domain.Message, 0)
				content := m.renderMessages()
				m.viewport.SetContent(content)
				m.textinput.Reset()
				return m, tea.Batch(cmds...)
			}
			if requestText != "" && !m.thinking {
				m.conversation = append(m.conversation, domain.NewUserMessage(requestText))

				content := m.renderMessages()
				m.viewport.SetContent(content)

				// cmds = append(cmds, m.client.GenerateInstructTextCmd(requestText))
				cmds = append(cmds, m.client.GenerateTextFromChatCmd(m.conversation, requestText))
				cmds = append(cmds, m.client.WaitResponseCmd())

				m.textinput.Reset()
				m.viewport.GotoBottom()
				m.thinking = true
			}
		}
	case api.TextGenerationMsg:
		m.thinking = false
		m.conversation = append(m.conversation, domain.NewAssistantMessage(msg))
		content := m.renderMessages()
		m.viewport.SetContent(content)
		m.viewport.GotoBottom()
	}

	return m, tea.Batch(cmds...)
}
