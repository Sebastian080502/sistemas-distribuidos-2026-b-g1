<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       03-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 03

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Apply the course documentation structure to the SpaceHub reservation platform and start Corte 1 HUs on the official docs repo.
<!-- CONFIG-END -->

## Project Board

| Board Name | URL |
| ---------- | --- |
| Distributed Reservation Platform — Backlog | https://github.com/orgs/code-corhuila/projects/19 |

## 1. User stories worked this week

| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-01 | Select and document the technology stack for the MVP | doing | [drp-docs#1](https://github.com/code-corhuila/drp-docs/issues/1) — ADR-004 drafted locally in `drp-docs` (not pushed yet) |
| HU-02 | Project discovery for the SpaceHub MVP | doing | [drp-docs#2](https://github.com/code-corhuila/drp-docs/issues/2) — context, scope, and domain drafted locally in `drp-docs` (not pushed yet) |
| HU-DOC-001 | Apply the SDD `docs/` folder structure (`00`–`15`) to DRP | doing | Session notes: [`session-summary.md`](./session-summary.md) · Official docs repo: https://github.com/code-corhuila/drp-docs |

## 2. My individual contribution

- Reviewed the official documentation scaffold (`00-governance` through `15-project-control` plus `99-archive`) and the rule that `drp-docs` updates go on `main`.
- Confirmed the product repos: `code-corhuila/drp-docs` (documentation) and `code-corhuila/distributed-reservation-platform` (code MVP, still empty).
- Continued SpaceHub discovery in the local `drp-docs` copy: problem, actors, in/out of MVP, and five bounded contexts (identity, space, reservation, payment, notification).
- Drafted the stack decision **Angular + Go + PostgreSQL** (ADR-004) as the HU-01 deliverable. It is not closed until the docs are pushed to `main`.
- Did not start implementation in the code MVP repository. Week 03 work stays on documentation.

## 3. Blockers and risks

- `drp-docs` local `main` is ahead of and behind `origin/main` (missing the remote `board-sync` workflow). A push without reconciling would drop the Project sync.
- HU-01 and HU-02 remain Open on the org board until the local docs are uploaded.
- The course fork HU-status for Week 03 was late; this file recovers that delivery.
- One developer: the docs fill must stay limited to Corte 1 (stack + discovery), not the full SDD.

## 4. Plan for next week

- Reconcile and push `drp-docs` to `origin/main`.
- Finish the HU-01 and HU-02 acceptance criteria and move both cards on the org board.
- Keep filling Week 04 HU-status in this fork.
- Leave `distributed-reservation-platform` empty until the docs HUs are accepted.

## 5. Compliance self-check

- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes on the unchecked items:

- Docs work uses `main` only, as required for `drp-docs`. Environment HU branches belong to the code MVP, which has not started.
- No production code or automated tests were added this week.

## 6. Evidence links

- Session summary: [`session-summary.md`](./session-summary.md)
- Instructor structure (do not edit): [`../01-session/estructura-repositorio-docs.md`](../01-session/estructura-repositorio-docs.md)
- Org issues: https://github.com/code-corhuila/drp-docs/issues/1 · https://github.com/code-corhuila/drp-docs/issues/2
- Docs repo: https://github.com/code-corhuila/drp-docs
- Code MVP (not used this week): https://github.com/code-corhuila/distributed-reservation-platform
- Course fork: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
