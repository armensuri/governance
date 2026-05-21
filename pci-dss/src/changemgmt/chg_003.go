// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package changemgmt

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// CHG003ProductionDeploymentApprovals — CHG-003: Production Deployment Approvals
// Category: Requirement 6 - Change Management
type CHG003ProductionDeploymentApprovals struct{}

func NewCHG003ProductionDeploymentApprovals() *CHG003ProductionDeploymentApprovals {
	return &CHG003ProductionDeploymentApprovals{}
}

func (c *CHG003ProductionDeploymentApprovals) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "CHG-003",
		Name:              "Production Deployment Approvals",
		Description:       `All production CDE deployments require code review and explicit sign-off before deployment`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "6.5.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 6 - Change Management`,
		Resource:          "process",
		Evidence:          []string{"CI/CD pipeline configuration", "PR approval records", "Deployment logs"},
	}
}

func (c *CHG003ProductionDeploymentApprovals) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "CHG-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for CHG-003",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
