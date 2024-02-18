package tabs

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/remshams/common/tui/styles"
)

type SelectTabAction struct {
	Index int
}

func CreateSelectTabAction(index int) tea.Cmd {
	return func() tea.Msg {
		return SelectTabAction{
			Index: index,
		}
	}
}

type TabKeyMap struct {
	Tab key.Binding
}

var TabKeys = TabKeyMap{
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "Switch tab"),
	),
}

type TabSelectedMsg = int

type Model struct {
	names        []string
	active       int
	tabSeparator lipgloss.Style
	tabActive    lipgloss.Style
	tabInactive  lipgloss.Style
}

func New(names []string) Model {
	return Model{
		names,
		0,
		lipgloss.NewStyle().Padding(0, 1).Foreground(styles.DisabledColor),
		lipgloss.NewStyle().Underline(true).Foreground(styles.SelectedColor),
		lipgloss.NewStyle(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case SelectTabAction:
		m.active = msg.Index
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, TabKeys.Tab):
			cmd = m.navigate()
		}
	}
	return m, cmd
}

func (m *Model) navigate() tea.Cmd {
	m.active++
	if m.active >= len(m.names) {
		m.active = 0
	}
	return m.createTabSelectedMsg()
}

func (m Model) View() string {
	tabs := []string{}
	for i, name := range m.names {
		tabs = append(tabs, m.renderTab(name, i == m.active, i == 0))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (m Model) renderTab(
	name string,
	active bool,
	isFirst bool) string {
	startSeparator := ""
	if isFirst {
		startSeparator = m.tabSeparator.Render("|")
	}
	endSeparator := m.tabSeparator.Render("|")
	styledName := m.tabInactive.Render(name)
	if active {
		styledName = m.tabActive.Render(name)
	}
	return fmt.Sprintf("%s%s%s", startSeparator, styledName, endSeparator)

}

func (m Model) createTabSelectedMsg() tea.Cmd {
	return func() tea.Msg {
		return TabSelectedMsg(m.active)
	}
}
