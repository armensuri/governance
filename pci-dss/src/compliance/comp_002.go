// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package compliance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// COMP002CompensatingControlsDocumented — COMP-002: Compensating Controls Documented
// Category: Requirement 12 - Compliance Validation
type COMP002CompensatingControlsDocumented struct{}

func NewCOMP002CompensatingControlsDocumented() *COMP002CompensatingControlsDocumented {
	return &COMP002CompensatingControlsDocumented{}
}

func (c *COMP002CompensatingControlsDocumented) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "COMP-002",
		Name:              "Compensating Controls Documented",
		Description:       `Any compensating controls used in lieu of PCI DSS requirements must be formally documented and approved`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.4.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Compliance Validation`,
		Resource:          "process",
		Evidence:          []string{"Compensating controls worksheet", "QSA approval", "Risk analysis supporting the control"},
	}
}

func (c *COMP002CompensatingControlsDocumented) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "COMP-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for COMP-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
