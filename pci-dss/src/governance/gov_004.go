// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package governance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// GOV004PciDssResponsibilityMatrixDefined — GOV-004: PCI DSS Responsibility Matrix Defined
// Category: Requirement 12 - Information Security Policy
type GOV004PciDssResponsibilityMatrixDefined struct{}

func NewGOV004PciDssResponsibilityMatrixDefined() *GOV004PciDssResponsibilityMatrixDefined {
	return &GOV004PciDssResponsibilityMatrixDefined{}
}

func (c *GOV004PciDssResponsibilityMatrixDefined) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "GOV-004",
		Name:              "PCI DSS Responsibility Matrix Defined",
		Description:       `A responsibility matrix defining PCI DSS responsibility between AWS and the organization must be maintained`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.1.3",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Information Security Policy`,
		Resource:          "process",
		Evidence:          []string{"AWS Shared Responsibility documentation", "RACI matrix", "AWS PCI DSS responsibility summary"},
	}
}

func (c *GOV004PciDssResponsibilityMatrixDefined) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "GOV-004",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for GOV-004",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
