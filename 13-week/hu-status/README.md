<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       13-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 13

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 SpaceHub
- SPRINT_GOAL: threat model TB6 + security backlog
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-W13-000 | Claim TB6: `drp-payment-api` → `drp-reservation-api` `POST /internal/events` | done | https://github.com/code-corhuila/drp-docs/issues/13 |
| HU-W13-001 | Map attack surface and TB1–TB7 (mine is TB6) | done | https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/commit/2db1825 |
| HU-W13-002 | STRIDE on TB6; CP6 review response | done | https://github.com/code-corhuila/drp-docs/pull/14 |
| HU-W13-003 | Three `[security]` stories; CP7 same issues | done | https://github.com/code-corhuila/drp-docs/issues/15 · https://github.com/code-corhuila/drp-docs/issues/16 · https://github.com/code-corhuila/drp-docs/issues/17 |
| HU-W13-004 | Security CI + live 401 on `:8083` (CP8) | done | `13-week/workshop/04-security-ci/HOW-TO-RUN.md` |
| HU-W13-005 | D1 Moodle 28 Oct; D2 close 30 Oct | done | this file |

## 2. My individual contribution
- Solo SpaceHub. Confirm is E-23 on `:8083`, not `POST /api/v1/reservations/{id}/confirm` (D-C14).
- D1 (28 Oct): CP0–CP5, STRIDE PR 14, issues 13/15/16/17, Moodle `c3-w13-workshop-d1`.
- D2 (30 Oct): CP6 answered the solo review (kept S, changed T). CP7 kept the same three stories on Project 19. CP8: `go test` PASS and live `POST /internal/events` without `X-Internal-Token` is **401**. CP9: implementation plan in the Google Doc.
- I did not Approve or merge PR 14.

## 3. Blockers and risks
- LIBRARY_BOOK still needs the physical CORHUILA call number from the 28 Oct desk record. Rule 11 caps the final delivery at 2.5 without it.
- S=9 must not enter qa (shared `INTERNAL_EVENT_TOKEN`).

## 4. Plan for next week
- Issue 15 first (split the walking token) before qa. Then issue 16 (sign the body). Then issue 17 (no `userId` in logs).
- Product confirm stays on the event hop.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes: class fork uses `docs/w13-workshop`. Product PRs are not this workshop’s evidence. I did not merge my own PR.

## 6. Evidence links
- Fork workshop: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/tree/docs/w13-workshop/13-week/workshop
- STRIDE PR: https://github.com/code-corhuila/drp-docs/pull/14
- Claim issue: https://github.com/code-corhuila/drp-docs/issues/13
- Stories: https://github.com/code-corhuila/drp-docs/issues/15 · https://github.com/code-corhuila/drp-docs/issues/16 · https://github.com/code-corhuila/drp-docs/issues/17
- Reflection Doc: https://docs.google.com/document/d/1O_fyxr0Brvmoq5AAY_EEUGefafkgS16C0jww7czzBHo/edit?usp=sharing
- Review comment: written in section 2 (solo team; no teammate PR)
