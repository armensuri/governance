// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package check

import "context"

// Severity levels for PCI DSS checks.
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"
)

// Status of a check execution.
type Status string

const (
	StatusPass    Status = "pass"
	StatusFail    Status = "fail"
	StatusError   Status = "error"
	StatusManual  Status = "manual"
	StatusSkipped Status = "skipped"
)

// Metadata describes a PCI DSS compliance check from pci_dss_aws_checks.json.
type Metadata struct {
	ID                string
	Name              string
	Description       string
	Severity          Severity
	PCIDSSRequirement string
	PCIDSSVersion     string
	Remediation       string
	AWSConfigRule     string
	Category          string
	Resource          string
	Evidence          []string
}

// Result is the outcome of running a check.
type Result struct {
	CheckID  string
	Status   Status
	Message  string
	Details  map[string]string
	Findings []string
}

// AWSCheck is implemented by automated AWS resource checks.
type AWSCheck interface {
	Metadata() Metadata
	Run(ctx context.Context) Result
}

// ProcessCheck is implemented by company process / evidence checks.
type ProcessCheck interface {
	Metadata() Metadata
	Run(ctx context.Context) Result
}
