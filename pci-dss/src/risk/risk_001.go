// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package risk

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// RISK001AnnualFormalRiskAssessment — RISK-001: Annual Formal Risk Assessment
// Category: Requirement 12 - Risk Management
type RISK001AnnualFormalRiskAssessment struct{}

func NewRISK001AnnualFormalRiskAssessment() *RISK001AnnualFormalRiskAssessment {
	return &RISK001AnnualFormalRiskAssessment{}
}

func (c *RISK001AnnualFormalRiskAssessment) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "RISK-001",
		Name:              "Annual Formal Risk Assessment",
		Description:       `A formal risk assessment that identifies critical assets, threats, and vulnerabilities related to the CDE must be performed annually`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.3.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Risk Management`,
		Resource:          "process",
		Evidence:          []string{"Risk assessment report", "Asset inventory", "Threat/vulnerability analysis", "Risk register"},
	}
}

func (c *RISK001AnnualFormalRiskAssessment) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "RISK-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for RISK-001",
		Details: map[string]string{
			"evidence_count": "4",
		},
	}
}
