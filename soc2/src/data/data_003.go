// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package data

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// DATA003DataEncryptionInTransit — DATA-003: Data Encryption in Transit
// Category: C1 - Confidentiality / P - Privacy
type DATA003DataEncryptionInTransit struct{}

func NewDATA003DataEncryptionInTransit() *DATA003DataEncryptionInTransit {
	return &DATA003DataEncryptionInTransit{}
}

func (c *DATA003DataEncryptionInTransit) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "DATA-003",
		Name:         "Data Encryption in Transit",
		Description:  `All data in transit uses TLS 1.2 or higher`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC6.7", "C1.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `C1 - Confidentiality / P - Privacy`,
		Resource:     "process",
		Evidence:     []string{"TLS configuration records", "Certificate inventory", "SSL Labs scan results"},
	}
}

func (c *DATA003DataEncryptionInTransit) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for DATA-003
	return check.Result{
		CheckID: "DATA-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for DATA-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
