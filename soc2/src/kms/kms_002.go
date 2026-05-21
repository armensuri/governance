// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package kms

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// KMS002KmsKeyPoliciesRestrictAccess — KMS-002: KMS Key Policies Restrict Access
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: N/A
type KMS002KmsKeyPoliciesRestrictAccess struct{}

func NewKMS002KmsKeyPoliciesRestrictAccess() *KMS002KmsKeyPoliciesRestrictAccess {
	return &KMS002KmsKeyPoliciesRestrictAccess{}
}

func (c *KMS002KmsKeyPoliciesRestrictAccess) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "KMS-002",
		Name:          "KMS Key Policies Restrict Access",
		Description:   `Ensure KMS key policies do not allow wildcard principal access`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.1", "CC6.3"},
		Remediation:   `Review and restrict KMS key policies to specific principals`,
		AWSConfigRule: "",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "kms",
	}
}

func (c *KMS002KmsKeyPoliciesRestrictAccess) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for KMS-002
	return check.Result{
		CheckID: "KMS-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for KMS-002",
	}
}
