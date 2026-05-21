// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package monitoring

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// MON002SecurityAlertingAndResponse — MON-002: Security Alerting and Response
// Category: CC4 - Monitoring / CC7 - System Operations
type MON002SecurityAlertingAndResponse struct{}

func NewMON002SecurityAlertingAndResponse() *MON002SecurityAlertingAndResponse {
	return &MON002SecurityAlertingAndResponse{}
}

func (c *MON002SecurityAlertingAndResponse) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "MON-002",
		Name:         "Security Alerting and Response",
		Description:  `Alerts are configured for security events and reviewed by the security team`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC7.3", "CC7.4"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC4 - Monitoring / CC7 - System Operations`,
		Resource:     "process",
		Evidence:     []string{"Alert configuration records", "On-call runbooks", "Incident tickets"},
	}
}

func (c *MON002SecurityAlertingAndResponse) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for MON-002
	return check.Result{
		CheckID: "MON-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for MON-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
