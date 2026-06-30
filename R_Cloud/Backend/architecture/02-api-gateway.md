# 02 — API Gateway

> Single entry point for all client requests.

---

# Overview

The API Gateway is the single entry point for every request to the platform.

No service is directly accessible from outside.

All traffic goes through the API Gateway.

---

# Responsibilities

- JWT validation
- API key validation
- Rate limiting
- Request routing to internal services
- Runtime proxy for `/execute` and `/stream`
- CORS
- Request logging
- OpenTelemetry tracing

---

# Folder Structure

```
Backend/api-gateway/
├── router/
│   └── router.go        ← register all routes
├── middleware/
│   └── middleware.go    ← JWT auth, rate limiting, CORS
├── handlers/
│   └── handler.go       ← route requests to correct service
└── proxy/
    └── proxy.go         ← proxy /execute to Railway runtime URL
```

---

# Request Flow

```text
Client
  ↓
API Gateway
  → validate JWT or API key
  → check rate limit
  → route to service
  ↓
Internal Service
```

---

# Middleware Order

Every request passes through in this order:

```
1. CORS
2. Rate Limiter
3. JWT Validator or API Key Validator
4. Request Logger
5. OpenTelemetry Span
6. Router → Handler
```

---

# Routes

## Auth Routes

```
POST /api/v1/auth/login
POST /api/v1/auth/register
POST /api/v1/auth/refresh
POST /api/v1/auth/logout
```

## Project Routes

```
POST   /api/v1/projects
GET    /api/v1/projects
GET    /api/v1/projects/{projectId}
DELETE /api/v1/projects/{projectId}
```

## GitHub Routes

```
POST /api/v1/projects/{projectId}/github
POST /api/v1/projects/{projectId}/sync
POST /api/v1/webhooks/github
```

## Deployment Routes

```
POST   /api/v1/deployments
GET    /api/v1/deployments
GET    /api/v1/deployments/{deploymentId}
POST   /api/v1/deployments/{deploymentId}/redeploy
DELETE /api/v1/deployments/{deploymentId}
```

## Runtime Routes

```
GET    /api/v1/runtimes
GET    /api/v1/runtimes/{runtimeId}
POST   /api/v1/runtimes/{runtimeId}/restart
POST   /api/v1/runtimes/{runtimeId}/stop
DELETE /api/v1/runtimes/{runtimeId}
```

## Execute Proxy Routes

```
POST /api/v1/deployments/{deploymentId}/execute
POST /api/v1/deployments/{deploymentId}/stream
```

## AgentOps Routes

```
GET /api/v1/agentops/dashboard
GET /api/v1/agentops/metrics/{deploymentId}
GET /api/v1/agentops/logs/{deploymentId}
GET /api/v1/agentops/health/{deploymentId}
```

## WebSocket

```
/ws
```

---

# Runtime Proxy

The most critical responsibility of the API Gateway is proxying execute requests.

```text
Client
  → POST /api/v1/deployments/dep_123/execute
API Gateway
  → validate JWT
  → lookup runtimeUrl from PostgreSQL
     deploymentId = dep_123 → runtimeUrl = https://customer-support.up.railway.app
  → proxy to https://customer-support.up.railway.app/execute
  → return response to client
```

The API Gateway does not inspect or modify the request body.

It only adds auth context and proxies.

---

# Authentication

Two methods supported:

## JWT

Used by frontend and developer dashboard.

```
Authorization: Bearer <token>
```

## API Key

Used by developer's own applications calling the platform.

```
X-API-Key: <api-key>
```

---

# Rate Limiting

Per user, per endpoint.

Stored in Redis.

Default limits:

| Endpoint | Limit |
|---|---|
| /execute | 100 req/min |
| /auth | 10 req/min |
| /deployments | 20 req/min |

---

# Standard Response Format

Success:

```json
{
  "success": true,
  "data": {},
  "message": "Success"
}
```

Error:

```json
{
  "success": false,
  "error": {
    "code": "UNAUTHORIZED",
    "message": "Invalid token."
  }
}
```

---

# HTTP Status Codes

| Code | Meaning |
|---|---|
| 200 | Success |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 422 | Validation Failed |
| 429 | Rate Limited |
| 500 | Internal Error |

---

# WebSocket Events

Events pushed to connected frontend clients:

```
deployment.started
deployment.completed
deployment.failed
runtime.started
runtime.restarted
runtime.stopped
runtime.failed
health.changed
metrics.updated
```
