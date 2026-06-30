# R Agent Cloud — Backend Architecture

> Control plane for the R Agent Cloud platform.

---

# Overview

The backend is the control plane of R Agent Cloud.

It is responsible for:

- Authentication
- Project Management
- GitHub Integration
- AI Project Validation
- Deployment Orchestration
- Runtime Management
- Runtime Registry
- Agent Registry
- Notifications
- AgentOps
- Observability

The backend never executes AI logic.

AI applications execute inside user runtimes deployed on Railway.

---

# High Level Architecture

```text
                    Client
                       │
                       ▼
                  API Gateway
                       │
      ┌────────────────┼────────────────┐
      ▼                ▼                ▼
 Authentication   Project Service   Deployment Service
                                         │
                                         ▼
                                 Validation Service
                                         │
                                         ▼
                                  Runtime Service ──── gRPC ───→ Ai-Agent/runtime-service
                                         │
                                         ▼
                                  Railway Provider
                                         │
                                         ▼
                              AI Application on Railway
                              (Ai-Agent/runtime-engine)

                 ─────────────────────────────────────

                 Notification Service   ← WebSocket + Email

                 AgentOps Service       ← subscribes to NATS events

                 PostgreSQL             ← platform data

                 Auth PostgreSQL        ← users, tokens, api keys

                 Redis                  ← cache, sessions, rate limiting

                 OpenTelemetry          ← distributed tracing

                 NATS                   ← async event bus
```

---

# Services

| Service | Responsibility | Doc |
|---|---|---|
| API Gateway | Single entry point, JWT, rate limiting, routing, proxy | 02-api-gateway.md |
| Auth Service | Login, register, JWT, refresh tokens, API keys | 03-auth-service.md |
| Project Service | Projects, GitHub repos, branches, metadata | 04-project-service.md |
| Deployment Service | Clone repo, lifecycle, version management, deploy plan | 05-deployment-service.md |
| Validation Service | Validate ragent.yaml, entrypoints, runtime contract | 06-validation-service.md |
| Runtime Service | Deploy, restart, stop, delete, health, registry | 07-runtime-service.md |
| Notification Service | Email, WebSocket, real-time updates | 08-notification-service.md |
| AgentOps Service | Metrics, logs, traces, analytics, dashboard | 09-agentops-service.md |

---

# Databases

| Database | Purpose | Doc |
|---|---|---|
| Auth PostgreSQL | Users, sessions, API keys, refresh tokens | 11-database.md |
| Platform PostgreSQL | Projects, deployments, runtime registry, agent registry | 11-database.md |
| Redis | Sessions, cache, rate limiting, queues | 11-database.md |

---

# Event Bus

Technology: NATS

Every service publishes and subscribes to events asynchronously.

Documentation: 10-event-bus.md

---

# Service Communication

| From | To | Protocol |
|---|---|---|
| Client | API Gateway | REST HTTPS |
| Frontend | API Gateway | WebSocket |
| API Gateway | Deployment Service | Internal Go call |
| Deployment Service | Runtime Service | gRPC TCP:50051 |
| Runtime Service | Railway API | REST HTTPS |
| Runtime Service | AI Application | HTTP GET /health |
| API Gateway | AI Application | HTTP proxy /execute |
| Any service | NATS | Publish/Subscribe |
| Any service | OTel Collector | OpenTelemetry |

Documentation: 13-service-communication.md

---

# Design Principles

- Every service owns a single responsibility.
- Runtime Service is the only service allowed to communicate with Railway.
- Validation happens before deployment.
- Authentication is isolated in its own database.
- Platform metadata is stored separately from auth data.
- Services communicate asynchronously through NATS.
- OpenTelemetry provides distributed tracing across all backend services.
- gRPC is used only for Deployment Service → Runtime Service communication.
- The backend never executes AI logic.
