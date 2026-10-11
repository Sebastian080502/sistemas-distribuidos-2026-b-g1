<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       11-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 11

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Promote the Corte 2 freeze to main and tag v2.0.0 for Moodle c2-release; start Corte 3 walking HTTP/outbox on develop. Front stays failover for the freeze demo.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-C2-REL-001 | Corte 2 v2.0.0 freeze on product main | done | [drp-docs v2.0.0](https://github.com/code-corhuila/drp-docs/releases/tag/v2.0.0) · [drp-front#3](https://github.com/code-corhuila/drp-front/pull/3) · [drp-infra-postgres#3](https://github.com/code-corhuila/drp-infra-postgres/pull/3) · [drp-infra-mongo#3](https://github.com/code-corhuila/drp-infra-mongo/pull/3) · [drp-api-gateway#3](https://github.com/code-corhuila/drp-api-gateway/pull/3) · [reservation-api#3](https://github.com/code-corhuila/drp-reservation-api/pull/3) — `release(c2)` → `main` then tag. Not `drp-worker` / `drp-workflow` (empty, untagged) |
| HU-DOC-002 | Align OpenAPI login with E-01 (no user, 422) | done | [drp-docs#11](https://github.com/code-corhuila/drp-docs/pull/11) MERGED `main` 2026-10-09 — after the v2.0.0 tag; do not retag |
| HU-DOC-003 | Close G-06 now that `drp-infra-mongo` exists | done | [drp-docs#10](https://github.com/code-corhuila/drp-docs/pull/10) MERGED `main` |
| HU-UX-001 | SpaceHub mostrador design direction | done | [drp-docs#9](https://github.com/code-corhuila/drp-docs/pull/9) MERGED `main` (was OPEN in Week 10) |
| HU-GW-001 | NGINX presence-only gateway | done | [gateway#1](https://github.com/code-corhuila/drp-api-gateway/pull/1) · [#2](https://github.com/code-corhuila/drp-api-gateway/pull/2) · [#3](https://github.com/code-corhuila/drp-api-gateway/pull/3) + [tag](https://github.com/code-corhuila/drp-api-gateway/releases/tag/v2.0.0) (was todo in Week 10) |
| HU-INF-001 | Postgres composition root promoted and tagged | done | [infra-postgres#1](https://github.com/code-corhuila/drp-infra-postgres/pull/1) · [#2](https://github.com/code-corhuila/drp-infra-postgres/pull/2) · [#3](https://github.com/code-corhuila/drp-infra-postgres/pull/3) + [tag](https://github.com/code-corhuila/drp-infra-postgres/releases/tag/v2.0.0) |
| HU-INF-002 | Mongo instance promoted and tagged | done | [infra-mongo#1](https://github.com/code-corhuila/drp-infra-mongo/pull/1) · [#2](https://github.com/code-corhuila/drp-infra-mongo/pull/2) · [#3](https://github.com/code-corhuila/drp-infra-mongo/pull/3) + [tag](https://github.com/code-corhuila/drp-infra-mongo/releases/tag/v2.0.0) |
| HU-DB-001 | Remaining Flyway/Liquibase schemas promoted | done | Week 10 OPEN PRs promoted: develop → `hu-c2-qa` → `release.2.0.0` → `main` + tag. Example [reservation-db#1](https://github.com/code-corhuila/drp-reservation-db/pull/1) · [#2](https://github.com/code-corhuila/drp-reservation-db/pull/2) · [#3](https://github.com/code-corhuila/drp-reservation-db/pull/3) |
| HU-API-001 | Hexagonal domain models promoted | done | Same promotion on the six `-api` repos. Example [reservation-api#1](https://github.com/code-corhuila/drp-reservation-api/pull/1) · [#2](https://github.com/code-corhuila/drp-reservation-api/pull/2) · [#3](https://github.com/code-corhuila/drp-reservation-api/pull/3) |
| HU-ID-001 | RS256 login promoted; E-01 422 on develop | done | [identity-api#2](https://github.com/code-corhuila/drp-identity-api/pull/2) · [#3](https://github.com/code-corhuila/drp-identity-api/pull/3) · [#4](https://github.com/code-corhuila/drp-identity-api/pull/4) on the freeze; [#5](https://github.com/code-corhuila/drp-identity-api/pull/5) MERGED `develop` (drop user from login 422) |
| HU-RES-UI-001 | Front failover published and tagged | done | [drp-front#1](https://github.com/code-corhuila/drp-front/pull/1) · [#2](https://github.com/code-corhuila/drp-front/pull/2) · [#3](https://github.com/code-corhuila/drp-front/pull/3) + [tag](https://github.com/code-corhuila/drp-front/releases/tag/v2.0.0). Unpublished local redesign is **not** evidence |
| HU-RES-HTTP-001 | Walking HTTP E-11..E-14 (Corte 3, develop only) | doing | [reservation-api#4](https://github.com/code-corhuila/drp-reservation-api/pull/4) MERGED `develop` 2026-10-11 — **not** in v2.0.0; not promoted to qa/main |
| HU-OUTBOX-001 | `reservation.outbox` table (not a Camunda saga) | doing | [reservation-db#4](https://github.com/code-corhuila/drp-reservation-db/pull/4) MERGED `develop` 2026-10-11. Walking in-memory + Flyway table. No publisher/broker. `drp-workflow` empty/untagged |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Still needs Pact/CI; reservation-api has partial `go test` for outbox/create only |
| HU-PO-001 | Ask PO who owns `amountCents` | doing | [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) still OPEN — do not invent an owner |

## 2. My individual contribution
- Sprint planning: reviewed the Moodle **c2-release** / **c2-sprints** rubric (func / quality / arch / scrum + group / individual / ceremonies). Delivery is through GitHub repos, not a zip. Moodle due 13 Oct 2026.
- Review: after `ariel5253` Approve, I (student/agent) merged product PRs. Bot Approve is **not** a merge.
- Promotion this week: `hu/feat` → `develop`, then `hu-c2-qa` with `cherry-pick -x` → `qa`, then `release.2.0.0` → `main`, then tag `v2.0.0` only on repos with real work. GitHub Releases UI may look empty; git tags exist at `/releases/tag/v2.0.0`.
- Daily / retro: I do **not** invent standup minutes. Review and planning happened around the freeze and the Moodle checklist. Ceremony evidence is this hu-status plus the merged PRs per environment.
- After the freeze, started Corte 3 on `develop`: walking HTTP E-11..E-14 and `reservation.outbox`. Confirmation is event-driven (outbox), not a saga coordinator. Did not fake Camunda.
- Profile CONFIG is published on [Sebastian080502/Sebastian080502](https://github.com/Sebastian080502/Sebastian080502).
- Did **not** copy product SDD into `11-week/02-session/`. This class fork is weekly evidence only and is **not** tagged `v2.0.0`.

## 3. Blockers and risks
- Corte 3 walking HTTP / outbox is on `develop` only. Do **not** retag `v2.0.0` (docs#11 also landed after the tag).
- Outbox has no publisher/broker yet. `drp-workflow` is empty and correctly untagged. Not a running broker and not a fake Camunda saga.
- Local compose: `drp-postgres-dev` healthy with `POSTGRES_PORT=5436` because 5432 was occupied (env not committed). Gateway `:8080` blocked locally by pgAdmin.
- Front Corte 2 demo stays `environment.mode='failover'`. Classroom accounts: `member@spacehub.local` / `admin@spacehub.local` / `Spacehub1!`.
- [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) (`amountCents`) still OPEN.
- HU-C2-002 contract tests still todo.
- Week 12 hu-status is **not** started.

## 4. Plan for next week
- Paste Moodle **c2-release** (product `v2.0.0` tag links) and **c2-sprints** (this fork + profile CONFIG). Links, not a zip.
- Corte 3 later: promote reservation HTTP/outbox via `hu-c3-*` cherry-pick to `qa`; payment Java HTTP; notification Mongo HTTP via other domains' APIs, never their tables.
- Do **not** fill Week 12 until that academic week.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on remaining gaps:

- This week did promote `develop` → `qa` → `main` (`hu-c2-qa`, `release.2.0.0`). Corte 3 uses `hu-c3-*-dev` → `develop` for reservation (not in the freeze).
- Tests: reservation-api `go test` for outbox/create is partial. Full Pact/CI (HU-C2-002) is still todo. No fake saga tests.
- Outbox publisher/broker is not running yet.
- This file is the Week 11 **class** evidence. Product URLs are on `code-corhuila/drp-*`.

## 6. Evidence links

### Product (`code-corhuila`) — Corte 2 freeze tags
- https://github.com/code-corhuila/drp-docs/releases/tag/v2.0.0
- https://github.com/code-corhuila/drp-front/releases/tag/v2.0.0
- https://github.com/code-corhuila/drp-infra-postgres/releases/tag/v2.0.0
- https://github.com/code-corhuila/drp-infra-mongo/releases/tag/v2.0.0
- https://github.com/code-corhuila/drp-api-gateway/releases/tag/v2.0.0

### Product — representative MERGED PRs (not a full inventory)
- https://github.com/code-corhuila/drp-docs/pull/9 (MERGED) · https://github.com/code-corhuila/drp-docs/pull/10 (MERGED) · https://github.com/code-corhuila/drp-docs/pull/11 (MERGED after tag)
- https://github.com/code-corhuila/drp-front/pull/3 · https://github.com/code-corhuila/drp-infra-postgres/pull/3 · https://github.com/code-corhuila/drp-infra-mongo/pull/3 · https://github.com/code-corhuila/drp-api-gateway/pull/3
- https://github.com/code-corhuila/drp-reservation-api/pull/3 · https://github.com/code-corhuila/drp-reservation-db/pull/3 · https://github.com/code-corhuila/drp-identity-api/pull/4

### Product — Corte 3 on develop (not the freeze)
- https://github.com/code-corhuila/drp-reservation-api/pull/4 (MERGED develop)
- https://github.com/code-corhuila/drp-reservation-db/pull/4 (MERGED develop)
- https://github.com/code-corhuila/drp-identity-api/pull/5 (MERGED develop)
- https://github.com/code-corhuila/drp-docs/issues/7 (OPEN)

### Course
- Profile CONFIG: https://github.com/Sebastian080502/Sebastian080502
- This class fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- This class fork is **not** tagged `v2.0.0`.
