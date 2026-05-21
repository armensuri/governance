// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package awareness

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// SAT003AcceptableUsePolicyAcknowledgment — SAT-003: Acceptable Use Policy Acknowledgment
// Category: Requirement 12 - Security Awareness Training
type SAT003AcceptableUsePolicyAcknowledgment struct{}

func NewSAT003AcceptableUsePolicyAcknowledgment() *SAT003AcceptableUsePolicyAcknowledgment {
	return &SAT003AcceptableUsePolicyAcknowledgment{}
}

func (c *SAT003AcceptableUsePolicyAcknowledgment) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "SAT-003",
		Name:              "Acceptable Use Policy Acknowledgment",
		Description:       `All personnel must acknowledge the acceptable use policy for CDE systems and cardholder data annually`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "12.6.3",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Security Awareness Training`,
		Resource:          "process",
		Evidence:          []string{"Signed AUP acknowledgments", "HR records", "Annual re-attestation records"},
	}
}

func (c *SAT003AcceptableUsePolicyAcknowledgment) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "SAT-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for SAT-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
