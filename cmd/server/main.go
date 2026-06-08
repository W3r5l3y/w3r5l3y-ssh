package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	wishtea "charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"

	"github.com/W3r5l3y/w3r5l3y-ssh/internal/app"
)

func main() {
	addr := getEnv("SSH_ADDR", "0.0.0.0:23234")
	hostKeyPath := getEnv("SSH_HOST_KEY_PATH", ".ssh/w3r5l3y_ed25519")

	if err := os.MkdirAll(filepath.Dir(hostKeyPath), 0700); err != nil {
		log.Fatal("could not create host key directory", "error", err)
	}

	server, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			wishtea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal("could not create SSH server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	host, port, _ := net.SplitHostPort(addr)
	log.Info("starting w3r5l3y-ssh", "host", host, "port", port)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("SSH server failed", "error", err)
			done <- syscall.SIGTERM
		}
	}()

	<-done

	log.Info("stopping w3r5l3y-ssh")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("could not stop SSH server cleanly", "error", err)
	}
}

func teaHandler(session ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := session.Pty()

	model := app.NewModel(app.SessionInfo{
		User:   session.User(),
		Term:   pty.Term,
		Width:  pty.Window.Width,
		Height: pty.Window.Height,
	})

	return model, []tea.ProgramOption{}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
