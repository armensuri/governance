#!/usr/bin/env python3
# Copyright (c) 2026 Armen Suri
# SPDX-License-Identifier: MIT

"""Generate one Go file per check from soc2_aws_checks.json."""
from __future__ import annotations

import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]  # src/
JSON_PATH = ROOT.parent / "soc2_aws_checks.json"  # governance/soc2_aws_checks.json
OUT = ROOT

AWS_SERVICES = {
    "iam", "s3", "ec2", "rds", "cloudtrail", "kms", "vpc", "guardduty", "config",
}
PROCESS_DIRS = {
    "security_policies": "policies",
    "access_management": "access",
    "risk_management": "risk",
    "hr_security": "hr",
    "monitoring_and_logging": "monitoring",
    "sdlc_and_change_management": "sdlc",
    "data_management": "data",
}


def slug(name: str) -> str:
    s = re.sub(r"[^a-zA-Z0-9]+", "_", name).strip("_").lower()
    return s or "check"


def type_name(check_id: str, name: str) -> str:
    parts = check_id.replace("-", "_").split("_")
    suffix = "".join(p.capitalize() for p in slug(name).split("_"))
    return f"{parts[0]}{parts[1]}{suffix}"


def go_string_list(items: list[str]) -> str:
    if not items:
        return "nil"
    quoted = ", ".join(f'"{x}"' for x in items)
    return f"[]string{{{quoted}}}"


def aws_config_rule(rule) -> str:
    if rule is None:
        return '""'
    return f'"{rule}"'


GO_HEADER = "// Copyright (c) 2026 Armen Suri\n// SPDX-License-Identifier: MIT\n\n"


def generate_aws(resource: str, category: str, c: dict) -> str:
    cid = c["id"]
    tname = type_name(cid, c["name"])
    pkg = resource
    if resource == "cloudtrail":
        pkg = "cloudtrail"
    elif resource == "config":
        pkg = "awsconfig"

    return f'''{GO_HEADER}package {pkg}

import (
\t"context"

\t"github.com/learning/governance/soc2/src/pkg/check"
)

// {tname} — {cid}: {c["name"]}
// Category: {category}
// AWS Config rule: {c.get("aws_config_rule") or "N/A"}
type {tname} struct{{}}

func New{tname}() *{tname} {{
\treturn &{tname}{{}}
}}

func (c *{tname}) Metadata() check.Metadata {{
\treturn check.Metadata{{
\t\tID:            "{cid}",
\t\tName:          "{c["name"]}",
\t\tDescription:   `{c["description"]}`,
\t\tSeverity:      check.Severity{capitalize(c["severity"])},
\t\tSOC2Criteria:  {go_string_list(c["soc2_criteria"])},
\t\tRemediation:   `{c["remediation"]}`,
\t\tAWSConfigRule: {aws_config_rule(c.get("aws_config_rule"))},
\t\tCategory:      `{category}`,
\t\tResource:      "{resource}",
\t}}
}}

func (c *{tname}) Run(ctx context.Context) check.Result {{
\t// TODO: implement AWS API validation for {cid}
\treturn check.Result{{
\t\tCheckID: "{cid}",
\t\tStatus:  check.StatusManual,
\t\tMessage: "automated evaluation stub — wire AWS SDK logic for {cid}",
\t}}
}}
'''


def generate_process(dirname: str, category: str, c: dict) -> str:
    cid = c["id"]
    tname = type_name(cid, c["name"])
    pkg = dirname
    evidence = c.get("evidence", [])
    remediation = c.get("remediation", "Collect and review evidence listed in check metadata.")

    return f'''{GO_HEADER}package {pkg}

import (
\t"context"

\t"github.com/learning/governance/soc2/src/pkg/check"
)

// {tname} — {cid}: {c["name"]}
// Category: {category}
type {tname} struct{{}}

func New{tname}() *{tname} {{
\treturn &{tname}{{}}
}}

func (c *{tname}) Metadata() check.Metadata {{
\treturn check.Metadata{{
\t\tID:           "{cid}",
\t\tName:         "{c["name"]}",
\t\tDescription:  `{c["description"]}`,
\t\tSeverity:     check.Severity{capitalize(c["severity"])},
\t\tSOC2Criteria: {go_string_list(c["soc2_criteria"])},
\t\tRemediation:  `{remediation}`,
\t\tCategory:     `{category}`,
\t\tResource:     "process",
\t\tEvidence:     {go_string_list(evidence)},
\t}}
}}

func (c *{tname}) Run(ctx context.Context) check.Result {{
\t// Manual/process check — validate evidence artifacts for {cid}
\treturn check.Result{{
\t\tCheckID: "{cid}",
\t\tStatus:  check.StatusManual,
\t\tMessage: "process check requires manual evidence review for {cid}",
\t\tDetails: map[string]string{{
\t\t\t"evidence_count": "{len(evidence)}",
\t\t}},
\t}}
}}
'''


def capitalize(s: str) -> str:
    return s[:1].upper() + s[1:]


def main() -> None:
    data = json.loads(JSON_PATH.read_text())
    root = data["soc2_compliance_checks"]

    count = 0
    for resource, block in root["aws_resource_checks"].items():
        pkg = "awsconfig" if resource == "config" else resource
        dirpath = OUT / pkg
        dirpath.mkdir(parents=True, exist_ok=True)
        category = block["category"]
        for c in block["checks"]:
            fname = c["id"].replace("-", "_").lower() + ".go"
            content = generate_aws(resource, category, c)
            if resource == "config":
                content = content.replace("package config", "package awsconfig")
            (dirpath / fname).write_text(content)
            count += 1

    for section, block in root["company_process_checks"].items():
        dirname = PROCESS_DIRS[section]
        dirpath = OUT / dirname
        dirpath.mkdir(parents=True, exist_ok=True)
        category = block["category"]
        for c in block["checks"]:
            fname = c["id"].replace("-", "_").lower() + ".go"
            (dirpath / fname).write_text(generate_process(dirname, category, c))
            count += 1

    registry = generate_registry(count)
    reg_path = OUT / "registry" / "registry.go"
    reg_path.parent.mkdir(parents=True, exist_ok=True)
    reg_path.write_text(registry)
    print(f"Generated {count} check files + registry under {OUT}")


def generate_registry(_count: int) -> str:
    """Build registry.go that lists all AWS and process checks."""
    lines = [
        "// Code generated by scripts/generate_checks.py; DO NOT EDIT.",
        "package registry",
        "",
        "import (",
        '\t"github.com/learning/governance/soc2/src/pkg/check"',
    ]
    imports = []
    entries = []

    data = json.loads(JSON_PATH.read_text())
    root = data["soc2_compliance_checks"]

    for resource, block in root["aws_resource_checks"].items():
        pkg = "awsconfig" if resource == "config" else resource
        alias = pkg
        imports.append(f'\t{alias} "github.com/learning/governance/soc2/src/{pkg}"')
        for c in block["checks"]:
            tname = type_name(c["id"], c["name"])
            entries.append(f"\t\t{alias}.New{tname}(),")

    for section, block in root["company_process_checks"].items():
        dirname = PROCESS_DIRS[section]
        imports.append(f'\t{dirname} "github.com/learning/governance/soc2/src/{dirname}"')
        for c in block["checks"]:
            tname = type_name(c["id"], c["name"])
            entries.append(f"\t\t{dirname}.New{tname}(),")

    lines.extend(sorted(set(imports)))
    lines.append(")")
    lines.append("")
    lines.append("// AllAWS returns every automated AWS resource check.")
    lines.append("func AllAWS() []check.AWSCheck {")
    lines.append("\treturn []check.AWSCheck{")
    aws_entries = [e for e in entries if "awsconfig" in e or any(x in e for x in [".iam.", ".s3.", ".ec2.", ".rds.", ".cloudtrail.", ".kms.", ".vpc.", ".guardduty."])]
    # Simpler: regenerate entries separately
    aws_entries = []
    process_entries = []
    for resource, block in root["aws_resource_checks"].items():
        pkg = "awsconfig" if resource == "config" else resource
        for c in block["checks"]:
            tname = type_name(c["id"], c["name"])
            aws_entries.append(f"\t\t{pkg}.New{tname}(),")
    for section, block in root["company_process_checks"].items():
        dirname = PROCESS_DIRS[section]
        for c in block["checks"]:
            tname = type_name(c["id"], c["name"])
            process_entries.append(f"\t\t{dirname}.New{tname}(),")

    lines = [
        "// Code generated by scripts/generate_checks.py; DO NOT EDIT.",
        "package registry",
        "",
        "import (",
        '\t"github.com/learning/governance/soc2/src/pkg/check"',
    ]
    seen = set()
    for resource in root["aws_resource_checks"]:
        pkg = "awsconfig" if resource == "config" else resource
        imp = f'\t{pkg} "github.com/learning/governance/soc2/src/{pkg}"'
        if imp not in seen:
            lines.append(imp)
            seen.add(imp)
    for section in root["company_process_checks"]:
        dirname = PROCESS_DIRS[section]
        imp = f'\t{dirname} "github.com/learning/governance/soc2/src/{dirname}"'
        if imp not in seen:
            lines.append(imp)
            seen.add(imp)
    lines.append(")")
    lines.append("")
    lines.append("func AllAWS() []check.AWSCheck {")
    lines.append("\treturn []check.AWSCheck{")
    lines.extend(aws_entries)
    lines.append("\t}")
    lines.append("}")
    lines.append("")
    lines.append("func AllProcess() []check.ProcessCheck {")
    lines.append("\treturn []check.ProcessCheck{")
    lines.extend(process_entries)
    lines.append("\t}")
    lines.append("}")
    lines.append("")
    lines.append("func All() []check.AWSCheck {")
    lines.append("\treturn AllAWS()")
    lines.append("}")
    return "\n".join(lines) + "\n"


if __name__ == "__main__":
    main()
