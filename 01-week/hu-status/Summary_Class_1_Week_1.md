# Resumen de Clase 1 — Semana 1

**Fecha:** 4 de agosto de 2026  
**Profesor:** Jesús Ariel González Bonilla  
**Tema:** Sistemas distribuidos — modelos, tiempo y consistencia  

Durante la clase 1 de la semana 1 se introdujo el curso de Sistemas Distribuidos: qué es un sistema distribuido, por qué no es solo “varios computadores”, y qué problemas aparecen cuando no hay un reloj ni una memoria compartida.

Se revisaron modelos de sistema (síncrono vs asíncrono), relojes lógicos frente a relojes físicos, y la idea de que **el orden de los eventos no es el mismo en todos los nodos**.

Se habló de **consistencia** y de que no se puede pedir a la vez consistencia fuerte, disponibilidad total y tolerancia a particiones (trade-off tipo CAP). Operaciones de negocio distintas pueden exigir consistencia distinta.

## Cómo lo apliqué a SpaceHub / DRP

El producto (plataforma de reservas de espacios) no puede permitir dos reservas **CONFIRMADAS** sobre el mismo espacio y el mismo intervalo (consistencia fuerte en creación/pago). Las notificaciones sí pueden llegar tarde (consistencia eventual). Eso quedó en el brief [`prd.md`](./prd.md) y más tarde en `drp-docs/01-context/` y `drp-docs/02-domain/`.

## Diagrama

[`Diagram Session 1.svg`](./Diagram%20Session%201.svg)
