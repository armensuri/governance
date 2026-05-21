// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package vpc

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// VPC001VpcFlowLogsEnabled — VPC-001: VPC Flow Logs Enabled
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: vpc-flow-logs-enabled
type VPC001VpcFlowLogsEnabled struct{}

func NewVPC001VpcFlowLogsEnabled() *VPC001VpcFlowLogsEnabled {
	return &VPC001VpcFlowLogsEnabled{}
}

func (c *VPC001VpcFlowLogsEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "VPC-001",
		Name:          "VPC Flow Logs Enabled",
		Description:   `Enable VPC Flow Logs for all VPCs`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC7.2", "CC7.3"},
		Remediation:   `Enable VPC Flow Logs and send to CloudWatch Logs or S3`,
		AWSConfigRule: "vpc-flow-logs-enabled",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "vpc",
	}
}

func (c *VPC001VpcFlowLogsEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for VPC-001
	return check.Result{
		CheckID: "VPC-001",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for VPC-001",
	}
}
