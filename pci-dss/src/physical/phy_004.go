// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package physical

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// PHY004SecureDestructionOfMediaWithCardholderData — PHY-004: Secure Destruction of Media with Cardholder Data
// Category: Requirement 9 - Physical Access Controls
type PHY004SecureDestructionOfMediaWithCardholderData struct{}

func NewPHY004SecureDestructionOfMediaWithCardholderData() *PHY004SecureDestructionOfMediaWithCardholderData {
	return &PHY004SecureDestructionOfMediaWithCardholderData{}
}

func (c *PHY004SecureDestructionOfMediaWithCardholderData) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "PHY-004",
		Name:              "Secure Destruction of Media with Cardholder Data",
		Description:       `Media containing cardholder data must be destroyed securely when no longer needed`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "9.4.6",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 9 - Physical Access Controls`,
		Resource:          "process",
		Evidence:          []string{"Media destruction policy", "Data destruction certificates", "Disposal records"},
	}
}

func (c *PHY004SecureDestructionOfMediaWithCardholderData) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "PHY-004",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for PHY-004",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
