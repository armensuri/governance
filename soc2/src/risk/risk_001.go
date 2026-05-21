// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package risk

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RISK001AnnualRiskAssessment — RISK-001: Annual Risk Assessment
// Category: CC3 - Risk Assessment / CC9 - Risk Mitigation
type RISK001AnnualRiskAssessment struct{}

func NewRISK001AnnualRiskAssessment() *RISK001AnnualRiskAssessment {
	return &RISK001AnnualRiskAssessment{}
}

func (c *RISK001AnnualRiskAssessment) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "RISK-001",
		Name:         "Annual Risk Assessment",
		Description:  `Company performs a formal risk assessment at least annually`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC3.1", "CC3.2", "CC3.3"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC3 - Risk Assessment / CC9 - Risk Mitigation`,
		Resource:     "process",
		Evidence:     []string{"Risk assessment report", "Risk register", "Management sign-off"},
	}
}

func (c *RISK001AnnualRiskAssessment) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for RISK-001
	return check.Result{
		CheckID: "RISK-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for RISK-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
