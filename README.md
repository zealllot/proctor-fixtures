# proctor-fixtures

A live integration fixture for [PRoctor](https://github.com/zealllot/proctor). Every PR opened here runs through the real PRoctor GitHub Action — analyze diff → plan tests → run them → auto-fix failures → post a structured report comment with screenshots.

If you want to **see PRoctor in action without setting it up yourself**, browse the open and closed PRs here. If you're integrating PRoctor into your own project, this repo is the working reference for `.pr-test.yml` and the Action workflow.

## Quick Start (5 minutes, run locally)

```bash
git clone https://github.com/zealllot/proctor-fixtures
cd proctor-fixtures

# Frontend (Vite + React, port 5173)
corepack enable && corepack prepare pnpm@9 --activate
pnpm install
pnpm dev &

# Admin (QOR + Go + SQLite, port 7000)
go run ./cmd/admin &

# Probes
curl http://127.0.0.1:5173        # Vite app — Login button + Settings form
curl http://127.0.0.1:7000/admin  # QOR admin — User and Article resources
```

`.pr-test.yml` and `.github/workflows/proctor.yml` show how PRoctor consumes both services in CI.

## What's in here

```
.
├── src/                          ← Vite + React 18 frontend
│   ├── components/Login.tsx      ← LoginButton (icon, hover, tooltip, loading)
│   ├── pages/Settings.tsx        ← Settings form
│   └── main.tsx                  ← Mounts both
├── cmd/admin/main.go             ← QOR admin: User + Article resources
├── api/                          ← Go HTTP handlers (users, profile)
├── migrations/                   ← SQL migrations for the admin's DB
├── docs/architecture.md          ← Path → category mapping for PRoctor analyzer
├── go.mod / go.sum               ← Go deps (qor v1.2.0, gorm v1.9.16)
├── package.json / tsconfig.json  ← Node deps + TS config
├── .pr-test.yml                  ← PRoctor configuration (setup commands, base_url)
└── .github/workflows/proctor.yml ← Pinned to zealllot/proctor/github-action@v0.2.1
```

## Live PR catalog

Eight long-lived PRs serve as PRoctor's e2e test corpus — they're never merged, but their reports show what PRoctor produces for each category.

| PR | Category | What it demonstrates |
|---|---|---|
| [#1](https://github.com/zealllot/proctor-fixtures/pull/1) | frontend | Adding `aria-label` and `type="button"` to a React component (lint-only verifiable) |
| [#2](https://github.com/zealllot/proctor-fixtures/pull/2) | api | Adding `?limit` query parameter to a Go handler (curl-verifiable) |
| [#3](https://github.com/zealllot/proctor-fixtures/pull/3) | schema | Adding a `user_tags` SQL migration |
| [#4](https://github.com/zealllot/proctor-fixtures/pull/4) | mixed (frontend + api) | Profile display-name field — exercises the **e2e-flow** rule (planner generates a cross-layer test item) |
| [#5](https://github.com/zealllot/proctor-fixtures/pull/5) | bug (frontend) | Intentional typo (`labl` vs `label`) — auto-fix should catch and patch |
| [#6](https://github.com/zealllot/proctor-fixtures/pull/6) | bug (api) | Intentional MIME-type misspelling — auto-fix should patch |
| [#7](https://github.com/zealllot/proctor-fixtures/pull/7) | unfixable refactor | Removes type contract; auto-fix can't restore it without redesign — report says "needs human review" |
| [#8](https://github.com/zealllot/proctor-fixtures/pull/8) | docs-only | Architecture doc edit — planner picks `lint-only` exclusively |

Plus continuously-running demo PRs that exercise specific PRoctor versions:

- [#21 — Admin visual: site name + colored status badges](https://github.com/zealllot/proctor-fixtures/pull/21) — chrome-devtools tests with inline screenshots
- [#18 — LoginButton: icon, hover, tooltip](https://github.com/zealllot/proctor-fixtures/pull/18) — 5/5 chrome-devtools pass with `screenshot_focus` annotations
- [#17 — User.Phone field](https://github.com/zealllot/proctor-fixtures/pull/17) — exercises pr_context (PR body referencing fake Slack/Jira URLs)

Click any PR → scroll to the bot comment → see exactly what PRoctor produces.

## Want to use this as a template

1. Copy [`.pr-test.yml`](.pr-test.yml) and adapt `setup:` for your stack.
2. Copy [`.github/workflows/proctor.yml`](.github/workflows/proctor.yml) — already at the latest pin.
3. Set the `CLAUDE_CODE_OAUTH_TOKEN` secret on your repo.
4. Enable Actions PR-creation:
   ```bash
   gh api -X PUT "/repos/<owner>/<repo>/actions/permissions/workflow" \
     -f default_workflow_permissions=write -F can_approve_pull_request_reviews=true
   ```

Or skip the manual steps and run `claude /proctor-init` in your repo — see the [main PRoctor README](https://github.com/zealllot/proctor#readme).

## Why two services in one repo

PRoctor categorizes diffs across 8 categories (frontend, api, schema, infra, mobile, cli, e2e-flow, docs). To exercise each end-to-end, the fixture needs surface area in multiple stacks:

- **Vite + React** for `frontend` items (chrome-devtools, computed styles)
- **Go + QOR** for `api`, `cli`, `schema` items (curl, go test, migrations)
- Mixed PRs cross both → `e2e-flow`

Real consumer projects usually pick one stack — they don't need this complexity. The fixture has to demonstrate all of them.

## Heads up

This repo isn't a real product. It exists to prove PRoctor works against realistic code. Don't use it as a starting template for an actual app — copy [`/proctor-init`](https://github.com/zealllot/proctor/blob/main/plugins/proctor/commands/proctor-init.md)'s output into your real project instead.
