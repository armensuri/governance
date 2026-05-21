// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package policies

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// POL005BusinessContinuityDisasterRecoveryPlan — POL-005: Business Continuity / Disaster Recovery Plan
// Category: CC1 - Control Environment / CC2 - Communication and Information
type POL005BusinessContinuityDisasterRecoveryPlan struct{}

func NewPOL005BusinessContinuityDisasterRecoveryPlan() *POL005BusinessContinuityDisasterRecoveryPlan {
	return &POL005BusinessContinuityDisasterRecoveryPlan{}
}

func (c *POL005BusinessContinuityDisasterRecoveryPlan) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "POL-005",
		Name:         "Business Continuity / Disaster Recovery Plan",
		Description:  `BCP/DRP is documented, tested at least annually, and includes AWS failover procedures`,
		Severity:     check.SeverityCritical,
		SOC2Criteria: []string{"A1.2", "A1.3", "CC7.5"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC1 - Control Environment / CC2 - Communication and Information`,
		Resource:     "process",
		Evidence:     []string{"BCP/DRP documents", "DR test reports", "RTO/RPO documentation"},
	}
}

func (c *POL005BusinessContinuityDisasterRecoveryPlan) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for POL-005
	return check.Result{
		CheckID: "POL-005",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for POL-005",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
