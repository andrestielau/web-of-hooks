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
  - `serve` Run Application
- `demo` Examples
  - `grpc` Create Messages from gRPC
  - `pubsub` Create Messages from PubSub
  - `temporal` Create Messages from Temporal
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
    - `third` Third party ports
      - `temporal`
        - `client` Provider for temporal
        - `worker` Adapter for Activities and Workflows
      - `vault` Provider for Vault
  - `app` Application Utilities
    - `cfg` Configuration Utilities
    - `cmd` Command Line Utilities
    - `flag` Flag Utilities
  - `util` Generic Utilities
- `webhooks` Implementation
  - `adapt` Application Adapters
    - `cron` Cron for Workers
    - `grpc` Internal API
      - `v1` Grpc Definition
    - `http` External API
      - `v1` OAPI Definition
    - `pub` Publisher
      - `v1` Produced Message Definition
    - `subs` Subscriber
      - `v1` Consumed Message Definition 
    - `work` Workflows and Activities
      - `v1` Workflow or Activity Message Definition 
  - `provide` Provider Implementations
    - `pub` Publisher
    - `repo` Postgres Access
      - `migrations` Database DDL
      - `queries` Database DQL/DML
    - `secrets` Vault Access 
  - `render` HTML renders
    - `components` Reusable Atomic Components
      - `atom` Simple Composable Components
      - `molecule` Encapsulation of Components
    - `layouts` Reusable Aggregate Dispositions
    - `pages` Pages to Render
    - `scripts` Reusable JavaScript
    - `styles` Style Definitions
    - `utils` Utility Methods

## High Level Architecture
Since it's not easy to understand the organization of an application from the folder structure alone, here's a little drawing to help you get a high level idea of the pieces of this service and how they're conected.
```mermaid
graph LR
    Temporal
    Activity
    Topic
    Grpc
    API

    Activity --> Repo
    Activity --> PubSub
    Temporal --> Activity
    Customers --> API
    GraphQL --> Grpc
    PubSub --> Topic

    Activity --> Secrets
    Cron --> Secrets
    Grpc --> Secrets
    API --> Secrets

    Cron --> Repo
    Grpc --> Repo
    API --> Repo
    Topic --> Repo
    Cron --> Endpoints 
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
    Secret {
        string key
    }
    Message }o--|| EventType : "has one"
    Message }o--|| Channel : "has one"
    Message }o--|| Application : "has one"
    Attempt }o--|| Message : "has one"
    Attempt }o--|| Endpoint : "has one"
    Application ||--o{ Secret : "has"
    Endpoint }o--|| Application : "has one"
    Endpoint }o--o{ Channel : "has many"
    Endpoint }o--o{ EventType : "filter"
    Endpoint }o--o| Secret : "use"
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
- Run: `make dev` Unless migrations change, you only need to change this.
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


# Notes
Why separate secret management from endpoint configuration?
- Decoupled storage:
  instead of salting, hashing, or alike and store secrets on the database, these can be stored in vault so that they're never exposed.
- Reusability
  since some clients might want to receive webhooks in different endpoints on the same server, they might also want to use the same secret.


# Validation Process:
- Using Authorized API Key, customer requests latest validation keys for endpoints
- Whenever a customer endpoint receives a message, they can use the endpoint's key to 