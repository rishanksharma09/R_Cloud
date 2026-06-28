---

# Agent-to-Agent (A2A) Deployment

R Agent Cloud supports deploying both single-agent and multi-agent systems.

A deployment can consist of multiple collaborating agents that communicate through an A2A workflow while exposing a single public API endpoint.

Example:

```text
                Client
                   │
                   ▼
              POST /execute
                   │
                   ▼
            Planner Agent
             /           \
            ▼             ▼
   Research Agent   Database Agent
            \             /
             ▼           ▼
            Reviewer Agent
                   │
                   ▼
            Final Response
```

The platform validates the agent topology defined in `ragent.yaml`, deploys the complete agent application, manages runtime communication, and exposes a single endpoint for clients.

Example configuration:

```yaml
application:
  name: customer-support
  mode: microservices

agents:
  - id: planner
    entrypoint: planner.py

  - id: researcher
    entrypoint: researcher.py

  - id: reviewer
    entrypoint: reviewer.py

workflow:
  planner:
    - researcher
  researcher:
    - reviewer

routes:
  execute: /execute
  stream: /stream
  health: /health
  metadata: /metadata
```

## Deployment Mode Detection

The platform should parse `ragent.yaml` during AI Project Validation and use `application.mode` to decide deployment structure:

- `mode: monolith` — deploy a single Railway service for the whole application.
- `mode: microservices` — deploy one Railway service per agent and use each agent's `entrypoint` as the startup command.

This ensures the A2A orchestration reflects the actual repository structure.

## Frontend Preview UX

Instead of a blind deployment toggle, the frontend should show the detected mode before deployment.

1. User pastes the GitHub repo link.
2. Backend analyzes `ragent.yaml` in the background.
3. Frontend displays a confirmation card like:

> 📦 **Deployment Mode Detected: Microservices**
> *Your configuration specifies deploying 3 separate agent services (Planner, Researcher, Reviewer) on Railway. Estimated resource usage: 3 containers.*

Then the user confirms and deploys.

Benefits:

- Deploy complete multi-agent systems.
- Validate agent topology before deployment.
- Single public endpoint for the entire workflow.
- Centralized runtime management.
- AgentOps and observability for the complete application.