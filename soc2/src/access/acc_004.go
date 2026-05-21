// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// ACC004VendorThirdPartyAccessControls — ACC-004: Vendor / Third-Party Access Controls
// Category: CC6 - Logical and Physical Access Controls
type ACC004VendorThirdPartyAccessControls struct{}

func NewACC004VendorThirdPartyAccessControls() *ACC004VendorThirdPartyAccessControls {
	return &ACC004VendorThirdPartyAccessControls{}
}

func (c *ACC004VendorThirdPartyAccessControls) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "ACC-004",
		Name:         "Vendor / Third-Party Access Controls",
		Description:  `Third-party access to systems is formally approved, limited, and reviewed`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC6.2", "CC9.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC6 - Logical and Physical Access Controls`,
		Resource:     "process",
		Evidence:     []string{"Vendor access agreements", "NDA records", "Third-party access logs"},
	}
}

func (c *ACC004VendorThirdPartyAccessControls) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for ACC-004
	return check.Result{
		CheckID: "ACC-004",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for ACC-004",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
