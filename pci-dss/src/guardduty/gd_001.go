// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package guardduty

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// GD001GuarddutyEnabledInAllCdeRegions — GD-001: GuardDuty Enabled in All CDE Regions
// Category: Requirement 10 & 11 - Threat Detection
// PCI DSS: 10.7.1 (v4.0)
type GD001GuarddutyEnabledInAllCdeRegions struct{}

func NewGD001GuarddutyEnabledInAllCdeRegions() *GD001GuarddutyEnabledInAllCdeRegions {
	return &GD001GuarddutyEnabledInAllCdeRegions{}
}

func (c *GD001GuarddutyEnabledInAllCdeRegions) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "GD-001",
		Name:              "GuardDuty Enabled in All CDE Regions",
		Description:       `Enable GuardDuty in all regions containing CDE resources for continuous threat detection`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "10.7.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable GuardDuty in all active CDE regions; configure SNS alerts for high-severity findings`,
		AWSConfigRule:     "guardduty-enabled-centralized",
		Category:          `Requirement 10 & 11 - Threat Detection`,
		Resource:          "guardduty",
	}
}

func (c *GD001GuarddutyEnabledInAllCdeRegions) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for GD-001
	return check.Result{
		CheckID: "GD-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for GD-001",
	}
}
