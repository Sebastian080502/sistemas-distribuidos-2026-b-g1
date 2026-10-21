# Week 11 workshop — observability on SpaceHub

Individual work for Corte 3, week 11. The files in this folder follow one SpaceHub
journey from `drp-docs`. They are not a generic sample and they are not ParkFlow.

<!-- CONFIG-START -->
- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 — Distributed Reservation Platform (SpaceHub)
- DOCS_REPO: https://github.com/code-corhuila/drp-docs
- REFLECTION_DOC: (Google Doc URL — share with jesus.gonzalez@corhuila.edu.co as Editor)
- JOURNEY: (written in CP1)
<!-- CONFIG-END -->

## Scope I own

I work on `drp-identity-api` and on the class-fork evidence. For this workshop I
instrument the reservation confirm path because that is the journey that already
crosses `drp-api-gateway`, `drp-reservation-api`, `drp-payment-api` and
`drp-notification-api` in `drp-docs` (`03-product`, `05-architecture`, `07-api`).

## Folder

| File | Checkpoint |
| --- | --- |
| `README.md` | CP0 |
| `01-journey.md` | CP1 |
| `02-log-contract.md` | CP2 |
| `03-instrumentation/` | CP3 |
| `04-metrics.md` | CP4 |
| `05-slo.md` | CP6 |
| `06-alerts.yml` / `06-alert-policy.md` | CP7 |
| `07-runbook.md` | CP8 |
| `08-incident.md` | CP9 |

Reflection lives in Google Docs, not in this repository.
