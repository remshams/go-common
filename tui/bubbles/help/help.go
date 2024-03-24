package help

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	help   help.Model
	keyMap *help.KeyMap
}

func New() Model {
	return Model{
		help:   help.New(),
		keyMap: nil,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case SetKeyMapMsg:
		m.keyMap = &msg.keyMap
	case ResetKeyMapMsg:
		m.keyMap = nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, HelpKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, cmd
}

func (m Model) View() string {
	if m.keyMap == nil {
		return ""
	}
	return m.help.View(*m.keyMap)
}
