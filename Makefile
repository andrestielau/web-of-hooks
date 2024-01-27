include .env
export

all: boot down db gen

boot:
	go mod tidy
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/jschaf/pggen/cmd/pggen@latest
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

run: 
	go run . serve

dev:
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ."

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
gen: gen/grpc gen/api gen/db

gen/grpc: ${GRPC_DIR}/webhooks.proto webhooks/buf.gen.yaml 
	@echo Generating gRPC
	@cd webhooks && buf generate

gen/api: ${API_DIR}/webhooks.yaml webhooks/api.gen.yaml
	@echo Generating API
	@oapi-codegen --config webhooks/api.gen.yaml ${API_DIR}/webhooks.yaml > ${API_DIR}/webhooks.gen.go

gen/db:
	@echo Generating DB
	@pggen gen go --query-glob ${REPO_DIR}/queries/queries.sql --postgres-connection ${DB_URL} ${PGGEN_MAP}

# Database Commands
db: db/up db/push

db/up:
	@echo Starting DB
	@docker-compose up -d postgres
	@$(MAKE) wait CONTAINER=postgres

db/push:
	@echo Migrating DB
	@cd ${REPO_DIR} && migrate -database ${DB_URL} -path ./migrations up
