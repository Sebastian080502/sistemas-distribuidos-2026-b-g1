<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       13-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 13

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Start Corte 3 harden on a weekly cadence — promote identity E-01 422 to qa with a real review request; do not self-merge; do not retag v2.0.0.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-C3-ID-001 | Promote E-01 422 login envelope `develop` → `qa` | doing | [identity-api#7](https://github.com/code-corhuila/drp-identity-api/pull/7) OPEN against `qa`, review requested `ariel5253`. Cherry-pick `-x` of `35e0ee5` / `2a53f8e`. Not merged. |
| HU-C3-INF-001 | Composition root up (Postgres + Mongo + gateway :8080) | doing | Local `drp-infra-postgres` `up` reached `drp-postgres-dev` on host **5436**. Docker Desktop then lost the engine (OOM while pulling Mongo). Not claimed healthy. No compose PR (no committed change). |
| HU-C3-LIVE-001 | Front `environment.mode='live'` against gateway | todo | Still `failover` on `drp-front` `main`. Blocked on gateway + identity process on the compose network. |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Not this week. Spread work per Ariel Corte 3 note. |

## 2. My individual contribution
- Applied Ariel c2-sprints feedback (4.8/5.0, 2026-10-17): **one** product promotion this week, opened on **Tuesday**, review requested **before** merge, even as a one-person team. Did **not** merge #7.
- Did **not** cite closed-unmerged PRs (week 08 lesson: `drp-docs` #5 was closed without merge — not evidence).
- Did **not** backfill week 12 hu-status on 20 Oct (that would repeat “reports written late”). Week 12 template stays empty.
- Corte 3 scope this week is identity 422 on `qa` only. Reservation HTTP/outbox is already on `qa`/`main` from Corte 2 catch-up; 422 tests stay on `develop` until a later week.
- Confirmation stays event-driven (D-C14). Did not fake Camunda. Did not retag `v2.0.0`.

## 3. Blockers and risks
- Docker Desktop engine died (`dockerDesktopLinuxEngine` pipe missing) after Postgres start + Mongo pull + local `go test` OOM. Compose + gateway live path paused until the engine is restarted. Host port 5436 is the local override (5432 often busy); not committed.
- [identity-api#7](https://github.com/code-corhuila/drp-identity-api/pull/7) must wait for a **real** review. Bot Approve is not a merge.
- [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) (`amountCents` owner) still OPEN — do not invent an owner.
- Identity API still uses the in-memory Corte 2 seed; Flyway in `drp-identity-db` is not wired into `cmd/api` yet.

## 4. Plan for next week
- After #7 is reviewed: merge `qa`, then `hu-c3-*` cherry-pick `-x` → `main` (separate PR, separate week if review lands late).
- Restart Docker; Postgres + Mongo + identity-api + gateway on `drp-dev`; curl login through `:8080`.
- Next domain slice: reservation 422 tests `develop` → `qa` **or** identity Postgres adapter — not both in the same dump.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes:

- Branch `hu-c3-identity-422-qa` → PR to **`qa`**, not a direct commit on `qa`. Cherry-pick `-x` (not merge `develop` into `qa`).
- Tests in #7 are the develop contract tests re-applied; CI on the `qa` PR is the evidence, not a local `go test` (that OOM'd).
- This file is Week 13 **class** evidence. Product URLs are on `code-corhuila/drp-*`. No product SDD copied into `13-week/`.

## 6. Evidence links

### Product (`code-corhuila`) — this week
- https://github.com/code-corhuila/drp-identity-api/pull/7 (OPEN `qa`, review requested)

### Product — already on main from earlier (not new this week; not re-tagged)
- https://github.com/code-corhuila/drp-reservation-api/pull/7
- https://github.com/code-corhuila/drp-front/pull/10

### Course
- Profile CONFIG: https://github.com/Sebastian080502/Sebastian080502
- This class fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Do not retag product `v2.0.0`. This fork is not tagged `v2.0.0`.
