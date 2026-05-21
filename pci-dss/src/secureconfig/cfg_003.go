// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package secureconfig

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// CFG003UnnecessaryServicesDisabledOnCdeSystems — CFG-003: Unnecessary Services Disabled on CDE Systems
// Category: Requirement 2 - Secure Configurations
// PCI DSS: 2.2.4 (v4.0)
type CFG003UnnecessaryServicesDisabledOnCdeSystems struct{}

func NewCFG003UnnecessaryServicesDisabledOnCdeSystems() *CFG003UnnecessaryServicesDisabledOnCdeSystems {
	return &CFG003UnnecessaryServicesDisabledOnCdeSystems{}
}

func (c *CFG003UnnecessaryServicesDisabledOnCdeSystems) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "CFG-003",
		Name:              "Unnecessary Services Disabled on CDE Systems",
		Description:       `Only required ports, protocols, and services are enabled on CDE system components`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "2.2.4",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Disable unused services and ports via OS configuration and security groups`,
		AWSConfigRule:     "",
		Category:          `Requirement 2 - Secure Configurations`,
		Resource:          "secureconfig",
	}
}

func (c *CFG003UnnecessaryServicesDisabledOnCdeSystems) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CFG-003
	return check.Result{
		CheckID: "CFG-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CFG-003",
	}
}
