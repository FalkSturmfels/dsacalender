package root

import (
	"fmt"

	"github.com/FalkSturmfels/dsacalender/pkg/date"
	"github.com/FalkSturmfels/dsacalender/pkg/gui/dateform"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dateform dateform.Model

	//list   list.Model
	//detail detail.Model
	focus FocusedView
}

type FocusedView int

const (
	DateFormView FocusedView = iota
	ListView
	DetailView
)

/*Der Root:

hält alle Sub-Models

entscheidet wer Updates bekommt

kombiniert deren Views

-> Du rufst Update der Sub-Models manuell auf und sammelst deren Cmds.*/

func NewModel() model {
	return model{
		dateform: dateform.NewModel(),
	}
}

func (m model) Init() tea.Cmd {
	fmt.Println(date.CurrentDate.Year)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

/*
Wichtiges Pattern

Root entscheidet wer Update bekommt

Sub-Models bleiben vollständig unabhängig

Commands werden gesammelt und gebatcht
*/
/*func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    var cmds []tea.Cmd

    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {
        case "tab":
            if m.focus == ListView {
                m.focus = DetailView
            } else {
                m.focus = ListView
            }
        }
    }

    // 🔁 Update aktives Sub-Model
    switch m.focus {
    case ListView:
        m.list, cmd = m.list.Update(msg)
        cmds = append(cmds, cmd)

    case DetailView:
        m.detail, cmd = m.detail.Update(msg)
        cmds = append(cmds, cmd)
    }

    return m, tea.Batch(cmds...)
}*/

func (m model) View() string {
	return m.dateform.View()
}

// View: Komposition der Sub-Views
/*func (m RootModel) View() string {
    switch m.focus {
    case ListView:
        return m.list.View()
    case DetailView:
        return m.detail.View()
    default:
        return ""
    }
}*/

/* Oder Split-Layout:
func (m RootModel) View() string {
    return lipgloss.JoinHorizontal(
        lipgloss.Top,
        m.list.View(),
        m.detail.View(),
    )
}*/

/*
Optional: Messages zwischen View-Models

Sehr wichtig für saubere Architektur
Sub-Models sollten keine Referenzen aufeinander haben.

Beispiel: List sendet Selection-Msg
type ItemSelectedMsg struct {
    Item string
}

In List-Update:
return m, func() tea.Msg {
    return ItemSelectedMsg{Item: selected}
}

Im Root-Update:
case ItemSelectedMsg:
    m.detail.SetItem(msg.Item)

-> Root = Event-Broker

Wann mehrere Root-Models sinnvoll sind
Szenario	Lösung
Tabs / Views	Root delegiert
Modal / Dialog	Sub-Model temporär
Wizard / Flow	Root als State-Machine
Komplexe App	Tree aus Root → Sub → Sub
6️⃣ Typische Fehler (vermeiden!)

❌ Sub-Models direkt miteinander koppeln
❌ Root zu viel Logik geben
❌ Commands vergessen zu sammeln
❌ Update nicht vollständig weiterleiten

Mini-Faustregel

Root koordiniert – Sub-Models handeln – Messages verbinden

*/
