// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package s3

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// S3002S3BucketEncryptionEnabled — S3-002: S3 Bucket Encryption Enabled
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: s3-bucket-server-side-encryption-enabled
type S3002S3BucketEncryptionEnabled struct{}

func NewS3002S3BucketEncryptionEnabled() *S3002S3BucketEncryptionEnabled {
	return &S3002S3BucketEncryptionEnabled{}
}

func (c *S3002S3BucketEncryptionEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "S3-002",
		Name:          "S3 Bucket Encryption Enabled",
		Description:   `Ensure S3 buckets have server-side encryption enabled`,
		Severity:      check.SeverityHigh,
		SOC2Criteria:  []string{"CC6.7", "C1.1"},
		Remediation:   `Enable SSE-S3 or SSE-KMS on all S3 buckets`,
		AWSConfigRule: "s3-bucket-server-side-encryption-enabled",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "s3",
	}
}

func (c *S3002S3BucketEncryptionEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for S3-002
	return check.Result{
		CheckID: "S3-002",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for S3-002",
	}
}
