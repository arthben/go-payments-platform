GOBASE = $(shell pwd)
GOVERSION := 1.24.4
GOLANGCI_VERSION := 1.64.8
MOCKGEN_VERSION := 1.6.0

GOLINT_PATH = $(GOBASE)/build/lint
GOLANGCI_LINT := $(LINT_PATH)/golangci-lint

MIGRATE=migrate -path migrations -database postgres://postgres:12345@localhost:5432/gpp_payment_service?sslmode=disable

.PHONY: 

deps: ## fetch required dependencies
	go mod tidy -compat=$(GOVERSION)

cilint: ## Run linter for CI
	$(GOLANGCI_LINT) run --timeout=5m -c .golangci.yml \
		--out-format code-climate:gl-code-quality-report.json,html:golangci-lint-report.html,colored-line-number

lint: ## Run linter 
	$(GOLANGCI_LINT) run --timeout=5m -c .golangci.yml \
		--out-format colored-line-number,html:golangci-lint-report.html

lint-fix:## Run linter with auto-fix
	$(GOLANGCI_LINT) run --timeout=5m -c .golangci.yml --fix

submodule: ## Sync protos
	git submodule sync --recursive
	git submodule update --init --recursive

submodule-update: ## Sync protos with remote
	git submodule sync --recursive
	git submodule update --init --recursive --remote

proto: submodule proto-generate

proto-update: submodule-update proto-generate

proto-generate: ## Generate protobuf code from proto files
	rm -r gen || :
	docker run --rm -v `pwd`:/defs namely/protoc-all:1.44_0 -i protos -f protos/go-payments-platform/payment/v1/payment.proto -l go -o gen/payment

migrate-create: ## Create migration file with name
	migrate create -ext sql -dir migrations -seq -digits 3 $(NAME)

migrate-up: ## Run migrations up
	$(MIGRATE) up

migrate-down: ## Rollback 1 migration
	$(MIGRATE) down 1

test: ## Tests
	mkdir reports || true
	go install gotest.tools/gotestsum@v1.9.0
	gotestsum --junitfile=report.xml --format=testname -- -race -coverprofile=reports/coverage-report.out -tags=unit,$(goKafkaTag) ./...

install-golangci: ## Install the correct version of lint
	@if [ ! -f $(LINT_PATH)/golangci-lint ] ; then GOBIN=$(LINT_PATH) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_VERSION) ; fi

install-mockgen: ## Install the correct version of mockgen
	@if [ ! -f $(MOCKGEN_PATH)/mockgen ] ; then GOBIN=$(MOCKGEN_PATH) go install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION) ; fi

mock-gen: install-mockgen ## Generate mocks
	rm -r **/mock || :
	$(MOCKGEN_PATH)/mockgen -package mock -source service/services.go -destination service/mock/services.go
