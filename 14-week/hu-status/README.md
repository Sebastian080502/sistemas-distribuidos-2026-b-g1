<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       14-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 14

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Keep Corte 3 on a weekly cadence — submit this hu-status, leave Week 13 PR #15 open for a real review, and verify identity login against the Anexo J Postgres instance (schema identity, not a second engine). No invented Week 14 workshop.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-C3-ID-001 | Promote E-01 422 login envelope `develop` → `qa` | doing | Still [identity-api#7](https://github.com/code-corhuila/drp-identity-api/pull/7) OPEN against `qa`, review requested `ariel5253`. Not merged. Carried from Week 13; no new promotion this week. |
| HU-C3-W13-001 | Week 13 class evidence (threat-model workshop + hu-status) | doing | Fork [PR #15](https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/pull/15) OPEN (`docs/week-13-hu-status` → `main`). I do **not** self-merge. |
| HU-C3-INF-001 | Composition root Postgres (Anexo J, host 5436) | doing | Local `drp-infra-postgres` only. Engine was restarted this week after a previous Docker Desktop crash. Flyway `identity-migrate` and live login vs `identity.users` are the verification target — not a compose PR (no committed env). |
| HU-C3-ID-002 | Identity API reads `identity` schema (no DDL in `-api`) | doing | Adapter already on `develop` (`DATABASE_URL` → `identity_app`). Local verify this week: Flyway then `POST /api/v1/auth/login` on `:8081`. Unpublished local runs are **not** tagged. |
| HU-C3-LIVE-001 | Front `environment.mode='live'` against gateway | todo | Still `failover` on `drp-front` `main`. Not this week. |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Not this week. |

## 2. My individual contribution
- Filled **this** Week 14 `hu-status` on academic week 14 (2026-11-09). Did **not** invent a Week 14 workshop: upstream `14-week/01-session` and `14-week/02-session` are still `.gitkeep` only. The Sync of `code-corhuila/main` brought `certification/` templates; that is Corte 3 Huawei structure, not a Week 14 class session.
- Left [fork PR #15](https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/pull/15) OPEN. Ariel c2-sprints: request review, do not self-merge even as a one-person team.
- Product work this week stays on **identity persist + login** against the single Postgres instance (`drp-infra-postgres`, schema `identity`). No AMQP. No Camunda. No self-merge of product PRs.
- Did **not** fill weeks 12, 15, or 16. Did **not** push the local `docs/c3-certification` branch (Huawei screenshots stay unpublished until the cert README is complete).
- Did **not** copy product SDD into `14-week/`. This fork is weekly evidence only and is **not** tagged `v2.0.0`.

## 3. Blockers and risks
- No Week 14 session files from the instructor. Filling a fake `02-session` would be a grave fault. This README is the only Week 14 class deliverable until material appears.
- Docker Desktop previously died (`dockerDesktopLinuxEngine` pipe missing / OOM). Host publishes Postgres on **5436** because 5432 is often busy; that override is local (`env/dev.env`), not committed.
- [identity-api#7](https://github.com/code-corhuila/drp-identity-api/pull/7) still waits for a **real** review. Bot Approve is not a merge.
- [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) (`amountCents` owner) still OPEN — do not invent an owner.
- Anexo A probes 23–24 (notification report with payment down / export FAILED) need payment + notification up. Not claimed this week.
- Confirmation stays event-driven (D-C14). Outbox has no broker (Q-001).

## 4. Plan for next week
- After #7 is reviewed: merge `qa`, then a **separate** `hu-c3-*` cherry-pick `-x` → `main` (separate PR).
- If Flyway + login vs PG succeed: next walking slice is reservation 422 tests on `qa` **or** payment Java HTTP — not both in the same dump. Do not start AMQP.
- Fill Week 15 `hu-status` only when that academic week starts. Keep 12/16 empty unless the instructor publishes those weeks.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked item:

- No new product test PR this week. Identity contract tests already live on `develop` / identity-api#7. Local `go test ./...` is avoided on this machine (OOM).
- Branch `docs/week-14-hu-status` → PR to fork **`main`**. Class fork is not a product `feat/` repo.
- This file is Week 14 **class** evidence. Product URLs are on `code-corhuila/drp-*`.

## 6. Evidence links

### Product (`code-corhuila`) — carried (OPEN, not merged by me)
- https://github.com/code-corhuila/drp-identity-api/pull/7 (OPEN `qa`, review requested)

### Course
- Week 13 still OPEN on this fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/pull/15
- Profile CONFIG: https://github.com/Sebastian080502/Sebastian080502
- This class fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Instructor `certification/` templates landed on fork `main` via Sync (`aa0bc76`). Screenshots are **not** in this PR.
- Do not retag product `v2.0.0`. This fork is not tagged `v2.0.0`.
