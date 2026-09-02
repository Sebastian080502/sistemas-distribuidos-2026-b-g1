# Week 03 session notes — Distributed Reservation Platform

**Course:** Distributed Systems 2026-B · Group 1  
**Student:** Juan Sebastian Osorio Fierro  

## Session 1 — Documentation repository structure

The class reviewed the reusable `docs/` layout used as the single source of truth for each product. Folders `00`–`15` are a reading order (governance → operations). `99-archive` keeps history.

For this project that repo is **`code-corhuila/drp-docs`**. Documentation commits go **directly to `main`**. Feature branches are not used to version docs.

The code MVP is a **different** repository: `code-corhuila/distributed-reservation-platform`. It stays empty until HU-01 and HU-02 are closed in docs.

## How this maps to SpaceHub / DRP

| Folder | What it holds for this project |
|--------|--------------------------------|
| `00-governance` | Git, DoR/DoD, security |
| `01-context` | SpaceHub problem, scope, glossary |
| `02-domain` | Five bounded contexts and rules |
| `03-product` / `04-requirements` | Vision, HU-RES backlog, Gherkin |
| `05-architecture` | ADRs (stack + later microservices) |
| `06-data` | One PostgreSQL model per service |
| `07-api` | OpenAPI contracts |
| `09-microservices` | Service catalog (identity, space, reservation, payment, notification) |
| `12-ux-ui` | Navigation and iterative mockup |

## Corte 1 HUs on `drp-docs`

1. **HU-01** — Choose and justify Angular + Go + PostgreSQL.  
2. **HU-02** — Discovery: problem, users, MVP in/out, initial data model, main flows.

Both issues already exist on the org project. Closing them requires a push of the local docs, not code in the MVP repo.
