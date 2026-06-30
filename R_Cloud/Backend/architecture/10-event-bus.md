# 10 — Event Bus

> Asynchronous service communication via NATS.

---

# Overview

NATS is the event bus for the R Agent Cloud backend.

Every service communicates state changes asynchronously through NATS.

No service calls another service's HTTP endpoint to deliver events.

---

# Why NATS

- Decouples services from each other
- Services don't need to know which other services are running
- Events are durable and replayable
- Simple publish/subscribe model

---

# NATS Folder

```
Backend/events/
└── nats.go    ← NATS connection + publish helpers for Backend services

Ai-Agent/runtime-service/src/events/
            ← NATS client for Runtime Service (same NATS server)
```

Both connect to the same NATS server.

Each has its own connection and publish/subscribe logic.

---

# Topics

## Deployment Service publishes

| Topic | When | Payload |
|---|---|---|
| `deployment.created` | Deploy request received | deploymentId, projectId, userId |
| `deployment.completed` | Runtime is running and healthy | deploymentId, runtimeId, runtimeUrl |
| `deployment.failed` | Any step failed | deploymentId, error |

## Runtime Service publishes

| Topic | When | Payload |
|---|---|---|
| `runtime.started` | First successful health check | runtimeId, deploymentId, runtimeUrl |
| `runtime.restarted` | Auto-restart triggered and succeeded | runtimeId, restartCount |
| `runtime.stopped` | Runtime stopped by user | runtimeId |
| `runtime.failed` | Max restarts exceeded | runtimeId, reason |
| `health.failed` | Single health check failed | runtimeId, url |

## AgentOps Service subscribes to

All of the above topics.

---

# Event Payload Format

```json
{
  "event":     "runtime.started",
  "timestamp": "2026-06-30T09:00:00Z",
  "data": {
    "runtimeId":    "rtm_456",
    "deploymentId": "dep_123",
    "runtimeUrl":   "https://customer-support.up.railway.app"
  }
}
```

---

# Connection

Both Backend services and Runtime Service connect to the same NATS server.

NATS URL is provided via environment variable: `NATS_URL`

Both codebases independently manage their own NATS client connection.
