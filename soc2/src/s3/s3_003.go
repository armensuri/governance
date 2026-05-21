// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package s3

import (
	"context"

	"github.com/learning/governance/soc2/src/pkg/check"
)

// S3003S3BucketVersioningEnabled — S3-003: S3 Bucket Versioning Enabled
// Category: CC6 - Logical and Physical Access Controls / C1 - Confidentiality
// AWS Config rule: s3-bucket-versioning-enabled
type S3003S3BucketVersioningEnabled struct{}

func NewS3003S3BucketVersioningEnabled() *S3003S3BucketVersioningEnabled {
	return &S3003S3BucketVersioningEnabled{}
}

func (c *S3003S3BucketVersioningEnabled) Metadata() check.Metadata {
	return check.Metadata{
		ID:            "S3-003",
		Name:          "S3 Bucket Versioning Enabled",
		Description:   `Enable versioning on S3 buckets storing critical data`,
		Severity:      check.SeverityMedium,
		SOC2Criteria:  []string{"A1.2", "CC7.5"},
		Remediation:   `Enable versioning and configure lifecycle policies`,
		AWSConfigRule: "s3-bucket-versioning-enabled",
		Category:      `CC6 - Logical and Physical Access Controls / C1 - Confidentiality`,
		Resource:      "s3",
	}
}

func (c *S3003S3BucketVersioningEnabled) Run(ctx context.Context) check.Result {
	// TODO: implement AWS API validation for S3-003
	return check.Result{
		CheckID: "S3-003",
		Status:  check.StatusManual,
		Message: "automated evaluation stub — wire AWS SDK logic for S3-003",
	}
}
