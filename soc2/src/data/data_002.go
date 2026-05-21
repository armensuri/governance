// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package data

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// DATA002DataRetentionAndDisposalPolicy — DATA-002: Data Retention and Disposal Policy
// Category: C1 - Confidentiality / P - Privacy
type DATA002DataRetentionAndDisposalPolicy struct{}

func NewDATA002DataRetentionAndDisposalPolicy() *DATA002DataRetentionAndDisposalPolicy {
	return &DATA002DataRetentionAndDisposalPolicy{}
}

func (c *DATA002DataRetentionAndDisposalPolicy) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "DATA-002",
		Name:         "Data Retention and Disposal Policy",
		Description:  `Formal policies govern data retention periods and secure disposal procedures`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"C1.2", "P4.3"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `C1 - Confidentiality / P - Privacy`,
		Resource:     "process",
		Evidence:     []string{"Retention policy", "Disposal records", "S3 lifecycle policies"},
	}
}

func (c *DATA002DataRetentionAndDisposalPolicy) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for DATA-002
	return check.Result{
		CheckID: "DATA-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for DATA-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
