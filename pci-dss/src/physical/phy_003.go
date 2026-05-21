// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package physical

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// PHY003MediaContainingCardholderDataSecured — PHY-003: Media Containing Cardholder Data Secured
// Category: Requirement 9 - Physical Access Controls
type PHY003MediaContainingCardholderDataSecured struct{}

func NewPHY003MediaContainingCardholderDataSecured() *PHY003MediaContainingCardholderDataSecured {
	return &PHY003MediaContainingCardholderDataSecured{}
}

func (c *PHY003MediaContainingCardholderDataSecured) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "PHY-003",
		Name:              "Media Containing Cardholder Data Secured",
		Description:       `All media containing cardholder data must be physically secured and access controlled`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "9.4.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 9 - Physical Access Controls`,
		Resource:          "process",
		Evidence:          []string{"Media inventory", "Physical storage controls", "Media handling policy"},
	}
}

func (c *PHY003MediaContainingCardholderDataSecured) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "PHY-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for PHY-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
