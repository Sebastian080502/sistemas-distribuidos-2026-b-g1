# Resumen de Clase 2 — Semana 1

**Fecha:** 5 de agosto de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Trade-offs, semántica de entrega y organización del proyecto  

La sesión 2 (planning) se usó para bajar los conceptos de la sesión 1 a **decisiones de producto**: qué problema se resuelve, qué queda fuera del MVP, y cómo se entrega el reporte semanal (`NN-week/hu-status/README.md`).

Se recordó el flujo Git por ambiente (`hu-xxx-dev` → `develop`, etc.) y Conventional Commits. El `hu-status` se califica de forma automática: hay que respetar la plantilla.

## Cómo lo apliqué a DRP

- Problema: reservas solapadas, pago duplicado, notificación que no debe deshacer una confirmación.  
- Backlog inicial HU-RES-001 … HU-RES-008.  
- Entrega: at-least-once + consumidores idempotentes en eventos; REST para consultas y creación.

Evidencia de producto: [`prd.md`](./prd.md). Más adelante el mismo contenido vive en `drp-docs` (`03-product/`, `04-requirements/user-stories.md`).

## Diagrama

[`Diagram Session 2.svg`](./Diagram%20Session%202.svg)
