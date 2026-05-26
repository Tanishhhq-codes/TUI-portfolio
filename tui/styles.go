package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// ─────────────────────────────────────────────
//  DATA TYPES
// ─────────────────────────────────────────────

type SkillCategory struct {
	Name  string
	Icon  string
	Items []Skill
}

type Skill struct {
	Name  string
	Level int
}

type Project struct {
	Name  string
	Emoji string
	Tech  []string
	Desc  string
	Link  string
	Stars string
}

type ContactItem struct {
	Icon  string
	Label string
	Value string
	Color string
}

// ─────────────────────────────────────────────
//  COLOURS  (neon multi-colour palette)
// ─────────────────────────────────────────────

var (
	ColorBg      = lipgloss.Color("#0d0d0d")
	ColorSurface = lipgloss.Color("#141414")
	ColorBorder  = lipgloss.Color("#2a2a2a")
	ColorMuted   = lipgloss.Color("#555555")

	ColorCyan    = lipgloss.Color("#00f5ff")
	ColorMagenta = lipgloss.Color("#ff00ff")
	ColorGreen   = lipgloss.Color("#39ff14")
	ColorYellow  = lipgloss.Color("#ffe600")
	ColorOrange  = lipgloss.Color("#ff6b00")
	ColorPink    = lipgloss.Color("#ff2d78")
	ColorPurple  = lipgloss.Color("#bf5af2")
	ColorBlue    = lipgloss.Color("#2997ff")
	ColorWhite   = lipgloss.Color("#f5f5f5")

	// tab accent colours cycling through the rainbow
	TabColors = []lipgloss.Color{ColorCyan, ColorGreen, ColorYellow, ColorPink}
)

// ─────────────────────────────────────────────
//  STYLES
// ─────────────────────────────────────────────

var (
	StyleBase = lipgloss.NewStyle().
			Background(ColorBg).
			Foreground(ColorWhite)

	StyleTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCyan)

	StyleSubtitle = lipgloss.NewStyle().
			Italic(true).
			Foreground(ColorMagenta)

	StyleSection = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorYellow).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(ColorBorder).
			MarginBottom(1)

	StyleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(1, 2)

	StyleHighlight = lipgloss.NewStyle().
			Foreground(ColorCyan).
			Bold(true)

	StyleMuted = lipgloss.NewStyle().
			Foreground(ColorMuted)

	StyleKey = lipgloss.NewStyle().
			Foreground(ColorGreen).
			Bold(true)

	StyleHelp = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Italic(true)
)
