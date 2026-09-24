<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       07-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 07

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Complete the Week 07 OVA activities (inter-service communication and versioned contracts) for SpaceHub and keep drp-docs aligned. No student-created GitHub repositories.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-XXX-001 | Complete the Week 07 Session 1 communication activity | done | [`07-session-1-inter-service-communication.md`](./07-session-1-inter-service-communication.md) |
| HU-XXX-002 | Complete the Week 07 Session 2 contract testing activity | done | [`07-session-2-versioned-contracts-contract-testing.md`](./07-session-2-versioned-contracts-contract-testing.md) |
| HU-XXX-003 | Create the Session 1 diagram on inter-service communication | done | [`Diagram Session 1.png`](./Diagram%20Session%201.png) |
| HU-XXX-004 | Create the Session 2 diagram on versioned contracts and contract testing | done | [`Diagram Session 2.png`](./Diagram%20Session%202.png) |
| HU-COM-001 | SpaceHub sync/async matrix in drp-docs | done | `drp-docs/09-microservices/communication-patterns.md` |
| HU-CTR-001 | Keep OpenAPI `/api/v1` as contract source of truth | doing | `drp-docs/07-api/contracts/openapi/` · `07-api/guidelines.md` |
| HU-01 | Select and document the technology stack for the MVP | doing | [drp-docs#1](https://github.com/code-corhuila/drp-docs/issues/1) |
| HU-02 | Project discovery for the SpaceHub MVP | doing | [drp-docs#2](https://github.com/code-corhuila/drp-docs/issues/2) |
| HU-C2-002 | Consumer-driven contract test in CI | todo | Blocked until the instructor creates service repos |

## 2. My individual contribution
- Completed the Session 1 activity focused on **inter-service communication** (REST default, gRPC not MVP, events for payment/notification, at-least-once + idempotency).
- Completed the Session 2 activity focused on **versioned contracts and contract testing** (OpenAPI `/api/v1`, breaking change → `/api/v2`, Pact/CI later).
- Created Session 1 and Session 2 diagrams of **SpaceHub communication and contracts** (REST to identity/space/reservation/payment; events to notification; OpenAPI in `drp-docs`).
- Updated `drp-docs` communication matrix and Corte 2 backlog. Implementation waits for instructor-created per-service repositories.

## 3. Blockers and risks
- Without running services, Pact/CI contract tests cannot be evidenced (HU-C2-002 `todo`).
- Dual-write (database + broker) remains a risk until an outbox exists in code.
- A long synchronous chain to notification would cascade; notification stays asynchronous.

## 4. Plan for next week
- Week 08 OVA: Agile and DevOps for distributed teams.
- After the instructor opens the service repos, use git-flow there (`hu-xxx-dev` -> develop). Compose location remains Q-004.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- Code MVP, HU branches, and automated contract tests require repositories created by the instructor.

## 6. Evidence links

### Session activities
- [`07-session-1-inter-service-communication.md`](./07-session-1-inter-service-communication.md)
- [`07-session-2-versioned-contracts-contract-testing.md`](./07-session-2-versioned-contracts-contract-testing.md)
- [`Summary_Class_1_Week_7.md`](./Summary_Class_1_Week_7.md)
- [`Summary_Class_2_Week_7.md`](./Summary_Class_2_Week_7.md)

### Session diagrams
- [`Diagram Session 1.png`](./Diagram%20Session%201.png) — SpaceHub REST vs events
- [`Diagram Session 2.png`](./Diagram%20Session%202.png) — SpaceHub OpenAPI / Pact in `drp-docs`

### Product documentation (`drp-docs`)
- https://github.com/code-corhuila/drp-docs
- Communication matrix: `09-microservices/communication-patterns.md`
- OpenAPI: `07-api/contracts/openapi/`
- Catch-up backlog: `15-project-control/technical-backlog.md`
- Org board: https://github.com/orgs/code-corhuila/projects/19
- Org issues: https://github.com/code-corhuila/drp-docs/issues/1 · https://github.com/code-corhuila/drp-docs/issues/2
- Course OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
