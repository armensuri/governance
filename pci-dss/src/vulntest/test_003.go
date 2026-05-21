// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package vulntest

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// TEST003InternalVulnerabilityScansQuarterly — TEST-003: Internal Vulnerability Scans Quarterly
// Category: Requirement 11 - Test Security Regularly
// PCI DSS: 11.3.1 (v4.0)
type TEST003InternalVulnerabilityScansQuarterly struct{}

func NewTEST003InternalVulnerabilityScansQuarterly() *TEST003InternalVulnerabilityScansQuarterly {
	return &TEST003InternalVulnerabilityScansQuarterly{}
}

func (c *TEST003InternalVulnerabilityScansQuarterly) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "TEST-003",
		Name:              "Internal Vulnerability Scans Quarterly",
		Description:       `Quarterly internal vulnerability scans of the CDE and all systems connected to it`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "11.3.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Use AWS Inspector or third-party scanner for quarterly internal scans`,
		AWSConfigRule:     "",
		Category:          `Requirement 11 - Test Security Regularly`,
		Resource:          "vulntest",
	}
}

func (c *TEST003InternalVulnerabilityScansQuarterly) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for TEST-003
	return check.Result{
		CheckID: "TEST-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for TEST-003",
	}
}
