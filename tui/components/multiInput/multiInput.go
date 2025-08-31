package multiInput

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/program"
	"github.com/mahibulhaque/gofast/steps"
	"github.com/mahibulhaque/gofast/tui/styles"
)

type Selection struct {
	Choice string
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(value string) {
	s.Choice = value
}

type model struct {
	cursor   int
	choices  []steps.Item
	selected map[int]struct{}
	choice   *Selection
	header   string
	exit     *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

// InitialModelMulti initializes a multiInput step with
// the given data
func InitialModelMulti(choices []steps.Item, selection *Selection, header string, program *program.Project) model {
	theme := styles.CurrentTheme()
	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
		choice:   selection,
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.selected) == 1 {
				m.selected = make(map[int]struct{})
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			if len(m.selected) == 1 {
				for selectedKey := range m.selected {
					m.choice.Update(m.choices[selectedKey].Title)
					m.cursor = selectedKey
				}
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	themeStyles := styles.CurrentTheme().S()
	s := m.header + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = themeStyles.TextSelected.Render(">")
			choice.Title = themeStyles.Title.Render(choice.Title)
			choice.Desc = themeStyles.Subtitle.Render(choice.Desc)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = themeStyles.TextSelected.Render("x")
		}

		title := themeStyles.TextSelected.Render(choice.Title)
		description := themeStyles.Subtitle.Render(choice.Desc)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", cursor, checked, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n\n", themeStyles.TextSelected.Render("y"))
	return s
}
