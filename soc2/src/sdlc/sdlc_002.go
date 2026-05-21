// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package sdlc

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// SDLC002SeparationOfEnvironments — SDLC-002: Separation of Environments
// Category: CC8 - Change Management
type SDLC002SeparationOfEnvironments struct{}

func NewSDLC002SeparationOfEnvironments() *SDLC002SeparationOfEnvironments {
	return &SDLC002SeparationOfEnvironments{}
}

func (c *SDLC002SeparationOfEnvironments) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "SDLC-002",
		Name:         "Separation of Environments",
		Description:  `Development, staging, and production environments are logically separated`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC6.1", "CC8.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC8 - Change Management`,
		Resource:     "process",
		Evidence:     []string{"AWS account structure diagram", "Network architecture documentation"},
	}
}

func (c *SDLC002SeparationOfEnvironments) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for SDLC-002
	return check.Result{
		CheckID: "SDLC-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for SDLC-002",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
