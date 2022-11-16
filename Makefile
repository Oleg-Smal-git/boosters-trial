OVERRIDE_HOST = ""
OVERRIDE_PORT = 0

build:
	@go build -o bin/boosters-trial ./main

run:
	@make build
	@./bin/boosters-trial --host=$(OVERRIDE_HOST) --port=$(OVERRIDE_PORT)

unit_test:
	@go test ./...

integration_test:
	@make run
