OVERRIDE_HOST = ""
OVERRIDE_PORT = 0
OVERRIDE_DSN = ""
MIGRATION_NAME = ""

build_scripts:
	@go build -o ./bin/scripts/ ./scripts/...

build_integration_tests:
	@go build -o ./bin/integration_tests ./integration_tests

route:
	@make build_scripts
	@./bin/scripts/route

build_api:
	@make route
	@go build -o ./bin/run-api ./app/main

build_all: build_api build_scripts build_integration_tests

run:
	@make route
	@make build_api
	@./bin/run-api --host=$(OVERRIDE_HOST) --port=$(OVERRIDE_PORT)

unit_tests:
	go test $$(go list ./... | grep -v integration_tests)

integration_tests:
	@make build_integration_tests
	@make run OVERRIDE_HOST="127.0.0.1" OVERRIDE_PORT=9999 &
	@./bin/integration_tests --targetHost="127.0.0.1" --targetPort=9999
	@pkill run-api || true

new_migration:
	@make build_scripts
	@./bin/scripts/new_migration --name=$(MIGRATION_NAME)

migrate:
	@make build_scripts
	@./bin/scripts/migrate --dsn=$(OVERRIDE_DSN)

docker_down:
	@docker-compose down

docker_up:
	@export ENV="docker"
	@make build_all
	@make docker_down
	@docker-compose build
	@docker-compose up -d postgres; sleep 1
	@docker-compose up migrations
	@docker-compose up api
	@make docker_down

