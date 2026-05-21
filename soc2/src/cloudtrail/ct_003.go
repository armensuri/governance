// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package cloudtrail

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// CT003CloudtrailLogsEncrypted — CT-003: CloudTrail Logs Encrypted
// Category: CC7 - System Operations / CC4 - Monitoring
// AWS Config rule: cloud-trail-encryption-enabled
type CT003CloudtrailLogsEncrypted struct{}

func NewCT003CloudtrailLogsEncrypted() *CT003CloudtrailLogsEncrypted {
	return &CT003CloudtrailLogsEncrypted{}
}

func (c *CT003CloudtrailLogsEncrypted) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "CT-003",
		Name:          "CloudTrail Logs Encrypted",
		Description:   `Ensure CloudTrail logs are encrypted with KMS`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.7", "C1.1"},
		Remediation:   `Configure CloudTrail to use a KMS CMK for log encryption`,
		AWSConfigRule: "cloud-trail-encryption-enabled",
		Category:      `CC7 - System Operations / CC4 - Monitoring`,
		Resource:      "cloudtrail",
	}
}

func (c *CT003CloudtrailLogsEncrypted) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for CT-003
	return check.Result{
		CheckID: "CT-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for CT-003",
	}
}
