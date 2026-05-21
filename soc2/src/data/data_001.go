// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package data

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// DATA001DataClassificationPolicy — DATA-001: Data Classification Policy
// Category: C1 - Confidentiality / P - Privacy
type DATA001DataClassificationPolicy struct{}

func NewDATA001DataClassificationPolicy() *DATA001DataClassificationPolicy {
	return &DATA001DataClassificationPolicy{}
}

func (c *DATA001DataClassificationPolicy) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "DATA-001",
		Name:         "Data Classification Policy",
		Description:  `Data is classified by sensitivity (Public, Internal, Confidential, Restricted)`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"C1.1", "CC6.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `C1 - Confidentiality / P - Privacy`,
		Resource:     "process",
		Evidence:     []string{"Data classification policy", "Data inventory/map"},
	}
}

func (c *DATA001DataClassificationPolicy) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for DATA-001
	return check.Result{
		CheckID: "DATA-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for DATA-001",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
