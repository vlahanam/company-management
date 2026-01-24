# Design Guidelines

This document outlines the design principles, patterns, and architectural decisions for the Company Management System.

## Table of Contents

- [Architectural Principles](#architectural-principles)
- [Domain-Driven Design](#domain-driven-design)
- [API Design](#api-design)
- [Database Design](#database-design)
- [Security Design](#security-design)
- [Error Handling Design](#error-handling-design)
- [Performance Guidelines](#performance-guidelines)
- [Scalability Considerations](#scalability-considerations)

## Architectural Principles

### Clean Architecture

The application follows clean architecture principles with clear boundaries between layers:

```
┌───────────────────────────────────────────┐
│          External Interfaces              │
│     (HTTP, CLI, External Services)        │
└───────────────┬───────────────────────────┘
                │
┌───────────────▼───────────────────────────┐
│        Interface Adapters                 │
│  (Controllers, Presenters, Gateways)      │
└───────────────┬───────────────────────────┘
                │
┌───────────────▼───────────────────────────┐
│        Application Business Rules         │
│          (Use Cases/Services)             │
└───────────────┬───────────────────────────┘
                │
┌───────────────▼───────────────────────────┐
│       Enterprise Business Rules           │
│         (Entities/Models)                 │
└───────────────────────────────────────────┘
```

### Key Principles

#### 1. Separation of Concerns

- Each layer has a single, well-defined responsibility
- Dependencies flow inward (outer layers depend on inner layers, never vice versa)
- Business logic is isolated from infrastructure concerns

#### 2. Dependency Inversion

```go
// Bad: Service depends on concrete implementation
type UserService struct {
    db *gorm.DB  // Direct dependency on GORM
}

// Good: Service depends on abstraction
type UserService struct {
    userRepo UserRepository  // Depends on interface
}

type UserRepository interface {
    FindByID(id uint64) (*User, error)
    // ...
}
```

#### 3. Single Responsibility Principle

```go
// Bad: Controller handling too many responsibilities
func (c *UserController) CreateUser() {
    // Parse request
    // Validate input
    // Hash password
    // Save to database
    // Send email
    // Log action
}

// Good: Each layer handles its responsibility
func (c *UserController) CreateUser() {
    // Parse and validate request
    user, err := c.userService.CreateUser(req)
    // Format and return response
}

func (s *UserService) CreateUser(req *Request) (*User, error) {
    // Business logic only
    user := s.buildUser(req)
    return s.userRepo.Create(user)
}
```

## Domain-Driven Design

### Bounded Contexts

The system is organized into the following bounded contexts:

#### 1. Identity & Access Management Context

- **Entities**: User, Role, Permission
- **Value Objects**: Email, Password, Token
- **Services**: AuthService, UserService, RoleService
- **Responsibilities**: Authentication, authorization, user management

#### 2. Organization Context

- **Entities**: Company, Position
- **Aggregates**: Company (with child companies and positions)
- **Services**: CompanyService, PositionService
- **Responsibilities**: Organizational structure, hierarchy management

#### 3. Employment Context

- **Entities**: Contract
- **Value Objects**: ContractType, ContractStatus, Salary
- **Services**: ContractService
- **Responsibilities**: Employment relationships, contract lifecycle

### Aggregate Design

#### Company Aggregate

```
Company (Aggregate Root)
├── CompanyID (Identity)
├── ParentCompany (Reference to another Company)
├── Positions (Entities within aggregate)
└── Contracts (References to Contract aggregate)
```

**Rules**:

- External entities can only reference the aggregate root (Company)
- Positions are created/updated through the Company aggregate
- Consistency boundary is maintained within the aggregate

#### User Aggregate

```
User (Aggregate Root)
├── UserID (Identity)
├── Roles (References to Role entities)
├── Positions (References via UserPosition join)
└── Contracts (References to Contract aggregate)
```

### Entity vs Value Object

**Entities** (have identity, mutable):

```go
type User struct {
    ID uint64  // Identity
    Email string
    FullName string
}
```

**Value Objects** (no identity, immutable):

```go
type ContractType string
const (
    ContractTypeProbation  ContractType = "Probation"
    ContractTypePermanent  ContractType = "Permanent"
)
```

## API Design

### RESTful Design Principles

#### Resource Naming

```
✅ Good:
GET    /api/v1/users
GET    /api/v1/users/:id
POST   /api/v1/users
PUT    /api/v1/users/:id
DELETE /api/v1/users/:id

❌ Bad:
GET    /api/v1/getUsers
POST   /api/v1/createUser
POST   /api/v1/user/delete
```

#### Nested Resources

```
✅ Good:
GET /api/v1/companies/:company_id/positions
POST /api/v1/companies/:company_id/positions

✅ Also acceptable (for simpler queries):
GET /api/v1/positions?company_id=123

❌ Bad (too deep):
GET /api/v1/companies/:id/positions/:id/users/:id/contracts
```

### HTTP Method Usage

| Method | Purpose                | Idempotent | Safe |
| ------ | ---------------------- | ---------- | ---- |
| GET    | Retrieve resource(s)   | Yes        | Yes  |
| POST   | Create new resource    | No         | No   |
| PUT    | Update entire resource | Yes        | No   |
| PATCH  | Partial update         | No         | No   |
| DELETE | Remove resource        | Yes        | No   |

### Status Code Guidelines

```go
// Success codes
200 OK              // Successful GET, PUT, PATCH, DELETE
201 Created         // Successful POST with resource creation
204 No Content      // Successful DELETE with no response body

// Client error codes
400 Bad Request     // Invalid request payload
401 Unauthorized    // Missing or invalid authentication
403 Forbidden       // Authenticated but no permission
404 Not Found       // Resource doesn't exist
409 Conflict        // Resource conflict (e.g., duplicate email)
422 Unprocessable   // Validation errors

// Server error codes
500 Internal Server Error  // Unexpected server error
503 Service Unavailable    // Temporary unavailability
```

### Request/Response Design

#### Request Structure

```go
// Use clear, descriptive field names
type CreateUserRequest struct {
    FullName    string     `json:"full_name"`
    Email       string     `json:"email"`
    Password    string     `json:"password"`
    PhoneNumber *string    `json:"phone_number,omitempty"`
    DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
}
```

#### Response Structure

```go
// Success response with data
{
    "data": {
        "id": 1,
        "full_name": "John Doe",
        "email": "john@example.com",
        "created_at": "2026-01-20T10:00:00Z"
    }
}

// List response with metadata
{
    "data": [...],
    "meta": {
        "total": 100,
        "page": 1,
        "per_page": 20,
        "total_pages": 5
    }
}

// Error response
{
    "error": "User not found",
    "code": "USER_NOT_FOUND",  // Optional: machine-readable code
    "details": {                // Optional: validation details
        "email": ["email is required"],
        "password": ["password must be at least 8 characters"]
    }
}
```

### Versioning Strategy

```
Current: /api/v1/*

Future versioning options:
1. URL path: /api/v2/users
2. Header: API-Version: v2
3. Query param: /api/users?version=2

Recommendation: Use URL path versioning for clarity
```

## Database Design

### Schema Design Principles

#### 1. Normalization

- Apply 3NF (Third Normal Form) for most tables
- Avoid data duplication
- Use foreign keys to establish relationships

```sql
-- ✅ Good: Normalized
CREATE TABLE contracts (
    id BIGINT PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    company_id BIGINT REFERENCES companies(id),
    position_id BIGINT REFERENCES positions(id)
);

-- ❌ Bad: Denormalized (duplicating user/company data)
CREATE TABLE contracts (
    id BIGINT PRIMARY KEY,
    user_name VARCHAR(100),
    user_email VARCHAR(100),
    company_name VARCHAR(100)
);
```

#### 2. Indexing Strategy

```sql
-- Primary keys (automatic index)
id BIGINT PRIMARY KEY

-- Unique constraints (automatic index)
email VARCHAR(100) UNIQUE

-- Foreign keys (manual index for better performance)
CREATE INDEX idx_contracts_user_id ON contracts(user_id);
CREATE INDEX idx_contracts_company_id ON contracts(company_id);

-- Frequently queried columns
CREATE INDEX idx_users_phone_number ON users(phone_number);

-- Composite indexes for multi-column queries
CREATE INDEX idx_contracts_company_status ON contracts(company_id, status);
```

#### 3. Data Types

```sql
-- ✅ Use appropriate types
id BIGINT                    -- For IDs (supports large datasets)
email VARCHAR(100)           -- Fixed max length
description TEXT             -- Variable, unlimited length
salary DECIMAL(15,2)         -- Precise decimal values
created_at TIMESTAMP         -- Date+time with timezone
is_active BOOLEAN            -- True/false flags

-- ❌ Avoid
id VARCHAR(50)               -- IDs should be numeric
price FLOAT                  -- Use DECIMAL for money
date VARCHAR(20)             -- Use proper DATE/TIMESTAMP
```

#### 4. Naming Conventions

```sql
-- Tables: plural nouns
users, companies, contracts

-- Columns: snake_case
full_name, created_at, company_id

-- Foreign keys: {referenced_table}_id
company_id, user_id, position_id

-- Indexes: idx_{table}_{column(s)}
idx_users_email, idx_contracts_company_status

-- Constraints: {table}_{column}_{type}
users_email_unique, contracts_user_id_fk
```

### Migration Strategy

#### Migration File Structure

```sql
-- 000005_create_contracts_table.up.sql
CREATE TABLE contracts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Contract ID',
    user_id BIGINT NOT NULL COMMENT 'Reference to user',
    company_id BIGINT NOT NULL COMMENT 'Reference to company',

    CONSTRAINT contracts_user_id_fk
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT contracts_company_id_fk
        FOREIGN KEY (company_id) REFERENCES companies(id)
        ON DELETE CASCADE
) COMMENT='Employment contracts';

CREATE INDEX idx_contracts_user_id ON contracts(user_id);
CREATE INDEX idx_contracts_company_id ON contracts(company_id);

-- 000005_create_contracts_table.down.sql
DROP TABLE IF EXISTS contracts;
```

#### Migration Best Practices

- Always create reversible migrations (up + down)
- Add comments to explain schema decisions
- Include indexes in the same migration as table creation
- Test both up and down migrations
- Never modify existing migrations after deployment

### Relationship Patterns

#### One-to-Many

```sql
-- Company → Positions
CREATE TABLE positions (
    id BIGINT PRIMARY KEY,
    company_id BIGINT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);
```

#### Many-to-Many

```sql
-- Users ↔ Roles
CREATE TABLE user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);
```

#### Self-Referential

```sql
-- Company hierarchy
CREATE TABLE companies (
    id BIGINT PRIMARY KEY,
    parent_id BIGINT,
    FOREIGN KEY (parent_id) REFERENCES companies(id)
);
```

## Security Design

### Authentication Flow

```
┌──────────┐                 ┌──────────┐                ┌──────────┐
│  Client  │                 │  Server  │                │    DB    │
└────┬─────┘                 └────┬─────┘                └────┬─────┘
     │                            │                           │
     │  POST /login               │                           │
     │ (email, password)          │                           │
     ├───────────────────────────>│                           │
     │                            │                           │
     │                            │  Find user by email       │
     │                            ├──────────────────────────>│
     │                            │<──────────────────────────┤
     │                            │       User data           │
     │                            │                           │
     │                            │  Verify password (bcrypt) │
     │                            │                           │
     │                            │  Generate JWT tokens      │
     │                            │                           │
     │  200 OK                    │                           │
     │  {access_token, refresh}   │                           │
     │<───────────────────────────┤                           │
     │                            │                           │
```

### Authorization Design

#### RBAC (Role-Based Access Control)

```
User → Roles → Permissions → Resources

Example:
- User "John" has Role "Admin"
- Role "Admin" has Permission "users:write"
- Permission "users:write" allows editing users
```

#### Permission Naming Convention

```
{resource}:{action}

Examples:
- users:read
- users:write
- users:delete
- companies:read
- companies:write
- contracts:read
- contracts:write
```

#### Middleware Design

```go
func RequirePermission(permission string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Extract user from JWT
        user := c.Locals("user").(*models.User)

        // Check if user has permission
        if !user.HasPermission(permission) {
            return fiber.ErrForbidden
        }

        return c.Next()
    }
}

// Usage in routes
app.Get("/users",
    middleware.JWT(),
    middleware.RequirePermission("users:read"),
    userController.GetUsers,
)
```

### Password Security

```go
// Hashing (bcrypt)
func HashPassword(password string) (string, error) {
    // Cost factor: 10 (default) - adjustable based on security needs
    bytes, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )
    return string(bytes), err
}

// Verification
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword(
        []byte(hash),
        []byte(password),
    )
    return err == nil
}
```

### JWT Token Design

```go
// Access Token (short-lived: 15 minutes)
{
    "user_id": 123,
    "email": "john@example.com",
    "roles": ["admin"],
    "exp": 1674234567  // 15 minutes from now
}

// Refresh Token (long-lived: 7 days)
{
    "user_id": 123,
    "type": "refresh",
    "exp": 1674838367  // 7 days from now
}
```

## Error Handling Design

### Error Categories

#### 1. Domain Errors (Business Logic)

```go
var (
    ErrUserNotFound       = errors.New("user not found")
    ErrEmailAlreadyExists = errors.New("email already exists")
    ErrInvalidContract    = errors.New("invalid contract configuration")
)
```

#### 2. Infrastructure Errors

```go
// Database errors
ErrDatabaseConnection = errors.New("database connection failed")
ErrTransactionFailed  = errors.New("transaction failed")

// External service errors
ErrEmailServiceDown = errors.New("email service unavailable")
```

#### 3. Validation Errors

```go
type ValidationError struct {
    Field   string
    Message string
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
    // Format as JSON
}
```

### Error Response Design

```go
// Controller layer
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
    user, err := c.userService.GetByID(id)
    if err != nil {
        return c.errorResponse(err)
    }
    return c.successResponse(user)
}

func (c *BaseController) errorResponse(err error) error {
    switch {
    case errors.Is(err, models.ErrUserNotFound):
        return fiber.NewError(fiber.StatusNotFound, err.Error())
    case errors.Is(err, models.ErrEmailAlreadyExists):
        return fiber.NewError(fiber.StatusConflict, err.Error())
    default:
        log.Error(err)  // Log unexpected errors
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "Internal server error",
        )
    }
}
```

## Performance Guidelines

### Database Query Optimization

#### 1. Use Preloading (Avoid N+1)

```go
// ❌ Bad: N+1 queries
users := []User{}
db.Find(&users)
for _, user := range users {
    db.Model(&user).Association("Roles").Find(&user.Roles)  // N queries
}

// ✅ Good: Single query with JOIN
users := []User{}
db.Preload("Roles").Find(&users)  // 1 query
```

#### 2. Select Only Needed Columns

```go
// ❌ Bad: Select all columns
db.Find(&users)

// ✅ Good: Select specific columns
db.Select("id", "email", "full_name").Find(&users)
```

#### 3. Use Indexes Wisely

```go
// Query: SELECT * FROM users WHERE email = ?
// Ensure index exists:
// CREATE INDEX idx_users_email ON users(email);

// Query: SELECT * FROM contracts WHERE company_id = ? AND status = ?
// Composite index:
// CREATE INDEX idx_contracts_company_status ON contracts(company_id, status);
```

### Caching Strategy (Future Enhancement)

```
┌─────────┐      Cache Miss      ┌──────────┐     ┌──────────┐
│  Client │ ───────────────────> │   API    │ ──> │   DB     │
└─────────┘                       └──────────┘     └──────────┘
     ▲                                  │
     │         Cache Hit                │
     └──────────────────────────────────┘
                (Redis)

Recommended caching:
- User sessions
- Role/Permission lookups
- Company hierarchies
- Frequently accessed static data
```

### Connection Pooling

```go
// GORM automatically handles connection pooling
db.DB().SetMaxIdleConns(10)
db.DB().SetMaxOpenConns(100)
db.DB().SetConnMaxLifetime(time.Hour)
```

## Scalability Considerations

### Horizontal Scaling

```
┌───────────┐
│ Load      │
│ Balancer  │
└─────┬─────┘
      │
      ├────────┬────────┬────────┐
      │        │        │        │
  ┌───▼───┐ ┌─▼────┐ ┌─▼────┐ ┌─▼────┐
  │Server1│ │Server2│ │Server3│ │ServerN│
  └───┬───┘ └──┬───┘ └──┬───┘ └──┬───┘
      └────────┴────────┴────────┘
                   │
            ┌──────▼──────┐
            │   Database  │
            └─────────────┘
```

**Considerations**:

- Stateless server design (JWT tokens, no server-side sessions)
- Shared database (current) or read replicas (future)
- Load balancer (Nginx) already in architecture

### Database Scaling Strategies

#### 1. Read Replicas (Future)

```
┌────────────┐
│   Primary  │ (Writes)
└─────┬──────┘
      │
      ├─────────┬─────────┐
      │         │         │
  ┌───▼───┐ ┌──▼────┐ ┌──▼────┐
  │Replica1│ │Replica2│ │ReplicaN│ (Reads)
  └────────┘ └───────┘ └───────┘
```

#### 2. Partitioning (Future)

- Partition by `company_id` for multi-tenant isolation
- Partition contracts by date range for archival

### Microservices Consideration (Future Evolution)

```
Current: Monolithic
┌────────────────────┐
│   Single Service   │
│  (All Features)    │
└────────────────────┘

Future: Microservices
┌─────────┐  ┌──────────┐  ┌──────────┐
│  Auth   │  │ Company  │  │ Contract │
│ Service │  │ Service  │  │ Service  │
└─────────┘  └──────────┘  └──────────┘
```

## Design Decision Log

### ADR 001: Layered Architecture

**Status**: Accepted  
**Decision**: Use layered architecture with Controllers → Services → Repositories  
**Rationale**: Clear separation of concerns, testability, maintainability  
**Consequences**: More files, but clearer responsibility boundaries

### ADR 002: JWT for Authentication

**Status**: Accepted  
**Decision**: Use JWT tokens (access + refresh)  
**Rationale**: Stateless, scalable, works well with SPAs  
**Consequences**: Cannot revoke tokens easily, need short expiration times

### ADR 003: GORM as ORM

**Status**: Accepted  
**Decision**: Use GORM for database operations  
**Rationale**: Feature-rich, good Go support, automatic migrations  
**Consequences**: Some SQL complexity abstracted, learning curve

### ADR 004: RBAC for Authorization

**Status**: Accepted  
**Decision**: Implement Role-Based Access Control  
**Rationale**: Flexible permission system, industry standard  
**Consequences**: Complex setup, but highly maintainable

---

**Maintained by**: Development Team  
**Last Updated**: January 2026  
**Review Cycle**: Quarterly
