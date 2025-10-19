.PHONY: help dev prod build-dev build-prod up-dev up-prod down-dev down-prod logs-dev logs-prod clean restart-dev restart-prod db-dev db-prod

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
	@echo "$(BLUE)Utility Commands:$(NC)"
	@echo "  $(GREEN)make clean$(NC)        - Remove all containers, volumes, and images"
	@echo "  $(GREEN)make status$(NC)       - Show status of all containers"
	@echo "  $(GREEN)make ps$(NC)           - Show running containers"
	@echo ""

## dev: Start development environment
dev: build-dev up-dev
	@echo "$(GREEN)✓ Development environment started!$(NC)"
	@echo "$(YELLOW)Server:$(NC) http://localhost:8080"
	@echo "$(YELLOW)Nginx:$(NC) http://localhost"
	@echo "$(YELLOW)MySQL:$(NC) localhost:3306"

## prod: Start production environment
prod: build-prod up-prod
	@echo "$(GREEN)✓ Production environment started!$(NC)"
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

