<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       10-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 10

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Remaining Flyway/Liquibase `-db`, Mongo instance, hexagonal domain models, and identity RS256 login; front stays a synthetic failover for the demo.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-DB-001 | Remaining Flyway schemas (space, availability, reservation, payment) | doing | [space-db#1](https://github.com/code-corhuila/drp-space-db/pull/1) · [availability-db#1](https://github.com/code-corhuila/drp-availability-db/pull/1) · [reservation-db#1](https://github.com/code-corhuila/drp-reservation-db/pull/1) · [payment-db#1](https://github.com/code-corhuila/drp-payment-db/pull/1) — all OPEN |
| HU-DB-002 | Notification Liquibase Mongo changelog (migrate-only) | doing | [drp-notification-db#1](https://github.com/code-corhuila/drp-notification-db/pull/1) OPEN |
| HU-INF-002 | MongoDB instance per environment (document engine) | doing | [drp-infra-mongo#1](https://github.com/code-corhuila/drp-infra-mongo/pull/1) OPEN (`feat/mongo-instance` → `develop`) |
| HU-API-001 | Hexagonal domain models on the six `-api` repos | doing | [identity-api#1](https://github.com/code-corhuila/drp-identity-api/pull/1) · [space-api#1](https://github.com/code-corhuila/drp-space-api/pull/1) · [availability-api#1](https://github.com/code-corhuila/drp-availability-api/pull/1) · [reservation-api#1](https://github.com/code-corhuila/drp-reservation-api/pull/1) · [payment-api#1](https://github.com/code-corhuila/drp-payment-api/pull/1) · [notification-api#1](https://github.com/code-corhuila/drp-notification-api/pull/1) — all OPEN |
| HU-ID-001 | RS256 login and JWKS (front stays synthetic) | doing | [drp-identity-api#2](https://github.com/code-corhuila/drp-identity-api/pull/2) OPEN |
| HU-RES-UI-001 | Keep the published front as demo failover | doing | [drp-front#1](https://github.com/code-corhuila/drp-front/pull/1) only; later local redesign is **not** evidence until it has a PR |
| HU-UX-001 | Record the mostrador design direction on `drp-docs` | doing | [drp-docs#9](https://github.com/code-corhuila/drp-docs/pull/9) OPEN (`docs/design-direction`) |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Still needs promoted, running `-api` |
| HU-GW-001 | NGINX gateway on `drp-api-gateway` | todo | Repo has instructor scaffolding only; no student PR |

## 2. My individual contribution
- Continued Corte 2 after the Week 09 catch-up: opened the remaining `-db` migrate-only PRs (Flyway on Postgres domains, Liquibase on notification Mongo) and `drp-infra-mongo`.
- Published hexagonal domain models on all six `-api` repos and identity RS256 login/JWKS. Payment stays Java 21. No DDL inside `-api`.
- Front: the **published** Corte 2 shell remains the failover for the demo. Unpublished local redesign is not cited.
- Flow: `feat/` (or earlier `hu-*-dev`) → PR to `develop`. Teacher Approves via bot but has **not** merged most code PRs. Promotion to `qa`/`main` and tags are not started.
- Did **not** copy product SDD into `10-week/02-session/`. This fork is weekly evidence only and is **not** tagged `v2.0.0`.

## 3. Blockers and risks
- Almost all Corte 2 code is still OPEN on `develop`. Empty `main` must not be tagged (Jose Miguel: empty repos are not tagged).
- Dual-write / outbox and Pact/CI remain later. `amountCents` ownership is still [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7).
- `drp-api-gateway` has no NGINX PR yet — do not tag it until that lands on `main`.
- Six `*-portal`, `drp-worker`, `drp-workflow`, and `distributed-reservation-platform` stay out of the `v2.0.0` set (README-only or not MVP).

## 4. Plan for next week
- After this class evidence is on fork `main`: continue product work in the order docs → front → infra (postgres + mongo) → `-db`/`-api` promotion. Do not start tags until real MVP is on `main`.
- Merge [drp-docs#9](https://github.com/code-corhuila/drp-docs/pull/9) when the teacher accepts it. Keep asking for merges of the OPEN code PRs.
- Week 11+ hu-status is **not** filled yet (deadline Corte 2: 15 Oct 2026).

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- Product uses `feat/` → `develop` after the bot flagged `hu-*-dev`. No student commits on permanent `qa`/`main` yet.
- Domain-model PRs are published; broader unit/integration suites wait for merge and a running compose.
- This file is the Week 10 **class** evidence. Product URLs are on `code-corhuila/drp-*`.

## 6. Evidence links

### Product (`code-corhuila`) — published only
- https://github.com/code-corhuila/drp-docs/pull/9
- https://github.com/code-corhuila/drp-infra-mongo/pull/1
- https://github.com/code-corhuila/drp-space-db/pull/1 · https://github.com/code-corhuila/drp-availability-db/pull/1 · https://github.com/code-corhuila/drp-reservation-db/pull/1 · https://github.com/code-corhuila/drp-payment-db/pull/1 · https://github.com/code-corhuila/drp-notification-db/pull/1
- https://github.com/code-corhuila/drp-identity-api/pull/1 · https://github.com/code-corhuila/drp-identity-api/pull/2
- https://github.com/code-corhuila/drp-space-api/pull/1 · https://github.com/code-corhuila/drp-availability-api/pull/1 · https://github.com/code-corhuila/drp-reservation-api/pull/1 · https://github.com/code-corhuila/drp-payment-api/pull/1 · https://github.com/code-corhuila/drp-notification-api/pull/1
- Front failover (same published PR): https://github.com/code-corhuila/drp-front/pull/1

### Course
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- This class fork is **not** tagged `v2.0.0`.
