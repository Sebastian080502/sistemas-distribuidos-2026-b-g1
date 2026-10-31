# How to run the TB6 negative test

Covers story `[security] Reject confirm hops without X-Internal-Token`
(acceptance: `POST /internal/events` without the header is 401, even with a user JWT).

This folder is a **proposal** for CI. It is not a change to `drp-reservation-api` today.
The handler copy matches the walking consume on `:8083` (E-23, D-C14).

## Local module (httptest)

```bash
cd 13-week/workshop/04-security-ci
go test -v ./...
```

Console (2026-10-30, this machine):

```
=== RUN   TestInternalEventsWithoutTokenIs401
--- PASS: TestInternalEventsWithoutTokenIs401 (0.01s)
PASS
ok  	github.com/Sebastian080502/spacehub-w13-authz	0.744s
```

## Real service on :8083 (CP8)

`drp-reservation-api` with `INTERNAL_EVENT_TOKEN=spacehub-dev-internal` (no `DATABASE_URL`; memory store). User JWT in `Authorization` is ignored.

```bash
curl.exe -sS -D - -X POST http://127.0.0.1:8083/internal/events -H "Content-Type: application/json" -H "Authorization: Bearer user-jwt" -d "{}"
```

Console (2026-10-30 18:35 UTC−5):

```
HTTP/1.1 401 Unauthorized
Content-Type: application/json
X-Correlation-Id: f45157d5d595d26e485b0feee476e342
Date: Fri, 30 Oct 2026 23:35:02 GMT
Content-Length: 116

{"error":"UNAUTHORIZED","message":"Token de autenticación requerido","traceId":"f45157d5d595d26e485b0feee476e342"}
```

Same URL with `X-Internal-Token: spacehub-dev-internal` and `{}` is **400** `VALIDATION_ERROR` / `Evento de pago inválido` (token accepted; body rejected). Never a portal confirm (D-C14).

## YAML (CP8)

`security.yml` is a GitHub Actions sketch: gitleaks, govulncheck, then this test.

Checked 2026-10-30 with Python 3 `yaml.safe_load` on `security.yml`: **YAML_OK** (parse succeeded; no schema error). `actionlint` is not installed on this machine; the file is a three-step `pull_request` job on `ubuntu-latest`.
