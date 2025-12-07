package webserver

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	opts ServerOptions
	log  *slog.Logger
	mux  *chi.Mux
}

func NewServer(logger *slog.Logger, opts ServerOptions) (srv *Server) {
	opts.populateDefaults()

	srv = &Server{
		opts: opts,
		log:  logger,
		mux:  chi.NewRouter(),
	}

	for path, realpath := range opts.StaticDirs {
		srv.addStaticDir(path, realpath)
	}

	for path, handler := range opts.Handlers {
		srv.addHandlers(path, handler)
	}

	return
}

func (s *Server) addStaticDir(path, realpath string) {
	s.log.Info("Adding static directory", "path", path)
	s.mux.Handle(path+"/*", http.StripPrefix(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realpath := filepath.Join(realpath, r.URL.Path)

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			compressedPath := realpath + ".gz"
			_, err := os.Lstat(compressedPath)
			if err == nil {
				realpath = compressedPath
				w.Header().Set("Content-Encoding", "gzip")
			}
		}

		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		} else if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}

		s.log.Info("Serving static file", "path", r.URL.Path, "realpath", realpath)
		http.ServeFile(w, r, realpath)
	})))
}

func (s *Server) addHandlers(path string, handlers Handlers) {
	s.log.Info("Registering routes", "path", path)
	r := chi.NewRouter()
	handlers.Register(path, r)
	s.mux.Mount(path, r)
}

func (s *Server) Start() error {
	if s.opts.Certs != nil {
		s.log.Info("Server is starting with TLS...",
			"addr", s.opts.Addr, "cert", s.opts.Certs.Cert, "key", s.opts.Certs.Key)
		return http.ListenAndServeTLS(s.opts.Addr, s.opts.Certs.Cert, s.opts.Certs.Key, s.mux)
	}
	s.log.Info("Server is starting...", "addr", s.opts.Addr)
	return http.ListenAndServe(s.opts.Addr, s.mux)
}
