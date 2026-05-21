// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package awsconfig

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// CFG001AwsConfigEnabled — CFG-001: AWS Config Enabled
// Category: CC7 - System Operations / Change Management
// AWS Config rule: config-enabled-in-region
type CFG001AwsConfigEnabled struct{}

func NewCFG001AwsConfigEnabled() *CFG001AwsConfigEnabled {
	return &CFG001AwsConfigEnabled{}
}

func (c *CFG001AwsConfigEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "CFG-001",
		Name:          "AWS Config Enabled",
		Description:   `Ensure AWS Config is enabled in all regions to track configuration changes`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC7.1", "CC7.2"},
		Remediation:   `Enable AWS Config with a delivery channel in all regions`,
		AWSConfigRule: "config-enabled-in-region",
		Category:      `CC7 - System Operations / Change Management`,
		Resource:      "config",
	}
}

func (c *CFG001AwsConfigEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CFG-001
	return check.Result{
		CheckID: "CFG-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CFG-001",
	}
}
