package multiSelect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/internal/program"
	"github.com/mahibulhaque/gofast/internal/steps"
	"github.com/mahibulhaque/gofast/internal/tui/styles"
)

type Selection struct {
	Choices map[string]bool
}

func (s *Selection) Update(optionName string, value bool) {
	s.Choices[optionName] = value
}

type model struct {
	cursor   int
	options  []steps.Item
	selected map[int]struct{}
	choices  *Selection
	header   string
	exit     *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModelMultiSelect(options []steps.Item, selection *Selection, header string, program *program.Project) model {
	theme := styles.CurrentTheme()
	return model{
		options:  options,
		selected: make(map[int]struct{}),
		choices:  selection,
		header:   theme.S().Title.Render(header),
		exit:     &program.Exit,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			for selectedKey := range m.selected {
				m.choices.Update(m.options[selectedKey].Flag, true)
				m.cursor = selectedKey
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	themeStyles := styles.CurrentTheme().S()
	s := m.header + "\n\n"

	for i, option := range m.options {

		cursor := " "
		themeStyles.Text.Render(">")
		if m.cursor == i {
			cursor = themeStyles.TextSelected.Render(">")
			option.Title = themeStyles.Title.Render(option.Title)
			option.Desc = themeStyles.Subtitle.Render(option.Desc)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = themeStyles.TextSelected.Render("*")
		}

		title := themeStyles.TextSelected.Render(option.Title)
		description := themeStyles.Subtitle.Render(option.Desc)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", cursor, checked, title, description)
	}
	s += fmt.Sprintf("Press %s to confirm choice.\n", themeStyles.TextSelected.Render("y"))

	return s
}
