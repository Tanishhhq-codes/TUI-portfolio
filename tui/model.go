package tui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"portfolio/tui/imgrender"
)

type TickMsg time.Time

type Model struct {
	ActiveTab  int
	Width      int
	Height     int
	Tick       int
	AvatarMode imgrender.RenderMode
}

func InitialModel() Model {
	return Model{
		ActiveTab:  0,
		AvatarMode: imgrender.ModeBraille, // Default to Braille art
	}
}

func (m Model) Init() tea.Cmd {
	return TickCmd()
}

func TickCmd() tea.Cmd {
	return tea.Tick(120*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case TickMsg:
		m.Tick++
		return m, TickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "l", "right":
			m.ActiveTab = (m.ActiveTab + 1) % len(Tabs)
		case "shift+tab", "h", "left":
			m.ActiveTab = (m.ActiveTab - 1 + len(Tabs)) % len(Tabs)
		case "1":
			m.ActiveTab = 0
		case "2":
			m.ActiveTab = 1
		case "3":
			m.ActiveTab = 2
		case "4":
			m.ActiveTab = 3

		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.Width == 0 {
		return "Loading..."
	}

	w := m.Width
	if w > 120 {
		w = 120
	}

	var b strings.Builder

	// ── Header ──────────────────────────────
	b.WriteString(RenderHeader(w, m.Tick))
	b.WriteString("\n")

	// ── Tabs ────────────────────────────────
	b.WriteString(RenderTabs(w, m.ActiveTab))
	b.WriteString("\n\n")

	// ── Content ─────────────────────────────
	switch m.ActiveTab {
	case 0:
		b.WriteString(RenderAbout(w, m.AvatarMode))
	case 1:
		b.WriteString(RenderSkills(w, m.Tick))
	case 2:
		b.WriteString(RenderProjects(w))
	case 3:
		b.WriteString(RenderContact(w, m.Tick))
	}

	// ── Footer ──────────────────────────────
	b.WriteString("\n")
	b.WriteString(RenderFooter(w, m.AvatarMode))

	return b.String()
}
