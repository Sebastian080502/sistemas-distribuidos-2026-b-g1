<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers.

     Your weekly grade is read AUTOMATICALLY from this file:
       02-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 02

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->

- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Define the initial domain boundaries, architectural direction, Git workflow, and backlog required to prepare the first formal development Sprint.
<!-- CONFIG-END -->

## 1. User stories worked this week

| HU ID      | Title                                                                     | Status (todo/doing/done) | Evidence (PR or commit URL)         |
| ---------- | ------------------------------------------------------------------------- | ------------------------ | ----------------------------------- |
| HU-ARC-001 | Identify bounded contexts for the reservation platform                    | done                     | Pending - Week 02 recovery commit   |
| HU-ARC-002 | Define the initial context map and relationships between bounded contexts | doing                    | Pending - context map documentation |
| HU-ARC-003 | Document the initial microservices architectural decision                 | doing                    | Pending - ADR-001                   |
| HU-ARC-004 | Define testable acceptance criteria for the initial architectural backlog | doing                    | Pending - Week 02 recovery commit   |

## 2. My individual contribution

- Completed the Git and GitHub branching exercise from Week 02 Session 02.
- Practiced the branch-based development workflow using development and QA branches.
- Created and worked with user-story branches as part of the exercise.
- Practiced Pull Requests between branches.
- Practiced integrating changes between branches using cherry-pick.
- Reviewed the relationship between user-story branches, environment branches, Pull Requests, and integration.
- Applied the Git workflow concepts to the project organization that will be used during the development of the reservation platform.
- Analyzed the reservation platform domain using the Domain-Driven Design approach introduced during the session.
- Identified the initial bounded contexts based on business capabilities rather than technical layers.
- Identified Identity and Access, Space Management, Reservation Management, Payment Management, and Notification Management as initial bounded context candidates.
- Defined the initial responsibilities of each bounded context.
- Established that each future microservice must own the data required by its business capability.
- Established that services must not access another service's database directly.
- Defined the initial relationships between bounded contexts for the Context Map.
- Identified synchronous communication as appropriate for immediate queries and selected commands.
- Identified asynchronous communication as appropriate for decoupled events such as payment results, reservation state changes, and notifications.
- Defined the initial architectural direction as a microservices architecture based on business capabilities and bounded contexts.
- Identified the need to document the architectural decision through ADR-001.
- Refined the project backlog so that architectural decisions can be converted into testable user stories.
- Prepared the project structure required to begin the first formal development Sprint.

## 3. Blockers and risks

- The Week 02 activities are being recovered after the initial project definition was completed later than planned.
- The final boundaries of the microservices must be validated before implementation to avoid unnecessary service fragmentation.
- The exact communication technology and message broker have not yet been selected because the architectural analysis is still being refined.
- The complete reservation lifecycle and payment states still require validation before implementation.
- The project is being developed individually, so the number of services and features must remain controlled.
- The Git branching exercise was completed as a learning activity, but the complete workflow still needs to be applied to the real project HUs.
- The first formal development Sprint has not started yet because the current activities correspond to project initialization and architectural preparation.

## 4. Plan for next week

- Complete and publish the Context Map.
- Complete and publish ADR-001.
- Validate the responsibilities and boundaries of the initial bounded contexts.
- Refine the initial microservice candidates.
- Define the initial API and event contracts between services.
- Add acceptance criteria to the main MVP user stories.
- Configure the GitHub Projects board.
- Define the Sprint Goal and Sprint Backlog for the first formal development Sprint.
- Create the first HU development branch following the validated Git workflow.
- Begin implementation of the first microservice after the architectural boundaries have been validated.

## 5. Compliance self-check

- [x] Git branching and integration workflow practiced
- [x] Bounded contexts identified
- [x] Initial service boundaries analyzed
- [x] Data ownership principle defined
- [x] Synchronous and asynchronous communication identified
- [x] Initial architectural direction defined
- [x] Initial backlog refined
- [ ] Context Map published
- [ ] ADR-001 published
- [ ] Testable acceptance criteria completed
- [ ] GitHub Projects board configured
- [ ] First formal Sprint started

Notes on the unchecked items:

- The Context Map and ADR-001 are being completed as part of the Week 02 architectural recovery.
- Acceptance criteria will be finalized after the architectural boundaries are validated.
- GitHub Projects will be configured before the first formal development Sprint.
- The first Sprint will begin only after the initialization and architectural preparation activities are completed.

## 6. Evidence links

- Git/GitHub branching exercise: Pending - exercise repository evidence to be added.
- Product brief: [`../hu-status/prd.md`](../hu-status/prd.md).
- Context Map: Pending - to be published during Week 02.
- ADR-001: Pending - to be published during Week 02.
- Course learning material (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repository: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1

The project follows the principle of **splitting services for a reason rather than for fashion**. The initial microservice boundaries will be derived from meaningful business capabilities, bounded contexts, independent data ownership, explicit contracts, and justified scalability or deployment needs.
