# Architecture (stub)

This is a fixture repo with no real architecture. Files exist to give PRoctor's analyzer something representative of each category to look at.

## Layout

| Path | Category |
|---|---|
| `src/components/`, `src/pages/` | frontend |
| `api/` | api |
| `migrations/` | schema |
| `cmd/` | cli |
| `docs/` | docs |

Nothing here is meant to compile or run end-to-end. The PRs in this repo intentionally span — and sometimes break — these areas to exercise PRoctor's pipeline stages.
