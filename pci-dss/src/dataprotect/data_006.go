// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package dataprotect

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// DATA006RdsEncryptionEnabledForCdeDatabases — DATA-006: RDS Encryption Enabled for CDE Databases
// Category: Requirement 3 - Protect Stored Account Data
// PCI DSS: 3.5.1 (v4.0)
type DATA006RdsEncryptionEnabledForCdeDatabases struct{}

func NewDATA006RdsEncryptionEnabledForCdeDatabases() *DATA006RdsEncryptionEnabledForCdeDatabases {
	return &DATA006RdsEncryptionEnabledForCdeDatabases{}
}

func (c *DATA006RdsEncryptionEnabledForCdeDatabases) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "DATA-006",
		Name:              "RDS Encryption Enabled for CDE Databases",
		Description:       `All RDS/Aurora instances storing cardholder data must be encrypted at rest`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "3.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable RDS storage encryption with KMS CMK at instance creation`,
		AWSConfigRule:     "rds-storage-encrypted",
		Category:          `Requirement 3 - Protect Stored Account Data`,
		Resource:          "dataprotect",
	}
}

func (c *DATA006RdsEncryptionEnabledForCdeDatabases) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for DATA-006
	return check.Result{
		CheckID: "DATA-006",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for DATA-006",
	}
}
