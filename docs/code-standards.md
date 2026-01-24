# Code Standards

This document outlines the coding standards and best practices for the Company Management System project.

## Table of Contents

- [General Principles](#general-principles)
- [Go Language Standards](#go-language-standards)
- [Project Structure](#project-structure)
- [Naming Conventions](#naming-conventions)
- [Error Handling](#error-handling)
- [Database Practices](#database-practices)
- [API Design](#api-design)
- [Testing Standards](#testing-standards)
- [Documentation](#documentation)

## General Principles

### Clean Code

- Write code that is self-documenting and easy to understand
- Follow the principle of least surprise
- Keep functions small and focused on a single responsibility
- Avoid deep nesting (maximum 3-4 levels)
- Use meaningful variable and function names

### SOLID Principles

- **Single Responsibility**: Each module should have one reason to change
- **Open/Closed**: Open for extension, closed for modification
- **Liskov Substitution**: Subtypes must be substitutable for their base types
- **Interface Segregation**: Many specific interfaces are better than one general interface
- **Dependency Inversion**: Depend on abstractions, not concretions

## Go Language Standards

### Code Formatting

- Use `gofmt` to format all Go code
- Follow official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `golangci-lint` for static code analysis

### Package Organization

```go
// Package comment should explain the purpose of the package
package controllers

import (
    // Standard library imports first
    "context"
    "fmt"
    "time"

    // External dependencies second
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"

    // Internal packages last
    "github.com/vlahanam/company-management/internal/models"
    "github.com/vlahanam/company-management/internal/services"
)
```

### Variable Declaration

```go
// Use short variable declaration when possible
user := &models.User{}

// Use var for zero values
var count int
var name string

// Group related declarations
var (
    timeout = 30 * time.Second
    maxRetries = 3
)

// Constants should be grouped and typed
const (
    StatusActive   = "active"
    StatusInactive = "inactive"
)
```

### Function Design

```go
// Good: Clear function signature with meaningful names
func (s *UserService) CreateUser(ctx context.Context, req *requests.CreateUserRequest) (*models.User, error) {
    // Implementation
}

// Bad: Unclear parameters
func Create(c context.Context, r interface{}) (interface{}, error) {
    // Implementation
}
```

### Error Handling

```go
// Always check errors
user, err := s.userRepo.FindByID(ctx, userID)
if err != nil {
    if errors.Is(err, models.ErrUserNotFound) {
        return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
    }
    return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch user")
}

// Use custom error types
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
)

// Wrap errors with context
if err := s.db.Save(&user).Error; err != nil {
    return fmt.Errorf("failed to save user: %w", err)
}
```

### Concurrency

```go
// Always use context for cancellation
func (s *Service) ProcessData(ctx context.Context, data []string) error {
    for _, item := range data {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Process item
        }
    }
    return nil
}

// Use sync.WaitGroup for coordinating goroutines
var wg sync.WaitGroup
for _, user := range users {
    wg.Add(1)
    go func(u *models.User) {
        defer wg.Done()
        // Process user
    }(user)
}
wg.Wait()
```

## Project Structure

### Layered Architecture

The project follows a clean architecture with clear separation of concerns:

```
server/
├── cmd/                    # Application entrypoints
│   ├── main.go            # Main application
│   └── seed/              # Database seeder
├── common/                # Shared utilities
├── database/              # Database migrations
├── internal/              # Private application code
│   ├── controllers/       # HTTP handlers (presentation layer)
│   ├── dto/              # Data transfer objects
│   ├── initialize/       # Application initialization
│   ├── models/           # Domain models (data layer)
│   ├── repositories/     # Data access layer
│   ├── requests/         # Request validation
│   └── services/         # Business logic layer
└── utils/                # Utility functions and middleware
```

### Layer Responsibilities

1. **Controllers** (`internal/controllers/`):
   - Handle HTTP requests and responses
   - Validate request data
   - Call service layer
   - Format responses
   - Should NOT contain business logic

2. **Services** (`internal/services/`):
   - Implement business logic
   - Orchestrate operations between repositories
   - Handle transactions
   - Should NOT know about HTTP

3. **Repositories** (`internal/repositories/`):
   - Database operations only
   - CRUD operations
   - Query building
   - Should NOT contain business logic

4. **Models** (`internal/models/`):
   - Define database schema
   - Data structures
   - Model-specific constants and errors

## Naming Conventions

### Files

- Use snake_case for file names: `user_controller.go`, `auth_service.go`
- Test files: `user_controller_test.go`
- Keep file names descriptive and concise

### Variables

```go
// Use camelCase for variables
var userCount int
var firstName string

// Use descriptive names
// Good
var activeUserCount int
var totalRevenue float64

// Bad
var cnt int
var rev float64

// Boolean variables should be questions
var isActive bool
var hasPermission bool
var canEdit bool
```

### Functions and Methods

```go
// Use PascalCase for exported functions
func CreateUser() {}
func GetUserByID() {}

// Use camelCase for private functions
func validateEmail() {}
func hashPassword() {}

// Use verbs for function names
func (s *UserService) CreateUser() {}
func (s *UserService) UpdateUser() {}
func (s *UserService) DeleteUser() {}
func (r *UserRepository) FindByEmail() {}
```

### Interfaces

```go
// Interface names should be nouns or adjectives ending in -er
type Reader interface {
    Read(p []byte) (n int, err error)
}

type UserRepository interface {
    FindByID(ctx context.Context, id uint64) (*User, error)
    Create(ctx context.Context, user *User) error
}

// Single-method interfaces
type Validator interface {
    Validate() error
}
```

### Constants

```go
// Use PascalCase for exported constants
const (
    DefaultTimeout = 30 * time.Second
    MaxRetries     = 3
)

// Group related constants
const (
    RoleAdmin      = "admin"
    RoleUser       = "user"
    RoleSuperAdmin = "super_admin"
)
```

## Error Handling

### Custom Errors

```go
// Define package-level errors
var (
    ErrUserNotFound       = errors.New("user not found")
    ErrEmailAlreadyExists = errors.New("email already exists")
    ErrInvalidPassword    = errors.New("invalid password")
)
```

### Error Wrapping

```go
// Wrap errors with context
if err := r.db.Create(&user).Error; err != nil {
    return fmt.Errorf("failed to create user: %w", err)
}

// Check wrapped errors
if errors.Is(err, models.ErrUserNotFound) {
    // Handle not found error
}
```

### HTTP Error Responses

```go
// Use Fiber's error handling
if err != nil {
    if errors.Is(err, models.ErrUserNotFound) {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "Internal server error",
    })
}
```

## Database Practices

### GORM Best Practices

```go
// Use transactions for multiple operations
func (s *UserService) CreateUserWithRole(user *models.User, roleID uint64) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(&user).Error; err != nil {
            return err
        }

        userRole := &models.UserRole{
            UserID: user.ID,
            RoleID: roleID,
        }
        if err := tx.Create(&userRole).Error; err != nil {
            return err
        }

        return nil
    })
}

// Use preloading to avoid N+1 queries
func (r *UserRepository) FindByIDWithRoles(id uint64) (*models.User, error) {
    var user models.User
    err := r.db.Preload("Roles").First(&user, id).Error
    return &user, err
}

// Use proper indexing
type User struct {
    Email string `gorm:"uniqueIndex"`
    Phone string `gorm:"index"`
}
```

### Migrations

```sql
-- Use descriptive migration names
-- Format: XXXXXX_description.up.sql

-- Always include comments
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier',
    email VARCHAR(100) UNIQUE NOT NULL COMMENT 'User email address',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp'
);

-- Always create corresponding down migration
DROP TABLE IF EXISTS users;
```

## API Design

### RESTful Endpoints

```
GET    /api/v1/users          # List all users
GET    /api/v1/users/:id      # Get user by ID
POST   /api/v1/users          # Create new user
PUT    /api/v1/users/:id      # Update user
DELETE /api/v1/users/:id      # Delete user
```

### Request Validation

```go
type CreateUserRequest struct {
    FullName    string `json:"full_name"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    PhoneNumber string `json:"phone_number"`
}

func (r CreateUserRequest) Validate() error {
    return validation.ValidateStruct(&r,
        validation.Field(&r.FullName, validation.Required, validation.Length(1, 100)),
        validation.Field(&r.Email, validation.Required, is.Email),
        validation.Field(&r.Password, validation.Required, validation.Length(8, 100)),
    )
}
```

### Response Format

```go
// Success response
{
    "data": {
        "id": 1,
        "full_name": "John Doe",
        "email": "john@example.com"
    }
}

// Error response
{
    "error": "User not found"
}

// List response
{
    "data": [...],
    "meta": {
        "total": 100,
        "page": 1,
        "per_page": 20
    }
}
```

## Testing Standards

### Unit Tests

```go
func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    service := NewUserService(db)

    // Act
    user, err := service.CreateUser(&requests.CreateUserRequest{
        Email: "test@example.com",
    })

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, "test@example.com", user.Email)
}
```

### Table-Driven Tests

```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "test@example.com", false},
        {"invalid email", "invalid", true},
        {"empty email", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("validateEmail() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## Documentation

### Code Comments

```go
// Package controllers provides HTTP request handlers for the application.
package controllers

// UserController handles HTTP requests related to user management.
type UserController struct {
    userService *services.UserService
}

// CreateUser handles the creation of a new user.
// It validates the request, calls the service layer, and returns the created user.
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
    // Implementation
}
```

### README Updates

- Keep README.md up to date with project changes
- Document environment variables
- Update API endpoints when adding new routes
- Maintain clear setup instructions

### API Documentation

- Document all API endpoints
- Include request/response examples
- Specify authentication requirements
- Document error codes

## Security Standards

### Password Handling

```go
// Always hash passwords before storing
import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}
```

### Environment Variables

```go
// Never hardcode secrets
// Bad
const jwtSecret = "my-secret-key"

// Good
jwtSecret := os.Getenv("JWT_SECRET")
```

### Input Validation

- Always validate user input
- Use parameterized queries (GORM handles this)
- Sanitize file uploads
- Validate file types and sizes

## Git Commit Standards

### Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(auth): add JWT refresh token functionality

Implemented refresh token mechanism to allow users to obtain
new access tokens without re-authenticating.

Closes #123
```

```
fix(user): prevent duplicate email registration

Added unique constraint check before creating user to prevent
duplicate email addresses in the database.
```

## Code Review Checklist

- [ ] Code follows project structure and naming conventions
- [ ] All errors are properly handled
- [ ] Input validation is implemented
- [ ] No hardcoded secrets or credentials
- [ ] Comments explain complex logic
- [ ] Tests are written and passing
- [ ] No unused imports or variables
- [ ] Database queries are optimized
- [ ] API responses follow standard format
- [ ] Security best practices are followed
