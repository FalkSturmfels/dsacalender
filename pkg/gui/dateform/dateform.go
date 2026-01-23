package dateform

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/FalkSturmfels/dsacalender/data"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var monthId int
var dayId int
var year string
var submit bool

type Model struct {
	form *huh.Form
}

func NewModel() Model {
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Wähle einen Monat:").
					Options(createMonthOptions()...,
					).Value(&monthId),
				huh.NewSelect[int]().
					Value(&dayId).
					Height(20).
					Title("Wähle ein Tag:").
					OptionsFunc(func() []huh.Option[int] {
						return createDaysOptions(monthId)
					}, &monthId),
				huh.NewInput().
					Title("Gib ein Jahr ein: ").
					Description("Das Jahr muss größer oder gleich 3138 sein").
					Placeholder("3138").
					Suggestions([]string{"3138", "3139", "3140"}).
					Value(&year).
					Validate(func(str string) error {
						if y, err := strconv.Atoi(str); err != nil {
							return errors.New("Das Jahr muss eine Zahl größer als 3138 sein.")
						} else if y < 3138 {
							return errors.New("Das Jahr muss 3138 oder größer sein.")
						}
						return nil
					}),
				huh.NewConfirm().
					Title("Datum setzen?").
					Affirmative("Ja").
					Negative("Abbrechen").
					Value(&submit),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m Model) View() string {
	if m.form.State == huh.StateCompleted {
		return fmt.Sprintf("Ausgewähltes Datum: %d %d %s", dayId, monthId, year)
	}
	return m.form.View()
}

func createMonthOptions() []huh.Option[int] {
	var options []huh.Option[int]
	for _, mon := range data.Months {
		options = append(options, huh.NewOption(mon.Month, mon.Id))
	}
	return options
}

func createDaysOptions(monthId int) []huh.Option[int] {
	var options []huh.Option[int]
	month := data.MonthMap[monthId]

	for _, day := range month.Days {
		options = append(options, huh.NewOption(day.Name, day.Id))
	}
	return options
}
