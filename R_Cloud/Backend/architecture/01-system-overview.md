# 01 — System Overview

> End-to-end picture of R Agent Cloud including both Platform Team and AI Runtime Team.

---

# What is R Agent Cloud

R Agent Cloud is a cloud platform for deploying, managing, and monitoring AI agent applications.

Developers push their AI project to GitHub.

The platform deploys it, monitors it, and exposes it as a public API endpoint.

---

# Two Teams

| Team | Language | Owns |
|---|---|---|
| Platform Team | Go | Backend, Infrastructure, Deployment, Monitoring |
| AI Runtime Team | Python or TypeScript | Runtime Engine, Workflow, A2A, Runtime Contract |

---

# System Boundary

```text
Developer pushes GitHub repo
         │
         ▼
R Agent Cloud Platform (Backend — Go)
  → validates ragent.yaml
  → deploys to Railway
  → monitors health
  → exposes API endpoint
         │
         ▼
Railway Cloud
  → runs AI Runtime Engine (Ai-Agent/runtime-engine)
  → AI application handles POST /execute
         │
         ▼
Developer's end users call POST /execute via our API Gateway
```

---

# Team Boundary — One Rule

```text
Platform Team never implements AI workflows.

AI Runtime Team never manages deployment infrastructure.
```

---

# Interface Between Teams

| Contract | Platform Team | AI Runtime Team |
|---|---|---|
| `ragent.yaml` | Validates during deploy | Reads at startup |
| `GET /health` | Calls every 30s | Must implement |
| `GET /metadata` | Stores in Agent Registry | Must implement |
| `POST /execute` | Proxies user request | Must implement |
| `POST /stream` | Proxies user stream | Must implement |

---

# Full System Flow

```text
Phase 1 — Deploy

Developer
  → POST /api/v1/deployments
API Gateway (Go)
  → Deployment Service (Go)
  → Validation Service (Go)
  → validates ragent.yaml
  → gRPC → Runtime Service (TypeScript in Ai-Agent/runtime-service)
  → Runtime Service calls Railway API
  → Railway starts Ai-Agent/runtime-engine container
  → Runtime Service calls GET /health
  → Runtime Service calls GET /metadata → stores in Agent Registry
  → NATS event: runtime.started
  → AgentOps updates dashboard
  → WebSocket: deployment.completed to Frontend

Phase 2 — Execute

Developer's app
  → POST /api/v1/deployments/{id}/execute
API Gateway
  → looks up runtimeUrl from PostgreSQL
  → HTTP proxy to Railway container URL
Ai-Agent/runtime-engine (developer's code)
  → runs workflow defined in ragent.yaml
  → calls LLMs (developer's own API keys)
  → returns response
API Gateway
  → returns response to developer's app
AgentOps
  → records request count, latency in background
```

---

# Repository Structure

```text
R_Cloud/
│
├── proto/
│   └── runtime.proto              ← shared gRPC contract (neither team edits this alone)
│
├── Backend/                       ← Platform Team (Go)
│   ├── generated/                 ← auto-generated Go gRPC code
│   └── architecture/              ← these docs
│
└── Ai-Agent/
    ├── runtime-service/           ← Platform Team's gRPC server (TypeScript)
    │   └── src/generated/         ← auto-generated TypeScript gRPC code
    └── runtime-engine/            ← AI Team's execution engine (Python/TypeScript)
```

---

# Why gRPC Between Deployment and Runtime Service

The Deployment Service (Go) needs to command the Runtime Service (TypeScript) to deploy, restart, stop, or delete runtimes.

gRPC provides:
- Type-safe communication via Protobuf
- High performance binary protocol
- Clear interface contract via proto file

Both services generate code from the same `proto/runtime.proto`.

Neither team edits the generated code directly.

---

# Why NATS

Services need to communicate state changes without tight coupling.

NATS allows:
- Deployment Service to publish when a deployment completes
- Runtime Service to publish when a runtime starts or fails
- AgentOps to subscribe and build the observability dashboard

No service depends directly on another service's HTTP endpoint for events.

---

# Why the Runtime Service is in Ai-Agent Folder

The `Ai-Agent/runtime-service` is the Platform Team's gRPC server.

It was built by your friend (Platform Team) and placed in the Ai-Agent folder.

It is not AI logic. It is infrastructure management.

It handles: Railway API, health checking, runtime registry, provisioning.

It connects to the same PostgreSQL and NATS as the Backend.

---

# Deployment States

```text
Created → Validating → Deploying → Running → Restarting → Stopped → Deleted → Failed
```

# Runtime Health States

```text
Starting → Healthy → Unhealthy → Stopped
                        ↓
               Restart triggered (up to max retries)
```
