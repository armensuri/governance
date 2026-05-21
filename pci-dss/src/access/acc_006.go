// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// ACC006RdsNotPubliclyAccessible — ACC-006: RDS Not Publicly Accessible
// Category: Requirement 7 - Restrict Access by Business Need to Know
// PCI DSS: 7.2.1 (v4.0)
type ACC006RdsNotPubliclyAccessible struct{}

func NewACC006RdsNotPubliclyAccessible() *ACC006RdsNotPubliclyAccessible {
	return &ACC006RdsNotPubliclyAccessible{}
}

func (c *ACC006RdsNotPubliclyAccessible) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "ACC-006",
		Name:              "RDS Not Publicly Accessible",
		Description:       `RDS instances in the CDE must not be publicly accessible`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "7.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Set PubliclyAccessible to false on all CDE RDS instances`,
		AWSConfigRule:     "rds-instance-public-access-check",
		Category:          `Requirement 7 - Restrict Access by Business Need to Know`,
		Resource:          "access",
	}
}

func (c *ACC006RdsNotPubliclyAccessible) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for ACC-006
	return check.Result{
		CheckID: "ACC-006",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for ACC-006",
	}
}
