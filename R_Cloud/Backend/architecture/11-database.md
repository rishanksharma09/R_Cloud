# 11 — Database

> Storage architecture for the R Agent Cloud platform.

---

# Overview

The platform uses three storage technologies:

| Database | Purpose |
|---|---|
| Auth PostgreSQL | Users, sessions, API keys, refresh tokens |
| Platform PostgreSQL | Projects, deployments, runtime registry, agent registry, metrics |
| Redis | Sessions cache, rate limiting, queues, WebSocket state |

---

# Auth PostgreSQL

Isolated database. No other service accesses this database.

Schema: `Backend/db/schema/auth.sql`

## Tables

```
users
  id            UUID PRIMARY KEY
  email         TEXT UNIQUE NOT NULL
  password_hash TEXT NOT NULL
  created_at    TIMESTAMP
  updated_at    TIMESTAMP

sessions
  id                 UUID PRIMARY KEY
  user_id            UUID REFERENCES users(id)
  refresh_token_hash TEXT NOT NULL
  expires_at         TIMESTAMP
  created_at         TIMESTAMP

api_keys
  id           UUID PRIMARY KEY
  user_id      UUID REFERENCES users(id)
  key_hash     TEXT UNIQUE NOT NULL
  name         TEXT
  created_at   TIMESTAMP
  last_used_at TIMESTAMP
  revoked_at   TIMESTAMP
```

---

# Platform PostgreSQL

Shared by: Deployment Service, Runtime Service, AgentOps Service, API Gateway.

Schema: `Backend/db/schema/runtime.sql`

## Tables

```
projects
  id              UUID PRIMARY KEY
  user_id         UUID NOT NULL
  name            TEXT NOT NULL
  github_repo_url TEXT
  default_branch  TEXT
  created_at      TIMESTAMP
  updated_at      TIMESTAMP

deployments
  id           UUID PRIMARY KEY
  project_id   UUID REFERENCES projects(id)
  user_id      UUID NOT NULL
  branch       TEXT
  commit_hash  TEXT
  version      TEXT
  mode         TEXT        ← monolith | microservices
  status       TEXT        ← VALIDATING | DEPLOYING | RUNNING | FAILED | STOPPED | DELETED
  created_at   TIMESTAMP
  completed_at TIMESTAMP

runtime_registry
  id             UUID PRIMARY KEY
  deployment_id  UUID REFERENCES deployments(id)
  runtime_url    TEXT NOT NULL
  provider       TEXT        ← railway
  status         TEXT        ← RUNNING | STOPPED | FAILED | DELETED
  health         TEXT        ← STARTING | HEALTHY | UNHEALTHY | STOPPED
  restart_count  INTEGER DEFAULT 0
  created_at     TIMESTAMP
  updated_at     TIMESTAMP

agent_registry
  id           UUID PRIMARY KEY
  runtime_id   UUID REFERENCES runtime_registry(id)
  name         TEXT
  framework    TEXT
  version      TEXT
  capabilities TEXT[]
  created_at   TIMESTAMP

agentops_events
  id            UUID PRIMARY KEY
  event_type    TEXT        ← runtime.started | deployment.completed | health.failed...
  deployment_id UUID
  runtime_id    UUID
  payload       JSONB
  created_at    TIMESTAMP

request_metrics
  id            UUID PRIMARY KEY
  deployment_id UUID
  runtime_id    UUID
  status_code   INTEGER
  latency_ms    INTEGER
  timestamp     TIMESTAMP
```

---

# Redis

Used across all services for high-speed temporary storage.

| Key Pattern | Used By | Purpose |
|---|---|---|
| `session:{userId}` | Auth Service | Session cache |
| `ratelimit:{userId}:{endpoint}` | API Gateway | Rate limiting counter |
| `jwt:blacklist:{jti}` | Auth Service | Revoked tokens |
| `runtime:{runtimeId}` | API Gateway | Runtime URL cache (avoid DB lookup per request) |
| `ws:sessions` | Notification Service | Active WebSocket connections |

---

# Schema Ownership

All SQL schema files live in `Backend/db/schema/`.

The Runtime Service (`Ai-Agent/runtime-service`) connects to the same Platform PostgreSQL but does NOT define schema. It only reads and writes to existing tables.

If the Runtime Service needs a new column: update `Backend/db/schema/runtime.sql`.

---

# Database Connection

```
Backend/db/postgresql/postgresql.go    ← Go DB connection pool
Backend/db/redis/redis.go              ← Go Redis client

Ai-Agent/runtime-service/src/db/      ← TypeScript DB connection (Drizzle ORM)
                                          connects to same Platform PostgreSQL
```
