# Install pre-commit
.PHONY: pre-commit-install
pre-commit-install:
	pip3 install pre-commit
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
	pre-commit install

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down
