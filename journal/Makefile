COMPOSE_FILE := docker/docker-compose.yml
ENV_FILE := docker/.env.development
MIGRATOR_DOCKERFILE := docker/migrator.Dockerfile
MIGRATOR_IMAGE := acta-journal-migrator
MIGRATOR_NETWORK := acta-journal-network
COLOR_GREEN := \033[32m
COLOR_RESET := \033[0m
SERVICE ?= $(word 2,$(MAKECMDGOALS))

SERVICES := app db

CONTAINER_app := acta-journal-app-development
CONTAINER_db := acta-journal-db-development

ifeq ($(SERVICE),)
CONTAINER := $(CONTAINER_app)
else ifneq ($(filter $(SERVICE),$(SERVICES)),)
CONTAINER := $(CONTAINER_$(SERVICE))
else
$(error Unknown service '$(SERVICE)'. Supported: $(SERVICES))
endif

.PHONY: up down migrate rollback shell logs env $(SERVICES)

up:
	@docker network inspect acta-shared-network >/dev/null 2>&1 || docker network create acta-shared-network >/dev/null
	docker compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) up --build -d

down:
	docker compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) down

migrate:
	docker build -f $(MIGRATOR_DOCKERFILE) -t $(MIGRATOR_IMAGE) .
	docker run --rm --env-file $(ENV_FILE) --env MIGRATION_COMMAND=up --network $(MIGRATOR_NETWORK) $(MIGRATOR_IMAGE)

rollback:
	docker build -f $(MIGRATOR_DOCKERFILE) -t $(MIGRATOR_IMAGE) .
	docker run --rm --env-file $(ENV_FILE) --env MIGRATION_COMMAND=down --network $(MIGRATOR_NETWORK) $(MIGRATOR_IMAGE)

env:
	@find . -type f -name '.env.*.example' | while read -r example; do \
		target="$${example%.example}"; \
		if [ ! -f "$$target" ]; then \
			cp "$$example" "$$target"; \
			printf '$(COLOR_GREEN)created %s$(COLOR_RESET)\n' "$$target"; \
		else \
			while IFS= read -r line || [ -n "$$line" ]; do \
				case "$$line" in \
					''|\#*) continue ;; \
				esac; \
				key="$${line%%=*}"; \
				if ! grep -Eq "^[[:space:]]*$$key=" "$$target"; then \
					printf '\n%s\n' "$$line" >> "$$target"; \
					printf '$(COLOR_GREEN)added %s to %s$(COLOR_RESET)\n' "$$key" "$$target"; \
				fi; \
			done < "$$example"; \
		fi; \
	done

shell:
	docker exec -it $(CONTAINER) bash

logs:
	docker logs -f $(CONTAINER)

app:
	@:

db:
	@:
