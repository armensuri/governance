// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package risk

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RISK003VulnerabilityManagementProgram — RISK-003: Vulnerability Management Program
// Category: CC3 - Risk Assessment / CC9 - Risk Mitigation
type RISK003VulnerabilityManagementProgram struct{}

func NewRISK003VulnerabilityManagementProgram() *RISK003VulnerabilityManagementProgram {
	return &RISK003VulnerabilityManagementProgram{}
}

func (c *RISK003VulnerabilityManagementProgram) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "RISK-003",
		Name:         "Vulnerability Management Program",
		Description:  `Regular vulnerability scanning and penetration testing are conducted`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC7.1", "CC7.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC3 - Risk Assessment / CC9 - Risk Mitigation`,
		Resource:     "process",
		Evidence:     []string{"Vulnerability scan reports", "Penetration test reports", "Remediation tracking"},
	}
}

func (c *RISK003VulnerabilityManagementProgram) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for RISK-003
	return check.Result{
		CheckID: "RISK-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for RISK-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
