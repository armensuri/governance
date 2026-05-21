// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package dataprotect

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// DATA008CardholderDataDiscoveryScan — DATA-008: Cardholder Data Discovery Scan
// Category: Requirement 3 - Protect Stored Account Data
// PCI DSS: 3.2.1 (v4.0)
type DATA008CardholderDataDiscoveryScan struct{}

func NewDATA008CardholderDataDiscoveryScan() *DATA008CardholderDataDiscoveryScan {
	return &DATA008CardholderDataDiscoveryScan{}
}

func (c *DATA008CardholderDataDiscoveryScan) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "DATA-008",
		Name:              "Cardholder Data Discovery Scan",
		Description:       `Perform periodic scans to detect unprotected PAN stored in unexpected locations`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "3.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Use Amazon Macie to discover and classify PAN in S3; scan databases and logs quarterly`,
		AWSConfigRule:     "",
		Category:          `Requirement 3 - Protect Stored Account Data`,
		Resource:          "dataprotect",
	}
}

func (c *DATA008CardholderDataDiscoveryScan) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for DATA-008
	return check.Result{
		CheckID: "DATA-008",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for DATA-008",
	}
}
