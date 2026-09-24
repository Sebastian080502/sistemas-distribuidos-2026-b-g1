# Unit 2 · Weekly · Corte 2

# Docker Compose and Orchestration Basics — SpaceHub / DRP

## 1. Session overview

This session covers Docker Compose for a distributed system: shared network, service-name DNS, health checks, `depends_on`, named volumes, and environment-based configuration.

The executable Compose file is **not** in this course fork and **not** invented in a student-created GitHub repo. It will live where the **instructor** assigns it (likely an infra repo next to per-service repos). Until then, this file is the design.

## 2. Learning objectives

1. Compose a multi-service system with a shared network.
2. Order startup with health checks and `depends_on`.
3. Externalize configuration per environment.
4. Understand when Compose is enough and when an orchestrator is needed.

## 3. From a container to a system

Week 05 was image vs container. Week 06 is the **system**: Angular, API gateway, five Go services, five PostgreSQL instances, and a local AMQP broker on one private bridge network (`spacehub-net`).

## 4. Annotated Compose example (SpaceHub)

This is the **design** recorded in `drp-docs/10-devops/local-setup.md`. Runtime Compose waits for the instructor.

```yaml
# DESIGN ONLY — not committed as a student-created repo.
# Instructor will create per-service repositories and assign where this file lives (Q-004).
# Do not commit secrets. Copy .env.example -> .env locally.

services:
  postgres-space:
    image: postgres:16
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: space
    volumes:
      - space_pg_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U $$POSTGRES_USER -d space"]
      interval: 5s
    networks: [spacehub-net]

  space-service:
    build: ./services/space-service
    environment:
      DATABASE_URL: postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres-space:5432/space
      PORT: "8082"
    depends_on:
      postgres-space:
        condition: service_healthy
    networks: [spacehub-net]

  api-gateway:
    build: ./services/api-gateway
    environment:
      SPACE_URL: http://space-service:8082
      IDENTITY_URL: http://identity-service:8081
    ports:
      - "8080:8080"
    depends_on:
      space-service:
        condition: service_started
    networks: [spacehub-net]

volumes:
  space_pg_data: {}

networks:
  spacehub-net:
    driver: bridge
```

Key ideas:

- `services` lists the system components.
- Internal URLs use **service names**, never `localhost` or fixed IPs (`http://space-service:8082`).
- `depends_on` + `condition: service_healthy` waits for PostgreSQL readiness.
- Named volumes keep catalog data across container recreate.
- Host port 8080 is only for the gateway. Domain databases stay on the private network (optional host ports 5432–5436 for debugging).

## 5. Startup order and readiness

`depends_on` without a condition only starts containers in order. SpaceHub Go services must retry PostgreSQL on boot. Health checks belong on each `postgres-*` service.

## 6. Persistent data

Each bounded context has its own volume (`identity_pg_data`, `space_pg_data`, `reservation_pg_data`, `payment_pg_data`, `notification_pg_data`). Application containers stay disposable.

## 7. When Compose is not enough

Compose is enough for local Corte 2 and a single-host demo. Kubernetes is out of this course MVP (`drp-docs/01-context/scope.md`).

## 8. Practical scenario

If `docker compose up` fails because `space-service` connects before PostgreSQL is ready: add `pg_isready` healthcheck, `condition: service_healthy`, and backoff in the Go adapter.

## 9. Common mistakes

- Relying on `depends_on` alone.
- Using `localhost` from inside a container to reach another service.
- Committing `.env` secrets.
- One shared PostgreSQL for all contexts (violates ADR-005).
- Storing database files only in the container writable layer.

## 10. Self-check

### Question 1

**In Compose, reservation-service reaches space-service by:**  
**Answer:** The service name on `spacehub-net`, e.g. `http://space-service:8082`.

### Question 2

**`depends_on` without a condition guarantees:**  
**Answer:** Start order only, not database readiness.

### Question 3

**Space catalog data should live in:**  
**Answer:** Named volume `space_pg_data`.

### Question 4

**Per-environment differences are handled by:**  
**Answer:** The same image plus environment values / Compose overrides.

### Question 5

**You move from Compose to an orchestrator when you need:**  
**Answer:** Multiple hosts, self-healing, rolling updates, autoscaling.

### Question 6

**`docker compose up` fails in CI because the database is not ready. Fix:**  
**Answer:** Healthcheck + `condition: service_healthy` + boot retry.

## 11. This week's activity

Bring SpaceHub up with a single `docker compose up` **when the instructor assigns the Compose file**: shared network, health checks, env-based configuration, and persistent volumes. Until then, the plan is recorded in `drp-docs` (TD-001, TD-005, Q-004).
