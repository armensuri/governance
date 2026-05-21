# Governance & Compliance

A learning workspace for **cloud and organizational compliance frameworks** — automated checks, control catalogs, and runnable validators. Each framework lives in its own subdirectory so SOC 2, FedRAMP, HIPAA, and others can grow independently without colliding.

---

## Repository layout

```text
governance/
├── README.md                 # This file — index for all frameworks
├── LICENSE                   # MIT (repo root)
├── .gitignore                # Go / build artifacts
│
├── soc2/                     # SOC 2 Type II (implemented)
│   ├── soc2_aws_checks.json  # Control catalog (AWS + company process)
│   └── src/                  # Go checks (56) + soc2runner CLI
│
├── pci-dss/                  # PCI DSS v4.0 (implemented)
│   ├── pci_dss_aws_checks.json
│   ├── LICENSE
│   └── src/                  # Go checks (94) + pcidssrunner CLI
│
├── fedramp/                  # FedRAMP — planned
│   └── (reserved)
│
├── hipaa/                    # HIPAA — planned
│   └── (reserved)
│
└── (future frameworks)       # e.g. ISO 27001, NIST 800-53
```

---

## Framework overview

| Framework | Directory | Status | Focus |
| --- | --- | --- | --- |
| **SOC 2** | [`soc2/`](soc2/) | Active | Trust Service Criteria (CC, A, C, PI, P); AWS technical + company process checks |
| **PCI DSS** | [`pci-dss/`](pci-dss/) | Active | PCI DSS v4.0; CDE network, data protection, logging, access, 94 checks |
| **FedRAMP** | [`fedramp/`](fedramp/) | Planned | Federal cloud security (NIST 800-53 baselines, Moderate/High) |
| **HIPAA** | [`hipaa/`](hipaa/) | Planned | PHI safeguards (Security, Privacy, Breach Notification rules) |
| *ISO 27001* | — | Future | ISMS controls |
| *NIST 800-53* | — | Future | Shared control catalog (often overlaps FedRAMP) |

---

## Architecture

```mermaid
flowchart TB
    subgraph Governance["governance/"]
        SOC2[soc2/]
        PCI[pci-dss/]
        FedRAMP[fedramp/]
        HIPAA[hipaa/]
        Future[future frameworks]
    end

    subgraph SOC2Detail["soc2/"]
        S2JSON[soc2_aws_checks.json]
        S2Go[src/ Go checks]
        S2CLI[soc2runner]
    end

    subgraph PCIDetail["pci-dss/"]
        PJSON[pci_dss_aws_checks.json]
        PGo[src/ Go checks]
        PCLI[pcidssrunner]
    end

    S2JSON --> S2Go --> S2CLI
    PJSON --> PGo --> PCLI
    SOC2 --> SOC2Detail
    PCI --> PCIDetail
    FedRAMP -.->|planned| FedRAMP
    HIPAA -.->|planned| HIPAA
```

Each framework is expected to follow the same pattern:

1. **Catalog** — machine-readable controls (JSON/YAML)
2. **Code** — one implementation file per check (Go, or other language)
3. **Runner** — CLI or CI job to execute checks and gate releases

---

## SOC 2 (implemented)

SOC 2 covers security, availability, confidentiality, processing integrity, and privacy. This repo implements checks aligned with **SOC 2 Type II** Trust Service Criteria.

### Contents

| Path | Description |
| --- | --- |
| [`soc2/soc2_aws_checks.json`](soc2/soc2_aws_checks.json) | 56 controls: 31 AWS resource checks + 25 company process checks |
| [`soc2/src/`](soc2/src/) | Go package — one `.go` file per check |
| [`soc2/src/README.md`](soc2/src/README.md) | Build, run, and regenerate instructions |

### AWS services covered

IAM · S3 · EC2 · RDS · CloudTrail · KMS · VPC · GuardDuty · AWS Config

### Quick start

```bash
cd soc2/src
go mod tidy
go build -o bin/soc2runner ./cmd/soc2runner

./bin/soc2runner -aws                 # AWS checks
./bin/soc2runner -check IAM-001       # Single check (root MFA)
```

Regenerate Go files after editing the JSON catalog:

```bash
cd soc2/src
python3 scripts/generate_checks.py
```

---

## PCI DSS (implemented)

PCI DSS v4.0 controls for Cardholder Data Environment (CDE) workloads on AWS.

### Contents

| Path | Description |
| --- | --- |
| [`pci-dss/pci_dss_aws_checks.json`](pci-dss/pci_dss_aws_checks.json) | 94 controls: 69 AWS resource checks + 25 company process checks |
| [`pci-dss/src/`](pci-dss/src/) | Go package — one `.go` file per check |
| [`pci-dss/src/README.md`](pci-dss/src/README.md) | Build, run, and regenerate instructions |

### AWS services covered

Network (VPC, SG, NACL) · Config · S3/RDS encryption · TLS/ALB · GuardDuty · Security Hub · IAM · KMS · CloudTrail · logging · vulnerability scanning · secure development (stubs)

### Quick start

```bash
cd pci-dss/src
go mod tidy
go build -o bin/pcidssrunner ./cmd/pcidssrunner

./bin/pcidssrunner -aws                 # AWS checks only
./bin/pcidssrunner -process             # Process checks only
./bin/pcidssrunner -check AUTH-002      # Single check
./bin/pcidssrunner -json                # JSON output
```

Regenerate Go files after editing the JSON catalog:

```bash
cd pci-dss/src
python3 scripts/generate_checks.py
```

---

## FedRAMP (planned)

FedRAMP authorizes cloud services for U.S. federal use based on **NIST SP 800-53** control baselines (Low, Moderate, High).

### Intended layout

```text
fedramp/
├── fedramp_aws_checks.json     # Control catalog mapped to 800-53 / FedRAMP
├── src/                        # Per-control Go (or policy) implementations
├── baselines/
│   ├── low/
│   ├── moderate/
│   └── high/
└── README.md
```

### Planned focus areas

- AC — Access control (IAM, federation, PAM)
- AU — Audit and accountability (CloudTrail, centralized logging)
- CM — Configuration management (AWS Config, drift)
- CP — Contingency planning (backups, multi-region)
- IA — Identification and authentication (MFA, key rotation)
- SC — System and communications protection (encryption, boundary protection)

*Status: directory reserved — controls and code not yet added.*

---

## HIPAA (planned)

HIPAA Security Rule safeguards for **ePHI** (electronic protected health information) in cloud environments.

### Intended layout

```text
hipaa/
├── hipaa_aws_checks.json       # Technical safeguards + organizational mapping
├── src/                        # Per-check implementations
├── safeguards/
│   ├── administrative.md
│   ├── physical.md
│   └── technical.md
└── README.md
```

### Planned focus areas

- Access controls and audit logs for systems touching PHI
- Encryption at rest and in transit (RDS, S3, TLS)
- Integrity controls and transmission security
- Minimum necessary access and workforce procedures
- BAA-aligned vendor and logging retention practices

*Status: directory reserved — controls and code not yet added.*

---

## Adding a new framework

1. Create a top-level folder: `governance/<framework>/`
2. Add `<framework>_aws_checks.json` (or YAML) with control metadata
3. Add `src/` with one file per check, `pkg/check/`, `registry/`, and `scripts/generate_checks.py`
4. Add `cmd/<framework>runner/` CLI and `src/README.md` with run instructions
5. Copy or symlink MIT `LICENSE` at framework root and under `src/`
6. Update this file’s table, layout diagram, and roadmap

Suggested naming:

| Item | Pattern |
| --- | --- |
| Catalog | `{framework}_aws_checks.json` |
| Code | `src/` (Go module per framework) |
| CLI | `src/cmd/{framework}runner/` |
| Generator | `src/scripts/generate_checks.py` |

---

## Shared conventions

Across frameworks, checks should include:

| Field | Purpose |
| --- | --- |
| `id` | Stable identifier (e.g. `IAM-001`, `AC-2`) |
| `name` | Short title |
| `description` | What is being verified |
| `severity` | `critical` / `high` / `medium` / `low` |
| `remediation` | How to fix a failure |
| `criteria` | Framework-specific control mapping |
| `evidence` | Artifacts for manual/process checks |

Automated AWS checks should map to **AWS Config rules** or equivalent API validation where possible.

---

## CI integration (recommended)

```text
PR / push
  → unit tests per framework
  → run AWS checks (credentials via OIDC)
  → fail on critical/high findings
  → publish report artifact
```

Examples:

```bash
# SOC 2
cd soc2/src && go test ./... && ./bin/soc2runner -aws -json > report.json

# PCI DSS
cd pci-dss/src && go test ./... && ./bin/pcidssrunner -aws -json > report.json
```

FedRAMP and HIPAA runners will follow the same pattern when added.

---

## Related paths

| Document | Location |
| --- | --- |
| SOC 2 check catalog | [`soc2/soc2_aws_checks.json`](soc2/soc2_aws_checks.json) |
| SOC 2 Go implementation | [`soc2/src/README.md`](soc2/src/README.md) |
| PCI DSS check catalog | [`pci-dss/pci_dss_aws_checks.json`](pci-dss/pci_dss_aws_checks.json) |
| PCI DSS Go implementation | [`pci-dss/src/README.md`](pci-dss/src/README.md) |
| MIT license (bulk headers) | [`../scripts/add_mit_license.py`](../scripts/add_mit_license.py) |

---

## Roadmap

- [x] SOC 2 — JSON catalog + 56 Go checks + CLI
- [x] PCI DSS — JSON catalog + 94 Go checks + CLI
- [ ] FedRAMP — NIST 800-53 Moderate baseline catalog
- [ ] HIPAA — ePHI technical safeguard checks
- [ ] Shared `pkg/governance` types across frameworks (optional)
- [ ] Unified multi-framework runner CLI
- [ ] GitHub Actions workflow per framework

---

## License

MIT License — see [LICENSE](LICENSE).

| Path | Notes |
| --- | --- |
| `soc2/src/LICENSE` | Module-level license copy |
| `pci-dss/LICENSE`, `pci-dss/src/LICENSE` | Framework + module copies |
| All `.go` sources | `SPDX-License-Identifier: MIT` in file headers (via generator or `add_mit_license.py`) |
