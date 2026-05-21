// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package network

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// NET005CdeResourcesInPrivateSubnets — NET-005: CDE Resources in Private Subnets
// Category: Requirement 1 - Network Security Controls
// PCI DSS: 1.3.2 (v4.0)
type NET005CdeResourcesInPrivateSubnets struct{}

func NewNET005CdeResourcesInPrivateSubnets() *NET005CdeResourcesInPrivateSubnets {
	return &NET005CdeResourcesInPrivateSubnets{}
}

func (c *NET005CdeResourcesInPrivateSubnets) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "NET-005",
		Name:              "CDE Resources in Private Subnets",
		Description:       `All CDE components (application servers, databases) must be deployed in private subnets`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "1.3.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Move CDE resources to private subnets; use ALB/NLB in public subnets as entry point`,
		AWSConfigRule:     "",
		Category:          `Requirement 1 - Network Security Controls`,
		Resource:          "network",
	}
}

func (c *NET005CdeResourcesInPrivateSubnets) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for NET-005
	return check.Result{
		CheckID: "NET-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for NET-005",
	}
}
