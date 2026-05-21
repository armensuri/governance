// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package incident

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// IR002IrpTestedAnnually — IR-002: IRP Tested Annually
// Category: Requirement 12 - Incident Response
type IR002IrpTestedAnnually struct{}

func NewIR002IrpTestedAnnually() *IR002IrpTestedAnnually {
	return &IR002IrpTestedAnnually{}
}

func (c *IR002IrpTestedAnnually) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "IR-002",
		Name:              "IRP Tested Annually",
		Description:       `The incident response plan must be tested at least once per year`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.10.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Incident Response`,
		Resource:          "process",
		Evidence:          []string{"Tabletop exercise records", "Test results", "Lessons learned documentation"},
	}
}

func (c *IR002IrpTestedAnnually) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "IR-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for IR-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
