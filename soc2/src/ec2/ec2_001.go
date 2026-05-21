// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package ec2

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// EC2001NoDefaultVpcUsage — EC2-001: No Default VPC Usage
// Category: CC6 - Logical and Physical Access Controls / CC7 - System Operations
// AWS Config rule: vpc-default-security-group-closed
type EC2001NoDefaultVpcUsage struct{}

func NewEC2001NoDefaultVpcUsage() *EC2001NoDefaultVpcUsage {
	return &EC2001NoDefaultVpcUsage{}
}

func (c *EC2001NoDefaultVpcUsage) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "EC2-001",
		Name:          "No Default VPC Usage",
		Description:   `Avoid using the default VPC for production workloads`,
		Severity:      check.SeverityMedium,
		SOC2Criteria:  []string{"CC6.1", "CC6.6"},
		Remediation:   `Create a custom VPC with proper network segmentation`,
		AWSConfigRule: "vpc-default-security-group-closed",
		Category:      `CC6 - Logical and Physical Access Controls / CC7 - System Operations`,
		Resource:      "ec2",
	}
}

func (c *EC2001NoDefaultVpcUsage) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for EC2-001
	return check.Result{
		CheckID: "EC2-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for EC2-001",
	}
}
