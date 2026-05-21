// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package policies

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// POL004ChangeManagementPolicy — POL-004: Change Management Policy
// Category: CC1 - Control Environment / CC2 - Communication and Information
type POL004ChangeManagementPolicy struct{}

func NewPOL004ChangeManagementPolicy() *POL004ChangeManagementPolicy {
	return &POL004ChangeManagementPolicy{}
}

func (c *POL004ChangeManagementPolicy) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "POL-004",
		Name:         "Change Management Policy",
		Description:  `Formal change management process with approvals for production changes`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC8.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment / CC2 - Communication and Information`,
		Resource:     "process",
		Evidence:     []string{"Change management policy", "Change request tickets", "Approval records"},
	}
}

func (c *POL004ChangeManagementPolicy) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for POL-004
	return check.Result{
		CheckID: "POL-004",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for POL-004",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
