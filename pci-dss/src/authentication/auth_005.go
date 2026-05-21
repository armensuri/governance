// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package authentication

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// AUTH005NoUnusedIamCredentials — AUTH-005: No Unused IAM Credentials
// Category: Requirement 8 - Identify and Authenticate Users
// PCI DSS: 8.2.6 (v4.0)
type AUTH005NoUnusedIamCredentials struct{}

func NewAUTH005NoUnusedIamCredentials() *AUTH005NoUnusedIamCredentials {
	return &AUTH005NoUnusedIamCredentials{}
}

func (c *AUTH005NoUnusedIamCredentials) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "AUTH-005",
		Name:              "No Unused IAM Credentials",
		Description:       `Disable or remove IAM credentials that have not been used in 90 days`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "8.2.6",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Review IAM credential report monthly; disable stale credentials`,
		AWSConfigRule:     "iam-user-unused-credentials-check",
		Category:          `Requirement 8 - Identify and Authenticate Users`,
		Resource:          "authentication",
	}
}

func (c *AUTH005NoUnusedIamCredentials) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for AUTH-005
	return check.Result{
		CheckID: "AUTH-005",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for AUTH-005",
	}
}
