.PHONY: default run clean docs test
# This will include the variables from the .env file
# in the scope of the makefile, so they can be used by the make command
ifeq (,$(wildcard .env))
$(error .env file not found)
endif
include .env
export $(shell sed 's/=.*//' .env)

default:
	@make air

run:
	@go run ./cmd

air:
	@air --build.cmd "go build -o tmp/api cmd/main.go" --build.bin "./tmp/api"

clean:
	@echo "Cleaning up..."
	@rm -rf tmp
	@rm -rf temp

test:
	@go test -v ./...


docs:
	@swag init -g cmd/main.go