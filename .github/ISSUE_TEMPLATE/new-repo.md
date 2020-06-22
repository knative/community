---
name: Request a new repo
about: Request creation of a new repo in the knative or knative-sandbox orgs
title: "New Repo: $REPONAME"
---

Use this issue type to request a new repo in `knative-sandbox` (or
`knative`, which may require additional discussion).

<!-- Update the information below with your request -->

## Repo information

Org: knative-sandbox

Repo:

Purpose (Description):

Sponsoring WG:

## Actions to fulfil

This area is used for the TOC to track the [repo creation
process](/mechanics/CREATING-A-SANDBOX-REPO.md)

- [ ] Create a new repo in Github UI, with "Apache License 2.0"
- [ ] Create an OWNERS and README.md file
- [ ] Add entries in `/peribolos/knative-sandbox.yaml`
- [ ] (golang) [Add aliases](https://github.com/knative/website) for
  `knative.dev`
- [ ] Set up [test-infra](https://github.com/knative/test-infra) following the
  docs linked at the beginning.
- [ ] Create a PR with a `CODE-OF-CONDUCT.md` to verify Prow.
- [ ] Once Prow is enabled and PR checks are passing, add branch protection rules.

