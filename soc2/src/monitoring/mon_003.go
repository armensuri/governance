// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package monitoring

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// MON003LogRetentionPolicy — MON-003: Log Retention Policy
// Category: CC4 - Monitoring / CC7 - System Operations
type MON003LogRetentionPolicy struct{}

func NewMON003LogRetentionPolicy() *MON003LogRetentionPolicy {
	return &MON003LogRetentionPolicy{}
}

func (c *MON003LogRetentionPolicy) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "MON-003",
		Name:         "Log Retention Policy",
		Description:  `Logs are retained for a minimum of 12 months per policy`,
		Severity:     check.SeverityMedium,
		SOC2Criteria: []string{"CC7.2", "A1.2"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC4 - Monitoring / CC7 - System Operations`,
		Resource:     "process",
		Evidence:     []string{"Log retention policy", "CloudWatch/S3 retention settings"},
	}
}

func (c *MON003LogRetentionPolicy) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for MON-003
	return check.Result{
		CheckID: "MON-003",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for MON-003",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
