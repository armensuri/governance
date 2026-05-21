// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package network

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// NET002SecurityGroupsRestrictInboundToCde — NET-002: Security Groups Restrict Inbound to CDE
// Category: Requirement 1 - Network Security Controls
// PCI DSS: 1.3.2 (v4.0)
type NET002SecurityGroupsRestrictInboundToCde struct{}

func NewNET002SecurityGroupsRestrictInboundToCde() *NET002SecurityGroupsRestrictInboundToCde {
	return &NET002SecurityGroupsRestrictInboundToCde{}
}

func (c *NET002SecurityGroupsRestrictInboundToCde) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "NET-002",
		Name:              "Security Groups Restrict Inbound to CDE",
		Description:       `No security group allows unrestricted inbound access (0.0.0.0/0) to CDE systems on any port`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "1.3.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Restrict all security group inbound rules to known, required CIDRs and ports only`,
		AWSConfigRule:     "restricted-ssh",
		Category:          `Requirement 1 - Network Security Controls`,
		Resource:          "network",
	}
}

func (c *NET002SecurityGroupsRestrictInboundToCde) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for NET-002
	return check.Result{
		CheckID: "NET-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for NET-002",
	}
}
