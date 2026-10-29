# How to run the TB6 negative test

Covers story `[security] Reject confirm hops without X-Internal-Token`
(acceptance: `POST /internal/events` without the header is 401, even with a user JWT).

This folder is a **proposal** for CI. It is not a change to `drp-reservation-api` today.
The handler copy matches the walking consume on `:8083` (E-23, D-C14).

## Local

```bash
cd 13-week/workshop/04-security-ci
go test -v ./...
```

Expected: `TestInternalEventsWithoutTokenIs401` **PASS**, status 401.

On the real service, with `INTERNAL_EVENT_TOKEN` set:

```bash
curl -s -o NUL -w "%{http_code}" -X POST http://localhost:8083/internal/events -d "{}"
```

Expected: `401`. A call with a valid `X-Internal-Token` is `202` or `400` (body), never a confirm via the portal.

## YAML

`security.yml` is a GitHub Actions sketch: gitleaks, govulncheck, then this test.
Validate with [actionlint](https://rhysd.github.io/actionlint/) or a YAML checker before CP8.
