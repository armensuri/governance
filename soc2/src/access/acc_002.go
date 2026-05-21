// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package access

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// ACC002OnboardingOffboardingProcedures — ACC-002: Onboarding / Offboarding Procedures
// Category: CC6 - Logical and Physical Access Controls
type ACC002OnboardingOffboardingProcedures struct{}

func NewACC002OnboardingOffboardingProcedures() *ACC002OnboardingOffboardingProcedures {
	return &ACC002OnboardingOffboardingProcedures{}
}

func (c *ACC002OnboardingOffboardingProcedures) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "ACC-002",
		Name:         "Onboarding / Offboarding Procedures",
		Description:  `Formal procedures ensure timely provisioning and de-provisioning of access`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC6.2", "CC6.3"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC6 - Logical and Physical Access Controls`,
		Resource:     "process",
		Evidence:     []string{"HR onboarding checklists", "Termination checklists", "Ticket history"},
	}
}

func (c *ACC002OnboardingOffboardingProcedures) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for ACC-002
	return check.Result{
		CheckID: "ACC-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for ACC-002",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
