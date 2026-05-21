// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package rds

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RDS002RdsNotPubliclyAccessible — RDS-002: RDS Not Publicly Accessible
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: rds-instance-public-access-check
type RDS002RdsNotPubliclyAccessible struct{}

func NewRDS002RdsNotPubliclyAccessible() *RDS002RdsNotPubliclyAccessible {
	return &RDS002RdsNotPubliclyAccessible{}
}

func (c *RDS002RdsNotPubliclyAccessible) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "RDS-002",
		Name:          "RDS Not Publicly Accessible",
		Description:   `Ensure RDS instances are not publicly accessible`,
		Severity:      check.SeverityCritical,
		SOC2Criteria:  []string{"CC6.1", "CC6.6"},
		Remediation:   `Set PubliclyAccessible to false and place in private subnet`,
		AWSConfigRule: "rds-instance-public-access-check",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "rds",
	}
}

func (c *RDS002RdsNotPubliclyAccessible) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for RDS-002
	return check.Result{
		CheckID: "RDS-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for RDS-002",
	}
}
