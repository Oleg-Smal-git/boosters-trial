OVERRIDE_HOST = ""
OVERRIDE_PORT = 0

build_scripts:
	@go build -o ./bin/scripts/ ./scripts/...

build_integration_tests:
	@go build -o ./bin/integration-tests ./test

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

unit_test:
	@go test ./...

integration_test:
	@make build
	@exec -a testAppInstance ./bin/run-api --host="127.0.0.1" --port=9999 &
	@./bin/integration-tests --targetHost="127.0.0.1" --targetPort=9999
	@pkill -f testAppInstance

migrate:
	@make build_scripts


docker_down:
	@docker-compose down

docker_up:
	@make docker_down
	@docker-compose up -d postgres; sleep 1
	@docker-compose up migrations
	@docker-compose up api
	@make docker_down

