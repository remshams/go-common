package table

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/remshams/common/tui/styles"
	app_store "github.com/remshams/jira-control/tui/store"
)

type TableDataUpdatedAction[T any] struct {
	values T
}

func CreateTableDataUpdatedAction[T any](values T) tea.Cmd {
	return func() tea.Msg {
		return TableDataUpdatedAction[T]{
			values: values,
		}
	}
}

type CreateColumnsFunc func(width int) []table.Column
type CreateRowsFunc[T any] func(values T) []table.Row
type KeyMap = table.KeyMap
type Column = table.Column
type Row = table.Row

var DefaultKeyMap = table.DefaultKeyMap()

var DefaultKeyBindings = []key.Binding{
	DefaultKeyMap.LineUp,
	DefaultKeyMap.LineDown,
	DefaultKeyMap.PageUp,
	DefaultKeyMap.PageDown,
	DefaultKeyMap.HalfPageUp,
	DefaultKeyMap.HalfPageDown,
	DefaultKeyMap.GotoTop,
	DefaultKeyMap.GotoBottom,
}

type Model[T any] struct {
	table         table.Model
	createColumns CreateColumnsFunc
	createRows    CreateRowsFunc[T]
	values        T
	widthOffset   int
	heightOffset  int
	noDataMessage string
}

func New[T any](createColumns CreateColumnsFunc, createRows CreateRowsFunc[T], widthOffset int, heightOffset int) Model[T] {
	return Model[T]{
		table:         table.New(table.WithFocused(true)),
		createColumns: createColumns,
		createRows:    createRows,
		widthOffset:   widthOffset,
		heightOffset:  heightOffset,
	}
}

func (m Model[T]) WithNotDataMessage(message string) Model[T] {
	m.noDataMessage = message
	return m
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m Model[T]) Update(msg tea.Msg) (Model[T], tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.WindowSizeMsg:
		m.refreshTable()
	case TableDataUpdatedAction[T]:
		m.values = msg.(TableDataUpdatedAction[T]).values
		m.refreshTable()
	default:
		m.table, cmd = m.table.Update(msg)
	}
	return m, cmd
}

func (m Model[T]) View() string {
	if m.IsEmpty() {
		style := lipgloss.NewStyle().
			Foreground(styles.SelectedColor).
			Width(app_store.LayoutStore.Width).
			Align(lipgloss.Center)
		return style.Render(m.noDataMessage)
	}
	return m.table.View()
}

func (m Model[T]) calculateTableDimensions() (int, int) {
	width := app_store.LayoutStore.Width - m.widthOffset
	height := app_store.LayoutStore.Height - m.heightOffset
	if height < 0 {
		height = styles.CalculateDimensionsFromPercentage(80, app_store.LayoutStore.Height, styles.UnlimitedDimension)
	}
	return width, height
}

func (m *Model[T]) refreshTable() {
	width, height := m.calculateTableDimensions()
	m.table.SetWidth(width)
	m.table.SetHeight(height)
	m.table.SetColumns(m.createColumns(width))
	m.table.SetRows(m.createRows(m.values))
	m.table.GotoTop()
}

func (m Model[T]) IsEmpty() bool {
	return len(m.table.Rows()) == 0
}

func (m Model[T]) SelectedRowCell(column int) string {
	return m.table.SelectedRow()[column]
}
