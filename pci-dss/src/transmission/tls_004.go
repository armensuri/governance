// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package transmission

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// TLS004NoCardholderDataSentViaUnencryptedChannels — TLS-004: No Cardholder Data Sent via Unencrypted Channels
// Category: Requirement 4 - Protect Cardholder Data in Transit
// PCI DSS: 4.2.2 (v4.0)
type TLS004NoCardholderDataSentViaUnencryptedChannels struct{}

func NewTLS004NoCardholderDataSentViaUnencryptedChannels() *TLS004NoCardholderDataSentViaUnencryptedChannels {
	return &TLS004NoCardholderDataSentViaUnencryptedChannels{}
}

func (c *TLS004NoCardholderDataSentViaUnencryptedChannels) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "TLS-004",
		Name:              "No Cardholder Data Sent via Unencrypted Channels",
		Description:       `Cardholder data must never be transmitted via email, chat, or any unencrypted channel`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "4.2.2",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Implement DLP controls; audit application code for unencrypted data transmission`,
		AWSConfigRule:     "",
		Category:          `Requirement 4 - Protect Cardholder Data in Transit`,
		Resource:          "transmission",
	}
}

func (c *TLS004NoCardholderDataSentViaUnencryptedChannels) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for TLS-004
	return check.Result{
		CheckID: "TLS-004",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for TLS-004",
	}
}
