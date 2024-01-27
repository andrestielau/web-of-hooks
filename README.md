# Web-of-Hooks
Welcome to our 2024's Hackathon project!  
This week we'll be working on a new notification solution: WebHooks!! :tada:  
I hope you're all excited for this, we'll need motivation to endure the hardships ahead.  
We're facing some tough challenges, so we should prepare ourselves first.  
Here's a list of topics that will come up in our journey:

# Project
## Structure
Y'all found this repository with an initial structure, this doesn't mean that the structure won't change, but if it does please update this document to reflect the latest structure.
- `.vscode` VSCode configs
- `cmd` CLI Commands
  - `call`  Call Webhooks Grpc
  - `serve` Run Webhooks Manager
  - `worker` Run Webhooks Worker
- `demo` Examples
  - `grpc` Create Messages from gRPC
  - `pubsub` Create Messages from PubSub
  - `temporal` Create Messages from Temporal
- `internal` Implementation
  - `domain` Domain Logic
    - `manager` Service to Manage Data
    - `worker` Service to Send WebHook Messages
  - `provide` Provider Implementations
    - `repo` Postgres Access
      - `migrations` Database DDL
      - `queries` Database DQL/DML
    - `secrets` Vault Access 
- `package` Generic Code
  - `actor` Lifecycle Manager
    - `cron` Base Cron Adapter
    - `net` Network Modules
      - `grpc` Grpc Actors
        - `server` Grpc Adapter
        - `client` Grpc Provider
      - `http` Http Actors
        - `server` Http Adapter
        - `client` Http Provider
  - `app` Application Utilities
    - `cfg` Configuration Utilities
    - `cmd` Command Line Utilities
    - `flag` Flag Utilities
  - `util` Generic Utilities
- `webhooks` Application Adapters
  - `cron` Cron for Workers
  - `grpc` Internal API
    - `v1` Grpc Definition
  - `html` Browser API
    - `components` Reusable Atomic Components
      - `atom` Simple Composable Components
      - `molecule` Encapsulation of Components
    - `layouts` Reusable Aggregate Dispositions
    - `pages` Pages to Render
    - `scripts` Reusable JavaScript
    - `styles` Style Definitions
    - `utils` Utility Methods
  - `http` External API
    - `v1` OAPI Definition
  - `pub` Publisher
    - `v1` Produced Message Definition
  - `subs` Subscriber
    - `v1` Consumed Message Definition 

## High Level Architecture
Since it's not easy to understand the organization of an application from the folder structure alone, here's a little drawing to help you get a high level idea of the pieces of this service and how they're conected.
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

## Entity-Relation Diagram
Since you'll be persisting data you need to be aware of the structure and relations that the data has, for that purpose you can use the following diagram to refresh your memory.
TODO: This section is still incomplete, please update it when possible.
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

# Tools
In order to make your developement experience more pleasant this repository makes use of some third-party tools.  
You can check the following links for more documentation on each:

## [Make](https://makefiletutorial.com/)

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

## [Docker Compose](https://docs.docker.com/compose/)
- `postgres` database
- `vault` secret manager

## [OAPI CodeGen](https://github.com/deepmap/oapi-codegen)
OpenAPI Client and Server Code Generator
## [Migrate](https://github.com/golang-migrate/migrate)
Database migrations written in Go
## [PgGen](https://github.com/jschaf/pggen)
Generate type safe Go methods from Postgres SQL queries
## [Templ](https://templ.guide/)
An HTML templating language for Go that has great developer tooling

# Develop

- Clone: `git clone https://github.com/andrestielau/web-of-hooks`
- Enter: `cd web-of-hooks`
- Bootstrap: `make` (yes, just `make`)
- Run: `go run . serve`
- Test `https://port3000.[your-name].anchorlabs.dev/health`

## Tasks:
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
  
## Goals 
- [ ] onboard tenants (create applications) (grpc + db)
- [ ] register event-types (CRUD event-types) (grpc + db)
- [ ] manage configs (edit application + CRUD endpoint) (http/grpc + db)
- [ ] register messages (create message) (grpc/http/pubsub/temporal + db)
- [ ] worker calls (dequeue + submit)
## Bonus
- [ ] Integrate grpc with GraphQL (grpc + graphql)
- [ ] Integrate with existing workflows (temporal)
- [ ] Integrate with existing API (http)
## Cherry on Top
- [ ] Integrate with Backoffice Dashboard
- [ ] Integrate with Client Dashboard
  
## Other Ideas
- Reusable Secrets
- Publish Errors and/or Successes
- OTEL Metrics
- Store Responses (and maybe requests) in FileStorage
- `Content-Type` Negotiator (get preferred content type from `Accept` headers and response formats for each endpoint)

## Common issues
### Zombie process
When you get the error `listen tcp :3000: bind: address already in use`
Run: `netstat -nlp | grep 3000`
You'll see something like: `tcp6 0 0 :::3000 :::* LISTEN {PID}/web-of-hooks`
Then run: `kill -9 {PID}`