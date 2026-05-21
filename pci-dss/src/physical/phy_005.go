// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package physical

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// PHY005AwsDataCenterPhysicalSecurity — PHY-005: AWS Data Center Physical Security
// Category: Requirement 9 - Physical Access Controls
type PHY005AwsDataCenterPhysicalSecurity struct{}

func NewPHY005AwsDataCenterPhysicalSecurity() *PHY005AwsDataCenterPhysicalSecurity {
	return &PHY005AwsDataCenterPhysicalSecurity{}
}

func (c *PHY005AwsDataCenterPhysicalSecurity) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "PHY-005",
		Name:              "AWS Data Center Physical Security",
		Description:       `Verify AWS data center physical security via AWS PCI DSS AOC or SOC 2 report`,
		Severity:          check.SeverityMedium,
		PCIDSSRequirement: "9.1.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 9 - Physical Access Controls`,
		Resource:          "process",
		Evidence:          []string{"AWS PCI DSS AOC", "AWS SOC 2 Type II report"},
	}
}

func (c *PHY005AwsDataCenterPhysicalSecurity) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "PHY-005",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for PHY-005",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
