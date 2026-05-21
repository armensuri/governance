// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package policies

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// POL002AcceptableUsePolicy — POL-002: Acceptable Use Policy
// Category: CC1 - Control Environment / CC2 - Communication and Information
type POL002AcceptableUsePolicy struct{}

func NewPOL002AcceptableUsePolicy() *POL002AcceptableUsePolicy {
	return &POL002AcceptableUsePolicy{}
}

func (c *POL002AcceptableUsePolicy) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "POL-002",
		Name:         "Acceptable Use Policy",
		Description:  `Employees acknowledge an Acceptable Use Policy for company assets and systems`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC1.4", "CC6.3"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment / CC2 - Communication and Information`,
		Resource:     "process",
		Evidence:     []string{"Signed AUP acknowledgments", "Onboarding records"},
	}
}

func (c *POL002AcceptableUsePolicy) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for POL-002
	return check.Result{
		CheckID: "POL-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for POL-002",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
