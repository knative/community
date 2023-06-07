---
title: "Knative Repository Guidelines"
linkTitle: "Repository Guidelines"
weight: 25
type: "docs"
aliases:
  - /contributing/repository-guidelines/
---

This document outlines a structure for creating and associating code
repositories with the Knative project. It also describes how and when
repositories are removed.

## Core Repositories

Core repositories are considered core components of Knative. They are utilities,
tools, applications, or libraries that form or support the foundation of the
project.

Core repositories are all those located under the
[github.com/knative organization](https://github.com/knative).

### Core Repository Requirements

- Repository must live under github.com/knative/project-name.
- Must adopt the Knative Code of Conduct.
- All code projects must use the Apache License version 2.0.
- Documentation repositories must use Creative Commons License version 4.0.
- All OWNERS must be members of the Knative community.
- Repository creation must be approved by the Technical Oversight Committee.

### Requirements for Repository Donation

The Technical Oversight Committee may allow repositories to be donated to the
Knative project. In _addition_ to meeting the core repository requirements,
donated repositories must:

- Adopt the Google CLA bot.
- All contributors to the donated repository must have signed the Google CLA.
- All code must adopt the standard Knative header, attributing copyright as
  follows: `"Copyright <year> The Knative Authors"`.
- Additions of the standard Knative header to code created by the contributors
  can occur post-transfer, but should ideally occur shortly thereafter.
- An `AUTHORS` file should be created in the repo root, if one does not already
  exist, that includes the original authors of the code. See
  [knative/serving AUTHORS](https://github.com/knative/serving/blob/main/AUTHORS)
  for reference.
  - Note that copyright notices should only be modified or removed by the people
    or organizations named in the notice.

## Removing Repositories

As important as it is to add new repositories, it is equally important to prune
repositories when necessary. See [Grounds for removal](#grounds-for-removal).

It is important to the success of Knative that all Knative repositories stay
active, healthy, and aligned with the scope and mission of project.

### Grounds for removal

Core repositories may be removed from the project if they are deemed _inactive_.
Inactive repositories are those that meet any of the following criteria:

- There are no longer any active maintainers for the project, and no
  replacements can be found.
- All PRs and issues have gone un-addressed for longer than six months.
- There have been no new commits or other changes in more than a year.
- The contents have been folded into another actively maintained project.

### Procedure for removal

When a repository has been deemed eligible for removal, we take the following
steps:

- A proposal to remove the repository is brought to the attention of the
  Technical Oversight Committee through a GitHub issue posted in the
  [community](https://github.com/knative/community) repo.
  - Feedback is encouraged during a Technical Oversight Committee meeting before
    any action is taken.
- Once the TOC has approved of the removal, if the repo is not moving to another
  actively maintained project:
  - The repo description is edited to start with the phrase "[EOL]"
  - All open issues and PRs are closed
  - All external collaborates are removed
  - All webhooks, apps, integrations, or services are removed
  - GitHub pages are disabled
  - The repo is marked as archived using GitHub's archive feature
  - The removal is announced on the knative-dev mailing list
  - Documentation is removed / updated as appropriate from the knative/docs website

This procedure maintains the complete record of issues, PRs, and other
contributions. It leaves the repository read-only and makes it clear that the
repository is retired and no longer maintained.

## Additional / Experimental Repositories

It is beneficial for Knative working groups to be able to own code outside the
core `knative` organization, which is kept intentionally small and has a high
bar for entry.  The `knative-sandbox` GitHub organization was created for this
purpose.  Repositories in `knative-sandbox` are intended to give working groups
more latitude to experiment with new ideas and to self-organize and manage
contributions which may be important to the project but which do not need the
level of governance of the core project.

Note that `knative-sandbox` is **not** an incubation area for projects entering
the `knative` github org; in most cases, projects in the sandbox will remain
there for the entire duration. `knative-sandbox` also provides a location for
implementations of core interfaces which do not themselves need to be part of
every knative installation, such as networking or eventing integrations. In the
event that a working group wants a project in sandbox to be considered for
transfer to the `knative` org, the request will be considered on a case-by-case
basis by joint decision of the Steering Committee and the Technical Oversight
Committee.

### Mechanics of working-group owned repositories in `knative-sandbox`

#### Creation

Note: prior art here from the
[Kubernetes community](https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md).

Working groups are allowed to request repositories that meet the following
requirements. The Steering Committee and TOC will maintain the process and
perform the repository creation steps to ensure that the following requirements
have been met:

- Vendor dependencies must be clearly called out, and participation must be open
  to the community.

  - If the repository is for integration with a specific technology vendor,
    affiliation with that vendor must not be required for participation.

  - Non-integration repositories should aim to avoid dependencies on
    vendor-specific technology by publishing abstract interfaces which can be
    implemented by multiple vendors. Those implementations are not required to
    live in separate repositories.

- Must adopt the Knative CLA.

- Must adopt the Knative
  [code of conduct](https://github.com/knative/community/blob/main/CODE-OF-CONDUCT.md).

- Must adopt the Apache 2.0 license.

- Must be sponsored by a specific working group, which should be designated in
  the `README.md` for the repo. In addition, at least one WG lead must be in the
  root-level OWNERS file.

- Copyright assigned to The Knative Authors.

- Document clear release and install processes.

- Clearly document the stability and API guarantees (both in the top-level
  `README.md` and by using appropriate API versioning for Kubernetes apigroups).

  - The TOC will maintain a template which will be used during the creation of
    new repositories. Existing repositiories are encouraged to update their
    `README`s from time to time as the format changes.

- Kubernetes APIs must be homed in a valid group with a DNS prefix approved by
  the Steering Committee.

  - APIs under `knative.dev` should:

    1. Be under a subdomain managed by the owning Working Group (e.g.
       `sources.knative.dev`, not `knative.dev`). The TOC is considered the
       authority for assigning subdomains to working groups.

    1. Coordinate with the Working Group in API naming to avoid collisions or
       confusing naming. This also includes API review by the responsible
       Working Group to ensure consistency and alignment with duck-type field
       naming.

Working groups must designate the repositories they own (by prefix or in the README.md of the repo, at a minimum).

Repositories owned by working groups are encouraged to embrace the project value
of pluggability.

### Non-requirements

The following are not required to create a working-group-owned repository:

- Steering or Trademark approval (see ["the fine print"](#the-fine-print))
- TOC approval
  - TOC may request certain naming patterns (e.g. `kn-plugin` for client WG)
- Solving an unique problem (exploring different approaches to problems outside
  the core is actively encouraged!)

### The fine print

Steering and Trademark reserve the right to require that repos be removed or
transferred out of the `knative-sandbox` organization and API groups to be
renamed. This is intended to simplify the process in the common case, while
giving Steering or Trademark the ability to step in and rectify problems that
may arise.

### Deletion

Working groups are responsible for the upkeep of their repositories and are
expected to keep a high quality bar. Inactive or non-conforming repositories
should be archived, deleted, or moved out of the `knative-sandbox` organization,
following the general [Procedure for Removal](#procedure-for-removal) for
repositories.

## CLOMonitor and CLOTributor

[CLOMonitor](https://clomonitor.io/) is a CNCF tool that periodically checks open source projects repositories to verify they meet certain project health best practices.

[CLOTributor](https://clotributor.dev/) is a CNCF tool that makes it easier to discover great opportunities to become a Cloud Native contributor.

When a new repository is created, consider adding it to CLOMonitor [data file](https://github.com/cncf/clomonitor/blob/main/data/cncf.yaml), similar to:

```yaml
    - name: eventing
      url: https://github.com/knative/eventing
      check_sets:
        - code-lite
```

This will make the repository will be listed in CLOMonitor's Knative project [health report](https://clomonitor.io/projects/cncf/knative). It will also make CLOTributor to list `good-first-issue`s of the new repository.

If you would like to exclude the repository from the health report but have the `good-first-issues` in CLOTributor, add following the field:

```yaml
      exclude:
        - clomonitor
```

Based on the maturity of the repository, it is possible to use different `check_sets`. See CLOMonitor's [checks documentation](https://github.com/cncf/clomonitor/blob/main/docs/checks.md) for more information.

---

Contents of this page are adopted from the
[Kubernetes repository guidelines](https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md),
which is licensed under Apache License 2.0.

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
