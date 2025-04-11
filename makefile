# This Makefile is tested on Ubuntu 24.04
.PHONY: add-deferrable-foreign-key
add-deferrable-foreign-key:
	sed -i -r 's/ADD FOREIGN KEY (.*) REFERENCES (.*);/ADD FOREIGN KEY \1 REFERENCES \2 DEFERRABLE INITIALLY DEFERRED;/g' migrations/*.sql

.PHONY: install-prerequisites
install-prerequisites:
# Check and install golangci-lint if not present
	@command -v golangci-lint >/dev/null 2>&1 || brew install golangci-lint
# Check and install goimports if not present
	@command -v goimports >/dev/null 2>&1 || go install golang.org/x/tools/cmd/goimports@0.31.0
# Check and install pre-commit, hooks if not present
	@command -v pre-commit >/dev/null 2>&1 || pip3 install pre-commit
	@pre-commit install
# Check and install testfixtures if not present
	@dpkg -l | grep testfixtures >/dev/null 2>&1 || { \
		wget https://github.com/go-testfixtures/testfixtures/releases/download/v3.14.0/testfixtures_linux_amd64.deb && \
		sudo apt install ./testfixtures_linux_amd64.deb && \
		rm testfixtures_linux_amd64.deb; \
	}
# Check and install goose
	@command -v goose >/dev/null 2>&1 || brew install goose

.PHONY: migrate-up-db
migrate-up-db:
	@goose -dir migrations postgres "postgresql://postgres:strimqadmin_1234@localhost:5432/postgres?sslmode=disable" up

.PHONY: load-mock-db
load-mock-db:
	@testfixtures --dangerous-no-test-database-check -d postgres -c "postgresql://postgres:strimqadmin_1234@localhost:5432/postgres?sslmode=disable" -D testdata/fixtures

.PHONY: up
up:
	@docker compose up -d

.PHONY: down
down:
	@docker compose down

.PHONY: run-local
run-local:
	@export $(shell cat ./.env.local | xargs) && go run ./cmd/api/...
