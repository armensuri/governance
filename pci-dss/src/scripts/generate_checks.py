#!/usr/bin/env python3
# Copyright (c) 2026 Armen Suri
# SPDX-License-Identifier: MIT

"""Generate one Go file per PCI DSS check from pci_dss_aws_checks.json."""
from __future__ import annotations

import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
JSON_PATH = ROOT.parent / "pci_dss_aws_checks.json"
OUT = ROOT

GO_HEADER = "// Copyright (c) 2026 Armen Suri\n// SPDX-License-Identifier: MIT\n\n"
MODULE = "github.com/learning/governance/pci-dss/src"

AWS_SECTION_PACKAGES = {
    "network_security": "network",
    "secure_configuration": "secureconfig",
    "data_protection": "dataprotect",
    "transmission_security": "transmission",
    "malware_protection": "malware",
    "secure_development": "securedev",
    "access_control": "access",
    "authentication": "authentication",
    "logging_and_monitoring": "logging",
    "vulnerability_testing": "vulntest",
    "kms_and_secrets": "kmsecrets",
}

PROCESS_SECTION_PACKAGES = {
    "governance_and_policies": "governance",
    "risk_management": "risk",
    "incident_response": "incident",
    "security_awareness": "awareness",
    "change_management": "changemgmt",
    "physical_security": "physical",
    "compliance_reporting": "compliance",
}


def slug(name: str) -> str:
    s = re.sub(r"[^a-zA-Z0-9]+", "_", name).strip("_").lower()
    return s or "check"


def type_name(check_id: str, name: str) -> str:
    parts = check_id.replace("-", "_").split("_")
    prefix = "".join(p.upper() for p in parts[:2] if p)  # NET001 style from NET-001
    if len(parts) >= 2:
        prefix = parts[0] + parts[1]
    suffix = "".join(p.capitalize() for p in slug(name).split("_"))
    return f"{prefix}{suffix}"


def capitalize(s: str) -> str:
    return s[:1].upper() + s[1:]


def aws_config_rule(rule) -> str:
    if rule is None:
        return '""'
    return f'"{rule}"'


def pci_pkg_for_aws(section: str, check_id: str) -> str:
    if section == "guardduty_and_security_hub":
        if check_id.startswith("SH-"):
            return "securityhub"
        return "guardduty"
    return AWS_SECTION_PACKAGES[section]


def generate_aws(pkg: str, category: str, c: dict) -> str:
    cid = c["id"]
    tname = type_name(cid, c["name"])
    return f'''{GO_HEADER}package {pkg}

import (
\t"context"

\t"{MODULE}/pkg/check"
)

// {tname} — {cid}: {c["name"]}
// Category: {category}
// PCI DSS: {c["pci_dss_requirement"]} ({c.get("pci_dss_version", "v4.0")})
type {tname} struct{{}}

func New{tname}() *{tname} {{
\treturn &{tname}{{}}
}}

func (c *{tname}) Metadata() check.Metadata {{
\treturn check.Metadata{{
\t\tID:                "{cid}",
\t\tName:              "{c["name"]}",
\t\tDescription:       `{c["description"]}`,
\t\tSeverity:          check.Severity{capitalize(c["severity"])},
\t\tPCIDSSRequirement: "{c["pci_dss_requirement"]}",
\t\tPCIDSSVersion:     "{c.get("pci_dss_version", "v4.0")}",
\t\tRemediation:       `{c["remediation"]}`,
\t\tAWSConfigRule:     {aws_config_rule(c.get("aws_config_rule"))},
\t\tCategory:          `{category}`,
\t\tResource:          "{pkg}",
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


def go_string_list(items: list[str]) -> str:
    if not items:
        return "nil"
    quoted = ", ".join(f'"{x}"' for x in items)
    return f"[]string{{{quoted}}}"


def generate_process(pkg: str, category: str, c: dict) -> str:
    cid = c["id"]
    tname = type_name(cid, c["name"])
    evidence = c.get("evidence", [])
    remediation = c.get(
        "remediation",
        "Collect and review evidence artifacts listed in check metadata.",
    )
    return f'''{GO_HEADER}package {pkg}

import (
\t"context"

\t"{MODULE}/pkg/check"
)

// {tname} — {cid}: {c["name"]}
// Category: {category}
type {tname} struct{{}}

func New{tname}() *{tname} {{
\treturn &{tname}{{}}
}}

func (c *{tname}) Metadata() check.Metadata {{
\treturn check.Metadata{{
\t\tID:                "{cid}",
\t\tName:              "{c["name"]}",
\t\tDescription:       `{c["description"]}`,
\t\tSeverity:          check.Severity{capitalize(c["severity"])},
\t\tPCIDSSRequirement: "{c["pci_dss_requirement"]}",
\t\tPCIDSSVersion:     "{c.get("pci_dss_version", "v4.0")}",
\t\tRemediation:       `{remediation}`,
\t\tCategory:          `{category}`,
\t\tResource:          "process",
\t\tEvidence:          {go_string_list(evidence)},
\t}}
}}

func (c *{tname}) Run(ctx context.Context) check.Result {{
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


def generate_registry(aws_entries: list[tuple[str, str]], process_entries: list[tuple[str, str]]) -> str:
    imports = {f'\t"{MODULE}/pkg/check"'}
    for pkg, tname in aws_entries:
        imports.add(f'\t{pkg} "{MODULE}/{pkg}"')
    for pkg, tname in process_entries:
        imports.add(f'\t{pkg} "{MODULE}/{pkg}"')

    lines = [
        "// Code generated by scripts/generate_checks.py; DO NOT EDIT.",
        GO_HEADER.rstrip(),
        "package registry",
        "",
        "import (",
    ]
    lines.extend(sorted(imports))
    lines.append(")")
    lines.append("")
    lines.append("func AllAWS() []check.AWSCheck {")
    lines.append("\treturn []check.AWSCheck{")
    for pkg, tname in aws_entries:
        lines.append(f"\t\t{pkg}.New{tname}(),")
    lines.append("\t}")
    lines.append("}")
    lines.append("")
    lines.append("func AllProcess() []check.ProcessCheck {")
    lines.append("\treturn []check.ProcessCheck{")
    for pkg, tname in process_entries:
        lines.append(f"\t\t{pkg}.New{tname}(),")
    lines.append("\t}")
    lines.append("}")
    return "\n".join(lines) + "\n"


def main() -> None:
    data = json.loads(JSON_PATH.read_text(encoding="utf-8"))
    root = data["pci_dss_compliance_checks"]

    aws_registry: list[tuple[str, str]] = []
    process_registry: list[tuple[str, str]] = []
    count = 0

    for section, block in root["aws_resource_checks"].items():
        pkg = pci_pkg_for_aws(section, "")
        category = block["category"]
        for c in block["checks"]:
            pkg = pci_pkg_for_aws(section, c["id"])
            dirpath = OUT / pkg
            dirpath.mkdir(parents=True, exist_ok=True)
            tname = type_name(c["id"], c["name"])
            fname = c["id"].replace("-", "_").lower() + ".go"
            (dirpath / fname).write_text(generate_aws(pkg, category, c), encoding="utf-8")
            aws_registry.append((pkg, tname))
            count += 1

    for section, block in root["company_process_checks"].items():
        pkg = PROCESS_SECTION_PACKAGES[section]
        category = block["category"]
        dirpath = OUT / pkg
        dirpath.mkdir(parents=True, exist_ok=True)
        for c in block["checks"]:
            tname = type_name(c["id"], c["name"])
            fname = c["id"].replace("-", "_").lower() + ".go"
            (dirpath / fname).write_text(generate_process(pkg, category, c), encoding="utf-8")
            process_registry.append((pkg, tname))
            count += 1

    reg_path = OUT / "registry" / "registry.go"
    reg_path.parent.mkdir(parents=True, exist_ok=True)
    reg_path.write_text(generate_registry(aws_registry, process_registry), encoding="utf-8")
    print(f"Generated {count} check files under {OUT}")


if __name__ == "__main__":
    main()
