# Unit 2 · Planning · Corte 2

# Planning — Environments, Configuration Strategy and Orchestration — SpaceHub / DRP

## 1. Session overview

This planning session defines how SpaceHub runs across `develop`, `qa`, and `prod`, how configuration and secrets are managed, how branches map to environments, and how orchestration work is sliced for Corte 2.

Docs (`drp-docs`) stay on `main` only. The branch model below applies later to the **code repositories the instructor creates** (expected: one repo per microservice).

## 2. Learning objectives

1. Define environments and what differs between them.
2. Design a configuration and secrets strategy following 12-factor.
3. Map the course git-flow to each environment.
4. Slice orchestration work into Corte 2 backlog items.

## 3. Environments

The same application image should run in every environment while only the configuration changes.

| Environment | Git | Purpose |
|-------------|-----|---------|
| local | developer machine | `docker compose up` |
| develop | `develop` | Fast integration, synthetic data |
| qa | `qa` | Production-like tests and demo |
| prod | `main` + tag | Course evaluation / live increment |

Promotion means running the **tested artifact** with the target environment's configuration, not rebuilding for production.

## 4. Configuration and secrets

> Configuration lives in the environment, not in the code.

- Do not hard-code `localhost` or environment names in Go/Angular.
- Supply `DATABASE_URL`, JWT keys, and broker URLs through env.
- Never commit `.env`. Commit `.env.example` without real secrets.
- Fail fast if a required variable is missing.

SpaceHub variables (local names): see `drp-docs/10-devops/environments.md`.

## 5. Branch model (code repo only)

```text
hu-xxx-dev  -> develop
hu-xxx-qa   -> qa
hu-xxx-main -> main
```

Example: walking-skeleton Compose lands on `hu-c2-001-dev` and PRs to `develop`.

## 6. Corte 2 orchestration stories (blocked until instructor repos)

### Must — Start all services with one Compose command (TD-005)

Acceptance:

- Health checks on each PostgreSQL.
- Named volumes per database.
- Service-name DNS on `spacehub-net`.

### Must — Externalize configuration (TD-006)

Acceptance:

- `.env.example` lists required variables.
- No secrets in Git.
- Missing vars fail startup with the variable name.

### Must — Promote the same artifact to QA

Acceptance:

- Image digest unchanged.
- QA configuration injected.

### Must — Map branch to environment

Acceptance:

- Child PRs target the correct parent branch.

### Should — Environment smoke tests

Acceptance:

- `GET /health` on the gateway in QA.
- One business path (list spaces) in QA.

### Could — Deployment metrics

Acceptance:

- Lead time and rollback time recorded later (Week 08+).

## 7. Configuration drift scenario

If SpaceHub works in `develop` but fails in `qa`:

- Hard-coded `localhost:5432` inside a container.
- Mixed names (`DATABASE_URL` vs `DB_URL`).
- Missing `POSTGRES_PASSWORD` in QA secrets.

Prevention: one naming convention, validate at startup, keep `.env.example`.

## 8. Common mistakes

- Rebuilding a different image per environment.
- Committing secrets or baking them into images.
- One database for all services.
- Skipping QA and merging HU branches to `main`.

## 9. Self-check

### Question 1

**Across environments you should:**  
**Answer:** Run the same built image and change only configuration.

### Question 2

**The 12-factor rule says:**  
**Answer:** Configuration lives in the environment, not in the code.

### Question 3

**A change reaches production:**  
**Answer:** After `hu-xxx-qa` → `qa`, then `hu-xxx-main` → `main`.

### Question 4

**To prevent configuration drift:**  
**Answer:** Consistent variable names, fail-fast validation, `.env.example`.

### Question 5

**Secrets belong:**  
**Answer:** In environment injection or a secret store, never in Git.

### Question 6

**"Works in develop, breaks in QA" can be caused by:**  
**Answer:** Configuration drift.

## 10. Expected outcome

A written strategy for environments, secrets, branch promotion, and the Compose topology. The executable Compose file waits until the instructor creates the service repos and answers Q-004 (`drp-docs/15-project-control/open-questions.md`).
