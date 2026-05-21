// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// ACC003PrivilegedAccessManagement — ACC-003: Privileged Access Management
// Category: CC6 - Logical and Physical Access Controls
type ACC003PrivilegedAccessManagement struct{}

func NewACC003PrivilegedAccessManagement() *ACC003PrivilegedAccessManagement {
	return &ACC003PrivilegedAccessManagement{}
}

func (c *ACC003PrivilegedAccessManagement) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "ACC-003",
		Name:         "Privileged Access Management",
		Description:  `Privileged access (admin, root) is restricted, monitored, and reviewed`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC6.3", "CC6.8"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC6 - Logical and Physical Access Controls`,
		Resource:     "process",
		Evidence:     []string{"PAM policy", "Admin user list", "Privileged session logs"},
	}
}

func (c *ACC003PrivilegedAccessManagement) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for ACC-003
	return check.Result{
		CheckID: "ACC-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for ACC-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
