---
title: Knative Governance Model
linkTitle: Governance
weight: 20
type: "docs"
aliases:
  - /contributing/governance
---

Knative is a vendor-neutral project under the [CNCF](https://cncf.io/).
Knative has used a number of different governance models in the past. The
_current_ governance of the Knative project is a maintainer council composed
of all currently active maintainers.  The maintainer council delegates
technical work to different [working groups](./working-groups/WORKING-GROUPS.md),
which are focused on specific components of Knative.  The specifics of Knative
governance decisions and processes are detailed below.

## Roles

All participants in the Knative project are contributors, and are expected to
follow the [CNCF Code of
Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md);
contributors.  Other roles in the
project are awarded based on participation, see [ROLES.md](./ROLES.md) for
details on the permissions and requirements for other roles.  Contributors
who have reached the [Approver](./ROLES.md#approver) or
[Working Group Lead](./ROLES.md#working-group-lead) roles are considered
[Maintainers](./MAINTAINERS.md), and may participate in voting and other processes described here.

## Conflict Resolution and Voting

The goal is for issues to be resolved by [consensus](./CONSENSUS.md) or in
[Working Groups](./working-groups/WORKING-GROUPS.md) where possible.  When those
mechanisms fail, issues are decided by a vote of all Maintainers across the project.
Maintainer decisions should be recorded publicly in GitHub issues or meeting notes
and recordings.  All votes should be recorded in GitHub issues.

* A vote requires a simple majority of quorum to pass.
* A vote requires participation by a quorum or at least 30% of [eligible maintainers](#organization-non-dominance).
* Voting is done by adding +1 (:thumbsup:) / -1 (:thumbsdown:) on the issue or PR.
* Each maintainer gets a binding vote, subject to organization limitations.  Non-binding votes from the community are permitted.
* Votes should be open for at least one week, and broadcast on at least one [Slack channel](./SLACK-GUIDELINES.md)
  which all maintainers should monitor.

### Organization Non-Dominance

No organization should be permitted to dominate governance by simple numbers.  If one organization makes up more than 40% of the Maintainers pool, votes from that organization may only be counted up to the 40% member threshold / quorum threshold.  The organization must coordinate which votes beyond the 40% should be discarded.

## Governance Updates

Changes to Governance (including this document) require a vote of the maintainers.
Beyond the above voting requirements, the following additional requirements apply
to governance updates:

* A vote requires a 2/3 majority of quorum to pass.
* Voting should be open for at least two weeks.

## Meetings

The Maintainers may choose to schedule live meetings regularly or from time to time;
these meetings may be combined with Working Group meetings where appropriate.  All such
meetings should be posted to the [community calendar](./CALENDAR.MD).