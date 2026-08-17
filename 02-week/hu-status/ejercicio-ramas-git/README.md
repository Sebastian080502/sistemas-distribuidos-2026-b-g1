# Git branching exercise — Week 02

Session 2 practice of the course Git workflow applied to the Distributed Reservation Platform.

This exercise does not modify `02-week/01-session/` or `02-week/02-session/`. The evidence stays in this folder.

## Practiced flow

```text
main
  └── qa
        └── develop
              └── hu-xxx-dev
```

Per environment:

- `develop` ← HU branch `hu-xxx-dev` ← Pull Request to `develop`
- `qa` ← selective integration with cherry-pick
- `main` ← selective integration with cherry-pick

Commits must follow Conventional Commits: `type(scope): summary`.

## What was practiced

1. Create the base branches `qa` and `develop` from `main`.
2. Create a user-story branch from `develop`.
3. Record a change and push it to the HU branch.
4. Open a Pull Request to `develop`.
5. Verify that `develop` contains the HU commit.
6. Bring that commit to `qa` with cherry-pick, without merging all of `develop`.
7. Bring the same commit to `main` with cherry-pick.

Cherry-pick was used to apply a specific commit onto another branch. It is not equivalent to merging the whole source branch.

## Relationship with the project

This is the flow that will be used when the first Sprint of the Distributed Reservation Platform starts. Product stories (`HU-RES-001` onward) will be implemented on `hu-xxx-dev` branches, with a Pull Request to `develop` and controlled promotion to `qa` and `main`.

## Pending screenshots

Own screenshots of the exercise must be stored in `capturas/` with these names:

1. `01-ramas-base-main-qa-develop.png`
2. `02-rama-historia-usuario.png`
3. `03-commit-y-push-hu.png`
4. `04-pull-request-hacia-develop.png`
5. `05-verificacion-en-develop.png`
6. `06-cherry-pick-hacia-qa.png`
7. `07-cherry-pick-hacia-main.png`

The URL of the practice repository must also be added when it is available.
