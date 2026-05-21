// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package kms

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// KMS001KmsKeyRotationEnabled — KMS-001: KMS Key Rotation Enabled
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: cmk-backing-key-rotation-enabled
type KMS001KmsKeyRotationEnabled struct{}

func NewKMS001KmsKeyRotationEnabled() *KMS001KmsKeyRotationEnabled {
	return &KMS001KmsKeyRotationEnabled{}
}

func (c *KMS001KmsKeyRotationEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "KMS-001",
		Name:          "KMS Key Rotation Enabled",
		Description:   `Ensure automatic annual key rotation is enabled for KMS CMKs`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.1", "CC6.7"},
		Remediation:   `Enable key rotation on all applicable KMS CMKs`,
		AWSConfigRule: "cmk-backing-key-rotation-enabled",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "kms",
	}
}

func (c *KMS001KmsKeyRotationEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for KMS-001
	return check.Result{
		CheckID: "KMS-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for KMS-001",
	}
}
