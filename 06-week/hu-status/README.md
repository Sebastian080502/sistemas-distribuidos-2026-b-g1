<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       06-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 06

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Complete the Week 06 OVA activities (Docker Compose orchestration and environment/12-factor planning) for SpaceHub, with session diagrams, without creating GitHub code repositories.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-XXX-001 | Complete the Session 1 activity on Docker Compose orchestration | done | [`06-session-1-docker-compose-orchestration.md`](./06-session-1-docker-compose-orchestration.md) |
| HU-XXX-002 | Complete the Session 2 activity on environments, configuration, and orchestration | done | [`06-session-2-environments-config-orchestration.md`](./06-session-2-environments-config-orchestration.md) |
| HU-XXX-003 | Create the Session 1 Docker Compose orchestration diagram | done | [`Diagram Session 1.png`](./Diagram%20Session%201.png) |
| HU-XXX-004 | Create the Session 2 environments, configuration, and orchestration diagram | done | [`Diagram Session 2.png`](./Diagram%20Session%202.png) |
| HU-01 | Select and document the technology stack for the MVP | doing | [drp-docs#1](https://github.com/code-corhuila/drp-docs/issues/1) |
| HU-02 | Project discovery for the SpaceHub MVP | doing | [drp-docs#2](https://github.com/code-corhuila/drp-docs/issues/2) |
| HU-C2-001 | Walking skeleton Compose (instructor repos) | todo | Blocked — professor creates per-service repos; Compose path is Q-004 |

## 2. My individual contribution
- Completed the **Session 1 activity** focused on Docker Compose orchestration (shared network, health checks, named volumes, service-name DNS) applied to SpaceHub.
- Completed the **Session 2 activity** focused on environments, 12-factor configuration, and orchestration planning (`develop` / `qa` / `main`).
- Created Session 1 and Session 2 diagrams of **SpaceHub Compose topology** (`spacehub-net`, gateway, identity, space, reservation, payment, notification).
- Documented the same plan in `drp-docs` (`10-devops/`, `15-project-control/technical-backlog.md`). No application GitHub repository was created; the instructor creates those.

## 3. Blockers and risks
- Per-service code repositories do not exist yet. Runtime `docker compose up` cannot be evidenced.
- The instructor has not said where Compose lives (infra vs other). Tracked as Q-004 in `drp-docs`.
- `drp-docs` stays on `main` only. Environment HU branches apply later on the code repos the professor creates.

## 4. Plan for next week
- Complete Week 07 HU-status (REST, gRPC, messaging, versioned contracts) in this fork.
- Keep Compose implementation blocked until the instructor creates the service repositories.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- Environment HU branches and tests require code repos created by the instructor.

## 6. Evidence links

### Session activities
- [`06-session-1-docker-compose-orchestration.md`](./06-session-1-docker-compose-orchestration.md)
- [`06-session-2-environments-config-orchestration.md`](./06-session-2-environments-config-orchestration.md)
- [`Summary_Class_1_Week_6.md`](./Summary_Class_1_Week_6.md)
- [`Summary_Class_2_Week_6.md`](./Summary_Class_2_Week_6.md)

### Session diagrams
- [`Diagram Session 1.png`](./Diagram%20Session%201.png) — SpaceHub Compose (`spacehub-net`)
- [`Diagram Session 2.png`](./Diagram%20Session%202.png) — SpaceHub environments / 12-factor

### Product documentation (`drp-docs`)
- https://github.com/code-corhuila/drp-docs
- Org board: https://github.com/orgs/code-corhuila/projects/19
- Org issues: https://github.com/code-corhuila/drp-docs/issues/1 · https://github.com/code-corhuila/drp-docs/issues/2
- Course OVA: https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
