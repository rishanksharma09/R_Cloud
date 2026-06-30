# 04 — Project Service

> Manages projects and GitHub repository connections.

---

# Overview

The Project Service manages the concept of a Project in R Agent Cloud.

A Project is a container that holds a GitHub repository and its deployment history.

Users create a project, connect a GitHub repository, and deploy from it.

---

# Responsibilities

- Create, list, get, delete projects
- Connect GitHub repository to a project
- Validate GitHub repository access
- List branches
- Store repository metadata
- Trigger sync on GitHub push

---

# Folder Structure

```
Backend/project/
└── handler/
    └── project.go       ← create, list, get, delete project handlers
```

---

# Database

Tables stored in Platform PostgreSQL:

```
projects
  id
  user_id
  name
  description
  github_repo_url
  github_repo_name
  github_owner
  default_branch
  created_at
  updated_at
```

---

# API Endpoints

## Create Project

```http
POST /api/v1/projects

Body:
{
  "name": "customer-support",
  "description": "Customer support AI agent"
}

Response:
{
  "success": true,
  "data": {
    "projectId": "proj_001",
    "name": "customer-support"
  }
}
```

## List Projects

```http
GET /api/v1/projects

Response:
{
  "success": true,
  "data": [
    {
      "projectId": "proj_001",
      "name": "customer-support",
      "repoUrl": "https://github.com/user/customer-support",
      "deploymentCount": 3
    }
  ]
}
```

## Get Project

```http
GET /api/v1/projects/{projectId}
```

## Delete Project

```http
DELETE /api/v1/projects/{projectId}
```

---

# GitHub Integration

## Connect Repository

```http
POST /api/v1/projects/{projectId}/github

Body:
{
  "repoUrl": "https://github.com/user/customer-support",
  "branch": "main"
}
```

Platform validates:
- Repository exists
- Access token has read permissions
- Repository is public or access granted

## Sync Repository

```http
POST /api/v1/projects/{projectId}/sync
```

Manually trigger re-cloning and re-validation of the repository.
