---
title: "Creating a extensions repo"
linkTitle: "extensions repo process"
weight: 30
type: "docs"
---

<!-- NOTE: some portion of this document is also reproduced in
/.github/ISSUE_TEMPLATES/new-repo.md. If you are updating the steps here, you
may need to look there, too. -->

# Background

The [`knative-extensions`](https://github.com/knative-extensions) GitHub org exists to
hold "non-core" Knative components which are owned and sponsored by a Working
Group. See the [Knative Repository Guidelines](../REPOSITORY-GUIDELINES.md) for
more details on the requirements for the `knative-extensions` org.

## Criteria

A Working Group Lead (either
[technical](../ROLES.md#working-group-technical-lead) or
[execution](../ROLES.md#working-group-execution-lead)) may request a new repo in
`knative-extensions` by filing an issue in the
[knative/community](https://github.com/knative/community/issues/new?template=new-repo.md)
repo. Once filed, the TOC should handle these promptly, though it should also be
considered fine to ping members or the group on Slack if it hasn't been acted on
in a few days. Generally, the request will be granted, though the TOC may have
additional questions or suggest an alternate mechanism in some cases.

Some of the following steps may require permissions that are only available to
the TOC or Steering Committee, though others are largely self-service or require
other WGs to review and approve impacting changes.

### Technical requirements

- Any contributed code should be contributed by the original authors or
  copyright holders under an Apache 2.0 license. See also
  [this section of the repository guidelines](../REPOSITORY-GUIDELINES.md#creation).

- Names between `knative-extensions` and the main `knative` repo should not
  overlap. This simplifies promoting repos between the two orgs and setting up
  `knative.dev` import paths for golang.

- Prow automation for tests is encouraged but not required for
  `knative-extensions`, but the Google CLA bot and OWNERS files/tide merge should
  be enforced.

## Process (to be executed by TOC or Steering member)

1. (Requires Org owner) Create the new repo in
   https://github.com/knative-extensions using the "New" button. Set the repo to
   public and include an "Apache License 2.0" but no `.gitignore` or `README`.

1. (Requires repo write/org owner) Create:

   - `OWNERS` file listing TOC and WG members as approvers, and WG members as
     reviewers

   - `CODE-OF-CONDUCT.md` (that links to
     https://github.com/knative/community/blob/main/CODE-OF-CONDUCT.md)

   - `README.md` for the repo by cloning it and pushing directly to the repo.

   At the end of this step, you should have 4 files: `LICENSE`, `OWNERS`,
   `CODE-OF-CONDUCT.md`, and `README.md`

1. Add entries for the repo to
   [`../peribolos/knative-extensions.yaml` in knative/community](https://github.com/knative/community/peribolos/knative-extensions.yaml).
   As part of this, grant one or more sponsoring WGs the "write" permission on
   the repo ([sample PR](https://github.com/knative/community/pull/170))

1. (For golang repos) Set up an alias under `knative.dev` to enable
   `knative.dev/` golang imports by adding a file for the repo to
   https://github.com/knative/website/tree/main/static/golang and updating the
   [`_redirects`](https://github.com/knative/website/blob/main/static/_redirects)
   file.

1. Set up
   [test-infra using the automation](https://github.com/knative/test-infra/blob/main/guides/prow_setup.md#setting-up-prow-for-a-new-repo-reviewers-assignment-and-auto-merge),
   which probably involves updating `config/prod/prow/config_knative.yaml` and
   then running `./hack/generate-configs.sh`

1. Ensure that Prow is working correctly by creating a PR against the repo. One
   good way to do this is to
   [add a `test/presubmit-tests.sh`](https://github.com/knative-extensions/discovery/pull/1)
   via either the web UI or using a fork.

1. Once at least one PR has been created, you'll be able to select the branch
   protection checks which are required in the "Settings" > "Branches" branch
   protection rule:

   - Under "Branches" add a branch protection rule for `main`:
     - Require status checks to pass (except `...-go-coverage` checks)
     - Include administrators

1. Encourage repository sponsor to add the repository to [CLOMonitor](https://clomonitor.io/projects/cncf/knative) ([more information](/REPOSITORY-GUIDELINES.md#clomonitor-and-clotributor)).
