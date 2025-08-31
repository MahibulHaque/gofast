package logo

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/exp/slice"
	"github.com/mahibulhaque/gofast/tui/styles"
)

// letterform represents a letterform. It can be stretched horizontally by
// a given amount via the boolean argument.
type letterform func(bool) string

const diag = `╱`

// Opts are the options for rendering the Gofast title art.
type Opts struct {
	FieldColor   color.Color // diagonal lines
	TitleColorA  color.Color // left gradient ramp point
	TitleColorB  color.Color // right gradient ramp point
	CharmColor   color.Color // Charm™ text color (can be used for branding)
	VersionColor color.Color // Version text color
	Width        int         // width of the rendered logo, used for truncation
}

// DefaultOpts returns default color options for the logo matching Charm CLI colors
func DefaultOpts() Opts {
	return Opts{
		FieldColor:   color.RGBA{102, 102, 255, 255}, // Blue for diagonal lines
		TitleColorA:  color.RGBA{255, 20, 147, 255},  // Deep pink/magenta (left side)
		TitleColorB:  color.RGBA{138, 43, 226, 255},  // Blue violet (right side)
		CharmColor:   color.RGBA{255, 20, 147, 255},  // Magenta for branding
		VersionColor: color.RGBA{200, 200, 200, 255}, // Light gray for version
		Width:        80,
	}
}

// Render renders the Gofast logo. Set the compact argument to true to render the narrow
// version, intended for use in a sidebar.
//
// The compact argument determines whether it renders compact for the sidebar
// or wider for the main pane.
func Render(version string, compact bool, o Opts) string {
	const branding = " Gofast™"

	fg := func(c color.Color, s string) string {
		return lipgloss.NewStyle().Foreground(c).Render(s)
	}

	// Title - just "FAST"
	const spacing = 1
	letterforms := []letterform{
		letterF,
		letterA,
		letterS,
		letterT,
	}
	stretchIndex := -1 // -1 means no stretching.
	if !compact {
		stretchIndex = rand.IntN(len(letterforms))
	}

	fast := renderWord(spacing, stretchIndex, letterforms...)
	fastWidth := lipgloss.Width(fast)
	b := new(strings.Builder)
	for r := range strings.SplitSeq(fast, "\n") {
		fmt.Fprintln(b, styles.ApplyForegroundGrad(r, o.TitleColorA, o.TitleColorB))
	}
	fast = b.String()

	// Branding and version.
	metaRowGap := 1
	maxVersionWidth := fastWidth - lipgloss.Width(branding) - metaRowGap
	version = ansi.Truncate(version, maxVersionWidth, "…") // truncate version if too long.
	gap := max(0, fastWidth-lipgloss.Width(branding)-lipgloss.Width(version))
	metaRow := fg(o.CharmColor, branding) + strings.Repeat(" ", gap) + fg(o.VersionColor, version)

	// Join the meta row and big Fast title.
	fast = strings.TrimSpace(metaRow + "\n" + fast)

	// Narrow version.
	if compact {
		field := fg(o.FieldColor, strings.Repeat(diag, fastWidth))
		return strings.Join([]string{field, field, fast, field, ""}, "\n")
	}

	fieldHeight := lipgloss.Height(fast)

	// Left field.
	const leftWidth = 6
	leftFieldRow := fg(o.FieldColor, strings.Repeat(diag, leftWidth))
	leftField := new(strings.Builder)
	for range fieldHeight {
		fmt.Fprintln(leftField, leftFieldRow)
	}

	// Right field.
	rightWidth := max(15, o.Width-fastWidth-leftWidth-2) // 2 for the gap.
	const stepDownAt = 0
	rightField := new(strings.Builder)
	for i := range fieldHeight {
		width := rightWidth
		if i >= stepDownAt {
			width = rightWidth - (i - stepDownAt)
		}
		fmt.Fprint(rightField, fg(o.FieldColor, strings.Repeat(diag, width)), "\n")
	}

	// Return the wide version.
	const hGap = " "
	logo := lipgloss.JoinHorizontal(lipgloss.Top, leftField.String(), hGap, fast, hGap, rightField.String())
	if o.Width > 0 {
		// Truncate the logo to the specified width.
		lines := strings.Split(logo, "\n")
		for i, line := range lines {
			lines[i] = ansi.Truncate(line, o.Width, "")
		}
		logo = strings.Join(lines, "\n")
	}
	return logo
}

// SmallRender renders a smaller version of the Gofast logo, suitable for
// smaller windows or sidebar usage.
func SmallRender(width int, o Opts) string {
	t := styles.CurrentTheme()
	title := t.S().Base.Foreground(t.Secondary).Render("Gofast™")
	title = fmt.Sprintf("%s %s", title, styles.ApplyBoldForegroundGrad("Fast", t.Secondary, t.Primary))
	remainingWidth := width - lipgloss.Width(title) - 1 // 1 for the space after "Fast"
	if remainingWidth > 0 {
		lines := strings.Repeat("╱", remainingWidth)
		title = fmt.Sprintf("%s %s", title, t.S().Base.Foreground(t.Primary).Render(lines))
	}
	return title
}

// renderWord renders letterforms to form a word. stretchIndex is the index of
// the letter to stretch, or -1 if no letter should be stretched.
func renderWord(spacing int, stretchIndex int, letterforms ...letterform) string {
	if spacing < 0 {
		spacing = 0
	}

	renderedLetterforms := make([]string, len(letterforms))

	// pick one letter randomly to stretch
	for i, letter := range letterforms {
		renderedLetterforms[i] = letter(i == stretchIndex)
	}

	if spacing > 0 {
		// Add spaces between the letters and render.
		renderedLetterforms = slice.Intersperse(renderedLetterforms, strings.Repeat(" ", spacing))
	}
	return strings.TrimSpace(
		lipgloss.JoinHorizontal(lipgloss.Top, renderedLetterforms...),
	)
}

// letterF renders the letter F in a stylized way.
func letterF(stretch bool) string {
	// ▄▀▀▀▀
	// █▀▀▀
	// ▀

	left := heredoc.Doc(`
		▄
		█
		▀
	`)
	right := heredoc.Doc(`
		▀
		▀
	`)
	return joinLetterform(
		left,
		stretchLetterformPart(right, letterformProps{
			stretch:    stretch,
			width:      4,
			minStretch: 6,
			maxStretch: 10,
		}),
	)
}

// letterA renders the letter A in a stylized way.
func letterA(stretch bool) string {
	// ▄▀▀▀▄
	// █▀▀▀█
	// ▀   ▀

	side := heredoc.Doc(`
		▄
		█
		▀
	`)
	middle := heredoc.Doc(`
		▀
		▀
	`)
	return joinLetterform(
		side,
		stretchLetterformPart(middle, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 6,
			maxStretch: 10,
		}),
		side,
	)
}

// letterS renders the letter S in a stylized way.
func letterS(stretch bool) string {
	// ▄▀▀▀▀
	// ▀▀▀▀█
	// ▀▀▀▀▀

	left := heredoc.Doc(`
		▄
		▀
		▀
	`)
	center := heredoc.Doc(`
		▀
		▀
		▀
	`)
	right := heredoc.Doc(`
		▀
		█
	`)
	return joinLetterform(
		left,
		stretchLetterformPart(center, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 6,
			maxStretch: 10,
		}),
		right,
	)
}

// letterT renders the letter T in a stylized way.
func letterT(stretch bool) string {
	// ▀▀▀▀▀
	//   █
	//   ▀

	// top := heredoc.Doc(`
	// 	▀
	// `)
	// bottom := heredoc.Doc(`
	//
	//
	// 	█
	// 	▀
	// `)

	if stretch {
		width := rand.IntN(6) + 6 // 6-11 characters wide
		topRow := strings.Repeat("▀", width)
		padding := strings.Repeat(" ", width/2-1)
		return topRow + "\n" + padding + "█\n" + padding + "▀"
	}

	return "▀▀▀▀▀\n  █\n  ▀"
}

func joinLetterform(letters ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, letters...)
}

// letterformProps defines letterform stretching properties.
type letterformProps struct {
	width      int
	minStretch int
	maxStretch int
	stretch    bool
}

// stretchLetterformPart is a helper function for letter stretching.
func stretchLetterformPart(s string, p letterformProps) string {
	if p.maxStretch < p.minStretch {
		p.minStretch, p.maxStretch = p.maxStretch, p.minStretch
	}
	n := p.width
	if p.stretch {
		n = rand.IntN(p.maxStretch-p.minStretch) + p.minStretch
	}
	parts := make([]string, n)
	for i := range parts {
		parts[i] = s
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, parts...)
}
