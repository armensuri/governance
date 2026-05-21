// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package iam

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// IAM005NoInlineIamPolicies — IAM-005: No Inline IAM Policies
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: iam-no-inline-policy-check
type IAM005NoInlineIamPolicies struct{}

func NewIAM005NoInlineIamPolicies() *IAM005NoInlineIamPolicies {
	return &IAM005NoInlineIamPolicies{}
}

func (c *IAM005NoInlineIamPolicies) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "IAM-005",
		Name:          "No Inline IAM Policies",
		Description:   `Avoid use of inline IAM policies; use managed policies instead`,
		Severity:      check.SeverityMedium,
		SOC2Criteria:  []string{"CC6.3"},
		Remediation:   `Convert inline policies to managed policies`,
		AWSConfigRule: "iam-no-inline-policy-check",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "iam",
	}
}

func (c *IAM005NoInlineIamPolicies) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for IAM-005
	return check.Result{
		CheckID: "IAM-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for IAM-005",
	}
}
