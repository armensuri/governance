// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package logging

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// LOG003CloudtrailLogsEncryptedWithKms — LOG-003: CloudTrail Logs Encrypted with KMS
// Category: Requirement 10 - Log and Monitor All Access
// PCI DSS: 10.3.2 (v4.0)
type LOG003CloudtrailLogsEncryptedWithKms struct{}

func NewLOG003CloudtrailLogsEncryptedWithKms() *LOG003CloudtrailLogsEncryptedWithKms {
	return &LOG003CloudtrailLogsEncryptedWithKms{}
}

func (c *LOG003CloudtrailLogsEncryptedWithKms) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "LOG-003",
		Name:              "CloudTrail Logs Encrypted with KMS",
		Description:       `Encrypt CloudTrail logs at rest using a KMS CMK`,
		Severity:          check.SeverityHigh,
		PCIDSSRequirement: "10.3.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Configure CloudTrail to use a KMS CMK for log encryption`,
		AWSConfigRule:     "cloud-trail-encryption-enabled",
		Category:          `Requirement 10 - Log and Monitor All Access`,
		Resource:          "logging",
	}
}

func (c *LOG003CloudtrailLogsEncryptedWithKms) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for LOG-003
	return check.Result{
		CheckID: "LOG-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for LOG-003",
	}
}
