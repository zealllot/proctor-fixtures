# proctor-fixtures

Sibling repo for [PRoctor](https://github.com/zealllot/proctor) end-to-end tests.

Contains a minimal cross-stack scaffold (frontend / api / schema / cli / docs) and a stable set of pre-built PRs covering each change category and several intentionally-broken cases. PRoctor's `tests/run-e2e.sh` drives `/proctor` against these PRs in dry-run mode and checks structured outcomes.

See `tests/fixtures/E2E.md` in the proctor repo for the PR contract (numbers and expected outcomes).

This repo is not a real application — it's just enough scaffolding for PRoctor's analyzer/planner to recognize each category. Nothing here is meant to compile or run end-to-end.
