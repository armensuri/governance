// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package policies

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// POL003IncidentResponsePlan — POL-003: Incident Response Plan
// Category: CC1 - Control Environment / CC2 - Communication and Information
type POL003IncidentResponsePlan struct{}

func NewPOL003IncidentResponsePlan() *POL003IncidentResponsePlan {
	return &POL003IncidentResponsePlan{}
}

func (c *POL003IncidentResponsePlan) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "POL-003",
		Name:         "Incident Response Plan",
		Description:  `A documented Incident Response Plan (IRP) is in place and tested annually`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"CC7.3", "CC7.4", "CC7.5"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment / CC2 - Communication and Information`,
		Resource:     "process",
		Evidence:     []string{"IRP document", "Tabletop exercise records", "Post-incident reports"},
	}
}

func (c *POL003IncidentResponsePlan) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for POL-003
	return check.Result{
		CheckID: "POL-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for POL-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
