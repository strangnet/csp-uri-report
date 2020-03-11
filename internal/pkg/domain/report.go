package domain

import (
	"time"

	"github.com/google/uuid"
)

// Report is a struct
type Report struct {
	ID                 uuid.UUID `json:"id"`
	CreatedAt          time.Time `json:"createdAt"`
	UserAgent          string    `json:"userAgent"`
	BlockedURI         string    `json:"blockedUri"`
	Disposition        string    `json:"disposition"`
	DocumentURI        string    `json:"documentUri"`
	EffectiveDirective string    `json:"effectiveDirective"`
	OriginalPolicy     string    `json:"originalPolicy"`
	Referrer           string    `json:"referrer"`
	ViolatedDirective  string    `json:"violatedDirective"`
	StatusCode         int       `json:"statusCode"`
	ScriptSample       string    `json:"scriptSample"`
}

// ReportRepository is a repository interface
type ReportRepository interface {
	Create(report *Report) error
	GetByID(id uuid.UUID) (*Report, error)
	ListReports() ([]*Report, error)
}

// NewReport creates a new report
func NewReport() *Report {
	return &Report{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
