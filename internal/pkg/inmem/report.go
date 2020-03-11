package inmem

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/strangnet/csp-uri-report/internal/pkg/domain"
)

// ReportRepository is an inmem repository
type ReportRepository struct {
	reports map[uuid.UUID]*domain.Report
	mu      sync.RWMutex
}

// NewReportRepository creates a new repository object
func NewReportRepository() domain.ReportRepository {
	return &ReportRepository{
		reports: make(map[uuid.UUID]*domain.Report),
	}
}

// Create stores a new report
func (r *ReportRepository) Create(report *domain.Report) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.reports[report.ID] = report

	fmt.Printf("%#v", report)

	return nil
}

// GetByID gets a report by its ID
func (r *ReportRepository) GetByID(id uuid.UUID) (*domain.Report, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	report, ok := r.reports[id]
	if !ok {
		return nil, errors.New("Report not Found")
	}

	return report, nil
}

// ListReports returns a list of all reports
func (r *ReportRepository) ListReports() ([]*domain.Report, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var reports []*domain.Report

	for _, report := range r.reports {
		reports = append(reports, report)
	}

	return reports, nil
}
