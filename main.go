package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var baseStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#641882"))

type appState int

const (
	stateMenu appState = iota
	stateTimeZones
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

var argentinaTimezones = map[string]string{
	"-03 GMT • Buenos Aires": "https://time.now/developer/api/timezone/America/Argentina/Buenos_Aires",
}
var brasilTimezones = map[string]string{
	"-02 GMT • Fernando de Noronha":                      "https://time.now/developer/api/timezone/America/Noronha",
	"-03 GMT • Brasília (S, SE, NE, GO, DF, PA, TO, AP)": "https://time.now/developer/api/timezone/America/Sao_Paulo",
	"-04 GMT • Manaus (AM, RR, RO, MT, MS)":              "https://time.now/developer/api/timezone/America/Manaus",
	"-05 GMT • Rio Branco (AC)":                          "https://time.now/developer/api/timezone/America/Rio_Branco",
}
var russiaTimezones = map[string]string{
	"+02 GMT • Kaliningrado":                               "https://time.now/developer/api/timezone/Europe/Kaliningrad",
	"+03 GMT • Moscou (São Petersburgo, Kazan, Sochi)":     "https://time.now/developer/api/timezone/Europe/Moscow",
	"+04 GMT • Samara (Astracã, Izhevsk)":                  "https://time.now/developer/api/timezone/Europe/Samara",
	"+05 GMT • Ecaterimburgo (Chelyabinsk, Perm, Ufa)":     "https://time.now/developer/api/timezone/Asia/Yekaterinburg",
	"+06 GMT • Omsk":                                       "https://time.now/developer/api/timezone/Asia/Omsk",
	"+07 GMT • Krasnoyarsk / Novosibirsk (Tomsk, Barnaul)": "https://time.now/developer/api/timezone/Asia/Krasnoyarsk",
	"+08 GMT • Irkutsk (Ulan-Ude)":                         "https://time.now/developer/api/timezone/Asia/Irkutsk",
	"+09 GMT • Yakutsk / Chita":                            "https://time.now/developer/api/timezone/Asia/Chita",
	"+10 GMT • Vladivostok (Khabarovsk)":                   "https://time.now/developer/api/timezone/Asia/Vladivostok",
	"+11 GMT • Magadan / Sakhalin (Srednekolymsk)":         "https://time.now/developer/api/timezone/Asia/Magadan",
	"+12 GMT • Kamchatka / Anadyr":                         "https://time.now/developer/api/timezone/Asia/Kamchatka",
}
var espanhaTimeZones = map[string]string{
	"+00 GMT • Ilhas Canárias (Las Palmas, Tenerife)": "https://time.now/developer/api/timezone/Atlantic/Canary",
	"+01 GMT • Madrid (Barcelona, Sevilha, Valência)": "https://time.now/developer/api/timezone/Europe/Madrid",
}
var chinhaTimeZones = map[string]string{
	"+8 GMT • Pequim / Xangai": "https://time.now/developer/api/timezone/Asia/Shanghai",
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			if m.state == stateResult {
				m.state = stateTimeZones
				m.table = tableTimeZone(m.paisSelecionado)
				return m, nil
			} else if m.state == stateTimeZones {
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
				m.state = stateTimeZones
				m.table = tableTimeZone(m.paisSelecionado)
				return m, nil
			} else if m.state == stateTimeZones {
				timeSelecionado := m.table.SelectedRow()[0]
				mapaPais := fusoSelecionado(m.paisSelecionado)
				urlAPI := mapaPais[timeSelecionado]

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
	return tea.NewView(baseStyle.Render(m.table.View()) + "v1.1" + "\n " + m.table.HelpView() + "\n ")

}

func tableMenu() table.Model {

	columns := []table.Column{
		{Title: "clockTimer", Width: 42},
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
		BorderForeground(lipgloss.Color("#641882")).
		Foreground(lipgloss.Color("#ab0fe9")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("#ab0fe9"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#41882")).
		Background(lipgloss.Color("#222321")).
		BorderBottom(true).
		Bold(false)
	menuT.SetStyles(s)

	return menuT
}
func tableTimeZone(paisSelecionado string) table.Model {
	titulo := fmt.Sprintf("Selecione um fuso-horario: %s", paisSelecionado)
	columns := []table.Column{
		{Title: titulo, Width: 55},
	}

	var chaves []string
	for fuso := range fusoSelecionado(paisSelecionado) {
		chaves = append(chaves, fuso)
	}
	sort.Strings(chaves)

	var rows []table.Row
	for _, fuso := range chaves {
		rows = append(rows, table.Row{fuso})
	}

	zoneT := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithWidth(55),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#641882")).
		Foreground(lipgloss.Color("#ab0fe9")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("#ab0fe9"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#41882")).
		Background(lipgloss.Color("#222321")).
		BorderBottom(true).
		Bold(false)
	zoneT.SetStyles(s)

	return zoneT
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
		BorderForeground(lipgloss.Color("#641882")).
		Foreground(lipgloss.Color("#ab0fe9")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("#ab0fe9"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#41882")).
		Background(lipgloss.Color("#222321")).
		BorderBottom(true).
		Bold(false)
	resultTable.SetStyles(s)
	return resultTable
}

func buscaHorario(urlAPI string) tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get(urlAPI)
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
func fusoSelecionado(paisSelecionado string) map[string]string {
	switch paisSelecionado {
	case "🇦🇷 Argentina":
		return argentinaTimezones
	case "🇧🇷 Brasil":
		return brasilTimezones
	case "🇷🇺 Russia":
		return russiaTimezones
	case "🇪🇸 Espanha":
		return espanhaTimeZones
	case "🇨🇳 China":
		return chinhaTimeZones
	}
	return nil
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
