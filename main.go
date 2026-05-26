package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"
	
	"portfolio/tui"
)

const (
	host    = "0.0.0.0"
	port    = "23234"
	keyPath = ".ssh/id_ed25519"
)

func main() {
	// Ensure SSH host key dir exists
	if err := os.MkdirAll(".ssh", 0700); err != nil {
		log.Fatalf("could not create .ssh dir: %v", err)
	}

	_ = termenv.HasDarkBackground() // ensure termenv is used

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%s", host, port)),
		wish.WithHostKeyPath(keyPath),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("✨ SSH Portfolio running on %s:%s", host, port)
	log.Printf("👉 Connect with: ssh -p %s localhost", port)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error: %v", err)
	}

	log.Println("Server stopped.")
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		fmt.Fprintln(s, "No active terminal detected. Connect with a TTY.")
		return nil, nil
	}
	m := tui.InitialModel()
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
