// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package monitoring

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// MON001CentralizedLogManagement — MON-001: Centralized Log Management
// Category: CC4 - Monitoring / CC7 - System Operations
type MON001CentralizedLogManagement struct{}

func NewMON001CentralizedLogManagement() *MON001CentralizedLogManagement {
	return &MON001CentralizedLogManagement{}
}

func (c *MON001CentralizedLogManagement) Metadata() check.Metadata {
	return check.Metadata{
		ID:           "MON-001",
		Name:         "Centralized Log Management",
		Description:  `Logs from all systems (AWS, application, infrastructure) are aggregated centrally`,
		Severity:     check.SeverityHigh,
		SOC2Criteria: []string{"CC7.2", "CC7.3"},
		Remediation:  `Collect and review evidence listed in check metadata.`,
		Category:     `CC4 - Monitoring / CC7 - System Operations`,
		Resource:     "process",
		Evidence:     []string{"SIEM/log aggregation configuration", "Log retention policy"},
	}
}

func (c *MON001CentralizedLogManagement) Run(ctx context.Context) check.Result {
	// Manual/process check — validate evidence artifacts for MON-001
	return check.Result{
		CheckID: "MON-001",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for MON-001",
		Details: map[string]string{
			"evidence_count": "2",
		},
	}
}
