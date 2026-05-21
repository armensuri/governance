// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package sdlc

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// SDLC001SecureSdlcPractices — SDLC-001: Secure SDLC Practices
// Category: CC8 - Change Management
type SDLC001SecureSdlcPractices struct{}

func NewSDLC001SecureSdlcPractices() *SDLC001SecureSdlcPractices {
	return &SDLC001SecureSdlcPractices{}
}

func (c *SDLC001SecureSdlcPractices) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "SDLC-001",
		Name:         "Secure SDLC Practices",
		Description:  `Security is integrated into the software development lifecycle (SAST, DAST, code review)`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC8.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC8 - Change Management`,
		Resource:     "process",
		Evidence:     []string{"SAST/DAST scan results", "Code review policy", "PR approval records"},
	}
}

func (c *SDLC001SecureSdlcPractices) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for SDLC-001
	return check.Result{
		CheckID: "SDLC-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for SDLC-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
