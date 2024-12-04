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

migrate: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) up

rollback: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) down

migrate-validate: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) validate

migrate-status: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN) status

prepare-test: .bin-goose
	rm -f ./database/database_test.sqlite
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DSN_TEST) up

test: prepare-test
	go test -v -count=1 -race ./...

start:
	docker-compose up -d

stop:
	docker-compose down
	docker image rm tinvest-bot