// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package network

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// NET004VpcFlowLogsEnabled — NET-004: VPC Flow Logs Enabled
// Category: Requirement 1 - Network Security Controls
// PCI DSS: 10.2.1 (v4.0)
type NET004VpcFlowLogsEnabled struct{}

func NewNET004VpcFlowLogsEnabled() *NET004VpcFlowLogsEnabled {
	return &NET004VpcFlowLogsEnabled{}
}

func (c *NET004VpcFlowLogsEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "NET-004",
		Name:              "VPC Flow Logs Enabled",
		Description:       `Enable VPC Flow Logs for all VPCs containing CDE resources to capture network traffic`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "10.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable VPC Flow Logs and send to CloudWatch Logs or S3 with 12-month retention`,
		AWSConfigRule:     "vpc-flow-logs-enabled",
		Category:          `Requirement 1 - Network Security Controls`,
		Resource:          "network",
	}
}

func (c *NET004VpcFlowLogsEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for NET-004
	return check.Result{
		CheckID: "NET-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for NET-004",
	}
}
