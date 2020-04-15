---
title: "Creating a sandbox repo"
linkTitle: "Sandbox repo process"
weight: 30
type: "docs"
---

# Background

The [`knative-sandbox`](https://github.com/knative-sandbox) GitHub org exists to
hold "non-core" Knative components which are owned and sponsored by a Working
Group but are not part of an official Knative release.

## Criteria

A Working Group Lead (either
[technical](../ROLES.md#working-group-technical-lead) or
[execution](../ROLES.md#working-group-execution-lead)) may request a new repo in
`knative-sandbox` from any [TOC member](../TECH-OVERSIGHT-COMMITTEE.md).
Generally, the request will be granted, though the TOC may have additional
questions or suggest an alternate mechanism in some cases.

Some of the following steps may require permissions that are only available to
the TOC or Steering Committee, though others are largely self-service or require
other WGs to review and approve impacting changes.

### Technical requirements

- Any contributed code should be contributed by the original authors or
  copyright holders under an Apache 2.0 license.

- Names between `knative-sandbox` and the main `knative` repo should not
  overlap. This simplifies promoting repos between the two orgs and setting up
  `knative.dev` import paths for golang.

- Prow automation for tests is not necessarily required for `knative-sandbox`,
  but the Google CLA bot and OWNERS files/tide merge should be enforced.

## Process

1. (Requires Org owner) Create the new repo in
   https://github.com/knative-sandbox using the "New" button. Set the repo to
   public and include an "Apache License 2.0" but no `.gitignore` or `README`.

1. (Requires repo write/org owner) Create an `OWNERS` file and `README.md` for
   the repo by cloning it and pushing directly to the repo.

1. (For golang repos) Set up an alias under `knative.dev` to enable
   `knative.dev/` golang imports by adding a file for the repo to
   https://github.com/knative/website/tree/master/static/golang and updating the
   [`_redirects`](https://github.com/knative/website/blob/master/static/_redirects)
   file.

1. Set up
   [test-infra](https://github.com/knative/test-infra/blob/master/guides/prow_setup.md#setting-up-prow-for-a-new-repo-reviewers-assignment-and-auto-merge)
   using the automation, which probably involves updating
   `config/prow/config_knative.yaml` and then running
   `./hack/generate-configs.sh`

1. Set up the "Settings" on the repo as follows:

   - Disable "wiki"
   - Set merge to only support squash merging
   - Under "Manage Access":
     - set "@knative-admin" group as "Admin" role
     - set "\${WG}-writers" group as "Write" role
   - Under "Branches" add a branch protection rule for `master`:
     - Require status checks to pass (except `...-go-coverage` checks)
     - Include administrators

1. Ensure that Prow is working correctly by creating a PR against the repo.
