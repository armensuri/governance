// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package cloudtrail

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// CT001CloudtrailEnabledInAllRegions — CT-001: CloudTrail Enabled in All Regions
// Category: CC7 - System Operations / CC4 - Monitoring
// AWS Config rule: cloud-trail-enabled
type CT001CloudtrailEnabledInAllRegions struct{}

func NewCT001CloudtrailEnabledInAllRegions() *CT001CloudtrailEnabledInAllRegions {
	return &CT001CloudtrailEnabledInAllRegions{}
}

func (c *CT001CloudtrailEnabledInAllRegions) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "CT-001",
		Name:          "CloudTrail Enabled in All Regions",
		Description:   `Ensure CloudTrail is enabled across all AWS regions`,
		Severity:      check.SeverityCritical,
		SOC2Criteria:  []string{"CC7.2", "CC7.3"},
		Remediation:   `Create a multi-region CloudTrail trail`,
		AWSConfigRule: "cloud-trail-enabled",
		Category:      `CC7 - System Operations / CC4 - Monitoring`,
		Resource:      "cloudtrail",
	}
}

func (c *CT001CloudtrailEnabledInAllRegions) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CT-001
	return check.Result{
		CheckID: "CT-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CT-001",
	}
}
