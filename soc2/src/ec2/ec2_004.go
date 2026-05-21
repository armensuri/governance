// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package ec2

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// EC2004Ec2InstancesPatched — EC2-004: EC2 Instances Patched
// Category: CC6 - Logical and Physical Access Controls / CC7 - System Operations
// AWS Config rule: ec2-managedinstance-patch-compliance-status-check
type EC2004Ec2InstancesPatched struct{}

func NewEC2004Ec2InstancesPatched() *EC2004Ec2InstancesPatched {
	return &EC2004Ec2InstancesPatched{}
}

func (c *EC2004Ec2InstancesPatched) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "EC2-004",
		Name:          "EC2 Instances Patched",
		Description:   `Ensure EC2 instances are regularly patched using AWS Systems Manager Patch Manager`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC7.1", "CC7.2"},
		Remediation:   `Configure SSM Patch Manager with a patch baseline`,
		AWSConfigRule: "ec2-managedinstance-patch-compliance-status-check",
		Category:      `CC6 - Logical and Physical Access Controls / CC7 - System Operations`,
		Resource:      "ec2",
	}
}

func (c *EC2004Ec2InstancesPatched) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for EC2-004
	return check.Result{
		CheckID: "EC2-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for EC2-004",
	}
}
