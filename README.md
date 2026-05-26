# 🖥 SSH Terminal Portfolio

A personal portfolio that lives entirely in the terminal, served over SSH.
Anyone in the world can visit your portfolio by running:

```bash
ssh -p 23234 your-server-ip
# or with a domain:
ssh -p 23234 yourname.dev
```

## ✨ Features

- 🎨 **Multi-colour neon theme** with animated ASCII art header
- 📄 **4 sections**: About, Skills, Projects, Contact
- 📊 **Animated skill bars** that fill on load
- ⌨️ **Keyboard navigation** — tab, arrow keys, number shortcuts
- 🌐 **Global access** — anyone with SSH can visit (no account needed)
- 🔒 **No auth required** for visitors — read-only TUI

---

## 🚀 Quick Start (Docker — recommended)

### 1. Clone / copy files to your VPS
```bash
git clone https://github.com/yourname/ssh-portfolio
cd ssh-portfolio
```

### 2. Personalise your content
Edit `main.go` — scroll to the `CONTENT` section at the top:
- Change `asciiName` (or regenerate with figlet/toilet)
- Edit `tagline`, `aboutText`
- Fill in your real `skills`, `projects`, `contactInfo`

### 3. Run with Docker Compose
```bash
docker compose up -d
```

### 4. Open firewall port
```bash
# UFW (Ubuntu)
ufw allow 23234/tcp

# Or with iptables
iptables -A INPUT -p tcp --dport 23234 -j ACCEPT
```

### 5. Test it!
```bash
ssh -p 23234 localhost
```

---

## 🛠 Manual Setup (no Docker)

### Prerequisites
- Go 1.22+
- A Linux VPS (DigitalOcean, Hetzner, Linode, etc.)

### Build & run
```bash
go mod tidy
go build -o portfolio .
./portfolio
```

### Run as a systemd service (keeps it alive 24/7)
```bash
# Create a dedicated user
sudo useradd -r -s /bin/false portfolio
sudo mkdir -p /opt/portfolio
sudo cp portfolio /opt/portfolio/
sudo cp portfolio.service /etc/systemd/system/

sudo systemctl daemon-reload
sudo systemctl enable portfolio
sudo systemctl start portfolio
sudo systemctl status portfolio
```

---

## 🌐 Custom Domain Setup

Point an A record to your VPS IP, then visitors can SSH with:
```bash
ssh -p 23234 yourname.dev
```

Or, if you use port 22 (standard SSH port), they don't need `-p 23234`:
```bash
# Change port in main.go:
const port = "22"
```
> ⚠️ If you use port 22, make sure your own SSH access uses a different port
> or key-based auth so you don't lock yourself out!

---

## ⌨️ Navigation

| Key | Action |
|-----|--------|
| `Tab` / `→` / `l` | Next section |
| `Shift+Tab` / `←` / `h` | Previous section |
| `1` `2` `3` `4` | Jump to section |
| `j` / `↓` | Scroll down |
| `k` / `↑` | Scroll up |
| `q` / `Ctrl+C` | Quit |

---

## 🎨 Customising the ASCII Art

Use `figlet` or `toilet` to generate new ASCII art:
```bash
# Install figlet
apt install figlet toilet

# Generate 3D shadow style
figlet -f shadow "Your Name"

# Or use toilet for coloured fonts
toilet -f future "Your Name"

# Copy output into the asciiName variable in main.go
```

Popular figlet fonts for names: `banner3`, `shadow`, `big`, `slant`, `3d`

---

## 📁 Project Structure

```
ssh-portfolio/
├── main.go           # Everything — SSH server + TUI app
├── go.mod            # Go module file
├── go.sum            # Dependency checksums
├── Dockerfile        # Docker build
├── docker-compose.yml
├── portfolio.service # systemd service
└── README.md
```

---

## 🔧 Troubleshooting

**"Connection refused"** — make sure the service is running and firewall port is open

**"Host key verification failed"** on reconnect — this is normal if you rebuilt;
visitors just need to run `ssh-keygen -R [ip]:23234`

**Terminal looks garbled** — make sure your terminal supports UTF-8 and 256 colours.
Use a modern terminal: kitty, alacritty, iTerm2, Windows Terminal, etc.

---

## 📜 Tech Stack

| Library | Purpose |
|---------|---------|
| [Wish](https://github.com/charmbracelet/wish) | SSH server framework |
| [BubbleTea](https://github.com/charmbracelet/bubbletea) | TUI framework (Elm architecture) |
| [Lipgloss](https://github.com/charmbracelet/lipgloss) | Terminal styling (CSS for terminals) |

---

Made with ❤️ and too much coffee. Because why have a normal portfolio?
