package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sudeep9/webserver"
)

type HandlerImpl struct{}

func (h *HandlerImpl) Register(r *chi.Mux) {
	r.Get("/hello", h.handleHome)
}

func (h *HandlerImpl) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// Initialize the web server with default options
	opts := webserver.ServerOptions{
		Handlers: map[string]webserver.Handlers{
			"/test": &HandlerImpl{},
		},
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server := webserver.NewServer(logger, opts)

	// Start the server
	if err := server.Start(); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
