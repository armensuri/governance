// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package risk

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// RISK003ServiceProviderListMaintained — RISK-003: Service Provider List Maintained
// Category: Requirement 12 - Risk Management
type RISK003ServiceProviderListMaintained struct{}

func NewRISK003ServiceProviderListMaintained() *RISK003ServiceProviderListMaintained {
	return &RISK003ServiceProviderListMaintained{}
}

func (c *RISK003ServiceProviderListMaintained) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "RISK-003",
		Name:              "Service Provider List Maintained",
		Description:       `Maintain a list of all service providers with access to cardholder data, including their PCI DSS compliance status`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.8.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Risk Management`,
		Resource:          "process",
		Evidence:          []string{"Service provider register", "AOC/SAQ records from providers", "Annual compliance confirmation records"},
	}
}

func (c *RISK003ServiceProviderListMaintained) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "RISK-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for RISK-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
