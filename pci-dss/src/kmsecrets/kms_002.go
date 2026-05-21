// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package kmsecrets

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// KMS002KmsKeyRotationEnabled — KMS-002: KMS Key Rotation Enabled
// Category: Requirement 3 & 8 - Key and Secret Management
// PCI DSS: 3.7.4 (v4.0)
type KMS002KmsKeyRotationEnabled struct{}

func NewKMS002KmsKeyRotationEnabled() *KMS002KmsKeyRotationEnabled {
	return &KMS002KmsKeyRotationEnabled{}
}

func (c *KMS002KmsKeyRotationEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "KMS-002",
		Name:              "KMS Key Rotation Enabled",
		Description:       `Automatic annual key rotation must be enabled for all KMS CMKs used in the CDE`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "3.7.4",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable automatic key rotation on all CDE KMS CMKs`,
		AWSConfigRule:     "cmk-backing-key-rotation-enabled",
		Category:          `Requirement 3 & 8 - Key and Secret Management`,
		Resource:          "kmsecrets",
	}
}

func (c *KMS002KmsKeyRotationEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for KMS-002
	return check.Result{
		CheckID: "KMS-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for KMS-002",
	}
}
