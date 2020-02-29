package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/strangnet/csp-uri-report/internal/pkg/report"
)

type reportHandler struct {
	encoder *encoder
	rs      report.Service
}

func newReportHandler(encoder *encoder, rs report.Service) *reportHandler {
	return &reportHandler{
		encoder: encoder,
		rs:      rs,
	}
}

func (h *reportHandler) Routes(r chi.Router) {
	r.Post("/", h.postReport)
	r.Get("/{id}", h.getReport)
	r.Get("/", h.listReports)
}

func (h *reportHandler) postReport(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var data map[string]report.CreateReportData

	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		h.encoder.Error(ctx, writer, err)
		return
	}

	d := data["csp-report"]
	d.UserAgent = request.Header.Get("User-Agent")
	result, err := h.rs.Create(ctx, d)

	if err != nil {
		h.encoder.Error(ctx, writer, err)
		return
	}

	h.encoder.StatusResponse(writer, result, http.StatusCreated)
}

func (h *reportHandler) getReport(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	rawID := chi.URLParam(request, "id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		h.encoder.Error(ctx, writer, err)
		return
	}

	report, err := h.rs.GetByID(id)

	h.encoder.Response(writer, report)
}

func (h *reportHandler) listReports(writer http.ResponseWriter, request *http.Request) {

}
