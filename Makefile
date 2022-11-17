OVERRIDE_HOST = ""
OVERRIDE_PORT = 0

build:
	@go build -o ./bin/run-api ./app/main
	@go build -o ./bin/integration-tests ./test

run:
	@make build
	@./bin/run-api --host=$(OVERRIDE_HOST) --port=$(OVERRIDE_PORT)

unit_test:
	@go test ./...

integration_test:
	@make build
	@exec -a testAppInstance ./bin/run-api --host="127.0.0.1" --port=99999 &
	@./bin/integration-tests --targetHost="127.0.0.1" --targetPort=99999
	@pkill -f testAppInstance
	@wait
