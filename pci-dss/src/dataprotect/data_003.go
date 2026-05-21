// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package dataprotect

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// DATA003PanEncryptedAtRestWithStrongCryptography — DATA-003: PAN Encrypted at Rest with Strong Cryptography
// Category: Requirement 3 - Protect Stored Account Data
// PCI DSS: 3.5.1 (v4.0)
type DATA003PanEncryptedAtRestWithStrongCryptography struct{}

func NewDATA003PanEncryptedAtRestWithStrongCryptography() *DATA003PanEncryptedAtRestWithStrongCryptography {
	return &DATA003PanEncryptedAtRestWithStrongCryptography{}
}

func (c *DATA003PanEncryptedAtRestWithStrongCryptography) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "DATA-003",
		Name:              "PAN Encrypted at Rest with Strong Cryptography",
		Description:       `Stored PAN must be rendered unreadable using strong cryptography (AES-256 or equivalent)`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "3.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Encrypt PAN using AES-256 with KMS CMK; use tokenization where possible`,
		AWSConfigRule:     "",
		Category:          `Requirement 3 - Protect Stored Account Data`,
		Resource:          "dataprotect",
	}
}

func (c *DATA003PanEncryptedAtRestWithStrongCryptography) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for DATA-003
	return check.Result{
		CheckID: "DATA-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for DATA-003",
	}
}
