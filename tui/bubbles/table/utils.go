package table_utils

import (
	"math"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func CreateTable(columns []table.Column, rows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

func TableKeyBindings() []key.Binding {
	tableKeys := table.DefaultKeyMap()
	return []key.Binding{
		tableKeys.LineUp,
		tableKeys.LineDown,
		tableKeys.PageUp,
		tableKeys.PageDown,
		tableKeys.HalfPageUp,
		tableKeys.HalfPageDown,
		tableKeys.GotoTop,
		tableKeys.GotoBottom,
	}
}

func ColumnWidthFromPercent(percent int, totalWidth int) int {
	return int(math.Round(float64(totalWidth) * (float64(percent) / 100)))
}
