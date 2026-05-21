// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package hr

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// HR003CodeOfConduct — HR-003: Code of Conduct
// Category: CC1 - Control Environment
type HR003CodeOfConduct struct{}

func NewHR003CodeOfConduct() *HR003CodeOfConduct {
	return &HR003CodeOfConduct{}
}

func (c *HR003CodeOfConduct) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "HR-003",
		Name:         "Code of Conduct",
		Description:  `All employees review and sign a Code of Conduct upon hire and annually`,
		Severity:     check.SeverityMedium,
		SOC2Criteria: []string{"CC1.1", "CC1.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment`,
		Resource:     "process",
		Evidence:     []string{"Signed Code of Conduct records", "Onboarding documentation"},
	}
}

func (c *HR003CodeOfConduct) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for HR-003
	return check.Result{
		CheckID: "HR-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for HR-003",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
