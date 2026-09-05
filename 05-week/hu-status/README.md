<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.

     Your weekly grade is read AUTOMATICALLY from this file:
       05-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 05

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->

- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Fill remaining Corte 1 SDD templates in drp-docs (service packs and documents) after the Week 05 Docker OVA, without implementing containers in the code repo.
<!-- CONFIG-END -->

## 1. User stories worked this week

| HU ID      | Title                                                     | Status (todo/doing/done) | Evidence (PR or commit URL) |
| ---------- | --------------------------------------------------------- | ------------------------ | --------------------------- |
| HU-01      | Select and document the technology stack for the MVP      | doing                    | [drp-docs#1](https://github.com/code-corhuila/drp-docs/issues/1) · `drp-docs` ADR-004 (local until push) |
| HU-02      | Project discovery for the SpaceHub MVP                    | doing                    | [drp-docs#2](https://github.com/code-corhuila/drp-docs/issues/2) · `drp-docs` scope/models/mockup |
| HU-DOC-001 | Fill 07-api / 08-uml / 09-microservices from templates    | doing                    | [`Diagram Session 2.svg`](./Diagram%20Session%202.svg) · `drp-docs/09-microservices/services/` |
| HU-CLS-005 | Week 05 class notes (Docker image vs container)           | done                     | [`Summary_Class_1_Week_5.md`](./Summary_Class_1_Week_5.md) · [`Diagram Session 1.svg`](./Diagram%20Session%201.svg) |

## 2. My individual contribution

- Recovered Week 05 HU-status in this fork using the instructor template (this file). Course notes stay here; product documentation stays in `drp-docs`.
- Reviewed the Week 05 OVA topic (containerization with Docker: image vs container). No Dockerfile or Compose was added to the code MVP this week (Compose is Week 06; running system is `c1-release`).
- Continued HU-01 / HU-02 on `drp-docs`: stack ADR-004, discovery, data models, OpenAPI, service catalog, mockup.
- Copied `09-microservices/_template/service/` into space, reservation, payment, and notification (README, data-model, events, decisions, runbook) and filled the brackets with SpaceHub content.
- Did not start `distributed-reservation-platform`. Code waits until HU-01 and HU-02 are closed on the org board, as required.

## 3. Blockers and risks

- Local `drp-docs` `main` must still be reconciled with `origin/main` (remote `board-sync`) before push. Until then, board cards HU-01 / HU-02 stay open.
- One developer: SDD fill is limited to Corte 1 documents, not every optional ops file.
- Code MVP is empty, so per-environment HU branches and automated tests cannot be evidenced yet.

## 4. Plan for next week

- Reconcile and push `drp-docs` to `origin/main`, then close HU-01 and HU-02 on Project 19.
- Week 06 class is Docker Compose; that is when the walking skeleton Compose file belongs in the code repo.
- Keep filling this fork’s `06-week/hu-status/README.md` from the same template.

## 5. Compliance self-check

- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- `drp-docs` is committed on `main` only (instructor rule). Environment HU branches apply to the code MVP, which has not started.
- No production code or tests this week.

## 6. Evidence links

- Class session 1: [`Summary_Class_1_Week_5.md`](./Summary_Class_1_Week_5.md) · [`Diagram Session 1.svg`](./Diagram%20Session%201.svg)
- Class session 2: [`Summary_Class_2_Week_5.md`](./Summary_Class_2_Week_5.md) · [`Diagram Session 2.svg`](./Diagram%20Session%202.svg)
- `drp-docs` API: `07-api/contracts/openapi/`
- `drp-docs` UML: `08-uml/diagrams/source/`
- `drp-docs` services: `09-microservices/services/01-api-gateway` … `06-notification-service`
- Org issues: https://github.com/code-corhuila/drp-docs/issues/1 · https://github.com/code-corhuila/drp-docs/issues/2
- Org board: https://github.com/orgs/code-corhuila/projects/19
- Docs repo: https://github.com/code-corhuila/drp-docs
- Instructor docs structure (do not edit): [`../../03-week/01-session/estructura-repositorio-docs.md`](../../03-week/01-session/estructura-repositorio-docs.md)
- Course OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
