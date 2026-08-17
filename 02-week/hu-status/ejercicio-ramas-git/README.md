# Ejercicio de ramas Git — Semana 02

Práctica de la Sesión 2 para aplicar el flujo Git del curso sobre la Plataforma Distribuida de Reservas.

El ejercicio no modifica `02-week/01-session/` ni `02-week/02-session/`. La evidencia queda en esta carpeta.

## Flujo practicado

```text
main
  └── qa
        └── develop
              └── hu-xxx-dev
```

Por ambiente:

- `develop` ← rama de HU `hu-xxx-dev` ← Pull Request a `develop`
- `qa` ← integración selectiva con cherry-pick
- `main` ← integración selectiva con cherry-pick

Los commits deben seguir Conventional Commits: `type(scope): summary`.

## Qué se practicó

1. Crear las ramas base `qa` y `develop` a partir de `main`.
2. Crear una rama de historia de usuario desde `develop`.
3. Registrar un cambio y subirlo a la rama de la HU.
4. Abrir un Pull Request hacia `develop`.
5. Verificar que `develop` contiene el commit de la HU.
6. Llevar ese commit a `qa` con cherry-pick, sin fusionar toda `develop`.
7. Llevar el mismo commit a `main` con cherry-pick.

Cherry-pick se usó para aplicar un commit puntual en otra rama. No equivale a un merge de toda la rama origen.

## Relación con el proyecto

Este flujo es el que se usará cuando empiece el primer Sprint de la Plataforma Distribuida de Reservas. Las HUs de producto (`HU-RES-001` en adelante) se implementarán en ramas `hu-xxx-dev`, con PR hacia `develop` y promoción controlada a `qa` y `main`.

## Capturas pendientes

Las capturas del ejercicio propio deben guardarse en `capturas/` con estos nombres:

1. `01-ramas-base-main-qa-develop.png`
2. `02-rama-historia-usuario.png`
3. `03-commit-y-push-hu.png`
4. `04-pull-request-hacia-develop.png`
5. `05-verificacion-en-develop.png`
6. `06-cherry-pick-hacia-qa.png`
7. `07-cherry-pick-hacia-main.png`

También se debe agregar la URL del repositorio de práctica cuando esté disponible.
