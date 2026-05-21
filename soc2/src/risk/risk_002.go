// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package risk

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RISK002VendorRiskManagement — RISK-002: Vendor Risk Management
// Category: CC3 - Risk Assessment / CC9 - Risk Mitigation
type RISK002VendorRiskManagement struct{}

func NewRISK002VendorRiskManagement() *RISK002VendorRiskManagement {
	return &RISK002VendorRiskManagement{}
}

func (c *RISK002VendorRiskManagement) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "RISK-002",
		Name:         "Vendor Risk Management",
		Description:  `Third-party vendors handling company data are assessed for security controls`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC9.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC3 - Risk Assessment / CC9 - Risk Mitigation`,
		Resource:     "process",
		Evidence:     []string{"Vendor security questionnaires", "SOC 2 reports from vendors", "Vendor inventory"},
	}
}

func (c *RISK002VendorRiskManagement) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for RISK-002
	return check.Result{
		CheckID: "RISK-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for RISK-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
