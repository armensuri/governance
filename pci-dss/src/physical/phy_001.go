// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package physical

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// PHY001PhysicalAccessToCdeSystemsRestricted — PHY-001: Physical Access to CDE Systems Restricted
// Category: Requirement 9 - Physical Access Controls
type PHY001PhysicalAccessToCdeSystemsRestricted struct{}

func NewPHY001PhysicalAccessToCdeSystemsRestricted() *PHY001PhysicalAccessToCdeSystemsRestricted {
	return &PHY001PhysicalAccessToCdeSystemsRestricted{}
}

func (c *PHY001PhysicalAccessToCdeSystemsRestricted) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "PHY-001",
		Name:              "Physical Access to CDE Systems Restricted",
		Description:       `Physical access to any on-premises CDE components must be restricted to authorized personnel`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "9.1.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 9 - Physical Access Controls`,
		Resource:          "process",
		Evidence:          []string{"Badge access logs", "Visitor logs", "Physical access control policy"},
	}
}

func (c *PHY001PhysicalAccessToCdeSystemsRestricted) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "PHY-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for PHY-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
