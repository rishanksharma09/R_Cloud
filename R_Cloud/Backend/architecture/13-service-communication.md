# 13 — Service Communication

> How every service in R Agent Cloud communicates with each other.

---

# Overview

R Agent Cloud uses four communication protocols:

| Protocol | Used For |
|---|---|
| REST HTTPS | External clients → API Gateway |
| WebSocket | Frontend real-time updates |
| gRPC | Deployment Service → Runtime Service |
| NATS | Async events between all services |
| HTTP (internal) | Runtime Service → AI Application health checks |
| HTTP proxy | API Gateway → AI Application execute calls |

---

# Communication Map

```text
External Client
  → REST HTTPS → API Gateway

Frontend Dashboard
  ↔ WebSocket → API Gateway

API Gateway
  → internal Go call → Deployment Service
  → internal Go call → Auth Service
  → internal Go call → Project Service
  → internal Go call → AgentOps Service
  → HTTP proxy → AI Application on Railway (POST /execute, POST /stream)

Deployment Service
  → gRPC (TCP:50051) → Runtime Service

Runtime Service
  → REST HTTPS → Railway API (deploy, status, delete)
  → HTTP GET → AI Application on Railway (/health, /metadata)
  → NATS publish → runtime.started, runtime.failed, runtime.stopped

Deployment Service
  → NATS publish → deployment.created, deployment.completed, deployment.failed

AgentOps Service
  → NATS subscribe → all runtime and deployment events
  → WebSocket push → Frontend Dashboard

GitHub
  → Webhook POST → API Gateway → Webhook Handler → Deployment Service
```

---

# REST

Used for all external client communication.

Base URL: `/api/v1`

Format: JSON request and response bodies.

Authentication: JWT Bearer token or X-API-Key header.

---

# WebSocket

Used for real-time dashboard updates.

Endpoint: `/ws`

Frontend connects once and receives push events.

Events: deployment.completed, runtime.started, health.changed, etc.

---

# gRPC

Used exclusively between Deployment Service and Runtime Service.

Port: `50051`

Proto file: `R_Cloud/proto/runtime.proto`

Methods:

```
CreateRuntime
RestartRuntime
StopRuntime
DeleteRuntime
GetRuntimeStatus
```

Both services generate code from the same proto file.

Backend generates Go code into `Backend/generated/`.

Runtime Service generates TypeScript code into `Ai-Agent/runtime-service/src/generated/`.

---

# NATS

Used for async event propagation between services.

No service depends on another service's HTTP endpoint for state change notifications.

Both Go Backend services and TypeScript Runtime Service connect to the same NATS server.

---

# HTTP Proxy (API Gateway → AI Application)

When a developer calls `POST /api/v1/deployments/{id}/execute`:

```text
API Gateway
  → lookup runtimeUrl from PostgreSQL (cached in Redis)
  → HTTP POST to https://{runtimeUrl}/execute
  → return response to caller
```

The API Gateway does not modify the request body.

The AI application's response is returned as-is.

---

# HTTP Health Checks (Runtime Service → AI Application)

Every 30 seconds:

```text
Runtime Service health-scheduler.ts
  → for each RUNNING runtime:
      GET https://{runtimeUrl}/health
      → 200 + { status: "healthy" } → mark HEALTHY
      → error or timeout → mark UNHEALTHY → trigger restart
```

---

# Protocols by Direction

| From | To | Protocol | Purpose |
|---|---|---|---|
| Browser | API Gateway | REST HTTPS | All platform API calls |
| Frontend | API Gateway | WebSocket | Real-time event stream |
| GitHub | API Gateway | Webhook POST | Push event trigger |
| API Gateway | Auth Service | Go internal | JWT validation |
| API Gateway | Deployment Service | Go internal | Trigger deploy |
| API Gateway | AgentOps | Go internal | Dashboard data |
| API Gateway | AI Application | HTTP proxy | Forward /execute |
| Deployment Service | Runtime Service | gRPC TCP:50051 | Provision commands |
| Runtime Service | Railway | REST HTTPS | Deploy/manage containers |
| Runtime Service | AI Application | HTTP GET | /health, /metadata |
| Deployment Service | NATS | Publish | Deployment events |
| Runtime Service | NATS | Publish | Runtime events |
| AgentOps | NATS | Subscribe | All events |
| All services | OTel Collector | OTLP | Distributed tracing |
