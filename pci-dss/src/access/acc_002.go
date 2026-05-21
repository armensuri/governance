// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// ACC002NoSharedIamCredentialsForCde — ACC-002: No Shared IAM Credentials for CDE
// Category: Requirement 7 - Restrict Access by Business Need to Know
// PCI DSS: 8.2.1 (v4.0)
type ACC002NoSharedIamCredentialsForCde struct{}

func NewACC002NoSharedIamCredentialsForCde() *ACC002NoSharedIamCredentialsForCde {
	return &ACC002NoSharedIamCredentialsForCde{}
}

func (c *ACC002NoSharedIamCredentialsForCde) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "ACC-002",
		Name:              "No Shared IAM Credentials for CDE",
		Description:       `All access to CDE systems must use individual, unique credentials; no shared accounts`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "8.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Audit IAM users and roles; eliminate shared credentials; assign individual accounts`,
		AWSConfigRule:     "",
		Category:          `Requirement 7 - Restrict Access by Business Need to Know`,
		Resource:          "access",
	}
}

func (c *ACC002NoSharedIamCredentialsForCde) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for ACC-002
	return check.Result{
		CheckID: "ACC-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for ACC-002",
	}
}
