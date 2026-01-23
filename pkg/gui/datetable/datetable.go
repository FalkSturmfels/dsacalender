package datetable

import (
	"github.com/FalkSturmfels/dsacalender/data"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type cell struct {
	day   data.Day
	width int
}

type row []cell

type Model struct {
	month    data.Month
	dayId    int
	KeyMap   KeyMap
	Help     help.Model
	rowIndex int // 0..n
	colIndex int // 0..n
	maxCols  int // len(cols) - 1
	maxRows  int // len(rows) - 1

	firstDay  data.Day
	secondDay data.Day
	lastDay   data.Day

	rows  []row
	focus bool
}

type MonthChangeMsg struct {
	monthId int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func New() Model {
	m := Model{}

	return m
}

func (m Model) MoveUp() {
	if m.rowIndex > 0 {
		m.rowIndex--
	} else {
		m.rowIndex = m.maxRows
	}
}

func (m Model) MoveDown() {
	if m.rowIndex < m.maxRows {
		m.rowIndex++
	} else {
		m.rowIndex = m.maxRows
	}
}

func (m Model) MoveRight() {
	if m.colIndex < m.maxCols {
		m.colIndex++
	} else {
		m.colIndex = 0
	}
}
func (m Model) MoveLeft() {
	if m.colIndex > 0 {
		m.colIndex--
	} else {
		m.colIndex = m.maxCols
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case MonthChangeMsg:
		m.updateMonth(msg.monthId)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.RowUp):
			m.MoveUp()
		case key.Matches(msg, m.KeyMap.RowDown):
			m.MoveDown()
		case key.Matches(msg, m.KeyMap.ColRight):
			m.MoveRight()
		case key.Matches(msg, m.KeyMap.ColLeft):
			m.MoveLeft()
		}
	}

	return m, nil
}

func (m *Model) updateMonth(id int) {
	m.month = data.GetMonth(id)

	m.firstDay = m.month.FirstDay()
	m.lastDay = m.month.LastDay()

	if m.month.Id == 1 {
		m.secondDay = m.month.SecondDay()
	}

	m.maxCols = 7
	if m.month.IsEquinoxMonth() {
		m.maxCols = 8
	}

	startDays := m.month.WeekStartDays()

	for index := range startDays {
		cells := make([]cell, m.maxCols)

		var daysSlice []data.Day
		if index < len(startDays)-2 {
			// [0:1], [1:2], [2:3], [3:]
			daysSlice = m.month.Days[startDays[index]:startDays[index+1]]
		} else {
			daysSlice = m.month.Days[startDays[index]:]
		}

		for _, day := range daysSlice {
			cells = append(cells, cell{day: day})
		}
		if len(cells) < m.maxCols {
			cells = append(cells, cell{})
		}
	}
}

func (m Model) View() string {
	return createGrid(m)
}

// KeyMap defines keybindings. It satisfies to the help.KeyMap interface, which
// is used to render the help menu.
type KeyMap struct {
	RowUp    key.Binding
	RowDown  key.Binding
	ColRight key.Binding
	ColLeft  key.Binding
	Select   key.Binding
}

// ShortHelp implements the KeyMap interface.
func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.RowUp, km.RowDown, km.ColRight, km.ColLeft}
}

// FullHelp implements the KeyMap interface.
func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.RowUp, km.RowDown, km.ColRight, km.ColLeft},
	}
}

func (m Model) HelpView() string {
	return m.Help.View(m.KeyMap)
}

// DefaultKeyMap returns a default set of keybindings.
func DefaultKeyMap() KeyMap {
	const spacebar = " "
	return KeyMap{
		RowUp: key.NewBinding(
			key.WithKeys("up", "w"),
			key.WithHelp("↑/w", "up"),
		),
		RowDown: key.NewBinding(
			key.WithKeys("down", "s"),
			key.WithHelp("↓/s", "down"),
		),
		ColRight: key.NewBinding(
			key.WithKeys("right", "d"),
			key.WithHelp("→/d", "right"),
		),
		ColLeft: key.NewBinding(
			key.WithKeys("left", "a"),
			key.WithHelp("←/a", "left"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
	}
}

var normalStyle lipgloss.Style = lipgloss.NewStyle().Width(10).Height(1)
var selectedStyle lipgloss.Style = normalStyle.Foreground(lipgloss.Color("212"))

func createGrid(model Model) string {
	rowStrings := []string{}

	for rowIndex, cells := range model.rows {
		joinedRow := createRow(model, rowIndex, cells)
		rowStrings = append(rowStrings, joinedRow)
	}

	return lipgloss.JoinVertical(lipgloss.Top, rowStrings...)
}

func createRow(model Model, i int, cells []cell) string {
	cellStrings := []string{}
	for j, cell := range cells {
		var style lipgloss.Style
		if model.rowIndex == i && model.colIndex == j {
			style = selectedStyle
		} else {
			style = normalStyle
		}
		cellStrings = append(cellStrings, style.Render(cell.day.Name))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, cellStrings...)
}
