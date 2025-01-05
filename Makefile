LOCAL_BIN:=$(CURDIR)/bin

PHONY: generate
generate: genproto genmock

PHONY: .bin-buf
.bin-buf:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing binaries...)
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.47.2

PHONY: genproto
genproto: .bin-buf
	$(info $(shell printf "\033[34;1m▶\033[0m") Generate protos...)
	$(LOCAL_BIN)/buf generate

PHONY: .copy-proto
.copy-proto:	
	rm -rf ./proto
	mkdir -p ./proto
	wget -O ./proto/tinvest.zip https://github.com/RussianInvestments/investAPI/archive/refs/heads/main.zip
	unzip -o ./proto/tinvest.zip "investAPI-main/src/docs/contracts/*" -d ./proto/tinvest_tmp
	rm -f ./proto/tinvest.zip
	mv ./proto/tinvest_tmp/investAPI-main/src/docs/contracts/*.proto ./proto
	mv ./proto/tinvest_tmp/investAPI-main/src/docs/contracts/* ./proto
	rm -rf ./proto/tinvest_tmp

PHONY: .bin-mock
.bin-mock:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing binaries...)
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.50.0

PHONY: genmock
genmock: .bin-mock
	$(LOCAL_BIN)/mockery

GOOSE_MIGRATION_DIR=./migrations
GOOSE_DRIVER=sqlite3
GOOSE_DSN="./database/database.sqlite"
GOOSE_DSN_TEST="./database/database_test.sqlite"

PHONY: .bin-golangci-lint
.bin-golangci-lint:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing golangci-lint...)
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

PHONY: lint
lint: .bin-golangci-lint
	$(LOCAL_BIN)/golangci-lint run --config=.golangci.yaml ./...

PHONY: .bin-goose
.bin-goose:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing goose binary...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

PHONY: migrate
migrate: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) up

PHONY: rollback
rollback: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) down

PHONY: migrate-validate
migrate-validate: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) validate

PHONY: migrate-status
migrate-status: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) status

PHONY: prepare-test
prepare-test: .bin-goose
	mkdir -p ./database
	rm -f ./database/database_test.sqlite
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN_TEST) up

PHONY: prepare-test
test: prepare-test
	go test -v -count=1 -race ./...

PHONY: run
run:
	go run ./cmd/bot

PHONY: start
start:
	docker compose up -d

PHONY: stop
stop:
	docker compose down
	docker image rm tinvest-bot

PHONY: pull-and-run
pull-and-run:
	docker stop tinvest-go
	docker rm tinvest-go
	docker rmi ghcr.io/ivangurin/tinvest-go
	docker pull ghcr.io/ivangurin/tinvest-go
	docker run -d \
		--name tinvest-go \
		-e TINVEST_BOT_TOKEN \
		--env-file .env \
		-v ./database:/app/database \
		--restart always \
		ghcr.io/ivangurin/tinvest-go