// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package physical

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// PHY002VisitorManagementProcedures — PHY-002: Visitor Management Procedures
// Category: Requirement 9 - Physical Access Controls
type PHY002VisitorManagementProcedures struct{}

func NewPHY002VisitorManagementProcedures() *PHY002VisitorManagementProcedures {
	return &PHY002VisitorManagementProcedures{}
}

func (c *PHY002VisitorManagementProcedures) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "PHY-002",
		Name:              "Visitor Management Procedures",
		Description:       `Visitors accessing areas with CDE systems must be logged, badged, and escorted`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "9.2.3",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 9 - Physical Access Controls`,
		Resource:          "process",
		Evidence:          []string{"Visitor log", "Visitor badge policy", "Escort procedures"},
	}
}

func (c *PHY002VisitorManagementProcedures) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "PHY-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for PHY-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
