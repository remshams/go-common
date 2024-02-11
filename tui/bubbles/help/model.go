package help

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

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
