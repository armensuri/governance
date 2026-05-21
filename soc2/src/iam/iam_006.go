// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package iam

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// IAM006IamAccessKeysRotated — IAM-006: IAM Access Keys Rotated
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: access-keys-rotated
type IAM006IamAccessKeysRotated struct{}

func NewIAM006IamAccessKeysRotated() *IAM006IamAccessKeysRotated {
	return &IAM006IamAccessKeysRotated{}
}

func (c *IAM006IamAccessKeysRotated) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "IAM-006",
		Name:          "IAM Access Keys Rotated",
		Description:   `Ensure IAM access keys are rotated every 90 days`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.1", "CC6.2"},
		Remediation:   `Rotate access keys and enforce rotation policy`,
		AWSConfigRule: "access-keys-rotated",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "iam",
	}
}

func (c *IAM006IamAccessKeysRotated) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for IAM-006
	return check.Result{
		CheckID: "IAM-006",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for IAM-006",
	}
}
