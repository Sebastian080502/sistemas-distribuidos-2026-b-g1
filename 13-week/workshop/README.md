# Week 13 workshop — threat model on SpaceHub

Individual work for Corte 3, week 13. Names come from `drp-docs`.

<!-- CONFIG-START -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 — Distributed Reservation Platform (SpaceHub)
- DOCS_REPO: https://github.com/code-corhuila/drp-docs
- BOUNDARY: drp-payment-api → drp-reservation-api (`POST /internal/events`)
- CLAIM_ISSUE: https://github.com/code-corhuila/drp-docs/issues/13
- REFLECTION_DOC: https://docs.google.com/document/d/1O_fyxr0Brvmoq5AAY_EEUGefafkgS16C0jww7czzBHo/edit?usp=sharing
- LIBRARY_BOOK: (author, title, edition, year, call number from CORHUILA library)
- LIBRARY_RECORD_TIME: (HH:MM on 2026-10-28, library desk)
- STORIES:
<!-- CONFIG-END -->

## Scope I own

I am the only SpaceHub member. For this workshop I own the walking hop that
carries `PaymentConfirmed` from `drp-payment-api` to `drp-reservation-api`.
The public portal never calls `POST /api/v1/reservations/{id}/confirm` (D-C14).
Confirm is E-23 on `:8083`, not on `drp-api-gateway`.

## Folder

| File | Checkpoint |
| --- | --- |
| `README.md` | CP0 |
| `01-attack-surface.md` | CP1 |
| `04-security-ci/` | CP4 |

STRIDE lives in `drp-docs` on `security/stride-Sebastian080502`.
Reflection lives in Google Docs, not in this repository.
