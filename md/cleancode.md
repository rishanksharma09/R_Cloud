# CleanCode.md

# Clean Code Standards

> Write code that is easy to read, easy to maintain, and easy to extend.

---

# 1. General Principles

* Write code for humans, not the compiler.
* Prefer simple solutions over clever ones.
* Keep code consistent throughout the project.
* Every function, struct, and file should have a clear purpose.
* Refactor instead of adding unnecessary complexity.

---

# 2. Naming

## Variables

Use descriptive names.

Good

```go
userID
retryCount
conversationHistory
requestBody
```

Bad

```go
a
x
tmp
data
```

---

## Functions

Functions should describe **what they do**.

Good

```go
ValidateRequest()
GenerateResponse()
StoreMemory()
CalculateCost()
```

Bad

```go
Do()
Run()
Work()
Data()
```

---

## Boolean Variables

Always use readable names.

```go
isAuthenticated
hasPermission
canRetry
shouldCache
```

---

## Constants

Use constants instead of magic values.

Bad

```go
timeout := 30
```

Good

```go
const RequestTimeout = 30
```

---

# 3. Functions

A function should do **one thing**.

Keep functions small.

Avoid deeply nested conditions.

Prefer early returns.

Bad

```go
if valid {
    if authenticated {
        if allowed {
        }
    }
}
```

Good

```go
if !valid {
    return
}

if !authenticated {
    return
}

if !allowed {
    return
}
```

Avoid too many parameters.

Prefer passing a struct when related values belong together.

---

# 4. Code Organization

* Keep related logic together.
* Remove unused code immediately.
* Remove unused imports.
* Don't leave commented-out code.
* Avoid duplicate logic.
* Extract reusable functions when needed.

---

# 5. Comments

Good code should explain itself.

Use comments only when explaining:

* Why something exists
* Business rules
* Complex algorithms
* Non-obvious decisions

Don't comment obvious code.

Bad

```go
// Increment counter
count++
```

Good

```go
// Prevent infinite retries.
retryCount++
```

---

# 6. Error Handling

Always handle errors.

Never ignore them.

Bad

```go
result, _ := service.Call()
```

Good

```go
result, err := service.Call()
if err != nil {
    return err
}
```

Return meaningful errors.

---

# 7. Configuration

Never hardcode configuration.

Keep all configuration inside:

* `.env`
* Configuration files
* Environment variables

Examples:

* API Keys
* Database URLs
* Ports
* Secrets
* Tokens
* Timeouts

---

# 8. Validation

Validate every external input.

Never trust:

* User input
* API responses
* File contents
* Environment variables

Fail early when data is invalid.

---

# 9. Logging

Log important events.

Examples:

* Application started
* Request received
* External API failed
* Database error

Never log:

* Passwords
* API keys
* Tokens
* Sensitive user information

---

# 10. Formatting

* Use consistent indentation.
* Keep lines reasonably short.
* Group related code together.
* Add blank lines between logical sections.
* Follow the project's formatting rules.

---

# 11. Testing

Every new feature should be testable.

Test:

* Success cases
* Failure cases
* Edge cases

Fix bugs by adding tests whenever possible.

---

# 12. Security

* Never expose secrets in code.
* Sanitize untrusted input.
* Validate all external data.
* Use least-privilege access.
* Never trust client-side validation alone.

---

# 13. Refactoring

Whenever you modify code:

* Improve names.
* Remove duplication.
* Simplify logic.
* Delete dead code.
* Keep behavior unchanged.

Leave the code cleaner than you found it.

---

# Golden Rules

* Keep it simple.
* Use meaningful names.
* One function, one responsibility.
* Avoid duplicate code.
* Handle every error.
* Remove dead code.
* Don't hardcode configuration.
* Write code that another developer can understand without explanation.
