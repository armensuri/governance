// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package rds

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RDS005RdsMinorVersionAutoUpgrade — RDS-005: RDS Minor Version Auto-Upgrade
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: rds-automatic-minor-version-upgrade-enabled
type RDS005RdsMinorVersionAutoUpgrade struct{}

func NewRDS005RdsMinorVersionAutoUpgrade() *RDS005RdsMinorVersionAutoUpgrade {
	return &RDS005RdsMinorVersionAutoUpgrade{}
}

func (c *RDS005RdsMinorVersionAutoUpgrade) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "RDS-005",
		Name:          "RDS Minor Version Auto-Upgrade",
		Description:   `Enable automatic minor version upgrades for RDS`,
		Severity:      check.SeverityMedium,
		SOC2Criteria:  []string{"CC7.1"},
		Remediation:   `Enable AutoMinorVersionUpgrade on RDS instances`,
		AWSConfigRule: "rds-automatic-minor-version-upgrade-enabled",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "rds",
	}
}

func (c *RDS005RdsMinorVersionAutoUpgrade) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for RDS-005
	return check.Result{
		CheckID: "RDS-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for RDS-005",
	}
}
