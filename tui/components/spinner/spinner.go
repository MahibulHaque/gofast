package spinner

import (
	"fmt"

	"github.com/charmbracelet/bubbles/v2/spinner"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/tui/styles"
)

type errorMsg error

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func NewSpinnerModel() model {
	t := styles.CurrentTheme()
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = t.S().Base.Foreground(t.Accent).Bold(true)

	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errorMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {

	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("%s Preparing...", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}
