// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package compliance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// COMP003ExecutiveResponsibilityForPciDssCompliance — COMP-003: Executive Responsibility for PCI DSS Compliance
// Category: Requirement 12 - Compliance Validation
type COMP003ExecutiveResponsibilityForPciDssCompliance struct{}

func NewCOMP003ExecutiveResponsibilityForPciDssCompliance() *COMP003ExecutiveResponsibilityForPciDssCompliance {
	return &COMP003ExecutiveResponsibilityForPciDssCompliance{}
}

func (c *COMP003ExecutiveResponsibilityForPciDssCompliance) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "COMP-003",
		Name:              "Executive Responsibility for PCI DSS Compliance",
		Description:       `A designated executive must be responsible for overall PCI DSS compliance and program management`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.1.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Compliance Validation`,
		Resource:          "process",
		Evidence:          []string{"Executive designation letter", "Compliance charter", "Board-level reporting records"},
	}
}

func (c *COMP003ExecutiveResponsibilityForPciDssCompliance) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "COMP-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for COMP-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
