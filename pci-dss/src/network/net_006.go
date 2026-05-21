// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package network

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// NET006WafEnabledOnPublicFacingApplications — NET-006: WAF Enabled on Public-Facing Applications
// Category: Requirement 1 - Network Security Controls
// PCI DSS: 6.4.2 (v4.0)
type NET006WafEnabledOnPublicFacingApplications struct{}

func NewNET006WafEnabledOnPublicFacingApplications() *NET006WafEnabledOnPublicFacingApplications {
	return &NET006WafEnabledOnPublicFacingApplications{}
}

func (c *NET006WafEnabledOnPublicFacingApplications) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "NET-006",
		Name:              "WAF Enabled on Public-Facing Applications",
		Description:       `Deploy AWS WAF on all public-facing web applications that are part of or connected to the CDE`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "6.4.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Attach AWS WAF to CloudFront, ALB, or API Gateway endpoints serving CDE traffic`,
		AWSConfigRule:     "alb-waf-enabled",
		Category:          `Requirement 1 - Network Security Controls`,
		Resource:          "network",
	}
}

func (c *NET006WafEnabledOnPublicFacingApplications) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for NET-006
	return check.Result{
		CheckID: "NET-006",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for NET-006",
	}
}
