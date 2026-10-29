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
| HU-W13-000 | Claim TB6: `drp-payment-api` → `drp-reservation-api` `POST /internal/events` | doing | https://github.com/code-corhuila/drp-docs/issues/13 |
| HU-W13-001 | Map attack surface and TB1–TB7 (mine is TB6) | doing | https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/commit/2db1825 |
| HU-W13-002 | STRIDE on TB6 (S,T,R then I,D,E) in `drp-docs` | doing | https://github.com/code-corhuila/drp-docs/pull/14 |
| HU-W13-003 | Three `[security]` backlog issues (S=9, T=6, I=6) | doing | https://github.com/code-corhuila/drp-docs/issues/15 · https://github.com/code-corhuila/drp-docs/issues/16 · https://github.com/code-corhuila/drp-docs/issues/17 |
| HU-W13-004 | Security CI proposal + 401 test for `/internal/events` | doing | https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/commit/f12260b |
| HU-W13-005 | Close Week 13 progress (this file + Moodle d1) | doing | https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/blob/docs/w13-workshop/13-week/hu-status/README.md |

## 2. My individual contribution
- Solo SpaceHub. I claimed the confirm hop: `PaymentConfirmed` from `drp-payment-api` :8084 to `drp-reservation-api` :8083. D-C14 / E-23: no portal confirm; E-16 starts the payment.
- CP0–CP5 live in `13-week/workshop/` and `13-week/hu-status/` on `docs/w13-workshop`. STRIDE lives in `code-corhuila/drp-docs` branch `security/stride-Sebastian080502`.
- **Peer review (solo team):** I compared my table with the class example (a public gateway hop vs a domain service). That example treats “skip the gateway” as spoofing. TB6 is already off `drp-api-gateway`, so the same idea here is any docker-network process with the shared `X-Internal-Token` (`spacehub-dev-internal` in `.env.example`). Row E already covers a USER JWT on `/internal/events`. Gap vs the example’s “do not trust the body price”: reservation does not recompute `amountCents` from E-16; a tampered body can still confirm the wrong cents if the token is valid (row T). I did not Approve or merge any PR.

## 3. Blockers and risks
- LIBRARY_BOOK is still a placeholder until the CORHUILA desk call number exists. Without the physical record, CP2/CP3 can be 0.0.
- S=9 is not acceptable into qa: shared `INTERNAL_EVENT_TOKEN`.
- Walking outbox is in-memory. `drp-infra-mongo` is not requested from this workshop.

## 4. Plan for next week
- CP6–CP9 (take-home): answer review, refine issues, run the 401 test on real `:8083`, write the implementation plan in the Google Doc.
- Product: keep confirm on the event hop; do not add `POST /api/v1/reservations/{id}/confirm`.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

Notes: class fork uses `docs/w13-workshop`, not `hu-c3-*-dev`. The Go test is a pipeline proposal in the fork, not a change to `drp-reservation-api` today. Per-environment product PRs are not this week’s evidence.

## 6. Evidence links
- Fork workshop: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1/tree/docs/w13-workshop/13-week/workshop
- STRIDE PR: https://github.com/code-corhuila/drp-docs/pull/14
- Claim issue: https://github.com/code-corhuila/drp-docs/issues/13
- Stories: https://github.com/code-corhuila/drp-docs/issues/15 · https://github.com/code-corhuila/drp-docs/issues/16 · https://github.com/code-corhuila/drp-docs/issues/17
- Reflection Doc: https://docs.google.com/document/d/1O_fyxr0Brvmoq5AAY_EEUGefafkgS16C0jww7czzBHo/edit?usp=sharing
- Review comment: written in section 2 (solo team; no teammate PR)
