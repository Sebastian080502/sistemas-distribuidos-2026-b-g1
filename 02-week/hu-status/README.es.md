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

## Project Board

| Board Name | URL |
| ---------- | --- |
| Distributed Reservation Platform — Backlog | https://github.com/orgs/code-corhuila/projects/19 |

## 1. User stories worked this week

| HU ID      | Title                                                                     | Status (todo/doing/done) | Evidence (PR or commit URL)                            |
| ---------- | ------------------------------------------------------------------------- | ------------------------ | ------------------------------------------------------ |
| HU-ADR-001 | Formalize the architectural decision for the Distributed Reservation Platform | done                     | [`docs/adr-001-architecture.md`](./docs/adr-001-architecture.md) |
| HU-ARC-001 | Identify bounded contexts for the reservation platform                    | done                     | [`docs/context-map.md`](./docs/context-map.md)                   |
| HU-ARC-002 | Define the initial context map and relationships between bounded contexts | done                     | [`docs/context-map.md`](./docs/context-map.md)                   |
| HU-ARC-003 | Document the initial microservices architectural decision                 | done                     | [`docs/adr-001-architecture.md`](./docs/adr-001-architecture.md) |
| HU-ARC-004 | Define testable acceptance criteria for the initial architectural backlog | done                     | [`docs/user-stories.md`](./docs/user-stories.md)                 |

## 2. My individual contribution

- Completed the Git branching exercise in https://github.com/Sebastian080502/prueba_sistema_distribuidos: `main` / `qa` / `develop`, HU branches, Pull Requests #1–#4, and cherry-pick of `d4c1a2d` onto `qa` and `main`.
- Analyzed the reservation platform domain using Domain-Driven Design.
- Identified Identity and Access, Space Management, Reservation Management, Payment Management, and Notification Management as bounded contexts.
- Established that each future microservice owns its data and must not read another service's database.
- Documented synchronous REST calls for immediate queries and asynchronous events for payments, reservation changes, and notifications.
- Formalized the microservices decision in ADR-001, including alternatives, consequences, and the immutability rule.
- Converted the architectural decision into testable user stories with Given/When/Then acceptance criteria.
- Created the GitHub Project **Distributed Reservation Platform - Sprint** and linked it to this repository.

## 3. Blockers and risks

- Week 02 is being recovered after the product definition was completed later than planned.
- The message broker has not been selected; that decision belongs to a later ADR.
- The project is developed by one person, so the number of services must stay controlled.
- The first formal development Sprint has not started.

## 4. Plan for next week

- Group the GitHub Project board by the Board field and apply WIP limits: Ready = 5, In progress = 2, In review = 3.
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
- [x] GitHub Projects board configured
- [ ] First formal Sprint started

Notes on the unchecked items:

- The first Sprint starts after the Sprint Goal and a small Sprint Backlog are selected on the board.

## 6. Evidence links

- Session summaries: [`session-summaries/week-02-session-summary.svg`](./session-summaries/week-02-session-summary.svg)
- Git branching exercise: [`git-branching-exercise/README.md`](./git-branching-exercise/README.md)
- Practice repository: https://github.com/Sebastian080502/prueba_sistema_distribuidos
- Product brief: [`../../01-week/hu-status/prd.md`](../../01-week/hu-status/prd.md)
- Context Map: [`docs/context-map.md`](./docs/context-map.md)
- ADR-001: [`docs/adr-001-architecture.md`](./docs/adr-001-architecture.md)
- User stories and acceptance criteria: [`docs/user-stories.md`](./docs/user-stories.md)
- GitHub Project: https://github.com/orgs/code-corhuila/projects/19
- Course learning material (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repository: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1

The project follows the principle of **splitting services for a reason rather than for fashion**. The initial microservice boundaries are derived from business capabilities, bounded contexts, independent data ownership, explicit contracts, and justified scalability or deployment needs.
