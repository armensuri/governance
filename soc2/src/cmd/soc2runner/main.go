// Copyright (c) 2026 Armen Suri
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/learning/governance/soc2/src/pkg/check"
	"github.com/learning/governance/soc2/src/registry"
)

func main() {
	var (
		checkID = flag.String("check", "", "Run a single check by ID (e.g. IAM-001)")
		awsOnly = flag.Bool("aws", false, "Run AWS resource checks only")
		process = flag.Bool("process", false, "Run company process checks only")
		jsonOut = flag.Bool("json", false, "Output results as JSON")
	)
	flag.Parse()

	ctx := context.Background()
	var results []check.Result

	switch {
	case *checkID != "":
		results = []check.Result{runOne(ctx, *checkID)}
	case *awsOnly:
		results = runAWS(ctx)
	case *process:
		results = runProcess(ctx)
	default:
		results = append(runAWS(ctx), runProcess(ctx)...)
	}

	output(results, *jsonOut)
	os.Exit(exitCode(results))
}

func runAWS(ctx context.Context) []check.Result {
	out := make([]check.Result, 0, len(registry.AllAWS()))
	for _, c := range registry.AllAWS() {
		out = append(out, c.Run(ctx))
	}
	return out
}

func runProcess(ctx context.Context) []check.Result {
	out := make([]check.Result, 0, len(registry.AllProcess()))
	for _, c := range registry.AllProcess() {
		out = append(out, c.Run(ctx))
	}
	return out
}

func runOne(ctx context.Context, id string) check.Result {
	id = strings.ToUpper(id)
	for _, c := range registry.AllAWS() {
		if c.Metadata().ID == id {
			return c.Run(ctx)
		}
	}
	for _, c := range registry.AllProcess() {
		if c.Metadata().ID == id {
			return c.Run(ctx)
		}
	}
	return check.Result{CheckID: id, Status: check.StatusError, Message: "unknown check id"}
}

func output(results []check.Result, asJSON bool) {
	if asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(results)
		return
	}
	for _, r := range results {
		fmt.Printf("[%s] %s — %s\n", r.Status, r.CheckID, r.Message)
	}
}

func exitCode(results []check.Result) int {
	for _, r := range results {
		if r.Status == check.StatusFail || r.Status == check.StatusError {
			return 1
		}
	}
	return 0
}
