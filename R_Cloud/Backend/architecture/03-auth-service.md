# 03 — Auth Service

> Authentication and authorization for the platform.

---

# Overview

The Auth Service handles all authentication and authorization for the platform.

It is isolated with its own dedicated database.

No other service stores user credentials or tokens.

---

# Responsibilities

- User registration
- User login
- JWT issuance
- JWT validation
- Refresh tokens
- API key management
- Session management
- Logout

---

# Folder Structure

```
Backend/auth-service/
├── config/
│   └── config.go        ← env vars for auth service
├── handler/
│   └── handler.go       ← login, register, refresh, logout handlers
├── middleware/
│   └── middleware.go    ← JWT validation middleware
├── models/
│   └── models.go        ← User, Token, APIKey structs
└── routes/
    └── routes.go        ← route definitions
```

---

# Database

Auth Service has its own dedicated PostgreSQL database.

It does not share tables with the platform database.

Tables:

```
users
  id
  email
  password_hash
  created_at
  updated_at

sessions
  id
  user_id
  refresh_token_hash
  expires_at
  created_at

api_keys
  id
  user_id
  key_hash
  name
  created_at
  last_used_at
  revoked_at
```

Redis is used for:
- JWT blacklist (on logout)
- Rate limiting login attempts
- Session cache

---

# API Endpoints

## Register

```http
POST /api/v1/auth/register

Body:
{
  "email": "user@example.com",
  "password": "..."
}

Response:
{
  "success": true,
  "data": {
    "accessToken": "...",
    "refreshToken": "..."
  }
}
```

## Login

```http
POST /api/v1/auth/login

Body:
{
  "email": "user@example.com",
  "password": "..."
}

Response:
{
  "success": true,
  "data": {
    "accessToken": "...",
    "refreshToken": "..."
  }
}
```

## Refresh Token

```http
POST /api/v1/auth/refresh

Body:
{
  "refreshToken": "..."
}

Response:
{
  "success": true,
  "data": {
    "accessToken": "..."
  }
}
```

## Logout

```http
POST /api/v1/auth/logout

Header: Authorization: Bearer <accessToken>

Adds token to Redis blacklist.
```

---

# JWT

Access token lifetime: 15 minutes

Refresh token lifetime: 7 days

Payload:

```json
{
  "sub": "user_001",
  "email": "user@example.com",
  "iat": 1234567890,
  "exp": 1234568790
}
```

---

# API Keys

API keys are used by developer applications calling the platform from their backend.

Format: `rac_live_<random>`

Stored as SHA-256 hash in the database.

Never stored in plain text.

---

# Security

- Passwords hashed with bcrypt (cost 12)
- JWT signed with RS256 (asymmetric)
- API keys hashed with SHA-256
- Refresh tokens hashed before storage
- Redis blacklist for revoked tokens
- Rate limiting on login (10 req/min)
- Failed login attempts tracked in Redis
