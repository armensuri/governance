// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package awareness

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// SAT001AnnualPciDssSecurityAwarenessTraining — SAT-001: Annual PCI DSS Security Awareness Training
// Category: Requirement 12 - Security Awareness Training
type SAT001AnnualPciDssSecurityAwarenessTraining struct{}

func NewSAT001AnnualPciDssSecurityAwarenessTraining() *SAT001AnnualPciDssSecurityAwarenessTraining {
	return &SAT001AnnualPciDssSecurityAwarenessTraining{}
}

func (c *SAT001AnnualPciDssSecurityAwarenessTraining) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "SAT-001",
		Name:              "Annual PCI DSS Security Awareness Training",
		Description:       `All personnel with access to the CDE must complete security awareness training upon hire and annually`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.6.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Security Awareness Training`,
		Resource:          "process",
		Evidence:          []string{"Training completion records", "Training curriculum", "Training dates and acknowledgments"},
	}
}

func (c *SAT001AnnualPciDssSecurityAwarenessTraining) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "SAT-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for SAT-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
