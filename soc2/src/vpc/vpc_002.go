// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package vpc

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// VPC002NetworkAclsRestrictTraffic — VPC-002: Network ACLs Restrict Traffic
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: N/A
type VPC002NetworkAclsRestrictTraffic struct{}

func NewVPC002NetworkAclsRestrictTraffic() *VPC002NetworkAclsRestrictTraffic {
	return &VPC002NetworkAclsRestrictTraffic{}
}

func (c *VPC002NetworkAclsRestrictTraffic) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "VPC-002",
		Name:          "Network ACLs Restrict Traffic",
		Description:   `Ensure Network ACLs do not allow unrestricted inbound/outbound traffic`,
		Severity:      check.SeverityMedium,
		SOC2Criteria:  []string{"CC6.1", "CC6.6"},
		Remediation:   `Configure NACLs with explicit allow/deny rules`,
		AWSConfigRule: "",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "vpc",
	}
}

func (c *VPC002NetworkAclsRestrictTraffic) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for VPC-002
	return check.Result{
		CheckID: "VPC-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for VPC-002",
	}
}
