# ADR-C3-01: One RDS for PostgreSQL per environment (schema isolation)

## Status

Accepted

## Context

Section 2 (Cost and Reliability) of the SpaceHub design note. `05-architecture/overview.md` still draws one PostgreSQL container per service. Anexo J and [ADR-006](https://github.com/code-corhuila/drp-docs/blob/main/05-architecture/decisions/records/ADR-006-instructor-anexo-j.md) already forbade that topology: `drp-infra-postgres` is the composition root, each `drp-*-db` only migrates its schema, and `develop` / `qa` / `main` are three volumes, not three schemas instead of environments. Moving the same decision onto Huawei Cloud without an ADR would recreate five RDS HA pairs and fail the course check **and** the bill.

## Decision

SpaceHub runs **one** RDS for PostgreSQL 16 per environment on Huawei Cloud. Schemas `identity`, `space`, `availability`, `reservation` and `payment` share that instance. Each domain’s `*_app` user writes only its schema. `drp-*-db` repositories keep Flyway (`V001__…sql`); `-api` repositories contain no DDL. `notification` stays off this RDS (MongoDB / DDS, ADR-007). Java remains only on `drp-payment-api`.

## Alternatives considered

- **Five RDS instances (database-per-container as in ADR-004 / the old C4):** rejected because Anexo J.1 / J.3.1 and ADR-006 already supersede it, and five HA pairs multiply cost without changing isolation, which is schema + user, not engine count.
- **One RDS and cross-schema JOIN for reports:** rejected because cross-schema writes are a grave fault and even a SELECT needs a GRANT plus a debt ADR. ADR-007 already gathers reports over published HTTP from `drp-notification-api`.

## Consequences

- Positive: one backup, one warm standby, one parameter group, and a bill that matches a one-developer MVP; Flyway jobs stay aligned with `06-data/migration-strategy.md`.
- Negative: a noisy query in `availability` can contend for IOPS with `reservation` on the same instance; we accept that and scale **compute** (CCE) before we split engines.
