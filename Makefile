# Makefile
LOCAL_COMPOSE_FILES = -f docker-compose.yml -f docker/env/local/docker-compose.override.yml

# Migrate
SERVICE ?=
TABLE ?=
STEP ?=
SERVICES = auth post user

ifneq (,$(wildcard .env))
    include .env
    export
endif

URL_CONNECT_DB = $(if $(filter auth,$(SERVICE)),"mysql://$(AUTH_MYSQL_USER):$(AUTH_MYSQL_PASSWORD)@tcp($(DB_HOST):$(AUTH_MYSQL_PORT_CONTAINER))/$(AUTH_MYSQL_DATABASE)",\
             $(if $(filter user,$(SERVICE)),"mysql://$(USER_MYSQL_USER):$(USER_MYSQL_PASSWORD)@tcp($(DB_HOST):$(USER_MYSQL_PORT_CONTAINER))/$(USER_MYSQL_DATABASE)",\
             "mysql://$(POST_MYSQL_USER):$(POST_MYSQL_PASSWORD)@tcp($(DB_HOST):$(POST_MYSQL_PORT_CONTAINER))/$(POST_MYSQL_DATABASE)"))

MIGRATE_DIR = $(if $(filter auth,$(SERVICE)),auth-service/migrations,\
             $(if $(filter user,$(SERVICE)),user-service/migrations,\
             post-service/migrations))
migrate:
ifndef SERVICE
	$(error ERROR: SERVICE is not set! Please provide SERVICE variable)
endif
	migrate -path $(MIGRATE_DIR) -database $(URL_CONNECT_DB) up

migrate-all:
	@for service in $(SERVICES); do \
		echo "Migrating $$service..."; \
		$(MAKE) migrate SERVICE=$$service; \
	done

db-rollback:
ifndef SERVICE
	$(error ERROR: SERVICE is not set! Please provide SERVICE variable)
endif
	migrate -path $(MIGRATE_DIR) -database $(URL_CONNECT_DB) down

db-rollback-all:
	@for service in $(SERVICES); do \
		echo "Rolling back $$service..."; \
		$(MAKE) db-rollback SERVICE=$$service; \
	done

db-rollback-step:
ifndef SERVICE
	$(error ERROR: SERVICE is not set! Please provide SERVICE variable)
endif
ifndef STEP
	$(error ERROR: STEP is not set! Please provide STEP variable)
endif
	migrate -path $(MIGRATE_DIR) -database $(URL_CONNECT_DB) down $(STEP)

migration:
ifndef SERVICE
	$(error ERROR: SERVICE is not set! Please provide SERVICE variable)
endif
ifndef TABLE
	$(error ERROR: TABLE is not set! Please provide TABLE variable)
endif
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(TABLE)\

## Docker
build:
	docker-compose up -d --build

up:
	docker-compose up -d

down:
	docker-compose down -v --remove-orphans

build-local:
	docker-compose $(LOCAL_COMPOSE_FILES) up -d --build

up-local:
	docker-compose $(LOCAL_COMPOSE_FILES) up -d

down-local:
	docker-compose $(LOCAL_COMPOSE_FILES) down -v --remove-orphans
