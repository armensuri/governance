// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package network

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// NET003SecurityGroupsDenyAllByDefault — NET-003: Security Groups Deny All by Default
// Category: Requirement 1 - Network Security Controls
// PCI DSS: 1.3.1 (v4.0)
type NET003SecurityGroupsDenyAllByDefault struct{}

func NewNET003SecurityGroupsDenyAllByDefault() *NET003SecurityGroupsDenyAllByDefault {
	return &NET003SecurityGroupsDenyAllByDefault{}
}

func (c *NET003SecurityGroupsDenyAllByDefault) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "NET-003",
		Name:              "Security Groups Deny All by Default",
		Description:       `All security groups deny inbound traffic by default; only explicitly required traffic is allowed`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "1.3.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Audit all security groups; remove any overly permissive rules`,
		AWSConfigRule:     "vpc-default-security-group-closed",
		Category:          `Requirement 1 - Network Security Controls`,
		Resource:          "network",
	}
}

func (c *NET003SecurityGroupsDenyAllByDefault) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for NET-003
	return check.Result{
		CheckID: "NET-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for NET-003",
	}
}
