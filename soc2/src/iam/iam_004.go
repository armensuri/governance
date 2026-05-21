// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package iam

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// IAM004IamPoliciesLeastPrivilege — IAM-004: IAM Policies Least Privilege
// Category: CC6 - Logical and Physical Access Controls
// AWS Config rule: iam-policy-no-statements-with-admin-access
type IAM004IamPoliciesLeastPrivilege struct{}

func NewIAM004IamPoliciesLeastPrivilege() *IAM004IamPoliciesLeastPrivilege {
	return &IAM004IamPoliciesLeastPrivilege{}
}

func (c *IAM004IamPoliciesLeastPrivilege) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "IAM-004",
		Name:          "IAM Policies Least Privilege",
		Description:   `Ensure IAM policies follow principle of least privilege`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.3"},
		Remediation:   `Review and tighten IAM policies using IAM Access Analyzer`,
		AWSConfigRule: "iam-policy-no-statements-with-admin-access",
		Category:      `CC6 - Logical and Physical Access Controls`,
		Resource:      "iam",
	}
}

func (c *IAM004IamPoliciesLeastPrivilege) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for IAM-004
	return check.Result{
		CheckID: "IAM-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for IAM-004",
	}
}
