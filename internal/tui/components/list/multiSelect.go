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

// MultiSelection holds the user's choices
type MultiSelection struct {
	Choices   []string
	Flags     []string
	Selected  map[int]bool
	Confirmed bool
}

// MultiModel represents the multi-select list component model
type MultiModel struct {
	list      list.Model
	selection *MultiSelection
	header    string
	project   *program.Project
	quitting  bool
	keyMap    multiKeyMap
}

type multiKeyMap struct {
	toggle  key.Binding
	confirm key.Binding
	quit    key.Binding
}

func (k multiKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.toggle, k.confirm, k.quit}
}

func (k multiKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.toggle, k.confirm, k.quit},
	}
}

var multiKeys = multiKeyMap{
	toggle: key.NewBinding(
		key.WithKeys(" ", "space"),
		key.WithHelp("space", "toggle item"),
	),
	confirm: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "confirm selection"),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// NewMultiListModel creates a new multi-select list model
func NewMultiListModel(items []steps.Item, selection *MultiSelection, header string, project *program.Project) *MultiModel {
	theme := styles.CurrentTheme()

	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = ListItem{
			title:       item.Title,
			description: item.Desc,
			flag:        item.Flag,
		}
	}

	l := list.New(listItems, newMultiItemDelegate(theme, selection), 0, 0)
	l.Title = header
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = theme.S().Title
	l.Styles.PaginationStyle = theme.S().Muted.
		Padding(0, 1)
	l.Styles.HelpStyle = theme.S().Muted.
		Padding(1, 0, 0, 2)

	if selection.Selected == nil {
		selection.Selected = make(map[int]bool)
	}

	return &MultiModel{
		list:      l,
		selection: selection,
		header:    header,
		project:   project,
		keyMap:    multiKeys,
	}
}

// multiItemDelegate defines how list items are rendered (with checkbox)
type multiItemDelegate struct {
	theme     *styles.Theme
	selection *MultiSelection
}

// Height implements list.ItemDelegate.
func (d multiItemDelegate) Height() int {
	return 2
}

// Spacing implements list.ItemDelegate.
func (d multiItemDelegate) Spacing() int {
	return 1
}

// Update implements list.ItemDelegate.
func (d multiItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render implements list.ItemDelegate.
func (d multiItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(ListItem)
	if !ok {
		return
	}

	var (
		title = item.Title()
		desc  = item.Description()
		s     strings.Builder
	)

	checked := d.selection != nil && d.selection.Selected[index]
	checkbox := "[ ]"
	if checked {
		checkbox = "[x]"
	}

	// Selected row (cursor on it)
	if index == m.Index() {
		titleStyle := d.theme.S().TextSelected.
			Bold(true).
			Padding(0, 1)
		descStyle := d.theme.S().TextSelected.
			Padding(0, 1)

		s.WriteString(titleStyle.Render("▶ " + checkbox + " " + title))
		s.WriteString("\n")
		s.WriteString(descStyle.Render("  " + desc))
	} else {
		titleStyle := d.theme.S().Text.
			Bold(true).
			Padding(0, 1)
		descStyle := d.theme.S().Muted.
			Padding(0, 1)

		s.WriteString(titleStyle.Render("   " + checkbox + " " + title))
		s.WriteString("\n")
		s.WriteString(descStyle.Render("  " + desc))
	}

	io.WriteString(w, s.String())
}

func newMultiItemDelegate(theme *styles.Theme, selection *MultiSelection) multiItemDelegate {
	return multiItemDelegate{theme: theme, selection: selection}
}

// Init implements tea.Model
func (m *MultiModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m *MultiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keyMap.toggle):
			idx := m.list.Index()
			m.selection.Selected[idx] = !m.selection.Selected[idx]
			return m, nil

		case key.Matches(msg, m.keyMap.confirm):
			// Build final choices in the order of items
			items := m.list.Items()
			choices := make([]string, 0)
			flags := make([]string, 0)
			for i, it := range items {
				if m.selection.Selected[i] {
					if li, ok := it.(ListItem); ok {
						choices = append(choices, li.Title())
						flags = append(flags, li.Flag())
					}
				}
			}
			m.selection.Choices = choices
			m.selection.Flags = flags
			m.selection.Confirmed = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View implements tea.Model
func (m *MultiModel) View() string {
	if m.quitting {
		return ""
	}

	theme := styles.CurrentTheme()

	view := m.list.View()

	// Inject help controls footer
	help := theme.S().Muted.Render("\n" + m.list.Help.View(m.keyMap))
	return theme.S().Base.Render(view + help)
}

// NewMultiSelectFromStep constructs a MultiModel from a step schema
func NewMultiSelectFromStep(step steps.StepSchema, selection *MultiSelection, project *program.Project) *MultiModel {
	return NewMultiListModel(step.Options, selection, step.Headers, project)
}

// RunMultiSelect runs the multi-select list and returns the selection
func RunMultiSelect(step steps.StepSchema, project *program.Project) (*MultiSelection, error) {
	selection := &MultiSelection{}
	model := NewMultiSelectFromStep(step, selection, project)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		return nil, err
	}

	return selection, nil
}
