# PCI DSS AWS & Process Checks (Go)

Go implementation of every control in `../pci_dss_aws_checks.json` — **one `.go` file per check**.

## Layout

```text
src/
├── pkg/check/              # Shared types
├── registry/             # Generated registry (94 checks)
├── network/              # NET-001 … NET-008
├── secureconfig/         # CFG-001 … CFG-005
├── dataprotect/          # DATA-001 … DATA-009
├── transmission/         # TLS-001 … TLS-004
├── malware/              # MAL-001 … MAL-004
├── securedev/            # DEV-001 … DEV-006
├── access/               # ACC-001 … ACC-006
├── authentication/       # AUTH-001 … AUTH-007
├── logging/              # LOG-001 … LOG-009
├── vulntest/             # TEST-001 … TEST-006
├── kmsecrets/            # KMS-001 … KMS-003
├── guardduty/            # GD-001
├── securityhub/          # SH-001
├── governance/           # GOV-001 … GOV-004 (process)
├── risk/                 # RISK-001 … RISK-003
├── incident/             # IR-001 … IR-004
├── awareness/            # SAT-001 … SAT-003
├── changemgmt/           # CHG-001 … CHG-003
├── physical/             # PHY-001 … PHY-005
├── compliance/           # COMP-001 … COMP-003
├── cmd/pcidssrunner/     # CLI
└── scripts/generate_checks.py
```

**94 checks** = 69 AWS + 25 company process.

## Regenerate from JSON

```bash
python3 scripts/generate_checks.py
```

## Build & run

```bash
cd pci-dss/src   # from governance/
go mod tidy
go build -o bin/pcidssrunner ./cmd/pcidssrunner

./bin/pcidssrunner -aws              # AWS checks only
./bin/pcidssrunner -process          # Process checks only
./bin/pcidssrunner -check AUTH-002   # Single check
./bin/pcidssrunner -json             # JSON output
```

## Check file pattern

Each file exports:

- `New…()` constructor
- `Metadata() check.Metadata` — PCI DSS requirement, version, severity, remediation
- `Run(ctx context.Context) check.Result` — evaluation logic (stubs ready for AWS SDK)

**AWS checks** implement `check.AWSCheck`.  
**Process checks** implement `check.ProcessCheck` (manual evidence review).

## Requirements

- Go 1.22+
- AWS credentials for AWS checks (env, `~/.aws/credentials`, or IAM role)

## License

MIT License — see [LICENSE](LICENSE). All source files include `SPDX-License-Identifier: MIT` headers.
