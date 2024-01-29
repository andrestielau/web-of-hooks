include .env
export

all: boot down db gen

boot:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/jschaf/pggen/cmd/pggen@latest
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

run: 
	go run . serve

dev:
	air serve

# Compose Commands
up:
	@docker-compose up -d

down:
	@docker-compose down

wait: 
	@echo Waiting for ${CONTAINER}
	@while [ "`docker inspect -f {{.State.Health.Status}} ${CONTAINER}`" != "healthy" ]; do \
		sleep 0.1 ; \
	done; 

# Generator Commands
gen: gen/ui gen/db gen/grpc gen/api gen/go

gen/grpc: ${GRPC_DIR}/webhooks.proto buf.gen.yaml 
	@echo Generating gRPC
	@buf generate

gen/api: ${API_DIR}/webhooks.yaml api.gen.yaml
	@echo Generating API
	@oapi-codegen --config api.gen.yaml ${API_DIR}/webhooks.yaml > ${API_DIR}/webhooks.gen.go

gen/db:
	@echo Generating DB
	@pggen gen go --query-glob '${REPO_DIR}/queries/*.sql' --postgres-connection ${DB_URL} ${PGGEN_MAP}

gen/ui:
	templ generate

gen/go:
	@go mod tidy
	@go generate ./...
	@go mod tidy

# Database Commands
db: db/up db/push

db/up:
	@echo Starting DB
	@docker-compose up -d postgres
	@$(MAKE) wait CONTAINER=postgres

db/push:
	@echo Migrating DB
	@cd ${REPO_DIR} && migrate -database ${DB_URL} -path ./migrations up
