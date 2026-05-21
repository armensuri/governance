// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package ec2

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// EC2005Ec2Imdsv2Enforced — EC2-005: EC2 IMDSv2 Enforced
// Category: CC6 - Logical and Physical Access Controls / CC7 - System Operations
// AWS Config rule: ec2-imdsv2-check
type EC2005Ec2Imdsv2Enforced struct{}

func NewEC2005Ec2Imdsv2Enforced() *EC2005Ec2Imdsv2Enforced {
	return &EC2005Ec2Imdsv2Enforced{}
}

func (c *EC2005Ec2Imdsv2Enforced) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "EC2-005",
		Name:          "EC2 IMDSv2 Enforced",
		Description:   `Require IMDSv2 on all EC2 instances to prevent SSRF attacks`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.1", "CC6.6"},
		Remediation:   `Set HttpTokens to 'required' on all EC2 instances`,
		AWSConfigRule: "ec2-imdsv2-check",
		Category:      `CC6 - Logical and Physical Access Controls / CC7 - System Operations`,
		Resource:      "ec2",
	}
}

func (c *EC2005Ec2Imdsv2Enforced) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for EC2-005
	return check.Result{
		CheckID: "EC2-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for EC2-005",
	}
}
