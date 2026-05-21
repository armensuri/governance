// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package transmission

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// TLS003StrongCipherSuitesEnforced — TLS-003: Strong Cipher Suites Enforced
// Category: Requirement 4 - Protect Cardholder Data in Transit
// PCI DSS: 4.2.1 (v4.0)
type TLS003StrongCipherSuitesEnforced struct{}

func NewTLS003StrongCipherSuitesEnforced() *TLS003StrongCipherSuitesEnforced {
	return &TLS003StrongCipherSuitesEnforced{}
}

func (c *TLS003StrongCipherSuitesEnforced) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "TLS-003",
		Name:              "Strong Cipher Suites Enforced",
		Description:       `Only strong, approved cipher suites are used on all CDE endpoints; weak ciphers are disabled`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "4.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Configure ALB/CloudFront security policies to use ELBSecurityPolicy-TLS13-1-2 or equivalent`,
		AWSConfigRule:     "",
		Category:          `Requirement 4 - Protect Cardholder Data in Transit`,
		Resource:          "transmission",
	}
}

func (c *TLS003StrongCipherSuitesEnforced) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for TLS-003
	return check.Result{
		CheckID: "TLS-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for TLS-003",
	}
}
