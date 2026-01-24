.PHONY: help dev prod build-dev build-prod up-dev up-prod down-dev down-prod logs-dev logs-prod clean restart-dev restart-prod db-dev db-prod migrate-create migrate-up migrate-down migrate-force migrate-version migrate-drop seed seed-local server-dev server-prod client-dev client-prod nginx-dev nginx-prod status ps backup-db restore-db

# Default target
.DEFAULT_GOAL := help

# Colors for output
YELLOW := \033[1;33m
GREEN := \033[1;32m
BLUE := \033[1;34m
RED := \033[1;31m
NC := \033[0m # No Color

# Docker compose files
COMPOSE_DEV := docker/docker-compose.yml
COMPOSE_PROD := docker/docker-compose.prod.yml

# Migration settings
MIGRATION_DIR := server/database/migrations
SEED_DIR := server/database/seed
DB_USER ?= dev_user
DB_PASSWORD ?= dev_password
DB_NAME ?= company_db
DB_HOST ?= localhost
DB_PORT ?= 33066
DB_URL := "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true"

# Docker container names
DEV_SERVER_CONTAINER := company-management-server-dev
PROD_SERVER_CONTAINER := company-management-server-prod
DEV_CLIENT_CONTAINER := company-management-client-dev
PROD_CLIENT_CONTAINER := company-management-client-prod

# Database URL for Docker internal network
DOCKER_DB_HOST := company-management-mysql-dev
DOCKER_DB_PORT := 3306
DOCKER_DB_URL := "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DOCKER_DB_HOST):$(DOCKER_DB_PORT))/$(DB_NAME)?multiStatements=true"

## help: Display this help message
help:
	@echo "$(GREEN)Company Management - Docker Environment Manager$(NC)"
	@echo ""
	@echo "$(YELLOW)Available commands:$(NC)"
	@echo ""
	@echo "$(BLUE)Development Environment:$(NC)"
	@echo "  $(GREEN)make dev$(NC)          - Start development environment (build + up)"
	@echo "  $(GREEN)make build-dev$(NC)    - Build development Docker images"
	@echo "  $(GREEN)make up-dev$(NC)       - Start development containers"
	@echo "  $(GREEN)make down-dev$(NC)     - Stop development containers"
	@echo "  $(GREEN)make restart-dev$(NC)  - Restart development containers"
	@echo "  $(GREEN)make logs-dev$(NC)     - Show development logs"
	@echo "  $(GREEN)make db-dev$(NC)       - Connect to development MySQL"
	@echo ""
	@echo "$(BLUE)Production Environment:$(NC)"
	@echo "  $(GREEN)make prod$(NC)         - Start production environment (build + up)"
	@echo "  $(GREEN)make build-prod$(NC)   - Build production Docker images"
	@echo "  $(GREEN)make up-prod$(NC)      - Start production containers"
	@echo "  $(GREEN)make down-prod$(NC)    - Stop production containers"
	@echo "  $(GREEN)make restart-prod$(NC) - Restart production containers"
	@echo "  $(GREEN)make logs-prod$(NC)    - Show production logs"
	@echo "  $(GREEN)make db-prod$(NC)      - Connect to production MySQL"
	@echo ""
	@echo "$(BLUE)Shell Access:$(NC)"
	@echo "  $(GREEN)make server-dev$(NC)   - Access development server shell"
	@echo "  $(GREEN)make server-prod$(NC)  - Access production server shell"
	@echo "  $(GREEN)make client-dev$(NC)   - Access development client shell"
	@echo "  $(GREEN)make client-prod$(NC)  - Access production client shell"
	@echo "  $(GREEN)make nginx-dev$(NC)    - Access development nginx shell"
	@echo "  $(GREEN)make nginx-prod$(NC)   - Access production nginx shell"
	@echo ""
	@echo "$(BLUE)Database Migration:$(NC)"
	@echo "  $(GREEN)make migrate-create$(NC)              - Create new migration file"
	@echo "  $(GREEN)make migrate-up$(NC)                  - Apply all migrations"
	@echo "  $(GREEN)make migrate-down$(NC)                - Rollback last migration"
	@echo "  $(GREEN)make migrate-force$(NC)               - Force set migration version"
	@echo "  $(GREEN)make migrate-version$(NC)             - Show current migration version"
	@echo "  $(GREEN)make migrate-drop$(NC)                - Drop all tables (dangerous!)"
	@echo ""
	@echo "$(BLUE)Database Seeding:$(NC)"
	@echo "  $(GREEN)make seed$(NC)                        - Run database seeder in Docker container"
	@echo "  $(GREEN)make seed-local$(NC)                  - Run database seeder from local machine"
	@echo ""
	@echo "$(BLUE)Database Backup & Restore:$(NC)"
	@echo "  $(GREEN)make backup-db$(NC)                   - Backup MySQL database"
	@echo "  $(GREEN)make restore-db$(NC)                  - Restore MySQL database from backup"
	@echo ""
	@echo "$(BLUE)Utility Commands:$(NC)"
	@echo "  $(GREEN)make clean$(NC)        - Remove all containers, volumes, and images"
	@echo "  $(GREEN)make status$(NC)       - Show status of all containers"
	@echo "  $(GREEN)make ps$(NC)           - Show running containers"
	@echo ""

## dev: Start development environment
dev: build-dev up-dev
	@echo "$(GREEN)✓ Development environment started!$(NC)"
	@echo "$(YELLOW)Client:$(NC) http://localhost:3030"
	@echo "$(YELLOW)Server:$(NC) http://localhost:8808"
	@echo "$(YELLOW)Nginx:$(NC) http://localhost:880"
	@echo "$(YELLOW)MySQL:$(NC) localhost:3306"

## prod: Start production environment
prod: build-prod up-prod
	@echo "$(GREEN)✓ Production environment started!$(NC)"
	@echo "$(YELLOW)Client:$(NC) http://localhost:3000"
	@echo "$(YELLOW)Server:$(NC) http://localhost:8080"
	@echo "$(YELLOW)Nginx:$(NC) https://localhost"
	@echo "$(YELLOW)MySQL:$(NC) localhost:3306"

## build-dev: Build development Docker images
build-dev:
	@echo "$(BLUE)Building development images...$(NC)"
	docker compose -f $(COMPOSE_DEV) build

## build-prod: Build production Docker images
build-prod:
	@echo "$(BLUE)Building production images...$(NC)"
	docker compose -f $(COMPOSE_PROD) build

## up-dev: Start development containers
up-dev:
	@echo "$(BLUE)Starting development containers...$(NC)"
	docker compose -f $(COMPOSE_DEV) up -d
	@echo "$(GREEN)✓ Development containers started$(NC)"

## up-prod: Start production containers
up-prod:
	@echo "$(BLUE)Starting production containers...$(NC)"
	docker compose -f $(COMPOSE_PROD) up -d
	@echo "$(GREEN)✓ Production containers started$(NC)"

## down-dev: Stop development containers
down-dev:
	@echo "$(BLUE)Stopping development containers...$(NC)"
	docker compose -f $(COMPOSE_DEV) down
	@echo "$(GREEN)✓ Development containers stopped$(NC)"

## down-prod: Stop production containers
down-prod:
	@echo "$(BLUE)Stopping production containers...$(NC)"
	docker compose -f $(COMPOSE_PROD) down
	@echo "$(GREEN)✓ Production containers stopped$(NC)"

## restart-dev: Restart development containers
restart-dev: down-dev up-dev
	@echo "$(GREEN)✓ Development containers restarted$(NC)"

## restart-prod: Restart production containers
restart-prod: down-prod up-prod
	@echo "$(GREEN)✓ Production containers restarted$(NC)"

## logs-dev: Show development logs
logs-dev:
	docker compose -f $(COMPOSE_DEV) logs -f

## logs-prod: Show production logs
logs-prod:
	docker compose -f $(COMPOSE_PROD) logs -f

## db-dev: Connect to development MySQL
db-dev:
	@echo "$(BLUE)Connecting to development MySQL...$(NC)"
	docker exec -it company-management-mysql-dev mysql -udev_user -pdev_password company_db

## db-prod: Connect to production MySQL
db-prod:
	@echo "$(BLUE)Connecting to production MySQL...$(NC)"
	@echo "$(YELLOW)Note: Use environment variables for credentials$(NC)"
	docker exec -it company-management-mysql-prod mysql -u$$MYSQL_USER -p company_db

## clean: Remove all containers, volumes, and images
clean:
	@echo "$(RED)⚠ This will remove all containers, volumes, and images!$(NC)"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		echo "$(BLUE)Cleaning development environment...$(NC)"; \
		docker compose -f $(COMPOSE_DEV) down -v --rmi all; \
		echo "$(BLUE)Cleaning production environment...$(NC)"; \
		docker compose -f $(COMPOSE_PROD) down -v --rmi all; \
		echo "$(GREEN)✓ Cleanup completed$(NC)"; \
	fi

## status: Show status of all containers
status:
	@echo "$(BLUE)Development Containers:$(NC)"
	@docker compose -f $(COMPOSE_DEV) ps || true
	@echo ""
	@echo "$(BLUE)Production Containers:$(NC)"
	@docker compose -f $(COMPOSE_PROD) ps || true

## ps: Show running containers
ps:
	@docker ps --filter "name=company-management" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

## server-dev: Access development server shell
server-dev:
	@echo "$(BLUE)Accessing development server shell...$(NC)"
	docker exec -it company-management-server-dev sh

## server-prod: Access production server shell
server-prod:
	@echo "$(BLUE)Accessing production server shell...$(NC)"
	docker exec -it company-management-server-prod sh

## client-dev: Access development client shell
client-dev:
	@echo "$(BLUE)Accessing development client shell...$(NC)"
	docker exec -it $(DEV_CLIENT_CONTAINER) sh

## client-prod: Access production client shell
client-prod:
	@echo "$(BLUE)Accessing production client shell...$(NC)"
	docker exec -it $(PROD_CLIENT_CONTAINER) sh

## nginx-dev: Access development nginx shell
nginx-dev:
	@echo "$(BLUE)Accessing development nginx shell...$(NC)"
	docker exec -it company-management-nginx-dev sh

## nginx-prod: Access production nginx shell
nginx-prod:
	@echo "$(BLUE)Accessing production nginx shell...$(NC)"
	docker exec -it company-management-nginx-prod sh

## backup-db: Backup MySQL database
backup-db:
	@echo "$(BLUE)Creating database backup...$(NC)"
	@mkdir -p ./docker/mysql/backup
	@docker exec company-management-mysql-prod mysqldump -u$$MYSQL_USER -p$$MYSQL_PASSWORD company_db > ./docker/mysql/backup/backup_$$(date +%Y%m%d_%H%M%S).sql
	@echo "$(GREEN)✓ Backup completed$(NC)"

## restore-db: Restore MySQL database
restore-db:
	@echo "$(BLUE)Available backups:$(NC)"
	@ls -lht ./docker/mysql/backup/*.sql | head -5
	@echo ""
	@read -p "Enter backup filename: " backup; \
	docker exec -i company-management-mysql-prod mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD company_db < ./docker/mysql/backup/$$backup
	@echo "$(GREEN)✓ Restore completed$(NC)"

## migrate-create: Create a new migration file (inside Docker container)
migrate-create:
	@read -p "Enter migration name (e.g., create_users_table): " migration_name; \
	if [ -z "$$migration_name" ]; then \
		echo "$(RED)Error: Migration name cannot be empty$(NC)"; \
		exit 1; \
	fi; \
	if echo "$$migration_name" | LC_ALL=C grep -q '[^a-zA-Z0-9_]'; then \
		echo "$(RED)Error: Migration name must contain only letters, numbers, and underscores$(NC)"; \
		echo "$(YELLOW)Invalid characters detected. Please use only: a-z, A-Z, 0-9, _$(NC)"; \
		exit 1; \
	fi; \
	echo "$(BLUE)Creating migration: $$migration_name (in Docker container)$(NC)"; \
	docker exec $(DEV_SERVER_CONTAINER) migrate create -ext sql -dir database/migrations -seq $$migration_name; \
	echo "$(GREEN)✓ Migration files created in $(MIGRATION_DIR)$(NC)"

## migrate-up: Apply all pending migrations (inside Docker container)
migrate-up:
	@echo "$(BLUE)Applying migrations in Docker container...$(NC)"
	@docker exec $(DEV_SERVER_CONTAINER) migrate -path database/migrations -database $(DOCKER_DB_URL) -verbose up
	@echo "$(GREEN)✓ Migrations applied$(NC)"

## migrate-down: Rollback the last migration (inside Docker container)
migrate-down:
	@echo "$(YELLOW)⚠ Rolling back last migration in Docker container...$(NC)"
	@docker exec $(DEV_SERVER_CONTAINER) migrate -path database/migrations -database $(DOCKER_DB_URL) -verbose down 1
	@echo "$(GREEN)✓ Migration rolled back$(NC)"

## migrate-force: Force set migration version (inside Docker container)
migrate-force:
	@read -p "Enter version number to force: " version; \
	if [ -z "$$version" ]; then \
		echo "$(RED)Error: Version number cannot be empty$(NC)"; \
		exit 1; \
	fi; \
	echo "$(YELLOW)⚠ Force setting migration version to $$version in Docker container$(NC)"; \
	docker exec $(DEV_SERVER_CONTAINER) migrate -path database/migrations -database $(DOCKER_DB_URL) force $$version; \
	echo "$(GREEN)✓ Migration version set to $$version$(NC)"

## migrate-version: Show current migration version (inside Docker container)
migrate-version:
	@echo "$(BLUE)Current migration version (from Docker container):$(NC)"
	@docker exec $(DEV_SERVER_CONTAINER) migrate -path database/migrations -database $(DOCKER_DB_URL) version

## migrate-drop: Drop all tables (DANGEROUS!) (inside Docker container)
migrate-drop:
	@echo "$(RED)⚠⚠⚠ WARNING: This will drop all tables! ⚠⚠⚠$(NC)"
	@read -p "Are you ABSOLUTELY sure? Type 'yes' to confirm: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		echo "$(BLUE)Dropping all tables in Docker container...$(NC)"; \
		docker exec $(DEV_SERVER_CONTAINER) migrate -path database/migrations -database $(DOCKER_DB_URL) drop -f; \
		echo "$(GREEN)✓ All tables dropped$(NC)"; \
	else \
		echo "$(YELLOW)Cancelled$(NC)"; \
	fi

## seed: Run database seeder (inside Docker container)
seed:
	@echo "$(BLUE)Running database seeder...$(NC)"
	@docker exec company-management-server-dev sh -c "cd /app && go run cmd/seed/main.go"
	@echo "$(GREEN)✓ Database seeding completed$(NC)"

## seed-local: Run database seeder from local machine
seed-local:
	@echo "$(BLUE)Running database seeder locally...$(NC)"
	@cd server && \
		DB_HOST=localhost \
		DB_PORT=33066 \
		DB_USER=$(DB_USER) \
		DB_PASSWORD=$(DB_PASSWORD) \
		DB_NAME=$(DB_NAME) \
		go run cmd/seed/main.go
	@echo "$(GREEN)✓ Database seeding completed$(NC)"

