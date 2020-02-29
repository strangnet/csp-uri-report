package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/strangnet/csp-uri-report/internal/pkg/report"
)

// loginHandler handles login
type loginHandler struct {
	encoder *encoder
	rs      report.Service
}

// newLoginHandler creates a new login handler
func newLoginHandler(encoder *encoder, rs report.Service) *loginHandler {
	return &loginHandler{
		encoder: encoder,
		rs:      rs,
	}
}

func (h *loginHandler) Routes(r chi.Router) {
	r.Post("/login", h.login)
}

func (h *loginHandler) login(writer http.ResponseWriter, request *http.Request) {

}
