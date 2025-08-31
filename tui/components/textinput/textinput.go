package textinput

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/program"
	"github.com/mahibulhaque/gofast/tui/styles"
)

type (
	errorMsg error
)

// Output represents the text provided in a textinput step
type Output struct {
	Output string
}

// Output.update updates the value of the Output
func (o *Output) update(val string) {
	o.Output = val
}

type model struct {
	textInput textinput.Model
	err       error
	output    *Output
	header    string
	exit      *bool
}

func sanitizeTextInput(input string) error {
	matched, err := regexp.MatchString("^[a-zA-Z0-9_\\/.-]+$", input)

	if !matched {
		return fmt.Errorf("string violates the input regex pattern, err: %v", err)
	}
	return nil
}

func NewTextInputModel(output *Output, header string, program *program.Project) model {

	themeStyles := styles.CurrentTheme().S()
	ti := textinput.New()

	ti.Styles = themeStyles.TextInput
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)
	exit := true
	return model{
		textInput: ti,
		err:       nil,
		output:    output,
		header:    themeStyles.Title.Render(header), // <- use theme Title style
		exit:      &exit,
	}
}

func CreateErrorInputModel(err error) model {

	themeStyles := styles.CurrentTheme().S()
	ti := textinput.New()

	ti.Styles = themeStyles.TextInput
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)
	exit := true

	return model{
		textInput: ti,
		err:       errors.New(themeStyles.Error.Render(err.Error())),
		output:    nil,
		header:    "",
		exit:      &exit,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(m.textInput.Value()) > 1 {
				m.output.update(m.textInput.Value())
				return m, tea.Quit
			}
		case "ctrl+c", "esc":
			*m.exit = true
			return m, tea.Quit
		}
	case errorMsg:
		m.err = msg
		*m.exit = true
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n", m.header, m.textInput.View())
}

func (m model) Err() string {
	return m.err.Error()
}
