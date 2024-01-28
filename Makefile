include .env
export

all: down db gen

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
gen: gen/db gen/grpc gen/api gen/go

gen/grpc: ${GRPC_DIR}/webhooks.proto buf.gen.yaml 
	@echo Generating gRPC
	@buf generate

gen/api: ${API_DIR}/webhooks.yaml api.gen.yaml
	@echo Generating API
	@oapi-codegen --config api.gen.yaml ${API_DIR}/webhooks.yaml > ${API_DIR}/webhooks.gen.go

gen/db:
	@echo Generating DB
	@pggen gen go --query-glob ${REPO_DIR}/queries/queries.sql --postgres-connection ${DB_URL} ${PGGEN_MAP}

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
