# Install pre-commit
.PHONY: pre-commit-install
pre-commit-install:
	poetry sync
	pre-commit install
