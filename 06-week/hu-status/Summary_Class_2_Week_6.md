# Resumen de Clase 2 — Semana 6

**Fecha:** 9 de septiembre de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Entornos, configuración 12-factor y plan de orquestación  

Sesión de planning: la **misma imagen** corre en `develop`, `qa` y `prod`; solo cambia la configuración. Los secretos no van en Git ni en la imagen. Un `.env.example` lista variables sin valores reales.

Git-flow del **código** (no de `drp-docs`):

```
hu-xxx-dev  -> develop
hu-xxx-qa   -> qa
hu-xxx-main -> main
```

`drp-docs` sigue en `main` solamente.

## Cómo lo apliqué

Historias Must de Corte 2 para el repo de código (detalle en `drp-docs/15-project-control/technical-backlog.md`):

1. Un `docker compose up` con healthchecks y volúmenes
2. Config externa y arranque fail-fast si falta una variable
3. Promover el mismo digest de imagen a QA
4. PRs de HU al ambiente correcto

## Diagrama

[`Diagram Session 2.png`](./Diagram%20Session%202.png)
