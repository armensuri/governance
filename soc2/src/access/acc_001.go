// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// ACC001UserAccessReview — ACC-001: User Access Review
// Category: CC6 - Logical and Physical Access Controls
type ACC001UserAccessReview struct{}

func NewACC001UserAccessReview() *ACC001UserAccessReview {
	return &ACC001UserAccessReview{}
}

func (c *ACC001UserAccessReview) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "ACC-001",
		Name:         "User Access Review",
		Description:  `Conduct quarterly access reviews for all systems including AWS`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC6.2", "CC6.3"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC6 - Logical and Physical Access Controls`,
		Resource:     "process",
		Evidence:     []string{"Access review reports", "Attestation sign-offs", "Revocation records"},
	}
}

func (c *ACC001UserAccessReview) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for ACC-001
	return check.Result{
		CheckID: "ACC-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for ACC-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
