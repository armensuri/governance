// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package iam

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// IAM001MfaEnabledForRootAccount — IAM-001: MFA Enabled for Root Account
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: root-account-mfa-enabled
type IAM001MfaEnabledForRootAccount struct{}

func NewIAM001MfaEnabledForRootAccount() *IAM001MfaEnabledForRootAccount {
	return &IAM001MfaEnabledForRootAccount{}
}

func (c *IAM001MfaEnabledForRootAccount) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "IAM-001",
		Name:          "MFA Enabled for Root Account",
		Description:   `Ensure MFA is enabled on the root AWS account`,
		Severity:      check.SeverityCritical,
		SOC2Criteria:  []string{"CC6.1", "CC6.2"},
		Remediation:   `Enable MFA on root account via AWS IAM console`,
		AWSConfigRule: "root-account-mfa-enabled",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "iam",
	}
}

func (c *IAM001MfaEnabledForRootAccount) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for IAM-001
	return check.Result{
		CheckID: "IAM-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for IAM-001",
	}
}
