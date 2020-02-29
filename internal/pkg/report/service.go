package report

import (
	"context"

	"github.com/google/uuid"
	"github.com/strangnet/csp-uri-report/internal/pkg/domain"
)

// Service is the interface defining a report service
type Service interface {
	Create(ctx context.Context, data CreateReportData) (*CreateReportResult, error)
	GetByID(id uuid.UUID) (*domain.Report, error)
}

type service struct {
	reports domain.ReportRepository
}

// NewService creates a new Report Service object
func NewService(reports domain.ReportRepository) Service {
	return &service{
		reports: reports,
	}
}

func (s *service) Create(ctx context.Context, data CreateReportData) (*CreateReportResult, error) {
	var report *domain.Report

	report = domain.NewReport()
	report.UserAgent = data.UserAgent
	report.BlockedURI = data.BlockedURI
	report.Disposition = data.Disposition
	report.DocumentURI = data.DocumentURI
	report.EffectiveDirective = data.EffectiveDirective
	report.OriginalPolicy = data.OriginalPolicy
	report.Referrer = data.Referrer
	report.ScriptSample = data.ScriptSample
	report.StatusCode = data.StatusCode
	report.ViolatedDirective = data.ViolatedDirective

	if err := s.reports.Create(report); err != nil {
		return nil, err
	}

	return &CreateReportResult{Report: report}, nil
}

func (s *service) GetByID(id uuid.UUID) (*domain.Report, error) {
	return s.reports.GetByID(id)
}
