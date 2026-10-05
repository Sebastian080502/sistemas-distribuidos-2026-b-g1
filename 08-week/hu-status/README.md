<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       08-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 08

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Close SpaceHub HTTP behaviour (Week 8 contract density: decisions, endpoint sheets, declared gaps) on drp-docs, using the class 08-week/02-session/spec/api-contract.md as the rigor model — not as the product domain.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-CTR-001 | Mature OpenAPI `/api/v1` to Norma 5.3.5–5.3.12 (`{data, meta}`, 422, Idempotency-Key) | done | [drp-docs#5](https://github.com/code-corhuila/drp-docs/pull/5) · `07-api/contracts/openapi/` |
| HU-CTR-002 | Write the behaviour SSOT (D-C1–D-C20, E-01–E-23, gaps, declared debt) | done | [drp-docs#5](https://github.com/code-corhuila/drp-docs/pull/5) · `07-api/api-contract.md` |
| HU-RES-009 | Specify exportable reports (Mongo on notification; HTTP to owner APIs) | done | [drp-docs#5](https://github.com/code-corhuila/drp-docs/pull/5) · ADR-007 |
| HU-GIT-001 | Align `git-conventions.md` with the course branching policy | done | Closes [drp-docs#3](https://github.com/code-corhuila/drp-docs/issues/3) in PR #5 |
| HU-INF-001 | Request instructor-created `drp-infra-mongo` (Anexo J.8.3) | done | [drp-docs#6](https://github.com/code-corhuila/drp-docs/issues/6) |
| HU-PO-001 | Ask PO who owns `amountCents` | done | [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Still needs running `-api` repos |

## 2. My individual contribution
- Read the class `08-week/02-session/spec/api-contract.md` (D-C / E-nn density) and applied that rigor to **SpaceHub**, not Simple Stock Flow.
- Individual product PR on **`drp-docs`** (not this class repo): Anexo J stack (Java on `drp-payment-api`, React on `drp-notification-portal`, Mongo on `notification`), `{data, meta}` lists, 422 for overlap, confirmation as an event, HU-RES-009 reports.
- Did **not** copy product SDD into `08-week/02-session/` (instructor material). Product contracts stay in `code-corhuila/drp-docs`.

## 3. Blockers and risks
- `drp-infra-mongo` does not exist yet — requested in drp-docs#6; students must not create `drp-*` remotes (J.8.3).
- Annexes A–I on SharePoint were not readable from this workstation; Norma + Anexo J PDFs and `template.zip` were used instead.
- Pact/CI (HU-C2-002) still blocked until domain `-api` implementations exist.
- drp-docs PR #4 (weeks 6–7) is approved and should merge **before** PR #5 so GitHub does not mix the two diffs.

## 4. Plan for next week
- After the teacher merges drp-docs #5, start `drp-infra-postgres` (composition root) and `drp-identity-db` / `drp-identity-api` against the published contract.
- If `drp-infra-mongo` is created, add `drp-notification-db` migrations (Liquibase Mongo, own changelog table).

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- Code MVP, `hu-*-dev` branches, and automated tests live on instructor `drp-*` repos, not on this class fork.
- This file is the Week 08 **class** evidence. The product contract PR is drp-docs#5.

## 6. Evidence links

### Class material used as the rigor model (do not treat as SpaceHub)
- [`../02-session/spec/api-contract.md`](../02-session/spec/api-contract.md)
- [`../02-session/spec/constitution.md`](../02-session/spec/constitution.md)

### Product documentation (`drp-docs`) — Week 8 individual PR
- Pull request: https://github.com/code-corhuila/drp-docs/pull/5
- Behaviour SSOT: `07-api/api-contract.md`
- OpenAPI: `07-api/contracts/openapi/`
- ADR-006 Anexo J: `05-architecture/decisions/records/ADR-006-instructor-anexo-j.md`
- ADR-007 Mongo reports: `05-architecture/decisions/records/ADR-007-mongodb-notification-reports.md`
- Issues: https://github.com/code-corhuila/drp-docs/issues/6 · https://github.com/code-corhuila/drp-docs/issues/7
- Prior docs PR (weeks 6–7, already approved): https://github.com/code-corhuila/drp-docs/pull/4

### Course
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
