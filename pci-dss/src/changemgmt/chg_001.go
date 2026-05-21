// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package changemgmt

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// CHG001FormalChangeManagementProcessForCde — CHG-001: Formal Change Management Process for CDE
// Category: Requirement 6 - Change Management
type CHG001FormalChangeManagementProcessForCde struct{}

func NewCHG001FormalChangeManagementProcessForCde() *CHG001FormalChangeManagementProcessForCde {
	return &CHG001FormalChangeManagementProcessForCde{}
}

func (c *CHG001FormalChangeManagementProcessForCde) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "CHG-001",
		Name:              "Formal Change Management Process for CDE",
		Description:       `All changes to CDE systems must go through a formal approval process before implementation`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "6.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 6 - Change Management`,
		Resource:          "process",
		Evidence:          []string{"Change management policy", "Change request tickets with approvals", "CAB meeting records"},
	}
}

func (c *CHG001FormalChangeManagementProcessForCde) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "CHG-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for CHG-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
