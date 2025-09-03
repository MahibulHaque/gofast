package list

import (
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/internal/program"
	"github.com/mahibulhaque/gofast/internal/steps"
	"github.com/mahibulhaque/gofast/internal/tui/styles"
)

// ListItem represents an item in the list that implements list.Item interface
type ListItem struct {
	title       string
	description string
	flag        string
}

func (i ListItem) FilterValue() string { return i.title }
func (i ListItem) Title() string       { return i.title }
func (i ListItem) Description() string { return i.description }
func (i ListItem) Flag() string        { return i.flag }

// Selection holds the user's choice
type Selection struct {
	Choice     string
	Flag       string
	IsSelected bool
}

// Model represents the list component model
type Model struct {
	list      list.Model
	selection *Selection
	header    string
	project   *program.Project
	quitting  bool
	keyMap    keyMap
}

type keyMap struct {
	confirm key.Binding
	quit    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.confirm, k.quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.confirm, k.quit},
	}
}

var keys = keyMap{
	confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm selection"),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// NewListModel creates a new list model
func NewListModel(items []steps.Item, selection *Selection, header string, project *program.Project) *Model {
	theme := styles.CurrentTheme()

	// Convert steps.Item to ListItem
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = ListItem{
			title:       item.Title,
			description: item.Desc,
			flag:        item.Flag,
		}
	}

	// Create the list with custom styling
	l := list.New(listItems, newItemDelegate(theme), 0, 0)
	l.Title = header
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = theme.S().Title
	l.Styles.PaginationStyle = theme.S().Muted.
		Padding(0, 1)
	l.Styles.HelpStyle = theme.S().Muted.
		Padding(1, 0, 0, 2)

	return &Model{
		list:      l,
		selection: selection,
		header:    header,
		project:   project,
		keyMap:    keys,
	}
}

// itemDelegate defines how list items are rendered
type itemDelegate struct {
	theme *styles.Theme
}

// Height implements list.ItemDelegate.
func (d itemDelegate) Height() int {
	return 2
}



// Spacing implements list.ItemDelegate.
func (d itemDelegate) Spacing() int {
	return 1
}

// Update implements list.ItemDelegate.
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render implements list.ItemDelegate.
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(ListItem)
	if !ok {
		return
	}

	var (
		title = item.Title()
		desc  = item.Description()
		s     strings.Builder
	)

	if index == m.Index() {
		// Selected item styling
		titleStyle := d.theme.S().TextSelected.
			Bold(true).
			Padding(0, 1)
		descStyle := d.theme.S().TextSelected.
			Padding(0, 1)

		s.WriteString(titleStyle.Render("▶ " + title))
		s.WriteString("\n")
		s.WriteString(descStyle.Render("  " + desc))
	} else {
		// Normal item styling
		titleStyle := d.theme.S().Text.
			Bold(true).
			Padding(0, 1)
		descStyle := d.theme.S().Muted.
			Padding(0, 1)

		s.WriteString(titleStyle.Render("  " + title))
		s.WriteString("\n")
		s.WriteString(descStyle.Render("  " + desc))
	}

	// Write the final string to the writer
	io.WriteString(w, s.String())
}


func newItemDelegate(theme *styles.Theme) itemDelegate {
	return itemDelegate{theme:theme}
}

// Init implements tea.Model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4) // Account for header and footer
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keyMap.confirm):
			selectedItem := m.list.SelectedItem()
			if selectedItem != nil {
				if listItem, ok := selectedItem.(ListItem); ok {
					m.selection.Choice = listItem.Title()
					m.selection.Flag = listItem.Flag()
					m.selection.IsSelected = true
				}
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View implements tea.Model
func (m *Model) View() string {
	if m.quitting {
		return ""
	}

	theme := styles.CurrentTheme()

	// Create the main view
	content := strings.Builder{}

	// Add some spacing at the top
	content.WriteString("\n")

	// Render the list
	content.WriteString(m.list.View())

	// Add help text at the bottom
	helpText := theme.S().Muted.Render("\n" + m.list.Help.View(m.keyMap))
	content.WriteString(helpText)

	return theme.S().Base.
		Render(content.String())
}

// Helper functions

func NewSingleSelectFromStep(step steps.StepSchema, selection *Selection, project *program.Project) *Model {
	return NewListModel(step.Options, selection, step.Headers, project)
}

// RunSingleSelect runs a single-select list and returns the selection
func RunSingleSelect(step steps.StepSchema, project *program.Project) (*Selection, error) {
	selection := &Selection{}
	model := NewSingleSelectFromStep(step, selection, project)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		return nil, err
	}

	return selection, nil
}
