package help

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Help key.Binding
}

var HelpKeys = KeyMap{
	Help: key.NewBinding(
		key.WithKeys("h", "help"),
		key.WithHelp("h", "toggle help"),
	),
}

type SetKeyMapMsg struct {
	keyMap help.KeyMap
}

type ToggleFullHelp struct{}

func CreateSetKeyMapMsg(keyMap help.KeyMap) tea.Cmd {
	return func() tea.Msg {
		return SetKeyMapMsg{
			keyMap: keyMap,
		}
	}
}

type ResetKeyMapMsg struct{}

func CreateResetKeyMapMsg() tea.Cmd {
	return func() tea.Msg {
		return ResetKeyMapMsg{}
	}
}

func CreateToggleFullHelpMsg() tea.Cmd {
	return func() tea.Msg {
		return ToggleFullHelp{}
	}
}
