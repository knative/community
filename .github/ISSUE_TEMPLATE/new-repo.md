---
name: Request a new repo
about: Request creation of a new repo in the knative or knative-extensions orgs
title: "New Repo: `$REPONAME`"
---

Use this issue type to request a new repo in `knative-extensions` (or
`knative`, which may require additional discussion).

<!-- Update the information below with your request -->

## Repo information

Org: `knative-extensions`

Repo: `$REPONAME`

Go module (Optional): `knative.dev/$REPONAME`

Purpose (Description):

Sponsoring WG:

## Actions to fulfill

This area is used to track the [repo creation process](https://github.com/knative/community/blob/main/mechanics/CREATING-A-SANDBOX-REPO.md).
The _requestor_ and _sponsoring WG lead_ should perform the steps listed below and cross out the checkmarks when done.
The TOC is involved only in the **TOC Gate** steps.

- [ ] Add this issue to the [TOC project board](https://github.com/orgs/knative/projects/43) for review. You are responsible for moving your entry on the board to "Needs Discussion" or "In Progress" as you move forward in this checklist.

_You may not be able to use the Projects quick menu on this page. In that case, go to the project board and use the **Add cards** interface._

- [ ] Send a PR adding entries for this repo in [`/peribolos/knative-extensions.yaml`](https://github.com/knative/community/blob/main/peribolos/knative-extensions.yaml). Please mind the alphabetical order when adding to a list.
  - [ ] Add the repository and a description.
  - [ ] Grant `Knative Admin` the `admin` privilege.
  - [ ] Grant the sponsoring WG the `write` privilege.

**TOC Gate**: _Once the TOC has approved the above, it will merge and Peribolos will create an empty repository._

- [ ] Ask Steering to add the repo to EasyCLA in [Slack](https://cloud-native.slack.com/archives/C04LQCW0C03/p1676466607624469).

- [ ] (golang) Send a PR to add aliases for `knative.dev/$REPONAME` import paths ([sample](https://github.com/knative/docs/pull/4160)).

- [ ] Have a lead from the sponsoring WG bootstrap the Git repository by pushing an
  appropriate "template" repository ([basic](https://github.com/knative-extensions/wg-repository),
  [sample-controller](https://github.com/knative-extensions/sample-controller),
  [sample-source](https://github.com/knative-extensions/sample-source)) to the new repository as
  a git remote.  For example:

  ```shell
    git clone https://github.com/knative-extensions/sample-controller.git
    cd sample-controller
    git remote add newrepo https://github.com/knative-extensions/$REPONAME.git
    git push newrepo main
  ```

- [ ] Add your GitHub ID to the `OWNERS` file for your repo.

- [ ] [Set up](https://github.com/knative/test-infra/blob/main/guides/prow_knative_setup.md#setting-up-prow-for-a-new-repo-reviewers-assignment-and-auto-merge) prow for a new repo

- [ ] Bootstrap your CI jobs using [hack](https://github.com/knative/hack) project (look at other sandbox repos for reference)

- [ ] Create a sample PR to verify Prow (e.g., edit the boilerplate README)

- [ ] Verify that within 24 hours the appropriate branch protections have been applied
   requiring `tide` to pass before PRs are merged.

- [ ] (optional) Send a PR adding the repo to [knobots](https://github.com/knative-extensions/knobots).

- [ ] (optional) Send a PR adding the repo to [CLOMonitor](https://clomonitor.io/projects/cncf/knative) ([more information](/REPOSITORY-GUIDELINES.md#clomonitor-and-clotributor)).
