<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       09-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 09

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Corte 2 catch-up after weeks 6–8 — apply docs review findings, start the front synthetic shell, and open the infra Postgres composition root.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-DOC-001 | Apply PR #4 review findings (Anexo J records, naming, governance) | done | [drp-docs#8](https://github.com/code-corhuila/drp-docs/pull/8) MERGED |
| HU-CTR-001 | Land the Week 8 contract SSOT on `drp-docs` main | done | Same merge [drp-docs#8](https://github.com/code-corhuila/drp-docs/pull/8); [drp-docs#5](https://github.com/code-corhuila/drp-docs/pull/5) CLOSED after the content landed via #8 |
| HU-RES-UI-001 | Start the Angular synthetic contract shell for Corte 2 | doing | [drp-front#1](https://github.com/code-corhuila/drp-front/pull/1) OPEN (`hu-res-ui-dev` → `develop`; teacher Approve, not merged) |
| HU-INF-001 | Unique Postgres compose as composition root | doing | [drp-infra-postgres#1](https://github.com/code-corhuila/drp-infra-postgres/pull/1) OPEN (`hu-inf-001-dev` → `develop`) |
| HU-ID-001 | Start identity Flyway schema (migrate-only job) | doing | [drp-identity-db#1](https://github.com/code-corhuila/drp-identity-db/pull/1) OPEN |
| HU-INF-002 | Instructor created `drp-infra-mongo` after the request | done | [drp-docs#6](https://github.com/code-corhuila/drp-docs/issues/6) CLOSED |
| HU-PO-001 | Ask PO who owns `amountCents` | doing | [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) still OPEN |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Needs running `-api` on a promoted environment |

## 2. My individual contribution
- Sprint planning for Corte 2 catch-up after weeks 6–8: close the docs review loop first, then start published code on instructor `drp-*` repos.
- Merged product evidence on **`drp-docs`**: [PR #4](https://github.com/code-corhuila/drp-docs/pull/4) (weeks 6–7) and [PR #8](https://github.com/code-corhuila/drp-docs/pull/8) (review findings + Week 8 contract sheets). Did not treat closed [PR #5](https://github.com/code-corhuila/drp-docs/pull/5) as extra evidence.
- Opened the first code PRs: synthetic front shell and `drp-infra-postgres` composition root, plus identity Flyway. Teacher `ariel5253` often Approves via bot; those PRs are **not** on `main` yet.
- Did **not** copy product SDD into `09-week/02-session/` (empty instructor slot). This fork is weekly evidence only.

## 3. Blockers and risks
- Bot Approve is not a merge. Corte 2 still needs `develop` → `qa` → `main` with `cherry-pick -x` before a repo can be tagged.
- `hu-*-dev` branches were flagged by the course bot; later product PRs use `feat/` → `develop`.
- PO question on `amountCents` ([drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7)) is still open.
- Pact/CI (HU-C2-002) stays blocked until `-api` work is promoted and running.

## 4. Plan for next week
- Remaining Flyway/Liquibase `-db` schemas, Mongo instance, hexagonal domain models, and identity RS256 login.
- Keep the front as a synthetic failover for the demo; only cite published commits/PRs.
- Do not tag `v2.0.0` on this class fork.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- Product PRs target `develop` (not `qa`/`main` yet). Early branches used `hu-*-dev`; the bot flagged that pattern.
- Automated tests and promotion live on instructor `drp-*` repos, not on this class fork.
- This file is the Week 09 **class** evidence. Product URLs are on `code-corhuila/drp-*`.

## 6. Evidence links

### Product (`code-corhuila`) — published only
- https://github.com/code-corhuila/drp-docs/pull/4 (MERGED)
- https://github.com/code-corhuila/drp-docs/pull/8 (MERGED)
- https://github.com/code-corhuila/drp-docs/pull/5 (CLOSED; superseded by #8)
- https://github.com/code-corhuila/drp-front/pull/1
- https://github.com/code-corhuila/drp-infra-postgres/pull/1
- https://github.com/code-corhuila/drp-identity-db/pull/1
- https://github.com/code-corhuila/drp-docs/issues/6
- https://github.com/code-corhuila/drp-docs/issues/7

### Course
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- This class fork is **not** tagged `v2.0.0`.
