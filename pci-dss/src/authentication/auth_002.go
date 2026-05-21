// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package authentication

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// AUTH002MfaOnRootAccount — AUTH-002: MFA on Root Account
// Category: Requirement 8 - Identify and Authenticate Users
// PCI DSS: 8.4.2 (v4.0)
type AUTH002MfaOnRootAccount struct{}

func NewAUTH002MfaOnRootAccount() *AUTH002MfaOnRootAccount {
	return &AUTH002MfaOnRootAccount{}
}

func (c *AUTH002MfaOnRootAccount) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "AUTH-002",
		Name:              "MFA on Root Account",
		Description:       `MFA must be enabled on the AWS root account for all CDE accounts`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "8.4.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable hardware or virtual MFA on the root account; restrict root usage`,
		AWSConfigRule:     "root-account-mfa-enabled",
		Category:          `Requirement 8 - Identify and Authenticate Users`,
		Resource:          "authentication",
	}
}

func (c *AUTH002MfaOnRootAccount) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for AUTH-002
	return check.Result{
		CheckID: "AUTH-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for AUTH-002",
	}
}
