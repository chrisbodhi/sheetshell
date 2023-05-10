// Demo code for the Table primitive.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.design/x/clipboard"
)

type model struct {
	pairs  [][]string
	cursor int
}

func initialModel() model {
	fromEnv := os.Environ()
	sort.Strings(fromEnv)
	pairs := [][]string{}
	for _, pair := range fromEnv {
		pairs = append(pairs, strings.SplitN(pair, "=", 2))
	}
	return model{
		pairs: pairs,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	err := clipboard.Init()
	if err != nil {
		log.Panic(err)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.pairs)-1 {
				m.cursor++
			}
		case "enter", " ":
			ba := []byte(m.pairs[m.cursor][1])
			clipboard.Write(clipboard.FmtText, ba)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Your local env vars\n"

	for i, pair := range m.pairs {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, pair[0], pair[1])
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
