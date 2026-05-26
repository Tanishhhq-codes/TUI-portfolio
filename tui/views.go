package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"portfolio/tui/imgrender"
)

// Content configurations
var AsciiName = `
 _____  _    _   _ ___ ____  _   _  ___  
|_   _|/ \  | \ | |_ _/ ___|| | | |/ _ \ 
  | | / _ \ |  \| || |\___ \| |_| | | | |
  | |/ ___ \| |\  || | ___) |  _  | |_| |
  |_/_/   \_\_| \_|___|____/|_| |_|\__\_\
`

var Tagline = "Systems Programming  ·  Rust & Solana  ·  Arch Linux"

var AboutText = `Hey, I’m a CS major currently in my 3rd year who’s really into tech and building random cool stuff. I love exploring different areas like full stack dev, Linux, blockchain, cloud, AI, and low level/system stuff just because I find all of it genuinely fun and interesting.

Most of my learning comes from making projects, breaking things, fixing them again, and spending time gaming or listening to music. Recently I got super obsessed with Linux systems and have been spending way too much time customizing, learning, and understanding how things work internally.

I honestly have interests in almost every domain of computer science. I enjoy talking to people about tech, learning from others, sharing random ideas, and just being around people who genuinely love what they do.

Currently looking for internships and cool opportunities, so yeah… if you think I’m interesting, feel free to say hello or drop a DM 😉.`

var SkillsList = []SkillCategory{
	{
		Name: "Languages",
		Icon: "⟨/⟩",
		Items: []Skill{
			{Name: "Rust"},
			{Name: "C / C++"},
			{Name: "TypeScript"},
			{Name: "JavaScript"},
			{Name: "Solidity"},
			{Name: "Python"},
			{Name: "HTML/CSS"},
		},
	},
	{
		Name: "Frameworks & Web",
		Icon: "◈",
		Items: []Skill{
			{Name: "React.js"},
			{Name: "Three.js"},
			{Name: "Node.js & Express"},
			{Name: "MERN Stack"},
			{Name: "MongoDB & MySQL"},
		},
	},
	{
		Name: "Systems & Tools",
		Icon: "⚙",
		Items: []Skill{
			{Name: "Solana & Anchor"},
			{Name: "Linux (Arch / Hyprland)"},
			{Name: "Git & Version Control"},
			{Name: "AWS Basics"},
			{Name: "Security Auditing"},
		},
	},
}

var ProjectsList = []Project{
	{
		Name:  "OS-Mon Academy",
		Emoji: "🎮",
		Tech:  []string{"React", "TypeScript", "Three.js"},
		Desc:  "3D browser-based RPG teaching OS concepts (scheduling, memory, deadlocks) with Pokémon-style gameplay. Features 18 creatures & 7 areas.",
		Link:  "osmon-academy-4.vercel.app",
		Stars: "osmon-academy-4",
	},
	{
		Name:  "Rustyfetch",
		Emoji: "🦀",
		Tech:  []string{"Rust", "CLI"},
		Desc:  "A fast, Neofetch-inspired system info CLI. Displays OS, kernel, CPU, memory, and uptime. Published on crates.io.",
		Link:  "crates.io/crates/rustyfetch-cli",
		Stars: "rustyfetch-cli",
	},
	{
		Name:  "blockHarvest",
		Emoji: "🌾",
		Tech:  []string{"Solana", "Anchor", "Rust"},
		Desc:  "Decentralised application on Solana with smart contract integration, wallet connectivity, and collaborative on-chain logic.",
		Link:  "github.com/Tanishhhq-codes/blockHarvest",
		Stars: "blockHarvest",
	},
	{
		Name:  "Disaster Management Platform",
		Emoji: "🌊",
		Tech:  []string{"Full-stack", "React"},
		Desc:  "Full-stack disaster awareness platform for coastal regions with real-time alert features, built for the Smart India Hackathon.",
		Link:  "sih-hack-five.vercel.app",
		Stars: "sih-hack-five",
	},
}

var ContactInfo = []ContactItem{
	{Icon: "󰊤", Label: "GitHub", Value: "github.com/Tanishhhq-codes", Color: "#f0f6fc"},
	{Icon: "󰖟", Label: "Website", Value: "tanishhq.tech", Color: "#58a6ff"},
	{Icon: "󰌻", Label: "LinkedIn", Value: "linkedin.com/in/tanishq-jaiswal", Color: "#0a66c2"},
	{Icon: "󰇮", Label: "Email", Value: "jaiswaltanishq10@gmail.com", Color: "#ea4335"},
}

var Tabs = []string{"About", "Skills", "Projects", "Contact"}

// ── Header Renderer ──────────────────────────

func RenderHeader(w int, tick int) string {
	lines := strings.Split(AsciiName, "\n")
	neonColors := []lipgloss.Color{ColorCyan, ColorBlue, ColorPurple, ColorMagenta, ColorPink, ColorOrange}

	var colored []string
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		c := neonColors[(i+tick/4)%len(neonColors)]
		colored = append(colored, lipgloss.NewStyle().Foreground(c).Bold(true).Render(line))
	}
	asciiBlock := strings.Join(colored, "\n")
	asciiBlock = lipgloss.NewStyle().Width(w).Align(lipgloss.Center).Render(asciiBlock)

	taglineRendered := lipgloss.NewStyle().
		Width(w).
		Align(lipgloss.Center).
		Foreground(ColorMagenta).
		Italic(true).
		Render(Tagline)

	dots := strings.Repeat("·", (tick/3)%20+1)
	dotsRendered := lipgloss.NewStyle().
		Width(w).
		Align(lipgloss.Center).
		Foreground(ColorBorder).
		Render(dots)

	return asciiBlock + "\n" + taglineRendered + "\n" + dotsRendered
}

// ── Tabs Renderer ────────────────────────────

func RenderTabs(w int, activeTab int) string {
	var tabWidgets []string
	for i, t := range Tabs {
		label := fmt.Sprintf("  %d. %s  ", i+1, t)
		if i == activeTab {
			s := lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorBg).
				Background(TabColors[i]).
				Padding(0, 1)
			tabWidgets = append(tabWidgets, s.Render(label))
		} else {
			s := lipgloss.NewStyle().
				Foreground(TabColors[i]).
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(ColorBorder).
				Padding(0, 1)
			tabWidgets = append(tabWidgets, s.Render(label))
		}
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, tabWidgets...)
	sep := lipgloss.NewStyle().
		Width(w).
		Foreground(ColorBorder).
		Render(strings.Repeat("─", w))

	return lipgloss.NewStyle().Width(w).Align(lipgloss.Center).Render(row) + "\n" + sep
}

// ── About View ───────────────────────────────

func RenderAbout(w int, mode imgrender.RenderMode) string {
	// Status bar
	statusItems := []string{
		lipgloss.NewStyle().Foreground(ColorGreen).Render("● Online"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render("|"),
		lipgloss.NewStyle().Foreground(ColorYellow).Render("☕ Caffeinated"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render("|"),
		lipgloss.NewStyle().Foreground(ColorCyan).Render("🚀 Building TUIs"),
	}
	statusBar := lipgloss.NewStyle().
		Width(w).
		Align(lipgloss.Center).
		MarginBottom(1).
		Render(strings.Join(statusItems, "  "))

	// Determine layout based on terminal width and avatar mode
	var contentRow string
	avatarWidth := 20
	if mode == imgrender.ModeBraille || mode == imgrender.ModeASCII {
		if w >= 110 {
			avatarWidth = 48
		} else if w >= 90 {
			avatarWidth = 36
		} else {
			avatarWidth = 26
		}
	}

	// Load avatar image if mode is not hidden
	var avatarStr string
	if mode != imgrender.ModeHidden {
		var err error
		avatarStr, err = imgrender.ConvertImageToString("TJ pfp.jpg", avatarWidth, mode)
		if err != nil {
			avatarStr = fmt.Sprintf("\n[Error loading avatar:\n%v]", err)
		}
	}

	if mode == imgrender.ModeHidden || w < 75 || avatarStr == "" {
		// Single column layout
		contentRow = StyleBox.
			Width(w - 4).
			BorderForeground(ColorCyan).
			Render(
				StyleSection.Width(w-8).Render("  About Me") + "\n\n" +
					lipgloss.NewStyle().
						Foreground(ColorWhite).
						Width(w-10).
						Render(AboutText),
			)
	} else {
		// Split screen layout
		aboutWidth := w - avatarWidth - 8
		if aboutWidth < 30 {
			aboutWidth = 30
		}

		aboutBox := StyleBox.
			Width(aboutWidth).
			BorderForeground(ColorCyan).
			Render(
				StyleSection.Width(aboutWidth - 4).Render("  About Me") + "\n\n" +
					lipgloss.NewStyle().
						Foreground(ColorWhite).
						Width(aboutWidth - 6).
						Render(AboutText),
			)

		avatarModeNames := map[imgrender.RenderMode]string{
			imgrender.ModeHalfBlock: "Half-Block (Double Res)",
			imgrender.ModeASCII:     "ASCII Density Art",
			imgrender.ModeBraille:   "Unicode Braille Matrix",
		}

		avatarBox := StyleBox.
			Width(avatarWidth + 4).
			BorderForeground(ColorMagenta).
			Render(
				StyleSection.Width(avatarWidth).Foreground(ColorMagenta).Render("  Identity") + "\n" +
					avatarStr + "\n" +
					lipgloss.NewStyle().
						Width(avatarWidth).
						Align(lipgloss.Center).
						Foreground(ColorMuted).
						Italic(true).
						Render(avatarModeNames[mode]),
			)

		contentRow = lipgloss.JoinHorizontal(lipgloss.Top, aboutBox, avatarBox)
		contentRow = lipgloss.NewStyle().Width(w).Align(lipgloss.Center).Render(contentRow)
	}

	// Fun stats
	stats := [][]string{
		{"🏅", "Solana Bootcamp", string(ColorCyan), "Top 3%"},
		{"🛡️", "Bootcamp Rank", string(ColorYellow), "125/4000+"},
		{"🦀", "Primary Language", string(ColorGreen), "Rust / Go"},
		{"🐧", "Development OS", string(ColorPink), "Arch Linux"},
	}

	var statBoxes []string
	boxW := (w - 8) / len(stats)
	if boxW < 12 {
		boxW = 12
	}
	for _, s := range stats {
		sb := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(s[2])).
			Padding(0, 1).
			Width(boxW).
			Align(lipgloss.Center).
			Render(
				s[0] + "\n" +
					lipgloss.NewStyle().Foreground(lipgloss.Color(s[2])).Bold(true).Render(s[3]) + "\n" +
					lipgloss.NewStyle().Foreground(ColorMuted).Render(s[1]),
			)
		statBoxes = append(statBoxes, sb)
	}
	statsRow := lipgloss.JoinHorizontal(lipgloss.Top, statBoxes...)
	statsRow = lipgloss.NewStyle().Width(w).Align(lipgloss.Center).MarginTop(1).Render(statsRow)

	return statusBar + "\n" + contentRow + "\n" + statsRow
}

// ── Skills View ──────────────────────────────

func RenderSkills(w int, tick int) string {
	colW := (w - 8) / 3
	if colW < 20 {
		colW = 20
	}

	catColors := []lipgloss.Color{ColorCyan, ColorGreen, ColorPink}
	var cols []string

	for ci, cat := range SkillsList {
		cc := catColors[ci%len(catColors)]
		header := lipgloss.NewStyle().
			Foreground(cc).
			Bold(true).
			Render(cat.Icon + "  " + cat.Name)

		var skillLines strings.Builder
		for _, sk := range cat.Items {
			bullet := lipgloss.NewStyle().Foreground(cc).Render("•")
			name := lipgloss.NewStyle().Foreground(ColorWhite).Render(" " + sk.Name)
			skillLines.WriteString("  " + bullet + name + "\n")
		}

		col := StyleBox.
			Width(colW).
			BorderForeground(cc).
			Render(header + "\n\n" + skillLines.String())

		cols = append(cols, col)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, cols...)
	return lipgloss.NewStyle().Width(w).Align(lipgloss.Center).Render(row)
}


// ── Projects View ────────────────────────────

func RenderProjects(w int) string {
	projColors := []lipgloss.Color{ColorCyan, ColorPurple, ColorGreen, ColorOrange}
	
	// Single column for small screens, double for wider
	if w < 75 {
		var cards []string
		for j, p := range ProjectsList {
			c := projColors[j%len(projColors)]
			var tags []string
			for _, t := range p.Tech {
				tags = append(tags, lipgloss.NewStyle().Foreground(ColorBg).Background(c).Padding(0, 1).Render(t))
			}
			tagRow := strings.Join(tags, " ")

			content := lipgloss.NewStyle().Bold(true).Foreground(c).Render(p.Emoji+"  "+p.Name) + "\n\n" +
				lipgloss.NewStyle().Foreground(ColorWhite).Width(w-10).Render(p.Desc) + "\n\n" +
				tagRow + "\n\n" +
				lipgloss.NewStyle().Foreground(ColorMuted).Render("⎋ "+p.Link) + "   " +
				lipgloss.NewStyle().Foreground(ColorYellow).Render(p.Stars)

			card := StyleBox.Width(w - 6).BorderForeground(c).Render(content)
			cards = append(cards, card)
		}
		return lipgloss.NewStyle().Width(w).Align(lipgloss.Center).Render(strings.Join(cards, "\n"))
	}

	halfW := (w - 6) / 2
	var rows []string
	for i := 0; i < len(ProjectsList); i += 2 {
		var pair []string
		for j := i; j < i+2 && j < len(ProjectsList); j++ {
			p := ProjectsList[j]
			c := projColors[j%len(projColors)]

			var tags []string
			for _, t := range p.Tech {
				tags = append(tags,
					lipgloss.NewStyle().
						Foreground(ColorBg).
						Background(c).
						Padding(0, 1).
						Render(t))
			}
			tagRow := strings.Join(tags, " ")

			content := lipgloss.NewStyle().Bold(true).Foreground(c).Render(p.Emoji+"  "+p.Name) + "\n\n" +
				lipgloss.NewStyle().Foreground(ColorWhite).Width(halfW-6).Render(p.Desc) + "\n\n" +
				tagRow + "\n\n" +
				lipgloss.NewStyle().Foreground(ColorMuted).Render("⎋ "+p.Link) + "   " +
				lipgloss.NewStyle().Foreground(ColorYellow).Render(p.Stars)

			card := StyleBox.
				Width(halfW).
				BorderForeground(c).
				Render(content)
			pair = append(pair, card)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Top, pair...)
		rows = append(rows, row)
	}

	all := strings.Join(rows, "\n")
	return lipgloss.NewStyle().Width(w).Align(lipgloss.Center).Render(all)
}

// ── Contact View ─────────────────────────────

func RenderContact(w int, tick int) string {
	cw := w - 10
	if cw > 70 {
		cw = 70
	}

	header := lipgloss.NewStyle().
		Width(cw).
		Align(lipgloss.Center).
		Foreground(ColorCyan).
		Bold(true).
		MarginBottom(1).
		Render("✉  Let's Connect!")

	subheader := lipgloss.NewStyle().
		Width(cw).
		Align(lipgloss.Center).
		Foreground(ColorMuted).
		Italic(true).
		Render("Always happy to chat about code, projects, or opportunities.")

	var items []string
	for _, c := range ContactInfo {
		label := lipgloss.NewStyle().Foreground(ColorMuted).Width(16).Render(c.Label)
		value := lipgloss.NewStyle().Foreground(lipgloss.Color(c.Color)).Bold(true).Render(c.Value)
		row := "  " + label + "  " + value
		items = append(items, row)
	}

	contactBox := StyleBox.
		Width(cw).
		BorderForeground(ColorMagenta).
		Render(strings.Join(items, "\n\n"))

	cursors := []string{"▌", "▋", "█", "▋"}
	cursor := cursors[(tick/5)%len(cursors)]
	prompt := lipgloss.NewStyle().
		Width(cw).
		Align(lipgloss.Center).
		Foreground(ColorGreen).
		MarginTop(1).
		Render("$ echo 'hello' | mail hello@yourname.dev" + cursor)

	return lipgloss.NewStyle().Width(w).Align(lipgloss.Center).
		Render(header + "\n" + subheader + "\n\n" + contactBox + "\n" + prompt)
}

// ── Footer Renderer ──────────────────────────

func RenderFooter(w int, mode imgrender.RenderMode) string {
	// Enable 24-bit color in termenv if possible
	_ = termenv.HasDarkBackground()

	help := []string{
		lipgloss.NewStyle().Foreground(ColorCyan).Render("←/→ or Tab"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render("navigate"),
		lipgloss.NewStyle().Foreground(ColorGreen).Render("1-4"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render("jump"),
		lipgloss.NewStyle().Foreground(ColorPink).Render("q"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render("quit"),
	}
	helpStr := strings.Join(help, "  ")

	// Avoid printing uptime dynamically here, we will render it simply
	uptimeStr := lipgloss.NewStyle().Foreground(ColorMuted).Render("SSH TUI")

	sep := lipgloss.NewStyle().Width(w).Foreground(ColorBorder).Render(strings.Repeat("─", w))
	
	var footer string
	if w > 80 {
		footer = lipgloss.NewStyle().Width(w).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				lipgloss.NewStyle().Width(w-len("SSH TUI")-4).Render(helpStr),
				uptimeStr,
			),
		)
	} else {
		// Simpler format for small screens
		footer = lipgloss.NewStyle().Width(w).Render(helpStr)
	}

	return sep + "\n" + footer
}
