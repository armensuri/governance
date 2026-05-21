// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package risk

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// RISK002ThirdPartyVendorRiskManagement — RISK-002: Third-Party / Vendor Risk Management
// Category: Requirement 12 - Risk Management
type RISK002ThirdPartyVendorRiskManagement struct{}

func NewRISK002ThirdPartyVendorRiskManagement() *RISK002ThirdPartyVendorRiskManagement {
	return &RISK002ThirdPartyVendorRiskManagement{}
}

func (c *RISK002ThirdPartyVendorRiskManagement) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "RISK-002",
		Name:              "Third-Party / Vendor Risk Management",
		Description:       `All third-party vendors with access to the CDE or cardholder data must be assessed for PCI DSS compliance`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.8.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Risk Management`,
		Resource:          "process",
		Evidence:          []string{"Vendor inventory", "Vendor PCI DSS AOC or questionnaire", "Vendor contracts with security requirements"},
	}
}

func (c *RISK002ThirdPartyVendorRiskManagement) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "RISK-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for RISK-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
