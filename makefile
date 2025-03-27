.PHONY: install-pre-commit
install-pre-commit:
	pip3 install pre-commit
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
	pre-commit install

.PHONY: install-testfixtures
install-testfixtures:
	wget https://github.com/go-testfixtures/testfixtures/releases/download/v3.14.0/testfixtures_linux_amd64.deb
	sudo apt install testfixtures_linux_amd64.deb
	rm testfixtures_linux_amd64.deb

.PHONY: load-testfixtures
load-testfixtures:
	testfixtures --dangerous-no-test-database-check -d postgres -c "postgresql://postgresql:strimqadmin_1234@localhost:5432/postgres?sslmode=disable" -D testdata/fixtures

.PHONY: add-deferrable-foreign-key
add-deferrable-foreign-key:
	sed -i -r 's/ADD FOREIGN KEY (.*) REFERENCES (.*);/ADD FOREIGN KEY \1 REFERENCES \2 DEFERRABLE INITIALLY DEFERRED;/g' migrations/*.sql

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: run-local
run-local:
	export $(shell cat ./.env | xargs) && go run ./cmd/api/...
