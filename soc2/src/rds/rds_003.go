// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package rds

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RDS003RdsAutomatedBackupsEnabled — RDS-003: RDS Automated Backups Enabled
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: db-instance-backup-enabled
type RDS003RdsAutomatedBackupsEnabled struct{}

func NewRDS003RdsAutomatedBackupsEnabled() *RDS003RdsAutomatedBackupsEnabled {
	return &RDS003RdsAutomatedBackupsEnabled{}
}

func (c *RDS003RdsAutomatedBackupsEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "RDS-003",
		Name:          "RDS Automated Backups Enabled",
		Description:   `Ensure RDS automated backups are enabled with retention >= 7 days`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"A1.2", "CC7.5"},
		Remediation:   `Set backup retention period to at least 7 days`,
		AWSConfigRule: "db-instance-backup-enabled",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "rds",
	}
}

func (c *RDS003RdsAutomatedBackupsEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for RDS-003
	return check.Result{
		CheckID: "RDS-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for RDS-003",
	}
}
