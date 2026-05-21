// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package guardduty

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// GD001GuarddutyEnabled — GD-001: GuardDuty Enabled
// Category: CC7 - System Operations / Threat Detection
// AWS Config rule: guardduty-enabled-centralized
type GD001GuarddutyEnabled struct{}

func NewGD001GuarddutyEnabled() *GD001GuarddutyEnabled {
	return &GD001GuarddutyEnabled{}
}

func (c *GD001GuarddutyEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "GD-001",
		Name:          "GuardDuty Enabled",
		Description:   `Ensure GuardDuty is enabled in all regions`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC7.2", "CC7.3"},
		Remediation:   `Enable GuardDuty and configure alerting via SNS`,
		AWSConfigRule: "guardduty-enabled-centralized",
		Category:      `CC7 - System Operations / Threat Detection`,
		Resource:      "guardduty",
	}
}

func (c *GD001GuarddutyEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for GD-001
	return check.Result{
		CheckID: "GD-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for GD-001",
	}
}
