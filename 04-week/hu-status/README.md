<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       04-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 04

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Deepen the SDD fill for SpaceHub (data models, contracts, mockup) without starting the code MVP.
<!-- CONFIG-END -->

## Project Board

| Board Name | URL |
| ---------- | --- |
| Distributed Reservation Platform — Backlog (org) | https://github.com/orgs/code-corhuila/projects/19 |
| Distributed Reservation Platform - Sprint (personal) | https://github.com/users/Sebastian080502/projects/2 |

## 1. User stories worked this week

| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-01 | Select and document the technology stack for the MVP | doing | [drp-docs#1](https://github.com/code-corhuila/drp-docs/issues/1) — ADR-004 and C4 overview exist locally |
| HU-02 | Project discovery for the SpaceHub MVP | doing | [drp-docs#2](https://github.com/code-corhuila/drp-docs/issues/2) — models, user stories, and mockup exist locally |
| HU-DATA-001 | Document one data model per microservice | doing | Local `drp-docs/06-data/models.md` (identity, space, reservation, payment, notification) |

## 2. My individual contribution

- Continued the local SDD fill for SpaceHub: data models per service, OpenAPI drafts, service catalog, and the iterative mockup.
- Kept the stack consistent: Angular frontend, Go services, one PostgreSQL per bounded context.
- Did not push `drp-docs` yet because local `main` diverged from `origin/main` (board-sync workflow on remote).
- Did not add code to `distributed-reservation-platform`.

## 3. Blockers and risks

- Week 04 HU-status is recovered late. The work exists in the local docs repo, not yet on GitHub `drp-docs` `main`.
- Scaffold leftovers (`auth-service` vs `identity-service`) must be cleaned before the docs push.
- Until HU-01 and HU-02 are closed, the code MVP should stay empty.

## 4. Plan for next week

- Reconcile `drp-docs` with `origin/main` and push the SpaceHub documentation.
- Check off and close HU-01 and HU-02 on the org board.
- Fill Week 05 HU-status in this fork as the next weekly delivery.

## 5. Compliance self-check

- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- No code MVP branch or tests this week. Documentation remains on `drp-docs` `main` after the pending push.

## 6. Evidence links

- Org issues: https://github.com/code-corhuila/drp-docs/issues/1 · https://github.com/code-corhuila/drp-docs/issues/2
- Org board: https://github.com/orgs/code-corhuila/projects/19
- Docs repo: https://github.com/code-corhuila/drp-docs
- Week 03 notes: [`../03-week/hu-status/session-summary.md`](../03-week/hu-status/session-summary.md)
- Course fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
