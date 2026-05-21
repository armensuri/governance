// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package iam

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// IAM003NoUnusedIamCredentials — IAM-003: No Unused IAM Credentials
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: iam-user-unused-credentials-check
type IAM003NoUnusedIamCredentials struct{}

func NewIAM003NoUnusedIamCredentials() *IAM003NoUnusedIamCredentials {
	return &IAM003NoUnusedIamCredentials{}
}

func (c *IAM003NoUnusedIamCredentials) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "IAM-003",
		Name:          "No Unused IAM Credentials",
		Description:   `Disable or remove IAM credentials unused for 90+ days`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.2", "CC6.3"},
		Remediation:   `Audit IAM credential report and disable stale credentials`,
		AWSConfigRule: "iam-user-unused-credentials-check",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "iam",
	}
}

func (c *IAM003NoUnusedIamCredentials) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for IAM-003
	return check.Result{
		CheckID: "IAM-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for IAM-003",
	}
}
