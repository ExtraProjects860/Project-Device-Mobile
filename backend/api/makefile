.PHONY: default run build test docs clean

# Variables
APP_NAME=gobackendmobile

# Tasks
default: docs

build:
	@go build -d $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
seeder:
	@go run ./db/seed
migrate:
	@go run ./db/migrate
clean:
	@rm -f $$(APP_NAME)
	@rm -rf ./docs