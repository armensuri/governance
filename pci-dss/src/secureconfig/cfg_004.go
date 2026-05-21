// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package secureconfig

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// CFG004AwsConfigEnabledForCdeAccounts — CFG-004: AWS Config Enabled for CDE Accounts
// Category: Requirement 2 - Secure Configurations
// PCI DSS: 2.2.1 (v4.0)
type CFG004AwsConfigEnabledForCdeAccounts struct{}

func NewCFG004AwsConfigEnabledForCdeAccounts() *CFG004AwsConfigEnabledForCdeAccounts {
	return &CFG004AwsConfigEnabledForCdeAccounts{}
}

func (c *CFG004AwsConfigEnabledForCdeAccounts) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "CFG-004",
		Name:              "AWS Config Enabled for CDE Accounts",
		Description:       `AWS Config must be enabled to continuously monitor CDE resource configurations`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "2.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Enable AWS Config with all resource types recorded in all CDE regions`,
		AWSConfigRule:     "config-enabled-in-region",
		Category:          `Requirement 2 - Secure Configurations`,
		Resource:          "secureconfig",
	}
}

func (c *CFG004AwsConfigEnabledForCdeAccounts) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CFG-004
	return check.Result{
		CheckID: "CFG-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CFG-004",
	}
}
