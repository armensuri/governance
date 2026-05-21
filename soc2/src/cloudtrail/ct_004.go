// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package cloudtrail

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// CT004CloudwatchAlarmsForCloudtrailEvents — CT-004: CloudWatch Alarms for CloudTrail Events
// Category: CC7 - System Operations / CC4 - Monitoring
// AWS Config rule: cloudwatch-alarm-action-check
type CT004CloudwatchAlarmsForCloudtrailEvents struct{}

func NewCT004CloudwatchAlarmsForCloudtrailEvents() *CT004CloudwatchAlarmsForCloudtrailEvents {
	return &CT004CloudwatchAlarmsForCloudtrailEvents{}
}

func (c *CT004CloudwatchAlarmsForCloudtrailEvents) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "CT-004",
		Name:          "CloudWatch Alarms for CloudTrail Events",
		Description:   `Set up CloudWatch metric filters and alarms for key CloudTrail events`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC7.2", "CC7.3"},
		Remediation:   `Create metric filters for root login, unauthorized API calls, MFA changes`,
		AWSConfigRule: "cloudwatch-alarm-action-check",
		Category:      `CC7 - System Operations / CC4 - Monitoring`,
		Resource:      "cloudtrail",
	}
}

func (c *CT004CloudwatchAlarmsForCloudtrailEvents) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CT-004
	return check.Result{
		CheckID: "CT-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CT-004",
	}
}
