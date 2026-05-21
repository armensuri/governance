// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package awareness

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// SAT002SocialEngineeringPhishingTraining — SAT-002: Social Engineering / Phishing Training
// Category: Requirement 12 - Security Awareness Training
type SAT002SocialEngineeringPhishingTraining struct{}

func NewSAT002SocialEngineeringPhishingTraining() *SAT002SocialEngineeringPhishingTraining {
	return &SAT002SocialEngineeringPhishingTraining{}
}

func (c *SAT002SocialEngineeringPhishingTraining) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "SAT-002",
		Name:              "Social Engineering / Phishing Training",
		Description:       `Training must include awareness of threats such as phishing and social engineering targeting CDE access`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.6.3.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Security Awareness Training`,
		Resource:          "process",
		Evidence:          []string{"Phishing simulation results", "Social engineering training records"},
	}
}

func (c *SAT002SocialEngineeringPhishingTraining) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "SAT-002",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for SAT-002",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
