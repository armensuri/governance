// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package governance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// GOV003AnnualPciDssScopeReview — GOV-003: Annual PCI DSS Scope Review
// Category: Requirement 12 - Information Security Policy
type GOV003AnnualPciDssScopeReview struct{}

func NewGOV003AnnualPciDssScopeReview() *GOV003AnnualPciDssScopeReview {
	return &GOV003AnnualPciDssScopeReview{}
}

func (c *GOV003AnnualPciDssScopeReview) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "GOV-003",
		Name:              "Annual PCI DSS Scope Review",
		Description:       `PCI DSS scope must be reviewed and confirmed at least annually and after significant changes`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.5.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Information Security Policy`,
		Resource:          "process",
		Evidence:          []string{"Annual scope review report", "Change management records", "QSA confirmation"},
	}
}

func (c *GOV003AnnualPciDssScopeReview) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "GOV-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for GOV-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
