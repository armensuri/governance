// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package dataprotect

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// DATA005EbsEncryptionEnabledForCdeInstances — DATA-005: EBS Encryption Enabled for CDE Instances
// Category: Requirement 3 - Protect Stored Account Data
// PCI DSS: 3.5.1 (v4.0)
type DATA005EbsEncryptionEnabledForCdeInstances struct{}

func NewDATA005EbsEncryptionEnabledForCdeInstances() *DATA005EbsEncryptionEnabledForCdeInstances {
	return &DATA005EbsEncryptionEnabledForCdeInstances{}
}

func (c *DATA005EbsEncryptionEnabledForCdeInstances) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "DATA-005",
		Name:              "EBS Encryption Enabled for CDE Instances",
		Description:       `All EBS volumes attached to CDE EC2 instances must be encrypted at rest`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "3.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable EBS default encryption at account level; use KMS CMK`,
		AWSConfigRule:     "ec2-ebs-encryption-by-default",
		Category:          `Requirement 3 - Protect Stored Account Data`,
		Resource:          "dataprotect",
	}
}

func (c *DATA005EbsEncryptionEnabledForCdeInstances) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for DATA-005
	return check.Result{
		CheckID: "DATA-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for DATA-005",
	}
}
