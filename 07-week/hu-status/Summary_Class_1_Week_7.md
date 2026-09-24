# Resumen de Clase 1 — Semana 7

**Fecha:** 15 de septiembre de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Comunicación entre servicios — REST, gRPC y mensajería  

OVA semana 7, sesión 1: cada interacción elige **síncrono** (el llamador espera) o **asíncrono** (publica y sigue). REST/JSON es el default público. gRPC es interno y solo si hay latencia/throughput medido. El broker desacopla efectos secundarios.

En red no hay entrega exactly-once. El patrón práctico es **at-least-once + idempotencia**.

## Cómo lo apliqué a SpaceHub

| Interacción | Modo |
|-------------|------|
| Browser → gateway (login, listar espacios, crear reserva, pagar) | REST |
| reservation → space (¿el catálogo permite el slot?) | REST |
| PaymentConfirmed / ReservationConfirmed → notification | Evento |
| gRPC | No en Corte 2 (queda Could) |

Consumidor idempotente: `notification-service` por `eventId` / `source_event_id`.

## Diagrama

[`Diagram Session 1.png`](./Diagram%20Session%201.png)
