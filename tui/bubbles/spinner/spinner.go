package spinner

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/remshams/common/tui/styles"
)

type Model struct {
	spinner spinner.Model
	label   string
}

func New() Model {
	spinner := spinner.New(spinner.WithSpinner(spinner.Dot))
	spinner.Style = lipgloss.NewStyle().Foreground(styles.SelectedColor)
	return Model{
		spinner: spinner,
		label:   "",
	}
}

func (m Model) WithLabel(label string) Model {
	m.label = label
	return m
}

func Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	styles := lipgloss.NewStyle().Foreground(styles.SelectedColor)
	return fmt.Sprintf("%s %s", m.spinner.View(), styles.Render(m.label))
}

func (m Model) Tick() tea.Cmd {
	return m.spinner.Tick
}
