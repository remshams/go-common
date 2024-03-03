package table

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
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

type Model[T any] struct {
	Table         table.Model
	createColumns CreateColumnsFunc
	createRows    CreateRowsFunc[T]
	values        T
}

func New[T any](createColumns CreateColumnsFunc, createRows CreateRowsFunc[T]) Model[T] {
	return Model[T]{
		Table:         table.New(table.WithFocused(true)),
		createColumns: createColumns,
		createRows:    createRows,
	}
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
		m.Table, cmd = m.Table.Update(msg)
	}
	return m, cmd
}

func (m Model[T]) View() string {
	return m.Table.View()
}

func (m Model[T]) calculateTableDimensions() (int, int) {
	width := app_store.LayoutStore.Width - 5
	height := app_store.LayoutStore.Height - 8
	if height < 0 {
		height = styles.CalculateDimensionsFromPercentage(80, app_store.LayoutStore.Height, styles.UnlimitedDimension)
	}
	return width, height
}

func (m *Model[T]) refreshTable() {
	width, height := m.calculateTableDimensions()
	m.Table.SetWidth(width)
	m.Table.SetHeight(height)
	m.Table.SetColumns(m.createColumns(width))
	m.Table.SetRows(m.createRows(m.values))
}
