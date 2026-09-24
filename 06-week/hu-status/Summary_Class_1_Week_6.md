# Resumen de Clase 1 — Semana 6

**Fecha:** 8 de septiembre de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Docker Compose y orquestación básica  

OVA semana 6, sesión 1: pasar de **un contenedor** a **un sistema**. Compose describe servicios en una red privada, DNS por nombre de servicio, `depends_on` + healthcheck, volúmenes con nombre y configuración por variables de entorno.

`depends_on` sin condición solo ordena el arranque; no garantiza que PostgreSQL acepte conexiones. Hace falta `condition: service_healthy` y reintento en el proceso Go.

Compose sirve para local y un solo host. Kubernetes queda fuera del MVP de este curso.

## Cómo lo apliqué a SpaceHub / DRP

El `docker-compose.yml` **no va en este fork**. El profesor creará los repos de código (uno por microservicio) y dirá dónde vive Compose (Q-004). El diseño de topología quedó en `drp-docs/10-devops/local-setup.md`:

- Red `spacehub-net`
- Un PostgreSQL por contexto (identity, space, reservation, payment, notification)
- Gateway `:8080`, Angular `:4200`, broker AMQP `:5672`
- Volúmenes `*_pg_data`; contenedores de aplicación sin estado

## Diagrama

[`Diagram Session 1.png`](./Diagram%20Session%201.png)
