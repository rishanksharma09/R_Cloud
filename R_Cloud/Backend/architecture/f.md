

---

# gRPC

Only one synchronous communication exists.

```
Deployment Service

        │

        │ gRPC

        ▼

Runtime Service
```

---

# OpenTelemetry

Instrumented Services

* API Gateway
* Deployment Service
* Runtime Service
* AgentOps Service

---

# Team Ownership

## 👤 You (Go)

* API Gateway
* Project Service
* Deployment Service
* Validation Service
* Shared Events
* gRPC Client

---

## 👤 Backend Developer

* Runtime Service
* Railway Integration
* Runtime Registry
* Health Monitoring
* gRPC Server

---

## 👤 Backend Developer

* Authentication Service
* Notification Service

---

# Why is `shared/events` outside every service?

Because **NATS is your platform event bus**, not a feature of one service.

Every service imports the same event definitions.

Example

Deployment publishes

```
deployment.created
```

Runtime subscribes.

Runtime publishes

```
runtime.started
```

Notification subscribes.

Notification sends an email.

AgentOps also subscribes and updates the dashboard.

No service calls another service directly through NATS—the event bus coordinates them all.
