// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package dataprotect

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// DATA002PanMaskedWhenDisplayed — DATA-002: PAN Masked When Displayed
// Category: Requirement 3 - Protect Stored Account Data
// PCI DSS: 3.5.1 (v4.0)
type DATA002PanMaskedWhenDisplayed struct{}

func NewDATA002PanMaskedWhenDisplayed() *DATA002PanMaskedWhenDisplayed {
	return &DATA002PanMaskedWhenDisplayed{}
}

func (c *DATA002PanMaskedWhenDisplayed) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "DATA-002",
		Name:              "PAN Masked When Displayed",
		Description:       `Primary Account Numbers (PAN) must be masked when displayed, showing only first 6 and last 4 digits`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "3.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Implement PAN masking in application display layer and APIs`,
		AWSConfigRule:     "",
		Category:          `Requirement 3 - Protect Stored Account Data`,
		Resource:          "dataprotect",
	}
}

func (c *DATA002PanMaskedWhenDisplayed) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for DATA-002
	return check.Result{
		CheckID: "DATA-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for DATA-002",
	}
}
