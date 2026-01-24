# Codebase Summary

## Project Overview

The Company Management System is a comprehensive full-stack application designed to manage multiple companies, employees, positions, contracts, and permissions. It features a robust Role-Based Access Control (RBAC) system with a Go backend and Next.js frontend, built with modern web technologies for scalability and maintainability.

## Technology Stack

### Backend Technologies

- **Go** (1.25.3) - Primary backend language
- **Fiber** (v2) - Fast, Express-inspired web framework
- **GORM** - Feature-rich ORM for database operations
- **MySQL** (8.0) - Relational database

### Frontend Technologies

- **Next.js** (16.1.4) - React framework with server-side rendering
- **React** (19.2.3) - UI library
- **TypeScript** (5.x) - Type-safe JavaScript
- **Tailwind CSS** (4.x) - Utility-first CSS framework
- **pnpm** (9.15.4) - Fast, disk-efficient package manager

### Key Libraries

- **JWT** (`golang-jwt/jwt/v5`) - Authentication and authorization
- **ozzo-validation** (`go-ozzo/ozzo-validation/v4`) - Request validation
- **bcrypt** (`golang.org/x/crypto`) - Password hashing
- **godotenv** - Environment variable management
- **golang-migrate** - Database migration tool

### Infrastructure

- **Docker** & **Docker Compose** - Containerization and orchestration
- **Nginx** - Reverse proxy, load balancer, and static file serving
- **Air** - Hot-reload for backend development
- **Turbopack** - Fast bundler for Next.js development

## Architecture

### Layered Architecture Pattern

The application follows a clean, layered architecture with clear separation of concerns:

```
┌─────────────────────────────────────┐
│     Presentation Layer               │
│     (HTTP Handlers/Controllers)      │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Business Logic Layer             │
│     (Services/Use Cases)             │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Data Access Layer                │
│     (Repositories)                   │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Database Layer                   │
│     (MySQL + GORM)                   │
└─────────────────────────────────────┘
```

### Key Architectural Principles

1. **Separation of Concerns**: Each layer has a specific responsibility
2. **Dependency Injection**: Services and repositories are injected into controllers
3. **Interface-Based Design**: Repository interfaces allow for easy testing and mocking
4. **Transaction Management**: GORM transactions ensure data consistency
5. **Middleware Pattern**: Authentication, logging, and error handling via middleware

## Directory Structure

```
server/
├── cmd/                           # Application entry points
│   ├── main.go                   # Main server application
│   └── seed/                     # Database seeder utility
│       └── main.go
│
├── common/                        # Shared utilities and constants
│   ├── common.go                 # Common types and utilities
│   ├── constants.go              # Application-wide constants
│   ├── file_type.go             # File type definitions
│   └── pagination.go            # Pagination utilities
│
├── database/                      # Database-related files
│   └── migrations/               # SQL migration files (18 files)
│       ├── 000001_create_users_table.up.sql
│       ├── 000002_create_companies_table.up.sql
│       └── ... (and corresponding .down.sql files)
│
├── internal/                      # Private application code
│   ├── controllers/              # HTTP request handlers (7 files)
│   │   ├── auth_controller.go
│   │   ├── company_controller.go
│   │   ├── contract_controller.go
│   │   ├── permission_controller.go
│   │   ├── position_controller.go
│   │   ├── role_controller.go
│   │   └── user_controller.go
│   │
│   ├── dto/                      # Data Transfer Objects (2 files)
│   │   ├── auth_dto.go
│   │   └── user_dto.go
│   │
│   ├── initialize/               # Application initialization (5 files)
│   │   ├── loadconfig.go        # Environment configuration
│   │   ├── logger.go            # Logging setup
│   │   ├── mysql.go             # Database connection
│   │   ├── router.go            # Route registration
│   │   └── run.go               # Application startup
│   │
│   ├── models/                   # Database models (10 files)
│   │   ├── company.go
│   │   ├── contract.go
│   │   ├── permission.go
│   │   ├── position.go
│   │   ├── role.go
│   │   ├── role_permission.go
│   │   ├── sql_model.go         # Base model with timestamps
│   │   ├── user.go
│   │   ├── user_position.go
│   │   └── user_role.go
│   │
│   ├── repositories/             # Data access layer (7 files)
│   │   ├── company_repo.go
│   │   ├── contract_repo.go
│   │   ├── gorm.go              # Base repository interface
│   │   ├── permission_repo.go
│   │   ├── position_repo.go
│   │   ├── role_repo.go
│   │   └── user_repo.go
│   │
│   ├── requests/                 # Request validation structs (8 files)
│   │   ├── auth_request.go
│   │   ├── common.go
│   │   ├── company_request.go
│   │   ├── contract_request.go
│   │   ├── permission_request.go
│   │   ├── position_request.go
│   │   ├── role_request.go
│   │   └── user_request.go
│   │
│   └── services/                 # Business logic layer (7 files)
│       ├── auth_service.go
│       ├── company_service.go
│       ├── contract_service.go
│       ├── permission_service.go
│       ├── position_service.go
│       ├── role_service.go
│       └── user_service.go
│
├── utils/                         # Utility functions (3 files)
│   ├── jwt.go                    # JWT token generation/validation
│   ├── middleware.go             # HTTP middleware
│   └── response.go               # Response helpers
│
├── .air.toml                     # Air hot-reload configuration
├── .env                          # Environment variables (not in git)
├── .env.example                  # Environment template
├── Dockerfile                    # Docker image definition
├── go.mod                        # Go module dependencies
└── go.sum                        # Dependency checksums

client/
├── app/                          # Next.js app directory
│   ├── layout.tsx               # Root layout component
│   ├── page.tsx                 # Home page component
│   └── globals.css              # Global styles
│
├── public/                       # Static assets
│   ├── next.svg                 # Next.js logo
│   └── vercel.svg               # Vercel logo
│
├── .dockerignore                 # Docker build exclusions
├── .gitignore                    # Git exclusions
├── Dockerfile                    # Multi-stage Docker image
├── eslint.config.mjs             # ESLint configuration
├── next.config.ts                # Next.js configuration
├── next-env.d.ts                 # Next.js TypeScript declarations
├── package.json                  # Node.js dependencies
├── pnpm-lock.yaml                # pnpm lockfile
├── pnpm-workspace.yaml           # pnpm workspace config
├── postcss.config.mjs            # PostCSS configuration
├── README.md                     # Client documentation
└── tsconfig.json                 # TypeScript configuration
```

## Core Modules

### 1. Authentication & Authorization Module

**Files**: `auth_controller.go`, `auth_service.go`, `auth_request.go`, `jwt.go`

**Responsibilities**:

- User login and registration
- JWT token generation (access + refresh tokens)
- Token validation and refresh
- Password hashing with bcrypt

**Key Features**:

- Dual-token system (access token + refresh token)
- Secure password hashing
- Token expiration management

### 2. User Management Module

**Files**: `user_*.go` across controllers, services, repositories, models

**Responsibilities**:

- CRUD operations for users
- User profile management
- Avatar handling
- User-position associations
- User-role assignments

**Database Schema**:

```sql
- id (Primary Key)
- full_name
- hash_password
- email (Unique)
- phone_number (Unique)
- date_of_birth
- gender
- id_card_number (Unique)
- avatar
- created_at, updated_at
```

### 3. Company Management Module

**Files**: `company_*.go` across controllers, services, repositories, models

**Responsibilities**:

- Company CRUD operations
- Multi-company hierarchy (parent-child relationships)
- Company details management (address, contact info, founded date)

**Database Schema**:

```sql
- id (Primary Key)
- name
- parent_id (Self-referential foreign key)
- description
- founded_date
- address
- phone_number
- email
- created_at, updated_at
```

**Key Features**:

- Hierarchical company structure
- Company-specific positions and contracts

### 4. Position Management Module

**Files**: `position_*.go` across controllers, services, repositories, models

**Responsibilities**:

- Position CRUD within companies
- Position hierarchy tracking
- Salary range management
- User-position associations

**Database Schema**:

```sql
- id (Primary Key)
- company_id (Foreign Key)
- title
- description
- level
- min_salary, max_salary
- created_at, updated_at
```

### 5. Contract Management Module

**Files**: `contract_*.go` across controllers, services, repositories, models

**Responsibilities**:

- Employment contract tracking
- Contract lifecycle management
- Multiple contract type support
- Contract status tracking

**Contract Types**:

- Probation
- Fixed-term
- Permanent
- Freelance
- Internship

**Contract Statuses**:

- Active
- Expired
- Terminated
- Pending

**Database Schema**:

```sql
- id (Primary Key)
- user_id (Foreign Key)
- company_id (Foreign Key)
- position_id (Foreign Key)
- contract_number (Unique)
- contract_type (Enum)
- start_date
- end_date
- salary
- status (Enum)
- file_path
- notes
- created_at, updated_at
```

### 6. RBAC (Role-Based Access Control) Module

**Files**: `role_*.go`, `permission_*.go`, `user_role.go`, `role_permission.go`

**Responsibilities**:

- Role management
- Permission assignment
- Role-permission associations
- User-role assignments
- Access control enforcement

**Default Roles**:

- **Super Admin**: Full system access
- **Admin**: Company-level administration
- **User**: Basic user access

**Database Schema**:

**Roles Table**:

```sql
- id (Primary Key)
- name (Unique)
- description
- created_at, updated_at
```

**Permissions Table**:

```sql
- id (Primary Key)
- name (Unique)
- resource
- action
- description
- created_at, updated_at
```

**Join Tables**:

- `user_roles`: Many-to-many (users ↔ roles)
- `role_permissions`: Many-to-many (roles ↔ permissions)

## Data Flow

### Typical Request Flow

1. **HTTP Request** → Nginx → Fiber Server
2. **Middleware** → JWT Validation → Authentication Check
3. **Controller** → Parse Request → Validate Input
4. **Service** → Business Logic → Transaction Management
5. **Repository** → Database Query → GORM Operations
6. **Database** → MySQL → Data Persistence
7. **Response** → DTO Mapping → JSON Response

### Example: User Login Flow

```
1. POST /api/v1/login
   ↓
2. AuthController.Login()
   - Parse LoginRequest
   - Validate input (email, password)
   ↓
3. AuthService.Login()
   - Find user by email
   - Verify password with bcrypt
   - Generate JWT tokens
   ↓
4. UserRepository.FindByEmail()
   - Query database via GORM
   ↓
5. Return Auth Response
   - access_token
   - refresh_token
```

## Key Design Patterns

### 1. Repository Pattern

- Abstracts data access layer
- Enables easy testing with mocks
- Centralizes database queries

```go
type UserRepository interface {
    FindByID(id uint64) (*User, error)
    FindByEmail(email string) (*User, error)
    Create(user *User) error
    Update(user *User) error
    Delete(id uint64) error
}
```

### 2. Service Pattern

- Encapsulates business logic
- Orchestrates multiple repositories
- Manages transactions

```go
type UserService struct {
    userRepo     UserRepository
    roleRepo     RoleRepository
    db           *gorm.DB
}
```

### 3. DTO Pattern

- Separates internal models from API responses
- Controls data exposure
- Formats responses consistently

```go
type UserDTO struct {
    ID          uint64  `json:"id"`
    FullName    string  `json:"full_name"`
    Email       string  `json:"email"`
    // Password is never exposed
}
```

### 4. Middleware Pattern

- Cross-cutting concerns (auth, logging, CORS)
- Request/response manipulation
- Error handling

```go
// JWT Authentication Middleware
func JWTMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Verify JWT token
        // Attach user to context
        return c.Next()
    }
}

// CORS Middleware
app.Use(cors.New(cors.Config{
    AllowOrigins:     "http://localhost:3030",
    AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
    AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
    AllowCredentials: true,
}))
```

## Database Design

### Entity Relationships

```
Users ←→ Companies (via Contracts)
Users ←→ Positions (via user_positions, many-to-many)
Users ←→ Roles (via user_roles, many-to-many)
Roles ←→ Permissions (via role_permissions, many-to-many)
Companies → Positions (one-to-many)
Companies → Contracts (one-to-many)
Companies → Companies (self-referential, parent-child)
Positions → Contracts (one-to-many)
```

### Migration Strategy

- **Version-controlled migrations** using golang-migrate
- **Up/Down migrations** for all schema changes
- **Naming convention**: `XXXXXX_description.up.sql` / `.down.sql`
- **Comments** in SQL for documentation
- **Idempotent migrations** where possible

## Configuration Management

### Environment Variables

Located in `server/.env`:

```env
# Server
GO_ENV=development
SERVER_PORT=8080

# Database
DB_HOST=company-management-mysql-dev
DB_PORT=3306
DB_USER=dev_user
DB_PASSWORD=dev_password
DB_NAME=company_db

# JWT
ACCESS_SECRET=your-access-token-secret
REFRESH_SECRET=your-refresh-token-secret

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3030

# Timezone
TZ=Asia/Ho_Chi_Minh
```

### Configuration Loading

Handled in `initialize/loadconfig.go`:

- Loads `.env` file using `godotenv`
- Provides type-safe access to configuration
- Validates required variables on startup

## API Structure

### Versioning

- API version prefix: `/api/v1/`
- Allows for future API iterations without breaking changes

### Authentication

- Protected endpoints require `Authorization: Bearer <token>` header
- JWT middleware validates tokens and extracts user information

### Response Format

**Success Response**:

```json
{
  "data": { ... }
}
```

**Error Response**:

```json
{
  "error": "Description of error"
}
```

### Endpoint Grouping

- `/api/v1/auth/*` - Authentication (login, register, refresh)
- `/api/v1/users/*` - User management
- `/api/v1/companies/*` - Company management
- `/api/v1/positions/*` - Position management
- `/api/v1/contracts/*` - Contract management
- `/api/v1/roles/*` - Role management
- `/api/v1/permissions/*` - Permission management

## Deployment Architecture

### Development Environment

```
┌─────────────┐
│   Nginx     │ :880
│ (Dev Proxy) │
└──┬─────┬────┘
   │     │
   │     └────────────────┐
   │                      │
┌──▼──────┐        ┌──────▼──────┐
│ Next.js │        │  Go Server  │
│ Client  │ :3030  │  (with Air) │ :8808
└─────────┘        └──────┬──────┘
                          │
                   ┌──────▼──────┐
                   │  MySQL      │ :33066
                   │  (Dev DB)   │
                   └─────────────┘
```

### Production Environment

```
┌─────────────┐
│   Nginx     │ :80, :443
│(Prod Proxy) │
└──┬─────┬────┘
   │     │
   │     └────────────────┐
   │                      │
┌──▼──────┐        ┌──────▼──────┐
│ Next.js │        │  Go Server  │
│ Client  │ :3000  │(Production) │ :8080
└─────────┘        └──────┬──────┘
                          │
                   ┌──────▼──────┐
                   │  MySQL      │ :3306
                   │ (Prod DB)   │
                   └─────────────┘
```

### Docker Containers

- **company-management-client**: Next.js 16 application with Turbopack (dev) or standalone build (prod)
- **company-management-server**: Go application with Fiber
- **company-management-mysql**: MySQL 8.0 database
- **company-management-nginx**: Nginx reverse proxy

## Testing Strategy

### Current State

- Project structure supports unit testing
- Test files follow `*_test.go` convention
- Future implementation planned for:
  - Repository layer tests
  - Service layer tests
  - Controller integration tests

### Recommended Test Structure

```
internal/
├── repositories/
│   ├── user_repo.go
│   └── user_repo_test.go
├── services/
│   ├── user_service.go
│   └── user_service_test.go
└── controllers/
    ├── user_controller.go
    └── user_controller_test.go
```

## Development Workflow

### Hot Reload

- **Air** configured for automatic server restart on code changes
- Watches `.go` files and configuration
- Configuration in `.air.toml`

### Database Management

- **Migrations**: Create, apply, rollback via Makefile
- **Seeding**: Populate test data via `cmd/seed/main.go`
- **Backup/Restore**: Available through Makefile commands

### Code Organization Best Practices

1. **Models** define data structure only
2. **Repositories** handle database operations
3. **Services** contain business logic and orchestration
4. **Controllers** handle HTTP and validation
5. **DTOs** control response formatting
6. **Requests** validate incoming data

## Security Considerations

### Implemented

- ✅ Password hashing with bcrypt
- ✅ JWT-based authentication
- ✅ Environment variable configuration
- ✅ Input validation via ozzo-validation
- ✅ GORM SQL injection protection
- ✅ CORS middleware for cross-origin requests

### Recommended Enhancements

- [ ] Rate limiting for API endpoints
- [ ] HTTPS in production (Nginx SSL configuration)
- [ ] Database connection encryption
- [ ] Audit logging for sensitive operations
- [ ] API key management for external services

## Dependencies Summary

### Main Dependencies

```go
github.com/gofiber/fiber/v2       // Web framework
gorm.io/gorm                       // ORM
gorm.io/driver/mysql               // MySQL driver
github.com/golang-jwt/jwt/v5       // JWT
github.com/go-ozzo/ozzo-validation // Validation
golang.org/x/crypto                // Bcrypt
github.com/joho/godotenv           // Environment
```

## Performance Considerations

### Database

- Indexed columns: `email`, `phone_number`, `id_card_number`
- GORM preloading for avoiding N+1 queries
- Connection pooling via GORM

### API

- Fiber's fast HTTP routing
- Minimal middleware stack
- JSON serialization optimization

## Maintenance and Operations

### Makefile Commands

- Development: `make dev`, `make logs-dev`
- Production: `make prod`, `make logs-prod`
- Database: `make migrate-up`, `make seed`
- Cleanup: `make clean`

### Monitoring

- Health check endpoint: `GET /health`
- Container logs: `make logs-dev` / `make logs-prod`
- Database access: `make db-dev` / `make db-prod`

## Future Enhancement Opportunities

1. **API Documentation**: Integrate Swagger/OpenAPI
2. **Caching**: Redis for session and query caching
3. **Message Queue**: RabbitMQ/Kafka for async operations
4. **File Storage**: S3 integration for avatars and contract files
5. **Email Service**: SMTP integration for notifications
6. **Audit Logging**: Track all CRUD operations
7. **Multi-tenancy**: Enhanced company isolation
8. **GraphQL**: Alternative API layer
9. **WebSockets**: Real-time notifications
10. **Background Jobs**: Cron jobs for contract expiration checks

---

**Last Updated**: January 2026  
**Go Version**: 1.25.3  
**Database Version**: MySQL 8.0  
**Total Files**: 77+ in server directory
