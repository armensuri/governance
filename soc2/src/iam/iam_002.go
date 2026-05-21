// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package iam

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// IAM002MfaEnabledForIamUsers — IAM-002: MFA Enabled for IAM Users
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: iam-user-mfa-enabled
type IAM002MfaEnabledForIamUsers struct{}

func NewIAM002MfaEnabledForIamUsers() *IAM002MfaEnabledForIamUsers {
	return &IAM002MfaEnabledForIamUsers{}
}

func (c *IAM002MfaEnabledForIamUsers) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "IAM-002",
		Name:          "MFA Enabled for IAM Users",
		Description:   `Ensure MFA is enabled for all IAM users with console access`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.1", "CC6.2"},
		Remediation:   `Enforce MFA via IAM password policy or SCPs`,
		AWSConfigRule: "iam-user-mfa-enabled",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "iam",
	}
}

func (c *IAM002MfaEnabledForIamUsers) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for IAM-002
	return check.Result{
		CheckID: "IAM-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for IAM-002",
	}
}
