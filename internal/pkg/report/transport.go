package report

import "github.com/strangnet/csp-uri-report/internal/pkg/domain"

// CreateReportData is used when decoding and transforming the report data in the api
type CreateReportData struct {
	BlockedURI         string `json:"blocked-uri"`
	Disposition        string `json:"disposition"`
	DocumentURI        string `json:"document-uri"`
	EffectiveDirective string `json:"effective-directive"`
	OriginalPolicy     string `json:"original-policy"`
	Referrer           string `json:"referrer"`
	ViolatedDirective  string `json:"violated-directive"`
	StatusCode         int    `json:"status-code"`
	ScriptSample       string `json:"script-sample"`
	UserAgent          string `json:"user-agent"`
}

// CreateReportResult is returned from the report service
type CreateReportResult struct {
	*domain.Report
}
