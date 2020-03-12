package http

import (
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/strangnet/csp-uri-report/internal/pkg/report"

	log "github.com/sirupsen/logrus"
)

// Server interface defines operations over http
type Server interface {
	Open() error

	Close()

	Handler() http.Handler
}

type server struct {
	listener net.Listener
	logger   *log.Logger
	encoder  *encoder
	addr     string
	rs       report.Service
}

func (s *server) Open() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.listener = listener
	server := http.Server{
		Handler: s.Handler(),
	}

	return server.Serve(s.listener)
}

func (s *server) Close() {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "web")
		fs := http.StripPrefix("/", http.FileServer(http.Dir(filesDir)))
		w.Header().Add("Content-Security-Policy", "default-src 'none'; img-src 'self'; script-src 'self'; report-uri http://localhost:8080/api/report;")
		fs.ServeHTTP(w, req)
	}))

	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Route("/login", newLoginHandler(s.encoder, s.rs).Routes)
			r.Route("/report", newReportHandler(s.encoder, s.rs).Routes)
		})
	})

	return r
}

// NewServer initializes a server object
func NewServer(logger *log.Logger, addr string, rs report.Service) Server {
	return &server{
		logger:  logger,
		addr:    addr,
		rs:      rs,
		encoder: newEncoder(logger),
	}
}
