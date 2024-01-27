# Web-of-Hooks
# Project Structure
- `.vscode` VSCode configs
- `cmd` CLI Commands
  - `call`  Call Webhooks Grpc
  - `serve` Run Webhooks Manager
  - `worker` Run Webhooks Worker
- `demo`
  - `grpc`
  - `pubsub`
  - `temporal`
- `internal`
  - `domain`
    - `manager`
    - `worker`
  - `provide`
    - `repo`
      - `migrations`
      - `queries`
    - `secrets`
- `package`
  - `actor`
  - `app`
  - `util` 
- `webhooks`
  - `cron`
    - `v1`
  - `grpc`
    - `v1`
  - `html`
    - `v1`
  - `http`
    - `v1`
  - `subs`
    - `v1`

# High Level Architecture
```mermaid
graph LR
    Temporal --> Activity
    Customers --> API
    GraphQL --> Grpc
    PubSub --> Topic

    API --> Manager
    Grpc --> Manager
    Topic --> Manager
    Activity --> Manager

    Manager --> Repo
    Manager --> Secrets

    Cron --> Worker

    Worker --> Repo
    Worker --> Secrets
    Worker --> Endpoints 
```

# Entity-Relation Diagram
```mermaid
erDiagram
    EventType {
        string key
    }
    Channel {
        string key
    }
    Application {
        string key
    }
    Endpoint {
        string key
    }
    Message {
        string key
    }
    Attempt {
        string key
    }
    Message }o--|| EventType : "has one"
    Message }o--|| Channel : "has one"
    Message }o--|| Application : "has one"
    Attempt }o--|| Message : "has one"
    Attempt }o--|| Endpoint : "has one"
    Endpoint }o--|| Application : "has one"
    Endpoint }o--o{ Channel : "has many"
    Endpoint }o--o{ EventType : "filter"
```

# Make
- `make` alias for `make down db gen`
- `make up` alias for `docker-compose up -d`
- `make down` alias for `docker-compose down`
- `make wait CONTAINER=...` waits for container to be healthy 
- `make gen` alias for `make gen/grpc gen/api gen/db`
  - `make gen/grpc` Generates gRPC structurs
  - `make gen/api` Generates HTTP structures
  - `make gen/db` Generates SQL structures
- `make db` alias for `make db/up db/push`
  - `make db/up` starts db and waits for it to be healthy
  - `make db/push` runs db migrations

# Compose
- Postgres
- Vault

# Develop

- Clone: `git clone https://github.com/andrestielau/web-of-hooks`
- Enter: `cd web-of-hooks`
- Bootstrap: `make` (yes, just `make`)
- Run: `go run . serve`
- Test `https://port3000.[your-name].anchorlabs.dev/health`

# Tasks:
- Adapters
  - [ ] Http
  - [ ] Grpc
  - [ ] Subs
  - [ ] Work
- Services
  - [ ] Manager
  - [ ] Worker
- Providers
  - [ ] Repo
  - [ ] Secrets
  
# Goals 
- [ ] onboard tenants (create applications) (grpc + db)
- [ ] register event-types (CRUD event-types) (grpc + db)
- [ ] manage configs (edit application + CRUD endpoint) (http/grpc + db)
- [ ] register messages (create message) (grpc/http/pubsub/temporal + db)
- [ ] worker calls (dequeue + submit)
# Bonus
- [ ] Integrate grpc with GraphQL (grpc + graphql)
- [ ] Integrate with existing workflows (temporal)
- [ ] Integrate with existing API (http)
# Cherry on Top
- [ ] Integrate with Backoffice Dashboard
- [ ] Integrate with Client Dashboard