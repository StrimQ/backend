# Install pre-commit
.PHONY: pre-commit-install
pre-commit-install:
	poetry sync
	pre-commit install

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down
