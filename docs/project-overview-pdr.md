# Project Overview - PDR (Project Definition Report)

## Executive Summary

The **Company Management System** is a comprehensive enterprise-grade backend application designed to streamline the management of multi-company organizations, employees, positions, contracts, and permissions. Built with Go and modern web technologies, the system provides a secure, scalable platform for organizations to manage their workforce across multiple companies with fine-grained access control.

### Project Information

- **Project Name**: Company Management System
- **Version**: 1.0.0
- **Status**: Active Development
- **Primary Language**: Go (1.25.3)
- **License**: [To be determined]
- **Repository**: `github.com/vlahanam/company-management`

### Key Stakeholders

- **Development Team**: Backend development and maintenance
- **System Administrators**: Deployment and operations
- **End Users**: HR managers, administrators, employees

## Problem Statement

### Business Challenge

Modern organizations, especially those managing multiple companies or subsidiaries, face significant challenges:

1. **Fragmented Systems**: Employee data scattered across multiple systems
2. **Complex Hierarchies**: Difficulty managing parent-child company relationships
3. **Access Control**: Limited granular permission control across organizational boundaries
4. **Contract Management**: Manual tracking of employment contracts and their lifecycles
5. **Position Management**: Inefficient tracking of organizational positions and hierarchies
6. **Scalability Issues**: Legacy systems unable to handle growth

### Target Users

- **HR Managers**: Manage employees, contracts, and positions
- **System Administrators**: Manage users, roles, and permissions
- **Company Executives**: Oversee organizational structure
- **Employees**: Access their own information and contracts

## Solution Overview

### Core Capabilities

The Company Management System addresses these challenges through:

#### 1. Multi-Company Management

- Hierarchical company structure with parent-child relationships
- Isolated data management per company
- Cross-company reporting and analytics capabilities

#### 2. Employee Management

- Complete employee lifecycle management
- Profile management with avatars
- Personal information tracking (DOB, gender, ID cards)
- Contact information management

#### 3. Contract Management

- Support for multiple contract types:
  - Probation contracts
  - Fixed-term contracts
  - Permanent contracts
  - Freelance agreements
  - Internship contracts
- Contract status tracking (Active, Expired, Terminated, Pending)
- Salary management
- Contract document storage
- Automated contract lifecycle management

#### 4. Position & Organizational Structure

- Position hierarchy within companies
- Salary range management per position
- Multi-position assignments for employees
- Position-based access control

#### 5. Role-Based Access Control (RBAC)

- Flexible permission system
- Customizable roles
- Permission-based resource access
- Three default roles:
  - **Super Admin**: Full system access
  - **Admin**: Company-level management
  - **User**: Basic access

#### 6. Security Features

- JWT-based authentication
- Bcrypt password hashing
- Access and refresh token mechanism
- Secure API endpoints
- Input validation and sanitization

## Technical Architecture

### Technology Stack

| Component        | Technology      | Version | Purpose                      |
| ---------------- | --------------- | ------- | ---------------------------- |
| Backend Language | Go              | 1.25.3  | Core application logic       |
| Web Framework    | Fiber           | v2      | HTTP routing and middleware  |
| ORM              | GORM            | Latest  | Database operations          |
| Database         | MySQL           | 8.0     | Data persistence             |
| Frontend         | Next.js         | 16.1.4  | React framework with SSR     |
| UI Library       | React           | 19.2.3  | User interface components    |
| Styling          | Tailwind CSS    | 4.x     | Utility-first CSS framework  |
| Type Safety      | TypeScript      | 5.x     | Static type checking         |
| Authentication   | JWT             | v5      | Token-based auth             |
| Validation       | ozzo-validation | v4      | Request validation           |
| Containerization | Docker          | Latest  | Application packaging        |
| Orchestration    | Docker Compose  | v2      | Multi-container management   |
| Reverse Proxy    | Nginx           | Latest  | Load balancing, SSL, routing |
| Hot Reload       | Air + Turbopack | Latest  | Development productivity     |

### System Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    Client Layer                          │
│              (Web/Mobile Applications)                   │
└────────────────────────┬────────────────────────────────┘
                         │ HTTPS
┌────────────────────────▼────────────────────────────────┐
│                   Nginx Reverse Proxy                    │
│              (Load Balancing, SSL, Caching)              │
└────────────────────────┬────────────────────────────────┘
                         │ HTTP
┌────────────────────────▼────────────────────────────────┐
│                  Go Fiber Server                         │
│  ┌──────────────────────────────────────────────────┐   │
│  │           Middleware Layer                       │   │
│  │  (JWT Auth, CORS, Logging, Error Handling)       │   │
│  └─────────────────────┬────────────────────────────┘   │
│  ┌─────────────────────▼────────────────────────────┐   │
│  │          Controller Layer                        │   │
│  │  (HTTP Handlers, Request Validation)             │   │
│  └─────────────────────┬────────────────────────────┘   │
│  ┌─────────────────────▼────────────────────────────┐   │
│  │          Service Layer                           │   │
│  │  (Business Logic, Orchestration)                 │   │
│  └─────────────────────┬────────────────────────────┘   │
│  ┌─────────────────────▼────────────────────────────┐   │
│  │          Repository Layer                        │   │
│  │  (Data Access, GORM Operations)                  │   │
│  └─────────────────────┬────────────────────────────┘   │
└────────────────────────┼────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────┐
│                   MySQL Database                         │
│            (Data Persistence, ACID Transactions)         │
└─────────────────────────────────────────────────────────┘
```

### Data Model

#### Core Entities

1. **Users**
   - Identity and authentication
   - Personal information
   - Profile management

2. **Companies**
   - Organization management
   - Hierarchical structure
   - Multi-tenancy support

3. **Positions**
   - Job roles and levels
   - Salary ranges
   - Company associations

4. **Contracts**
   - Employment agreements
   - Contract lifecycle
   - Salary and terms

5. **Roles**
   - Access control groups
   - Permission bundles

6. **Permissions**
   - Resource-action pairs
   - Granular access control

#### Entity Relationships

```
Users ←──→ Roles (many-to-many via user_roles)
Roles ←──→ Permissions (many-to-many via role_permissions)
Users ←──→ Positions (many-to-many via user_positions)
Users ──→ Contracts (one-to-many)
Companies ──→ Positions (one-to-many)
Companies ──→ Contracts (one-to-many)
Companies ──→ Companies (self-referential, parent-child)
Positions ──→ Contracts (one-to-many)
```

## Project Scope

### In Scope

✅ **Phase 1 (Current)**

- User authentication and authorization
- Company CRUD operations with hierarchy
- Position management
- Contract lifecycle management
- RBAC implementation
- RESTful API endpoints
- Database migrations
- Docker containerization (backend + frontend)
- Basic error handling and validation
- **Next.js 16 frontend setup with Docker**

### Out of Scope (Current Phase)

❌ **Not Included in Current Version**

- Full frontend application features (UI components, pages, forms)
- Email notifications
- File storage (S3 integration)
- Advanced reporting and analytics
- Export functionality (PDF, Excel)
- Audit logging
- Multi-language support
- GraphQL API
- Real-time notifications (WebSockets)
- Background job processing

## Success Criteria

### Functional Requirements

| Requirement                 | Status         | Priority |
| --------------------------- | -------------- | -------- |
| User registration and login | ✅ Implemented | P0       |
| JWT authentication          | ✅ Implemented | P0       |
| RBAC system                 | ✅ Implemented | P0       |
| Company management          | ✅ Implemented | P0       |
| Position management         | ✅ Implemented | P0       |
| Contract management         | ✅ Implemented | P0       |
| User profile management     | ✅ Implemented | P1       |
| Password security (bcrypt)  | ✅ Implemented | P0       |
| Input validation            | ✅ Implemented | P1       |

### Non-Functional Requirements

| Requirement         | Target                  | Status             |
| ------------------- | ----------------------- | ------------------ |
| API Response Time   | < 200ms (avg)           | ⚠️ To be measured  |
| Database Query Time | < 100ms (avg)           | ⚠️ To be measured  |
| Concurrent Users    | 1000+                   | ⚠️ To be tested    |
| Uptime              | 99.9%                   | ⚠️ To be monitored |
| Security            | OWASP Top 10 compliance | 🔄 In progress     |
| Code Coverage       | > 80%                   | ⏳ Planned         |

## Deployment Architecture

### Environments

#### Development

```
Ports:
- Nginx: 880
- API Server: 8808
- MySQL: 33066

Features:
- Hot reload (Air)
- Debug logging
- Development database
```

#### Production

```
Ports:
- Nginx: 80, 443 (HTTPS)
- API Server: 8080 (internal)
- MySQL: 3306 (internal)

Features:
- Optimized builds
- Error logging only
- Production database
- SSL/TLS encryption
```

### Infrastructure Requirements

| Component | Development | Production                |
| --------- | ----------- | ------------------------- |
| CPU       | 2 cores     | 4+ cores                  |
| RAM       | 2GB         | 8GB+                      |
| Storage   | 10GB        | 50GB+ (+ backups)         |
| Network   | 100Mbps     | 1Gbps                     |
| Database  | MySQL 8.0   | MySQL 8.0 (with replicas) |

## Security Considerations

### Implemented Security Measures

1. **Authentication**
   - JWT tokens with expiration
   - Refresh token mechanism
   - Bcrypt password hashing (cost: 10)

2. **Authorization**
   - Role-based access control
   - Permission-based resource access
   - Middleware enforcement

3. **Data Protection**
   - Environment variable configuration
   - No hardcoded secrets
   - .gitignore for sensitive files

4. **Input Validation**
   - ozzo-validation for requests
   - GORM SQL injection protection
   - Type-safe parameters

### Recommended Security Enhancements

- [ ] Rate limiting for API endpoints
- [ ] HTTPS enforcement (SSL certificates)
- [ ] Database connection encryption
- [ ] Audit logging for sensitive operations
- [ ] IP whitelisting for admin endpoints
- [ ] Two-factor authentication (2FA)
- [ ] Account lockout after failed attempts
- [ ] Password complexity requirements
- [ ] Regular security audits

## Development Workflow

### Team Collaboration

```
Developer → Git → Code Review → CI/CD → Deployment

Tools:
- Version Control: Git
- Code Review: Pull Requests
- CI/CD: [To be implemented]
- Deployment: Docker Compose
```

### Development Process

1. **Feature Development**
   - Create feature branch
   - Implement changes
   - Write tests
   - Submit pull request

2. **Code Review**
   - Peer review
   - Automated checks
   - Approval required

3. **Testing**
   - Unit tests
   - Integration tests
   - Manual testing

4. **Deployment**
   - Merge to main
   - Automated deployment (future)
   - Database migrations
   - Verification

### Database Management

- **Migrations**: golang-migrate with versioned SQL files
- **Seeding**: Automated seeder for development data
- **Backups**: Manual backups via Makefile (to be automated)

## Project Metrics

### Current Status

| Metric            | Value                         |
| ----------------- | ----------------------------- |
| Total Files       | 90+ (server + client)         |
| Go Files          | 46                            |
| TypeScript Files  | 5+ (client)                   |
| Database Tables   | 9                             |
| Migration Files   | 18 (9 up/down pairs)          |
| API Endpoints     | 30+                           |
| Middleware        | 3+                            |
| Docker Containers | 4 (server, client, db, nginx) |

### Code Quality Indicators

- **Go Version**: 1.25.3 (latest)
- **Node.js Version**: 20 (LTS)
- **Next.js Version**: 16.1.4 (latest)
- **Dependencies**: 20+ backend, 15+ frontend (well-maintained packages)
- **Architecture**: Layered backend (4 layers), Component-based frontend
- **Test Coverage**: [To be implemented]
- **Documentation**: Comprehensive (5 docs)

## Risks and Mitigation

### Technical Risks

| Risk                   | Impact   | Probability | Mitigation                         |
| ---------------------- | -------- | ----------- | ---------------------------------- |
| Database bottleneck    | High     | Medium      | Implement caching, read replicas   |
| Security breach        | Critical | Low         | Regular audits, security updates   |
| Data loss              | Critical | Low         | Automated backups, replication     |
| Scalability issues     | High     | Medium      | Horizontal scaling, load balancing |
| Third-party dependency | Medium   | Low         | Version pinning, fallback options  |

### Operational Risks

| Risk                    | Impact   | Probability | Mitigation                          |
| ----------------------- | -------- | ----------- | ----------------------------------- |
| Server downtime         | High     | Low         | High availability setup, monitoring |
| Deployment failures     | Medium   | Medium      | Rollback procedures, staging env    |
| Performance degradation | High     | Medium      | Monitoring, auto-scaling            |
| Data corruption         | Critical | Very Low    | Transactions, validation, backups   |

## Cost Analysis

### Infrastructure Costs (Estimated Monthly)

| Resource         | Development | Production   |
| ---------------- | ----------- | ------------ |
| Server (EC2/VPS) | $10-20      | $50-100      |
| Database (RDS)   | $15-25      | $100-200     |
| Storage          | $5          | $20-50       |
| Bandwidth        | Included    | $20-50       |
| Backup Storage   | $5          | $20-30       |
| Monitoring       | Free        | $0-50        |
| **Total**        | **$35-55**  | **$210-480** |

_Note: Costs vary based on cloud provider and usage_

### Development Costs

- **Initial Development**: [Completed]
- **Maintenance**: Ongoing
- **Feature Development**: As needed
- **Support**: Team-based

## Timeline

### Phase 1: Core Development ✅ (Completed)

- User authentication ✅
- Company management ✅
- Position management ✅
- Contract management ✅
- RBAC implementation ✅
- Docker setup ✅

### Phase 2: Enhancement 🔄 (In Progress)

- Documentation creation 🔄
- Testing implementation ⏳
- Performance optimization ⏳
- Security hardening ⏳

### Phase 3: Advanced Features ⏳ (Planned)

- File storage integration
- Email notifications
- Advanced reporting
- Audit logging
- Frontend application

## Support and Maintenance

### Maintenance Plan

- **Updates**: Regular dependency updates
- **Backups**: Daily database backups
- **Monitoring**: Server and application monitoring
- **Logs**: Centralized logging system
- **Security**: Monthly security reviews

### Support Channels

- **Issue Tracking**: GitHub Issues
- **Documentation**: README.md and docs/ folder
- **Code Repository**: GitHub
- **Team Communication**: [To be determined]

## Appendices

### Glossary

- **RBAC**: Role-Based Access Control
- **JWT**: JSON Web Token
- **ORM**: Object-Relational Mapping
- **CRUD**: Create, Read, Update, Delete
- **API**: Application Programming Interface
- **DTO**: Data Transfer Object

### References

- [Go Documentation](https://golang.org/doc/)
- [Fiber Framework](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [JWT Specification](https://jwt.io/)
- [Docker Documentation](https://docs.docker.com/)

### Contact Information

- **Project Repository**: `github.com/vlahanam/company-management`
- **Team**: [Contact information to be added]
- **Issues**: GitHub Issues

---

**Document Version**: 1.0  
**Last Updated**: January 2026  
**Next Review**: Quarterly  
**Status**: Living Document
