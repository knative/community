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

## Actions to fulfill

This area is used for the TOC to track the [repo creation
process](/mechanics/CREATING-A-SANDBOX-REPO.md)


- [ ] Add this issue to the [TOC project board](https://github.com/orgs/knative/projects/9) for review.

- [ ] Send a PR adding entries for this repo in `/peribolos/knative-sandbox.yaml`
  - [ ] Add the repository and a description.
  - [ ] Grant `Knative Admin` the `admin` privilege.
  - [ ] Grant the sponsoring WG the `write` privilege.

_Once the TOC has approved the above, it will merge and Peribolos will create an empty repository._

- [ ] (golang) Send a PR to add aliases for `knative.dev/$REPONAME` import paths ([sample](https://github.com/knative/website/pull/187)).

- [ ] Have a lead from the sponsoring WG bootstrap the Git repository by pushing an
  appropriate "template" repository ([basic](https://github.com/knative-sandbox/wg-repository),
  [sample-controller](https://github.com/knative-sandbox/sample-controller),
  [sample-source](https://github.com/knative-sandbox/sample-source)) to the new repository as
  a git remote.  For example:

  ```shell
    git clone https://github.com/knative-sandbox/sample-controller.git
    cd sample-controller
    git remote add newrepo https://github.com/knative-sandbox/$REPONAME.git
    git push newrepo master
  ```

- [ ] Set up [test-infra](https://github.com/knative/test-infra) following the
  docs linked at the beginning.

- [ ] Create a sample PR to verify Prow (e.g. edit the boilerplate README)

_Once Prow has been verified._

- [ ] Have a member of the TOC enable branch protections so that `tide` is a required presubmit check. 

