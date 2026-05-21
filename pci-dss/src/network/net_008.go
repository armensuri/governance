// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package network

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// NET008NoDirectInternetAccessToCde — NET-008: No Direct Internet Access to CDE
// Category: Requirement 1 - Network Security Controls
// PCI DSS: 1.3.2 (v4.0)
type NET008NoDirectInternetAccessToCde struct{}

func NewNET008NoDirectInternetAccessToCde() *NET008NoDirectInternetAccessToCde {
	return &NET008NoDirectInternetAccessToCde{}
}

func (c *NET008NoDirectInternetAccessToCde) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "NET-008",
		Name:              "No Direct Internet Access to CDE",
		Description:       `CDE systems must not have direct inbound internet access; all internet-facing traffic routes through controlled entry points`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "1.3.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Remove internet gateways from private subnets; use NAT Gateway for outbound-only internet access`,
		AWSConfigRule:     "",
		Category:          `Requirement 1 - Network Security Controls`,
		Resource:          "network",
	}
}

func (c *NET008NoDirectInternetAccessToCde) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for NET-008
	return check.Result{
		CheckID: "NET-008",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for NET-008",
	}
}
