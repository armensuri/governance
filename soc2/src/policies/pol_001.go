// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package policies

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// POL001InformationSecurityPolicyDocumented — POL-001: Information Security Policy Documented
// Category: CC1 - Control Environment / CC2 - Communication and Information
type POL001InformationSecurityPolicyDocumented struct{}

func NewPOL001InformationSecurityPolicyDocumented() *POL001InformationSecurityPolicyDocumented {
	return &POL001InformationSecurityPolicyDocumented{}
}

func (c *POL001InformationSecurityPolicyDocumented) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "POL-001",
		Name:         "Information Security Policy Documented",
		Description:  `Company maintains a written information security policy reviewed annually`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC1.1", "CC2.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment / CC2 - Communication and Information`,
		Resource:     "process",
		Evidence:     []string{"Information Security Policy document", "Annual review sign-off records"},
	}
}

func (c *POL001InformationSecurityPolicyDocumented) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for POL-001
	return check.Result{
		CheckID: "POL-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for POL-001",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
