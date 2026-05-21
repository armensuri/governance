// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package sdlc

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// SDLC003ProductionDeploymentApprovals — SDLC-003: Production Deployment Approvals
// Category: CC8 - Change Management
type SDLC003ProductionDeploymentApprovals struct{}

func NewSDLC003ProductionDeploymentApprovals() *SDLC003ProductionDeploymentApprovals {
	return &SDLC003ProductionDeploymentApprovals{}
}

func (c *SDLC003ProductionDeploymentApprovals) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "SDLC-003",
		Name:         "Production Deployment Approvals",
		Description:  `All production deployments require peer review and approval`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC8.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC8 - Change Management`,
		Resource:     "process",
		Evidence:     []string{"CI/CD pipeline configuration", "Deployment approval records"},
	}
}

func (c *SDLC003ProductionDeploymentApprovals) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for SDLC-003
	return check.Result{
		CheckID: "SDLC-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for SDLC-003",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
