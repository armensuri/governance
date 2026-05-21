// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package secureconfig

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// CFG005Ec2Imdsv2Enforced — CFG-005: EC2 IMDSv2 Enforced
// Category: Requirement 2 - Secure Configurations
// PCI DSS: 2.2.7 (v4.0)
type CFG005Ec2Imdsv2Enforced struct{}

func NewCFG005Ec2Imdsv2Enforced() *CFG005Ec2Imdsv2Enforced {
	return &CFG005Ec2Imdsv2Enforced{}
}

func (c *CFG005Ec2Imdsv2Enforced) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "CFG-005",
		Name:              "EC2 IMDSv2 Enforced",
		Description:       `Require IMDSv2 on all EC2 instances in the CDE to prevent SSRF-based credential theft`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "2.2.7",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Set HttpTokens to 'required' on all CDE EC2 instance metadata configurations`,
		AWSConfigRule:     "ec2-imdsv2-check",
		Category:          `Requirement 2 - Secure Configurations`,
		Resource:          "secureconfig",
	}
}

func (c *CFG005Ec2Imdsv2Enforced) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CFG-005
	return check.Result{
		CheckID: "CFG-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CFG-005",
	}
}
