// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package governance

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// GOV001InformationSecurityPolicyDocumentedAndReviewedAnnually — GOV-001: Information Security Policy Documented and Reviewed Annually
// Category: Requirement 12 - Information Security Policy
type GOV001InformationSecurityPolicyDocumentedAndReviewedAnnually struct{}

func NewGOV001InformationSecurityPolicyDocumentedAndReviewedAnnually() *GOV001InformationSecurityPolicyDocumentedAndReviewedAnnually {
	return &GOV001InformationSecurityPolicyDocumentedAndReviewedAnnually{}
}

func (c *GOV001InformationSecurityPolicyDocumentedAndReviewedAnnually) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "GOV-001",
		Name:              "Information Security Policy Documented and Reviewed Annually",
		Description:       `An overall information security policy that addresses all PCI DSS requirements must be documented and reviewed at least annually`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.1.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Information Security Policy`,
		Resource:          "process",
		Evidence:          []string{"Information security policy", "Annual review record", "Management approval sign-off"},
	}
}

func (c *GOV001InformationSecurityPolicyDocumentedAndReviewedAnnually) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "GOV-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for GOV-001",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
