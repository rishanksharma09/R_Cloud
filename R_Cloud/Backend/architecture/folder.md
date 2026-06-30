# R Agent Cloud - Complete Backend Folder Structure

```text
R_Cloud/
│
├── README.md
├── LICENSE
├── .gitignore
├── docker-compose.yml
├── Makefile
├── .env.example
│
├── docs/
│   ├── architecture/
│   │   ├── architecture.md
│   │   ├── backend.md
│   │   ├── deployment.md
│   │   ├── runtime.md
│   │   ├── validation.md
│   │   ├── notification.md
│   │   ├── auth.md
│   │   ├── event-bus.md
│   │   ├── grpc.md
│   │   ├── database.md
│   │   └── observability.md
│   ├── api/
│   │   ├── rest.md
│   │   ├── gateway.md
│   │   └── runtime-contract.md
│   ├── protobuf/
│   └── database/
│
│
├── proto/
│   ├── runtime.proto              ← ONLY gRPC contract: Deployment Service → Runtime Service
│   ├── common.proto               ← shared proto message types
│   └── README.md
│
│
├── shared/                            ← Go only — imported by Go services
│   ├── events/
│   │   ├── subjects.go               ← NATS topic name constants
│   │   ├── deployment.go             ← DeploymentCreatedEvent, DeploymentFailedEvent structs
│   │   ├── runtime.go                ← RuntimeStartedEvent, RuntimeFailedEvent structs
│   │   └── README.md
│   │
│   ├── constants/
│   │   ├── status.go
│   │   ├── errors.go
│   │   └── constants.go
│   │
│   ├── models/
│   │   ├── project.go                ← Project struct (project-service, deployment-service)
│   │   ├── deployment.go             ← Deployment struct (deployment-service, api-gateway)
│   │   ├── runtime.go                ← Runtime struct (api-gateway proxy lookup)
│   │   └── agent.go                  ← Agent struct (api-gateway agent registry)
│   │
│   └── utils/
│       ├── retry.go          ← shared retry with backoff (all Go services use this)
│       ├── timeout.go        ← shared deadline enforcement
│       └── response.go       ← standard { success, data, message } JSON helper
│
│
├── infrastructure/
│   ├── postgres/
│   │   └── postgres.go                 ← DB connection pool (imported by Go services)
│   │
│   ├── redis/
│   │   └── redis.go
│   │
│   ├── nats/
│   │   ├── nats.go
│   │   ├── docker-compose.yml
│   │   ├── nats.conf
│   │   └── README.md
│   │
│   └── otel/
│       ├── collector.yaml
│       ├── jaeger.yaml
│       └── README.md
│
│
└── backend/
    │
    ├── db/                                 ← ONE central database folder
    │   ├── schema/
    │   │   ├── platform.sql                ← projects, deployments, runtime_registry, agent_registry
    │   │   └── auth.sql                    ← users, sessions, api_keys
    │   └── migrations/
    │       ├── 001_init.sql
    │       └── 002_runtime.sql
    │
    ├── api-gateway/                        [Go]
    │   ├── cmd/
    │   │   └── main.go
    │   ├── internal/
    │   │   ├── server.go
    │   │   └── app.go
    │   ├── handlers/
    │   │   ├── project.go
    │   │   ├── deployment.go
    │   │   ├── runtime.go
    │   │   ├── agentops.go
    │   │   └── proxy.go
    │   ├── middleware/
    │   │   ├── jwt.go
    │   │   ├── apikey.go
    │   │   └── ratelimit.go
    │   ├── routes/
    │   │   └── routes.go
    │   ├── telemetry/
    │   │   ├── tracer.go
    │   │   ├── metrics.go
    │   │   └── logger.go
    │   ├── config/
    │   │   └── config.go
    │   ├── go.mod
    │   ├── Dockerfile
    │   └── README.md
    │
    │
    ├── project-service/                    [Go]
    │   ├── cmd/
    │   │   └── main.go
    │   ├── internal/
    │   │   ├── server.go
    │   │   └── app.go
    │   ├── handlers/
    │   │   └── project.go
    │   ├── service/
    │   │   └── project.service.go
    │   ├── repository/
    │   │   └── project.repository.go
    │   ├── github/
    │   │   ├── oauth.go
    │   │   ├── clone.go
    │   │   ├── webhook.go
    │   │   └── client.go
    │   ├── config/
    │   │   └── config.go
    │   ├── go.mod
    │   ├── Dockerfile
    │   └── README.md
    │
    │
    ├── deployment-service/                 [Go]
    │   ├── cmd/
    │   │   └── main.go
    │   ├── internal/
    │   │   ├── server.go
    │   │   └── app.go
    │   ├── handlers/
    │   │   └── deployment.go
    │   ├── service/
    │   │   └── deployment.service.go
    │   ├── planner/
    │   │   ├── planner.go
    │   │   └── deployment-plan.go
    │   ├── clients/
    │   │   ├── github/
    │   │   │   └── github.client.go
    │   │   ├── grpc/
    │   │   │   └── runtime.client.go       ← gRPC CLIENT calls Runtime Service
    │   │   └── postgres/
    │   │       └── postgres.client.go
    │   ├── publisher/
    │   │   └── nats.go
    │   ├── config/
    │   │   └── config.go
    │   ├── go.mod
    │   ├── Dockerfile
    │   └── README.md
    │
    │
    ├── validation-service/                 [Go]
    │   ├── cmd/
    │   │   └── main.go
    │   ├── internal/
    │   │   ├── server.go
    │   │   └── app.go
    │   ├── service/
    │   │   └── validation.service.go
    │   ├── validator/
    │   │   ├── validator.go
    │   │   ├── repository.go
    │   │   ├── yaml.go
    │   │   ├── workflow.go
    │   │   ├── endpoints.go
    │   │   ├── environment.go
    │   │   ├── railway.go
    │   │   └── dependency.go
    │   ├── config/
    │   │   └── config.go
    │   ├── go.mod
    │   ├── Dockerfile
    │   └── README.md
    │
    │
    ├── runtime-service/                    [TypeScript]
    │   ├── src/
    │   │   ├── index.ts
    │   │   └── app.ts
    │   ├── grpc/
    │   │   ├── server.ts                   ← gRPC SERVER implements runtime.proto
    │   │   └── handlers/
    │   │       ├── create-runtime.ts
    │   │       ├── restart-runtime.ts
    │   │       ├── stop-runtime.ts
    │   │       ├── delete-runtime.ts
    │   │       └── health-runtime.ts
    │   ├── railway/
    │   │   ├── client.ts
    │   │   ├── deploy.ts
    │   │   ├── restart.ts
    │   │   ├── delete.ts
    │   │   └── environment.ts
    │   ├── registry/
    │   │   ├── runtime.service.ts
    │   │   ├── runtime.repository.ts
    │   │   └── runtime.model.ts
    │   ├── health/
    │   │   ├── checker.ts
    │   │   ├── scheduler.ts
    │   │   └── restart-manager.ts
    │   ├── events/
    │   │   ├── publisher.ts
    │   │   ├── subscriber.ts
    │   │   └── subjects.ts
    │   ├── telemetry/
    │   │   ├── tracer.ts
    │   │   ├── metrics.ts
    │   │   └── logger.ts
    │   ├── config/
    │   │   └── config.ts
    │   ├── index.ts
    │   ├── Dockerfile
    │   ├── package.json
    │   ├── tsconfig.json
    │   └── README.md
    │
    │
    ├── auth-service/                       [JavaScript — Google OAuth + Session]
    │   ├── src/
    │   │   ├── index.js
    │   │   └── app.js
    │   ├── auth/
    │   │   ├── google.js                   ← Google OAuth redirect
    │   │   ├── callback.js                 ← Google OAuth callback handler
    │   │   └── logout.js
    │   ├── session/
    │   │   ├── session.service.js          ← create, validate, destroy session
    │   │   └── session.middleware.js       ← attach session to request
    │   ├── users/
    │   │   ├── users.model.js              ← User: googleId, email, name, avatar
    │   │   ├── users.service.js            ← find or create user on Google login
    │   │   └── users.repository.js
    │   ├── apikeys/
    │   │   ├── apikeys.service.js          ← generate, validate, revoke API keys
    │   │   └── apikeys.repository.js
    │   ├── middleware/
    │   │   ├── auth.middleware.js          ← check session or API key on each request
    │   │   └── cors.middleware.js
    │   ├── routes/
    │   │   └── routes.js
    │   ├── events/
    │   │   └── publisher.js                ← publish auth.login, auth.logout to NATS
    │   ├── config/
    │   │   └── config.js                   ← GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, SESSION_SECRET
    │   ├── Dockerfile
    │   ├── package.json
    │   └── README.md
    │
    │
    ├── notification-service/               [JavaScript]
    │   ├── src/
    │   │   ├── index.js
    │   │   └── app.js
    │   ├── consumers/
    │   │   ├── deployment.consumer.js      ← subscribe to deployment.created / failed
    │   │   └── runtime.consumer.js         ← subscribe to runtime.started / failed
    │   ├── websocket/
    │   │   ├── ws.gateway.js               ← push real-time events to connected frontend
    │   │   └── ws.manager.js               ← track connected clients
    │   ├── email/
    │   │   ├── sender.js
    │   │   └── templates/
    │   │       ├── deployment-failed.js
    │   │       ├── runtime-failed.js
    │   │       └── welcome.js
    │   ├── discord/
    │   │   └── discord.service.js
    │   ├── slack/
    │   │   └── slack.service.js
    │   ├── events/
    │   │   └── subjects.js
    │   ├── config/
    │   │   └── config.js
    │   ├── Dockerfile
    │   ├── package.json
    │   └── README.md
    │
    │
    └── agentops-service/                   [TypeScript]
        ├── src/
        │   ├── index.ts
        │   └── app.ts
        ├── consumers/
        │   ├── deployment.consumer.ts
        │   └── runtime.consumer.ts
        ├── metrics/
        │   ├── metrics.service.ts
        │   └── metrics.repository.ts
        ├── logs/
        │   ├── logs.service.ts
        │   └── logs.repository.ts
        ├── traces/
        │   └── traces.service.ts
        ├── analytics/
        │   ├── analytics.service.ts
        │   └── analytics.repository.ts
        ├── dashboard/
        │   ├── dashboard.handler.ts
        │   └── dashboard.service.ts
        ├── config/
        │   └── config.ts
        ├── Dockerfile
        ├── package.json
        ├── tsconfig.json
        └── README.md


└── scripts/
    ├── proto-gen.sh
    ├── build.sh
    ├── dev.sh
    ├── deploy.sh
    ├── test.sh
    └── lint.sh
```

---

# Language Per Service

| Service | Language | Owner |
|---|---|---|
| api-gateway | Go | Dev 1 |
| project-service | Go | Dev 1 |
| deployment-service | Go | Dev 1 |
| validation-service | Go | Dev 1 |
| runtime-service | TypeScript | Dev 2 |
| auth-service | JavaScript | Dev 3 |
| notification-service | JavaScript | Dev 3 |
| agentops-service | TypeScript | Dev 3 |

---

# Auth — Google OAuth + Session Only

No email/password login. No JWT issued by platform.

```text
User clicks "Login with Google"
        │
        ▼
GET /auth/google
        │
        ▼
Google OAuth Consent Screen
        │
        ▼
GET /auth/google/callback
        │
        ▼
auth-service: find or create user by googleId
        │
        ▼
create session → store in Redis
        │
        ▼
set session cookie on response
        │
        ▼
redirect to frontend dashboard
```

Session stored in Redis. Every request checks session cookie.

API keys are also supported for programmatic access.

---

# gRPC

One synchronous connection only.

```text
deployment-service (Go)
    clients/grpc/runtime.client.go    ← gRPC CLIENT

        │  gRPC TCP:50051

        ▼

runtime-service (TypeScript)
    grpc/server.ts                    ← gRPC SERVER
    grpc/handlers/
```

Proto file lives in: `proto/runtime.proto`

Dev 1 writes the proto. Dev 2 implements the server from it.

---

# NATS Event Flow

```text
deployment-service
    ├── deployment.created
    ├── deployment.validated
    └── deployment.failed

runtime-service
    ├── runtime.started
    ├── runtime.failed
    └── runtime.deleted

notification-service
    subscribes to above
    ├── sends email
    └── pushes WebSocket to frontend

agentops-service
    subscribes to above
    ├── stores metrics
    ├── stores logs
    └── updates dashboard
```

---

# OpenTelemetry

| Service | Tracer Location |
|---|---|
| api-gateway | api-gateway/telemetry/tracer.go |
| deployment-service | deployment-service/telemetry/ |
| runtime-service | runtime-service/telemetry/tracer.ts |
| agentops-service | agentops-service/src/telemetry/ |
