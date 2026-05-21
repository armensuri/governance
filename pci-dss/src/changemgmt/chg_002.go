// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package changemgmt

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// CHG002SecurityImpactAssessmentForCdeChanges — CHG-002: Security Impact Assessment for CDE Changes
// Category: Requirement 6 - Change Management
type CHG002SecurityImpactAssessmentForCdeChanges struct{}

func NewCHG002SecurityImpactAssessmentForCdeChanges() *CHG002SecurityImpactAssessmentForCdeChanges {
	return &CHG002SecurityImpactAssessmentForCdeChanges{}
}

func (c *CHG002SecurityImpactAssessmentForCdeChanges) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "CHG-002",
		Name:              "Security Impact Assessment for CDE Changes",
		Description:       `Security impact of CDE changes must be assessed and documented before implementation`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "6.5.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 6 - Change Management`,
		Resource:          "process",
		Evidence:          []string{"Security impact assessment records", "Change tickets with security review sign-off"},
	}
}

func (c *CHG002SecurityImpactAssessmentForCdeChanges) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "CHG-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for CHG-002",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
