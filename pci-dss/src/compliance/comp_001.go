// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package compliance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// COMP001AnnualPciDssAssessmentQsaOrSaq — COMP-001: Annual PCI DSS Assessment (QSA or SAQ)
// Category: Requirement 12 - Compliance Validation
type COMP001AnnualPciDssAssessmentQsaOrSaq struct{}

func NewCOMP001AnnualPciDssAssessmentQsaOrSaq() *COMP001AnnualPciDssAssessmentQsaOrSaq {
	return &COMP001AnnualPciDssAssessmentQsaOrSaq{}
}

func (c *COMP001AnnualPciDssAssessmentQsaOrSaq) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "COMP-001",
		Name:              "Annual PCI DSS Assessment (QSA or SAQ)",
		Description:       `Complete an annual PCI DSS assessment via a Qualified Security Assessor (QSA) or Self-Assessment Questionnaire (SAQ) as applicable`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.4.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Compliance Validation`,
		Resource:          "process",
		Evidence:          []string{"Report on Compliance (ROC) or SAQ", "Attestation of Compliance (AOC)", "QSA engagement records"},
	}
}

func (c *COMP001AnnualPciDssAssessmentQsaOrSaq) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "COMP-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for COMP-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
