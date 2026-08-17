# Git branching exercise — Week 02

Practice repository: https://github.com/Sebastian080502/prueba_sistema_distribuidos

Local practice folder: `C:\Users\USER\prueba`

This exercise applied the course Git workflow for the Distributed Reservation Platform. It does not modify `02-week/01-session/` or `02-week/02-session/`.

## Environment branches

```text
main
  └── qa
        └── develop
              ├── HU-01-dev
              ├── HU-01-qa
              ├── HU-02
              └── release/1.0.0
```

Commits follow Conventional Commits: `type(scope): summary`.

## What was practiced

### 1. HU-01 on develop

Branch `HU-01-dev` was created from `develop`. The work was committed as `df6675a` with message `feat: implement HU-01`, pushed, and merged through [Pull Request #1](https://github.com/Sebastian080502/prueba_sistema_distribuidos/pull/1) into `develop`.

![HU-01-dev commit and push](./screenshots/01-hu-01-dev-commit-push.png)

### 2. HU-01 on qa

Branch `HU-01-qa` was created. The validation commit `test:validate HU-01-qa` was pushed and integrated through [Pull Request #2](https://github.com/Sebastian080502/prueba_sistema_distribuidos/pull/2) into `qa`.

![Create HU-01-qa branch](./screenshots/02-create-hu-01-qa-branch.png)

![HU-01-qa commit and push](./screenshots/03-hu-01-qa-commit-push.png)

### 3. Release 1.0.0 on main

Branch `release/1.0.0` prepared the release with `release: prepare version 1.0.0` and [Pull Request #3](https://github.com/Sebastian080502/prueba_sistema_distribuidos/pull/3) into `main`.

![Create release/1.0.0](./screenshots/04-create-release-1-0-0.png)

![Fetch and branch graph](./screenshots/05-fetch-and-branch-graph.png)

![Release commit and push](./screenshots/06-release-commit-push.png)

### 4. HU-02 from develop

HU-02 started from `develop` with `feat: implement HU-02` and [Pull Request #4](https://github.com/Sebastian080502/prueba_sistema_distribuidos/pull/4) into `develop`.

![Create HU-02 from develop](./screenshots/07-create-hu-02-from-develop.png)

![HU-02 commit](./screenshots/08-hu-02-commit.png)

### 5. Cherry-pick HU-02 onto qa and main

Commit `d4c1a2d` was cherry-picked onto `qa`, then onto `main` (`62666b4` / `9f992fe`). Cherry-pick applied one specific commit without merging the whole source branch.

![Cherry-pick HU-02 to qa](./screenshots/09-cherry-pick-hu-02-to-qa.png)

![Pull main](./screenshots/10-pull-main.png)

![Cherry-pick on qa then switch to main](./screenshots/11-cherry-pick-on-qa-then-switch-main.png)

![Cherry-pick to main](./screenshots/12-cherry-pick-to-main.png)

![Push main](./screenshots/13-push-main.png)

![Full branch graph](./screenshots/14-full-branch-graph.png)

## Relationship with the project

The same flow will be used in the first Sprint of the Distributed Reservation Platform: product stories on `hu-xxx-dev`, Pull Request to `develop`, and controlled promotion to `qa` and `main`.
