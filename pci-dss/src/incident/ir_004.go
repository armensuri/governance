// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package incident

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// IR004BreachNotificationProcedures — IR-004: Breach Notification Procedures
// Category: Requirement 12 - Incident Response
type IR004BreachNotificationProcedures struct{}

func NewIR004BreachNotificationProcedures() *IR004BreachNotificationProcedures {
	return &IR004BreachNotificationProcedures{}
}

func (c *IR004BreachNotificationProcedures) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "IR-004",
		Name:              "Breach Notification Procedures",
		Description:       `Procedures for notifying card brands (Visa, Mastercard, etc.) and acquirer within required timeframes`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "12.10.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Collect and review evidence artifacts listed in check metadata.`,
		Category:          `Requirement 12 - Incident Response`,
		Resource:          "process",
		Evidence:          []string{"Card brand contact list", "Notification SLA documentation", "Breach log"},
	}
}

func (c *IR004BreachNotificationProcedures) Run(ctx context.Context) check.Result {
	return check.Result{
		CheckID: "IR-004",
		Status:  check.StatusManual,
		Message: "process check requires manual evidence review for IR-004",
		Details: map[string]string{
			"evidence_count": "3",
		},
	}
}
