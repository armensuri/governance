// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// ACC004QuarterlyAccessReviewsForCde — ACC-004: Quarterly Access Reviews for CDE
// Category: Requirement 7 - Restrict Access by Business Need to Know
// PCI DSS: 7.2.4 (v4.0)
type ACC004QuarterlyAccessReviewsForCde struct{}

func NewACC004QuarterlyAccessReviewsForCde() *ACC004QuarterlyAccessReviewsForCde {
	return &ACC004QuarterlyAccessReviewsForCde{}
}

func (c *ACC004QuarterlyAccessReviewsForCde) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "ACC-004",
		Name:              "Quarterly Access Reviews for CDE",
		Description:       `Review all user access to CDE systems and cardholder data at least every 6 months`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "7.2.4",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Conduct formal access reviews; document and action revocations`,
		AWSConfigRule:     "",
		Category:          `Requirement 7 - Restrict Access by Business Need to Know`,
		Resource:          "access",
	}
}

func (c *ACC004QuarterlyAccessReviewsForCde) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for ACC-004
	return check.Result{
		CheckID: "ACC-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for ACC-004",
	}
}
