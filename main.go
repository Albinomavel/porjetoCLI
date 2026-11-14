package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var baseStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("100"))

type appState int

const (
	stateMenu appState = iota
	stateResult
)

type horarioMsg struct {
	conteudo string
	err      error
}
type timeResponse struct {
	Datetime string `json:"datetime"`
}

type model struct {
	state           appState
	table           table.Model
	paisSelecionado string
}

var urls = map[string]string{
	"🇦🇷 Argentina": "https://time.now/developer/api/timezone/America/Argentina/Buenos_Aires",
	"🇧🇷 Brasil":    "https://time.now/developer/api/timezone/America/Sao_Paulo",
	"🇷🇺 Russia":    "https://time.now/developer/api/timezone/Europe/Moscow",
	"🇪🇸 Espanha":   "https://time.now/developer/api/timezone/Europe/Madrid",
	"🇨🇳 China":     "https://time.now/developer/api/timezone/Asia/Shanghai",
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			if m.state == stateResult {
				m.state = stateMenu
				m.table = tableMenu()
				return m, nil
			}
			return m, tea.Quit
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.state == stateMenu {
				m.paisSelecionado = m.table.SelectedRow()[0]
				urlAPI := urls[m.paisSelecionado]
				m.table = tableResult(m.paisSelecionado, "⏳ Buscando horário...")
				m.state = stateResult
				return m, buscaHorario(urlAPI)
			}
		case "space":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		}
	case horarioMsg:
		if msg.err != nil {
			m.table = tableResult("Erro", "Falha ao uscar o horário")
			return m, nil
		}
		m.table = tableResult(m.paisSelecionado, msg.conteudo)
		return m, nil
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func (m model) View() tea.View {
	return tea.NewView(baseStyle.Render(m.table.View()) + "\n " + m.table.HelpView() + "\n")
}

func tableMenu() table.Model {

	columns := []table.Column{
		{Title: "horário mundial", Width: 42},
	}

	rows := []table.Row{
		{"🇦🇷 Argentina"},
		{"🇧🇷 Brasil"},
		{"🇷🇺 Russia"},
		{"🇪🇸 Espanha"},
		{"🇨🇳 China"},
	}
	menuT := table.New(
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
	menuT.SetStyles(s)

	return menuT
}

func tableResult(paisSelecionado string, conteudo string) table.Model {
	titulo := fmt.Sprintf("horario mundial: %s", paisSelecionado)
	columns := []table.Column{
		{Title: titulo, Width: 42},
	}
	rows := []table.Row{
		{conteudo},
	}
	resultTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(3),
		table.WithWidth(30),
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
	resultTable.SetStyles(s)
	return resultTable
}

func buscaHorario(urls string) tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get(urls)
		if err != nil {
			return horarioMsg{err: err}
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return horarioMsg{err: err}
		}
		var timeResp timeResponse
		err = json.Unmarshal(body, &timeResp)
		if err != nil {
			return horarioMsg{err: err}
		}
		parsedTime, err := time.Parse(time.RFC3339, timeResp.Datetime)
		if err != nil {
			return horarioMsg{conteudo: timeResp.Datetime}
		}
		timeFormat := parsedTime.Format("15:04:05")
		return horarioMsg{conteudo: timeFormat}
	}
}

func main() {

	m := model{table: tableMenu()}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

//SHANGAI
//BUEINOS AIRES
//MADRI
//MOSCOU
