# SOC2 AWS & Process Checks (Go)

Go implementation of every check defined in `../soc2_aws_checks.json` — **one `.go` file per check** under `src/`.

## Layout

```text
src/
├── pkg/check/           # Shared types (Metadata, Result, interfaces)
├── registry/            # Generated registry of all checks
├── iam/                 # IAM-001 … IAM-006
├── s3/                  # S3-001 … S3-005
├── ec2/                 # EC2-001 … EC2-005
├── rds/                 # RDS-001 … RDS-005
├── cloudtrail/          # CT-001 … CT-004
├── kms/                 # KMS-001 … KMS-002
├── vpc/                 # VPC-001 … VPC-002
├── guardduty/           # GD-001
├── awsconfig/           # CFG-001
├── policies/            # POL-001 … POL-005
├── access/              # ACC-001 … ACC-004
├── risk/                # RISK-001 … RISK-003
├── hr/                  # HR-001 … HR-003
├── monitoring/          # MON-001 … MON-003
├── sdlc/                # SDLC-001 … SDLC-003
├── data/                # DATA-001 … DATA-004
├── cmd/soc2runner/      # CLI to run checks
└── scripts/generate_checks.py
```

**56 checks** = 31 AWS + 25 company process.

## Regenerate from JSON

After editing `soc2_aws_checks.json`:

```bash
python3 scripts/generate_checks.py
```

Re-apply any hand-written implementations (e.g. `iam/iam_001.go`) if overwritten.

## Build & run

```bash
cd src
go mod tidy
go build -o bin/soc2runner ./cmd/soc2runner
```

Run all checks:

```bash
./bin/soc2runner
```

AWS only:

```bash
./bin/soc2runner -aws
```

Single check:

```bash
./bin/soc2runner -check IAM-001
```

JSON output:

```bash
./bin/soc2runner -check IAM-001 -json
```

## Check file pattern

Each file exports:

- A struct type (e.g. `IAM001MfaEnabledForRootAccount`)
- `New…()` constructor
- `Metadata() check.Metadata` — fields from JSON
- `Run(ctx context.Context) check.Result` — evaluation logic

**AWS checks** implement `check.AWSCheck`.  
**Process checks** implement `check.ProcessCheck` (manual evidence review).

`iam/iam_001.go` includes a full AWS SDK implementation for root MFA; other AWS checks are stubs ready for SDK logic.

## Requirements

- Go 1.22+
- AWS credentials for AWS checks (env, `~/.aws/credentials`, or IAM role)

## License

MIT License — see [LICENSE](LICENSE).
