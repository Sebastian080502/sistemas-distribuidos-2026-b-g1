<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers.

     The weekly grade is read AUTOMATICALLY from:
       02-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 02 (mirror)

The official graded delivery is [`README.md`](./README.md). This file is kept because the course template includes `README.es.md`; its content is English.

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->

- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Define the initial domain boundaries, architectural direction, Git workflow, and backlog required to prepare the first formal development Sprint.
<!-- CONFIG-END -->

## 1. User stories worked this week

| HU ID      | Title                                                                     | Status (todo/doing/done) | Evidence (PR or commit URL)                            |
| ---------- | ------------------------------------------------------------------------- | ------------------------ | ------------------------------------------------------ |
| HU-ADR-001 | Formalize the architectural decision for the Distributed Reservation Platform | done                     | [`adr-001-architecture.md`](./adr-001-architecture.md) |
| HU-ARC-001 | Identify bounded contexts for the reservation platform                    | done                     | [`context-map.md`](./context-map.md)                   |
| HU-ARC-002 | Define the initial context map and relationships between bounded contexts | done                     | [`context-map.md`](./context-map.md)                   |
| HU-ARC-003 | Document the initial microservices architectural decision                 | done                     | [`adr-001-architecture.md`](./adr-001-architecture.md) |
| HU-ARC-004 | Define testable acceptance criteria for the initial architectural backlog | done                     | [`user-stories.md`](./user-stories.md)                 |

## 2. My individual contribution

- Completed the Git branching exercise of Week 02 Session 02: `main` → `qa` → `develop`, HU branch, Pull Request, and cherry-pick.
- Analyzed the reservation platform domain using Domain-Driven Design.
- Identified Identity and Access, Space Management, Reservation Management, Payment Management, and Notification Management as bounded contexts.
- Established that each future microservice owns its data and must not read another service's database.
- Documented synchronous REST calls for immediate queries and asynchronous events for payments, reservation changes, and notifications.
- Formalized the microservices decision in ADR-001, including alternatives, consequences, and the immutability rule.
- Converted the architectural decision into testable user stories with Given/When/Then acceptance criteria.

## 3. Blockers and risks

- Week 02 is being recovered after the product definition was completed later than planned.
- The message broker has not been selected; that decision belongs to a later ADR.
- The project is developed by one person, so the number of services must stay controlled.
- GitHub Projects is not published yet because GitHub CLI is not available in this environment.
- The first formal development Sprint has not started.

## 4. Plan for next week

- Create the GitHub Project **Distributed Reservation Platform - Sprint** with columns Backlog, Ready, In progress, In review, and Done.
- Set WIP limits: Ready = 5, In progress = 2, In review = 3.
- Define the Sprint Goal and a small Sprint Backlog, starting with Space Management.
- Create the first HU branch (`hu-xxx-dev`) and open a Pull Request to `develop`.
- Begin the first microservice after the architectural boundaries remain stable.

## 5. Compliance self-check

- [x] Git branching and integration workflow practiced
- [x] Bounded contexts identified
- [x] Initial service boundaries analyzed
- [x] Data ownership principle defined
- [x] Synchronous and asynchronous communication identified
- [x] Initial architectural direction defined
- [x] Initial backlog refined
- [x] Context Map published
- [x] ADR-001 published
- [x] Testable acceptance criteria completed
- [ ] GitHub Projects board configured
- [ ] First formal Sprint started

Notes on the unchecked items:

- GitHub Projects remains pending: the intended board name is **Distributed Reservation Platform - Sprint**. No board URL is published yet.
- The first Sprint starts after the board exists and the Git exercise screenshots are attached.

## 6. Evidence links

- Session summaries: [`resumen-sesiones/resumen_sistemas_distribuidos_semana_2.svg`](./resumen-sesiones/resumen_sistemas_distribuidos_semana_2.svg)
- Git branching exercise: [`ejercicio-ramas-git/README.md`](./ejercicio-ramas-git/README.md)
- Product brief: [`../../01-week/hu-status/prd.md`](../../01-week/hu-status/prd.md)
- Context Map: [`context-map.md`](./context-map.md)
- ADR-001: [`adr-001-architecture.md`](./adr-001-architecture.md)
- User stories and acceptance criteria: [`user-stories.md`](./user-stories.md)
- GitHub Project: pending — board not created yet
- Course learning material (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repository: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1

The project follows the principle of **splitting services for a reason rather than for fashion**. The initial microservice boundaries are derived from business capabilities, bounded contexts, independent data ownership, explicit contracts, and justified scalability or deployment needs.
