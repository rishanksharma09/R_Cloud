# 08 — Notification Service

> Real-time and async notifications to users.

---

# Overview

The Notification Service is responsible for delivering notifications to users about deployment and runtime events.

---

# Responsibilities

- WebSocket real-time updates to frontend
- Email notifications for critical events
- Future: Discord, Slack

---

# Notification Types

## WebSocket (Real-time)

Used by the dashboard frontend.

Events pushed to connected clients:

```
deployment.started
deployment.completed
deployment.failed
runtime.started
runtime.restarted
runtime.stopped
runtime.failed
health.changed
```

## Email

Sent for critical events:

- Deployment failed
- Runtime failed after max restarts
- Runtime stopped unexpectedly

---

# Flow

```text
NATS event received
  → runtime.failed
  → Notification Service
  → WebSocket push to connected user
  → Email to user
```

---

# Future

- Discord webhooks
- Slack integration
- Custom webhooks
- SMS alerts
