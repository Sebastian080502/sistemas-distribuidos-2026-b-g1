# Resumen de Clase 1 — Semana 2

**Fecha:** 11 de agosto de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Arquitecturas distribuidas (cliente-servidor, P2P, SOA, microservicios)  

Se compararon estilos: monolito, cliente-servidor, SOA y microservicios. El profesor insistió en **no partir servicios por moda**: cada servicio necesita frontera de negocio, datos propios y un motivo para fallar o escalar aparte.

Un “microservicio” que comparte la misma base de datos con otros es un **monolito distribuido**.

## Cómo lo apliqué a DRP

Cinco bounded contexts: Identity, Space, Reservation, Payment, Notification. Cada uno con su PostgreSQL. Gateway y Angular son transversales.

Evidencia en este fork: [`docs/context-map.md`](./docs/context-map.md), [`docs/adr-001-architecture.md`](./docs/adr-001-architecture.md).  
Evidencia en producto: `drp-docs/02-domain/domain-map.md`, `drp-docs/05-architecture/decisions/records/ADR-005-microservices-architecture.md`.

## Diagrama

[`Diagram Session 1.svg`](./Diagram%20Session%201.svg)
