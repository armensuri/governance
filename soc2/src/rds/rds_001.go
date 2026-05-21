// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package rds

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// RDS001RdsEncryptionAtRest — RDS-001: RDS Encryption at Rest
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: rds-storage-encrypted
type RDS001RdsEncryptionAtRest struct{}

func NewRDS001RdsEncryptionAtRest() *RDS001RdsEncryptionAtRest {
	return &RDS001RdsEncryptionAtRest{}
}

func (c *RDS001RdsEncryptionAtRest) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "RDS-001",
		Name:          "RDS Encryption at Rest",
		Description:   `Ensure all RDS instances are encrypted at rest`,
		Severity:      check.SeverityCritical,
		SOC2Criteria:  []string{"CC6.7", "C1.1"},
		Remediation:   `Enable encryption when creating RDS instances; snapshot + restore for existing`,
		AWSConfigRule: "rds-storage-encrypted",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "rds",
	}
}

func (c *RDS001RdsEncryptionAtRest) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for RDS-001
	return check.Result{
		CheckID: "RDS-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for RDS-001",
	}
}
