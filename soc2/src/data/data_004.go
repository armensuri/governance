// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package data

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// DATA004PrivacyPolicyAndGdprCcpaCompliance — DATA-004: Privacy Policy and GDPR/CCPA Compliance
// Category: C1 - Confidentiality / P - Privacy
type DATA004PrivacyPolicyAndGdprCcpaCompliance struct{}

func NewDATA004PrivacyPolicyAndGdprCcpaCompliance() *DATA004PrivacyPolicyAndGdprCcpaCompliance {
	return &DATA004PrivacyPolicyAndGdprCcpaCompliance{}
}

func (c *DATA004PrivacyPolicyAndGdprCcpaCompliance) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "DATA-004",
		Name:         "Privacy Policy and GDPR/CCPA Compliance",
		Description:  `Company maintains a privacy policy and complies with applicable data privacy regulations`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"P1.1", "P2.1", "P3.1"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `C1 - Confidentiality / P - Privacy`,
		Resource:     "process",
		Evidence:     []string{"Privacy policy", "Privacy impact assessments", "Data subject request procedures"},
	}
}

func (c *DATA004PrivacyPolicyAndGdprCcpaCompliance) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for DATA-004
	return check.Result{
		CheckID: "DATA-004",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for DATA-004",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
