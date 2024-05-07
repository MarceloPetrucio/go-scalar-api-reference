default: run-with-docs

run:
	@go run *.go

docs:
	@swag init -g *.go

run-with-docs:
	@make docs
	@make run

.PHONY: docs test build 