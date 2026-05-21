// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package ec2

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// EC2002SecurityGroupsRestrictInbound — EC2-002: Security Groups Restrict Inbound
// Category: CC6 - Logical and Physical Access Controls / CC7 - System Operations
// AWS Config rule: restricted-ssh
type EC2002SecurityGroupsRestrictInbound struct{}

func NewEC2002SecurityGroupsRestrictInbound() *EC2002SecurityGroupsRestrictInbound {
	return &EC2002SecurityGroupsRestrictInbound{}
}

func (c *EC2002SecurityGroupsRestrictInbound) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "EC2-002",
		Name:          "Security Groups Restrict Inbound",
		Description:   `Ensure no security group allows unrestricted inbound access (0.0.0.0/0) on sensitive ports`,
		Severity:      check.SeverityCritical,
		SOC2Criteria:  []string{"CC6.1", "CC6.6"},
		Remediation:   `Restrict security group inbound rules to known CIDRs`,
		AWSConfigRule: "restricted-ssh",
		Category:      `CC6 - Logical and Physical Access Controls / CC7 - System Operations`,
		Resource:      "ec2",
	}
}

func (c *EC2002SecurityGroupsRestrictInbound) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for EC2-002
	return check.Result{
		CheckID: "EC2-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for EC2-002",
	}
}
