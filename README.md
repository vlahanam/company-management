# Company Management System

A comprehensive company management system built with Go, featuring employee management, contract tracking, role-based access control (RBAC), and multi-company support.

## 📋 Overview

This system provides a complete solution for managing companies, employees, positions, contracts, and permissions. It's designed with scalability and security in mind, featuring a clean architecture with separation of concerns.

### Key Features

- 👥 **User Management**: Complete employee lifecycle management with profiles, avatars, and personal information
- 🏢 **Multi-Company Support**: Manage multiple companies within a single system
- 📝 **Contract Management**: Track employment contracts with various types (Probation, Fixed-term, Permanent, Freelance, Internship)
- 🔐 **Role-Based Access Control (RBAC)**: Fine-grained permission system with customizable roles and permissions
- 📍 **Position Management**: Define and assign positions within companies
- 🔒 **JWT Authentication**: Secure authentication with access and refresh tokens
- 🐳 **Docker Support**: Complete containerization for development and production environments
- 🔄 **Database Migrations**: Version-controlled schema migrations using golang-migrate
- 🌱 **Database Seeding**: Automated data seeding for development and testing

## 🏗️ Architecture

```mermaid
graph TB
    subgraph "Client Layer"
        Client[Client Application]
    end

    subgraph "Infrastructure Layer"
        Nginx[Nginx Reverse Proxy]
    end

    subgraph "Application Layer"
        Server[Go Fiber Server]
        Auth[JWT Authentication]
        Middleware[Middleware Layer]
    end

    subgraph "Business Layer"
        Controllers[Controllers]
        Services[Services/Use Cases]
        Repos[Repositories]
    end

    subgraph "Data Layer"
        MySQL[(MySQL Database)]
        Migrations[Database Migrations]
    end

    Client --> Nginx
    Nginx --> Server
    Server --> Auth
    Auth --> Middleware
    Middleware --> Controllers
    Controllers --> Services
    Services --> Repos
    Repos --> MySQL
    Migrations --> MySQL
```

### Project Structure

```
company-management/
├── client/                     # Client application (placeholder)
├── server/                     # Go backend application
│   ├── cmd/                   # Application entrypoints
│   │   ├── main.go           # Main server application
│   │   └── seed/             # Database seeder
│   ├── common/               # Shared utilities and constants
│   ├── database/             # Database-related files
│   │   └── migrations/       # SQL migration files
│   ├── internal/             # Private application code
│   │   ├── controllers/      # HTTP request handlers
│   │   ├── dto/              # Data Transfer Objects
│   │   ├── initialize/       # App initialization (router, config, db)
│   │   ├── models/           # Database models
│   │   ├── repositories/     # Data access layer
│   │   ├── requests/         # Request validation structs
│   │   └── services/         # Business logic layer
│   ├── utils/                # Utility functions and middleware
│   ├── .air.toml             # Hot-reload configuration
│   ├── Dockerfile            # Docker image definition
│   ├── go.mod                # Go module dependencies
│   └── .env                  # Environment variables (not in git)
├── docker/                    # Docker configuration
│   ├── docker-compose.yml    # Development environment
│   ├── docker-compose.prod.yml # Production environment
│   └── mysql/                # MySQL initialization scripts
├── nginx/                     # Nginx configuration
│   ├── nginx.dev.conf        # Development config
│   └── conf.d/               # Additional nginx configs
├── Makefile                   # Build and deployment automation
└── README.md                  # This file
```

## 🛠️ Technology Stack

### Backend

- **[Go](https://golang.org/)** (1.25.3) - Primary programming language
- **[Fiber](https://gofiber.io/)** (v2) - Fast HTTP web framework
- **[GORM](https://gorm.io/)** - ORM library for database operations
- **[JWT](https://github.com/golang-jwt/jwt)** - JSON Web Tokens for authentication
- **[ozzo-validation](https://github.com/go-ozzo/ozzo-validation)** - Request validation
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable management

### Database

- **[MySQL](https://www.mysql.com/)** (8.0) - Primary database
- **[golang-migrate](https://github.com/golang-migrate/migrate)** - Database migrations

### Infrastructure

- **[Docker](https://www.docker.com/)** - Containerization
- **[Docker Compose](https://docs.docker.com/compose/)** - Multi-container orchestration
- **[Nginx](https://nginx.org/)** - Reverse proxy and load balancer
- **[Air](https://github.com/cosmtrek/air)** - Hot-reload for Go applications (development)

## 📊 Database Schema

### Core Tables

- **users**: Employee information and authentication credentials
- **companies**: Company details
- **positions**: Job positions within companies
- **user_positions**: Many-to-many relationship between users and positions
- **contracts**: Employment contracts linking users, companies, and positions
- **roles**: User roles for RBAC
- **permissions**: System permissions
- **user_roles**: User role assignments
- **role_permissions**: Permission assignments to roles

### Default Roles

The system includes three predefined roles:

- **Super Admin**: Full system access
- **Admin**: Company-level administrative access
- **User**: Basic user access

## 🚀 Getting Started

### Prerequisites

- **Docker** (20.10+) and **Docker Compose** (2.0+)
- **Make** (for using Makefile commands)
- **Go** (1.25.3+) - Only if running locally without Docker
- **MySQL** (8.0+) - Only if running locally without Docker

### Quick Start

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd company-management
   ```

2. **Configure environment variables**

   ```bash
   # Copy example environment files
   cp server/.env.example server/.env
   cp docker/.env.example docker/.env

   # Edit the .env files with your configuration
   # At minimum, set database credentials and JWT secrets
   ```

3. **Start the development environment**

   ```bash
   make dev
   ```

   This command will:
   - Build Docker images
   - Start MySQL, Go server, and Nginx containers
   - Run database migrations
   - Start the server with hot-reload enabled

4. **Access the application**
   - API Server: http://localhost:8808
   - Nginx Proxy: http://localhost:880
   - MySQL: localhost:33066

5. **Seed the database (optional)**
   ```bash
   make seed
   ```

### Environment Variables

Key environment variables in `server/.env`:

```env
# Server Configuration
GO_ENV=development
SERVER_PORT=8080

# Database Configuration
DB_HOST=company-management-mysql-dev
DB_PORT=3306
DB_USER=dev_user
DB_PASSWORD=dev_password
DB_NAME=company_db

# JWT Configuration
ACCESS_SECRET=your-access-token-secret
REFRESH_SECRET=your-refresh-token-secret

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3030

# Timezone
TZ=Asia/Ho_Chi_Minh
```

## 📖 Usage

### Makefile Commands

The project includes a comprehensive Makefile for common operations:

#### Development Environment

```bash
make dev              # Build and start development environment
make build-dev        # Build development Docker images
make up-dev          # Start development containers
make down-dev        # Stop development containers
make restart-dev     # Restart development containers
make logs-dev        # Show development logs
make db-dev          # Connect to development MySQL
```

#### Production Environment

```bash
make prod            # Build and start production environment
make build-prod      # Build production Docker images
make up-prod        # Start production containers
make down-prod      # Stop production containers
make restart-prod   # Restart production containers
make logs-prod      # Show production logs
make db-prod        # Connect to production MySQL
```

#### Database Migrations

```bash
make migrate-create      # Create new migration file
make migrate-up         # Apply all pending migrations
make migrate-down       # Rollback last migration
make migrate-version    # Show current migration version
make migrate-force      # Force set migration version
make migrate-drop       # Drop all tables (DANGEROUS!)
```

#### Database Seeding

```bash
make seed           # Run seeder in Docker container
make seed-local     # Run seeder from local machine
```

#### Utility Commands

```bash
make clean          # Remove all containers, volumes, and images
make status         # Show status of all containers
make ps             # Show running containers
make help           # Display all available commands
```

### API Endpoints

#### Authentication

- `POST /api/v1/login` - User login
- `POST /api/v1/register` - User registration
- `POST /api/v1/refresh` - Refresh access token

#### Users (Protected)

- `GET /api/v1/users` - List all users (Super Admin only)
- `GET /api/v1/users/:id` - Get user details
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

#### Companies (Protected)

- `POST /api/v1/companies` - Create company
- `GET /api/v1/companies` - List all companies
- `GET /api/v1/companies/:id` - Get company details
- `PUT /api/v1/companies/:id` - Update company
- `DELETE /api/v1/companies/:id` - Delete company

#### Positions (Protected)

- `POST /api/v1/positions/:company_id` - Create position
- `GET /api/v1/positions/:company_id` - List positions by company
- `GET /api/v1/positions/:id` - Get position details
- `PUT /api/v1/positions/:id` - Update position
- `DELETE /api/v1/positions/:id` - Delete position

#### Contracts (Protected)

- `POST /api/v1/contracts` - Create contract
- `GET /api/v1/contracts` - List all contracts
- `GET /api/v1/contracts/:id` - Get contract details
- `PUT /api/v1/contracts/:id` - Update contract
- `DELETE /api/v1/contracts/:id` - Delete contract

#### Roles (Protected)

- `POST /api/v1/roles` - Create role
- `GET /api/v1/roles` - List all roles
- `GET /api/v1/roles/:id` - Get role details
- `PUT /api/v1/roles/:id` - Update role
- `DELETE /api/v1/roles/:id` - Delete role

#### Permissions (Protected)

- `POST /api/v1/permissions` - Create permission
- `GET /api/v1/permissions` - List all permissions
- `GET /api/v1/permissions/:id` - Get permission details
- `PUT /api/v1/permissions/:id` - Update permission
- `DELETE /api/v1/permissions/:id` - Delete permission

#### Health Check

- `GET /health` - API health check

> **Note**: All protected endpoints require a valid JWT token in the Authorization header: `Authorization: Bearer <token>`

## 🔧 Development

### Running Without Docker

1. **Start MySQL**

   ```bash
   # Using your local MySQL installation
   mysql -u root -p
   CREATE DATABASE company_db;
   ```

2. **Run migrations**

   ```bash
   cd server
   migrate -path database/migrations -database "mysql://user:password@tcp(localhost:3306)/company_db?multiStatements=true" up
   ```

3. **Run the server**
   ```bash
   cd server
   go run cmd/main.go
   ```

### Hot Reload (Development)

The development environment uses [Air](https://github.com/cosmtrek/air) for hot-reload functionality. Any changes to `.go` files will automatically rebuild and restart the server.

Configuration is in `server/.air.toml`.

### Creating Database Migrations

1. **Create migration files**

   ```bash
   make migrate-create
   # Enter migration name when prompted (e.g., add_column_to_users)
   ```

2. **Edit the generated files**
   - `server/database/migrations/XXXXXX_<name>.up.sql` - Schema changes
   - `server/database/migrations/XXXXXX_<name>.down.sql` - Rollback changes

3. **Apply migrations**
   ```bash
   make migrate-up
   ```

### Adding New Features

The application follows a layered architecture:

1. **Model** (`internal/models/`) - Define the database model
2. **Migration** (`database/migrations/`) - Create database schema
3. **Repository** (`internal/repositories/`) - Data access layer
4. **Service** (`internal/services/`) - Business logic
5. **Request/DTO** (`internal/requests/`, `internal/dto/`) - Request validation and response formatting
6. **Controller** (`internal/controllers/`) - HTTP handlers
7. **Router** (`internal/initialize/router.go`) - Register routes

## 🚢 Deployment

### Production Deployment

1. **Configure production environment**

   ```bash
   # Edit docker/.env for production settings
   # Set strong passwords and secure JWT secrets
   ```

2. **Build and start production environment**

   ```bash
   make prod
   ```

3. **Run migrations**

   ```bash
   make migrate-up
   ```

4. **Optional: Seed initial data**
   ```bash
   make seed
   ```

### Production Ports

- API Server: 8080 (internal)
- Nginx HTTP: 80
- Nginx HTTPS: 443
- MySQL: 3306 (internal)

### Database Backup & Restore

```bash
# Create backup
make backup-db

# Restore from backup
make restore-db
```

## 🔒 Security Considerations

- Change default database credentials in production
- Use strong JWT secrets (minimum 32 characters)
- Configure CORS allowed origins appropriately for your environment
  - Development: `CORS_ALLOWED_ORIGINS=http://localhost:3030`
  - Production: Set to your actual frontend domain
- Enable HTTPS in production (configure SSL certificates in Nginx)
- Implement rate limiting for API endpoints
- Regular security updates for dependencies
- Use environment variables for sensitive data (never commit `.env` files)

## 📝 License

[Specify your license here]

## 👥 Contributors

[Add your team members and contributors here]

## 📞 Support

For issues and questions:

- Create an issue in the repository
- Contact: [Your contact information]

---

**Built with ❤️ using Go and Docker**
