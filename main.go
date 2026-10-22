package main

import (
	"fmt"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var baseStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("100"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(tea.Printf("Selected %s!", m.table.SelectedRow()[0]))
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func (m model) View() tea.View {
	return tea.NewView(baseStyle.Render(m.table.View()) + "\n " + m.table.HelpView() + "\n")
}

func main() {
	columns := []table.Column{
		{Title: "Country clock", Width: 42},
	}

	rows := []table.Row{
		{"🇦🇷 Argentina"},
		{"🇧🇷 Brasil"},
		{"🇷🇺 Russia"},
		{"🇪🇸 Espanha"},
		{"🇨🇳 China"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
		table.WithWidth(42),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("30")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("16")).
		BorderBottom(true).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

//PYONGYANG
//BUEINOS AIRES
//MADRI
//MOSCOU
