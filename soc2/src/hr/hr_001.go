// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package hr

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// HR001BackgroundChecks — HR-001: Background Checks
// Category: CC1 - Control Environment
type HR001BackgroundChecks struct{}

func NewHR001BackgroundChecks() *HR001BackgroundChecks {
	return &HR001BackgroundChecks{}
}

func (c *HR001BackgroundChecks) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "HR-001",
		Name:         "Background Checks",
		Description:  `Background checks are performed for all employees prior to system access`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC1.4"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment`,
		Resource:     "process",
		Evidence:     []string{"Background check policy", "HR records showing completion"},
	}
}

func (c *HR001BackgroundChecks) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for HR-001
	return check.Result{
		CheckID: "HR-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for HR-001",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
