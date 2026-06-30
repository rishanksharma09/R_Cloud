# 06 — Validation Service

> Validates AI repositories before deployment.

---

# Overview

The Validation Service validates that an AI repository is correctly structured and ready to deploy.

Validation runs before any infrastructure is provisioned.

If validation fails, deployment is rejected immediately.

---

# Responsibilities

- Validate ragent.yaml exists
- Validate mode value (monolith or microservices)
- Validate entrypoints exist in the repository
- Validate runtime contract routes are defined
- Validate environment variables are declared

---

# Folder Structure

```
Backend/deployment/validator/
└── ragent_validator.go    ← all validation logic
```

---

# Validation Checks

## ragent.yaml Exists

Must be in repository root. If missing: fail.

## Mode is Valid

Allowed values: `monolith` or `microservices`. Anything else: fail.

## Entrypoints Exist

For microservices mode, every agent entrypoint file must exist in the repository.

## Routes Defined

The following routes must be declared in ragent.yaml:

```yaml
routes:
  execute: /execute
  health: /health
  metadata: /metadata
```

---

# Validation Flow

```text
ragent.yaml exists? → No: FAIL
mode is valid?      → No: FAIL
entrypoints exist?  → No: FAIL
routes defined?     → No: FAIL
                    → Yes: PASS → proceed to deployment
```

---

# Validation Response

Success:

```json
{
  "valid": true,
  "mode": "microservices",
  "agents": ["planner", "researcher"]
}
```

Failure:

```json
{
  "valid": false,
  "errors": ["Missing ragent.yaml", "Missing /health route"]
}
```
