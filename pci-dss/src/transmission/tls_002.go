// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package transmission

import (
	"context"

	"github.com/learning/governance/pci-dss/src/pkg/check"
)

// TLS002S3BucketsEnforceSslOnlyAccess — TLS-002: S3 Buckets Enforce SSL-Only Access
// Category: Requirement 4 - Protect Cardholder Data in Transit
// PCI DSS: 4.2.1 (v4.0)
type TLS002S3BucketsEnforceSslOnlyAccess struct{}

func NewTLS002S3BucketsEnforceSslOnlyAccess() *TLS002S3BucketsEnforceSslOnlyAccess {
	return &TLS002S3BucketsEnforceSslOnlyAccess{}
}

func (c *TLS002S3BucketsEnforceSslOnlyAccess) Metadata() check.Metadata {
	return check.Metadata{
		ID:                "TLS-002",
		Name:              "S3 Buckets Enforce SSL-Only Access",
		Description:       `Deny HTTP (non-SSL) requests to CDE S3 buckets via bucket policy`,
		Severity:          check.SeverityCritical,
		PCIDSSRequirement: "4.2.1",
		PCIDSSVersion:     "v4.0",
		Remediation:       `Add bucket policy denying s3:* when aws:SecureTransport is false`,
		AWSConfigRule:     "s3-bucket-ssl-requests-only",
		Category:          `Requirement 4 - Protect Cardholder Data in Transit`,
		Resource:          "transmission",
	}
}

func (c *TLS002S3BucketsEnforceSslOnlyAccess) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for TLS-002
	return check.Result{
		CheckID: "TLS-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for TLS-002",
	}
}
