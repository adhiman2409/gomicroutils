# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

# Overview

`github.com/adhiman2409/gomicroutils` is a shared Go library (module, no `main` package) imported by ~20+ UNIRMS microservices. Changes here ripple into every consumer, so prefer purely additive changes to exported structs/functions (see `ProjectTasks.md` for past examples of the checks done before altering a shared struct).

## Commands

```bash
go build ./...            # compile everything (the only routine check; CI does not build or test)
go vet ./...
gofmt -l -w <files>       # formatting is expected on every edited file
go test ./...             # only crypto_utils_test.go and distance/osrm_test.go exist
go test -run TestName ./distance/
```

Regenerating protobuf code (`proto/**/*.proto` -> `genproto/**`, protoc v3.20.3 / protoc-gen-go v1.33.0 per file headers; no Makefile or script exists, run from repo root since the generated `source:` paths are `proto/<svc>/<file>.proto`):

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/<svc>/<file>.proto
```

Never hand-edit `genproto/` (generated, `DO NOT EDIT`).

## Architecture

Packages are independent utilities; the non-obvious wiring:

- **`hrms/`** (largest package, ~100 files): shared domain models (`bson`/`json` tagged structs) and enums for employees, attendance, leave, payroll, assets, monitoring, etc. `hrms/collections.go` holds the Mongo collection-name constants, prefixed per owning service (`AUTH_COLLECTION_*`, `ORG_COLLECTION_*`, ...), plus helpers that derive per-tenant collection names. Add new collection names there rather than inline in services. Bson tag/zero-value semantics matter because existing documents are decoded into these structs (e.g. `IsInactive` rather than `IsActive` so legacy docs default to active).
- **`proto/` -> `genproto/` -> `grpcclient/`**: `.proto` sources, generated stubs, and hand-written client wrappers for the auth, org, mail, notification, logger, and monitor services. Each `*_client.go` has `Start<X>Client()`/`Stop...` and a package-level singleton accessed via `Get<X>Client()`; `*_requester.go`/`*_info.go` files wrap request/response conversion so services don't touch protobuf types. All clients dial with mutual-TLS-style credentials from `loadTLSCredentials()` (CA at `/app/cert/$CA_CERT_FILE`, default `ca-cert.pem`) and read the service address from env vars (e.g. `AUTH_SERVICE_URL`, default `auth-srv:50051`). `grpcclient/README.md` documents the monitor client.
- **`middleware/`**: gorilla/mux HTTP middleware. `RequestAuth` extracts the Bearer token, calls `grpcclient.GetAuthClient().Verify(token, routeName)` (route name = `mux.CurrentRoute(r).GetName()`, so routes must be named), and places claims JSON in the request context under `"claims"`; unauthenticated requests continue as role `anonymous`. Requires `StartAuthClient()` to have been called.
- **`amqp/`**: RabbitMQ singleton (`amqp.Get()`) plus publishers for logging, tracking, mail, SMS, and worker request/response queues.
- **`gcs/`**: Google Cloud Storage helpers (uploads for logos, salary slips, policies, tickets, inventory, etc.).
- **Small utilities**: `logger/` (zap), `tracer/` (Zipkin), `errs/` (REST error types), `timex/`, `distance/` (OSRM client), `ampl/`, `models/`, root `crypto_utils.go`.

## CI

`.github/workflows/timesheet.yml` only logs commit timesheets (`timesheet.csv`); there is no build/test pipeline.

# Project Management

This repo is shared across multiple developers. It has a resident **project-manager** agent (`.claude/agents/project-manager.md`) that owns:

- `ProjectTasks.md` — the task ledger (Task ID, Summary, Status, Start/End Date, Remarks, Created By)
- All documentation upkeep
- Code review before any task is marked `Finished`

**Before making any code or docs change in this repo, invoke the `project-manager` agent** rather than doing the work directly in the main thread. This applies whether the request comes from a developer typing in Claude Code or from another AI agent operating on this repo. The project-manager agent will log the task and delegate to the right specialist subagent (code review, build fixes) based on this repo's tech stack. It does **not** write or run tests, and does **not** run `git commit` / `git push` / PR operations — both stay with the developer.

If you're unsure whether a request qualifies (e.g. a read-only question, or pure exploration with no changes), it's fine to answer directly — only changes to the working tree need to route through project-manager.
