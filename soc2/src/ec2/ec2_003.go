// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package ec2

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// EC2003EbsVolumeEncryption — EC2-003: EBS Volume Encryption
// Category: CC6 - Logical and Physical Access Controls / CC7 - System Operations
// AWS Config rule: ec2-ebs-encryption-by-default
type EC2003EbsVolumeEncryption struct{}

func NewEC2003EbsVolumeEncryption() *EC2003EbsVolumeEncryption {
	return &EC2003EbsVolumeEncryption{}
}

func (c *EC2003EbsVolumeEncryption) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "EC2-003",
		Name:          "EBS Volume Encryption",
		Description:   `Ensure all EBS volumes are encrypted at rest`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.7", "C1.1"},
		Remediation:   `Enable EBS default encryption at the account level`,
		AWSConfigRule: "ec2-ebs-encryption-by-default",
		Category:      `CC6 - Logical and Physical Access Controls / CC7 - System Operations`,
		Resource:      "ec2",
	}
}

func (c *EC2003EbsVolumeEncryption) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for EC2-003
	return check.Result{
		CheckID: "EC2-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for EC2-003",
	}
}
