OVERRIDE_HOST = ""
OVERRIDE_PORT = 0

build:
	@go build -o ./bin/run-api ./app/main
	@go build -o ./bin/integration-tests ./test
	@go build -o ./bin/scripts/ ./scripts/...

route:
	@make build

run:
	@make build
	@./bin/run-api --host=$(OVERRIDE_HOST) --port=$(OVERRIDE_PORT)

unit_test:
	@go test ./...

integration_test:
	@make build
	@exec -a testAppInstance ./bin/run-api --host="127.0.0.1" --port=9999 &
	@./bin/integration-tests --targetHost="127.0.0.1" --targetPort=9999
	@pkill -f testAppInstance

setup_test_db:
	@make build

migrate:
	@make build

