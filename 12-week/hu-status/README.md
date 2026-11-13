<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       12-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 12

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: After the Corte 2 v2.0.0 freeze, keep walking HTTP/outbox on develop only. No cloud workshop: instructor 12-week sessions are still .gitkeep.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-RES-HTTP-001 | Walking HTTP E-11..E-14 (Corte 3, develop only) | doing | Carried from Week 11: [reservation-api#4](https://github.com/code-corhuila/drp-reservation-api/pull/4) MERGED `develop` 2026-10-11. Not in `v2.0.0`. No new Week 12 promotion. |
| HU-OUTBOX-001 | `reservation.outbox` table (not a Camunda saga) | doing | [reservation-db#4](https://github.com/code-corhuila/drp-reservation-db/pull/4) MERGED `develop` 2026-10-11. No publisher/broker (Q-001). |
| HU-ID-001 | E-01 422 login envelope on develop | doing | [identity-api#5](https://github.com/code-corhuila/drp-identity-api/pull/5) MERGED `develop`. Not promoted to `qa` this week. |
| HU-C2-002 | Consumer-driven contract tests in CI | todo | Not this week. |
| HU-PO-001 | Ask PO who owns `amountCents` | doing | [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) still OPEN — do not invent an owner. |

## 2. My individual contribution
- **Written late on 2026-11-12** so Ariel does not see a silent zero on Week 12. This is honest catch-up, not a backdated workshop.
- Academic Week 12 product work was the **post-C2 freeze**: walking HTTP and outbox already opened on `develop` in Week 11. I did **not** invent a Week 12 cloud/lab workshop. Upstream `12-week/01-session` and `12-week/02-session` are still `.gitkeep` only.
- Did **not** retag product `v2.0.0`. Confirmation stays event-driven (D-C14). No AMQP (Q-001 unanswered). No Camunda.
- Did **not** copy product SDD into `12-week/`. This class fork is weekly evidence only.

## 3. Blockers and risks
- No Week 12 session files from the instructor. Filling a fake `02-session` would be a grave fault.
- Walking HTTP / outbox stayed on `develop` during that week; promoting to `qa`/`main` was later Corte 3 work (see Weeks 13–14), not invented here as a Week 12 merge.
- Outbox has no broker. `drp-workflow` empty/untagged.
- [drp-docs#7](https://github.com/code-corhuila/drp-docs/issues/7) (`amountCents`) still OPEN.

## 4. Plan for next week
- Week 13 as actually executed: one product promotion (identity 422 `develop` → `qa`) with a real review request; do not self-merge.
- Keep 01-session / 02-session empty until the instructor publishes them.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked item:

- No **new** product test PR in academic Week 12. Walking HTTP tests already lived on `develop` from Week 11 (`reservation-api#4`).
- This file is Week 12 **class** evidence, filled 2026-11-12. Product URLs are on `code-corhuila/drp-*`.

## 6. Evidence links

### Product (`code-corhuila`) — carried from Week 11, still the Week 12 develop line
- https://github.com/code-corhuila/drp-reservation-api/pull/4 (MERGED `develop`)
- https://github.com/code-corhuila/drp-reservation-db/pull/4 (MERGED `develop`)
- https://github.com/code-corhuila/drp-identity-api/pull/5 (MERGED `develop`)
- https://github.com/code-corhuila/drp-docs/issues/7 (OPEN)

### Product — Corte 2 freeze (Week 11, not re-tagged in Week 12)
- https://github.com/code-corhuila/drp-docs/releases/tag/v2.0.0

### Course
- Profile CONFIG: https://github.com/Sebastian080502/Sebastian080502
- This class fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
- OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Do not retag product `v2.0.0`. This fork is not tagged `v2.0.0`.
