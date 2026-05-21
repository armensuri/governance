// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package hr

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// HR002SecurityAwarenessTraining — HR-002: Security Awareness Training
// Category: CC1 - Control Environment
type HR002SecurityAwarenessTraining struct{}

func NewHR002SecurityAwarenessTraining() *HR002SecurityAwarenessTraining {
	return &HR002SecurityAwarenessTraining{}
}

func (c *HR002SecurityAwarenessTraining) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "HR-002",
		Name:         "Security Awareness Training",
		Description:  `All employees complete security awareness training annually`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC1.4", "CC2.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment`,
		Resource:     "process",
		Evidence:     []string{"Training completion records", "Training content", "Phishing simulation results"},
	}
}

func (c *HR002SecurityAwarenessTraining) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for HR-002
	return check.Result{
		CheckID: "HR-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for HR-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
