// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package governance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// GOV002PciDssScopeDefinedAndDocumented — GOV-002: PCI DSS Scope Defined and Documented
// Category: Requirement 12 - Information Security Policy
type GOV002PciDssScopeDefinedAndDocumented struct{}

func NewGOV002PciDssScopeDefinedAndDocumented() *GOV002PciDssScopeDefinedAndDocumented {
	return &GOV002PciDssScopeDefinedAndDocumented{}
}

func (c *GOV002PciDssScopeDefinedAndDocumented) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "GOV-002",
		Name:              "PCI DSS Scope Defined and Documented",
		Description:       `The CDE scope, including all in-scope systems, people, and processes, must be formally defined and documented`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Information Security Policy`,
		Resource:          "process",
		Evidence:          []string{"Scope definition document", "Network diagram", "Data flow diagram", "System inventory"},
	}
}

func (c *GOV002PciDssScopeDefinedAndDocumented) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "GOV-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for GOV-002",
		Details: map[string]string{
			"evidence_count": "4",
		},
	}
}
