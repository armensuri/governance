// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package securedev

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// DEV006NoLivePansInTestEnvironments — DEV-006: No Live PANs in Test Environments
// Category: Requirement 6 - Secure Systems and Software
// PCI DSS: 6.2.3.1 (v4.0)
type DEV006NoLivePansInTestEnvironments struct{}

func NewDEV006NoLivePansInTestEnvironments() *DEV006NoLivePansInTestEnvironments {
	return &DEV006NoLivePansInTestEnvironments{}
}

func (c *DEV006NoLivePansInTestEnvironments) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "DEV-006",
		Name:              "No Live PANs in Test Environments",
		Description:       `Real cardholder data (live PAN) must not be used in non-production environments`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "6.2.3.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Use synthetic or tokenized data in test environments; scan with Macie to detect real PAN`,
		AWSConfigRule:     "",
		Category:          `Requirement 6 - Secure Systems and Software`,
		Resource:          "securedev",
	}
}

func (c *DEV006NoLivePansInTestEnvironments) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for DEV-006
	return check.Result{
		CheckID: "DEV-006",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for DEV-006",
	}
}
