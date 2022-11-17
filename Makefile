OVERRIDE_HOST = ""
OVERRIDE_PORT = 0

build:
	@go build -o bin/boosters-trial ./main
	@go build

run:
	@make build
	@./bin/boosters-trial --host=$(OVERRIDE_HOST) --port=$(OVERRIDE_PORT)

unit_test:
	@go test ./...

integration_test:
	@make run OVERRIDE_HOST="127.0.0.1" OVERRIDE_PORT=99999
