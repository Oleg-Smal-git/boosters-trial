build:
	@go build -o bin/boosters-trial main/run.go

run:
	@build
	@./bin/boosters-trial

unit_test:
	@go test ./...

integration_test:
	@run
