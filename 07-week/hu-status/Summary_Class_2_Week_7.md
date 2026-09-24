# Resumen de Clase 2 — Semana 7

**Fecha:** 16 de septiembre de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Contratos versionados y contract testing  

El contrato machine-readable (OpenAPI, proto, JSON Schema) es la fuente de verdad, no el código. Cambio compatible: campo opcional nuevo. Breaking: renombrar/borrar/cambiar tipo → `/api/v2` o `PaymentConfirmed.v2`, con deprecación (`Sunset`).

Pact (consumer-driven): el consumidor publica expectativas; el productor las verifica en CI antes de promover a QA. Los unit tests no detectan ese drift.

## Cómo lo apliqué

OpenAPI `/api/v1` ya está en `drp-docs/07-api/contracts/openapi/`. Envelope de error en `_shared.yaml`. El test de contrato en CI espera a que el profesor cree los repos de cada servicio (TD-004).

## Diagrama

[`Diagram Session 2.png`](./Diagram%20Session%202.png)
