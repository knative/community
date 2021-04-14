---
title: "Knative TOC election process"
linkTitle: "TOC election process"
weight: 30
type: "docs"
aliases:
  - /contributing/mechanics/toc/
---

# Guiding Principles

- The TOC seat should be earned by individuals and not company participation
- Members of the TOC should be exemplary members of the community who have
  demonstrated our community values
- The TOC should be made up of individuals from different working groups and
  vendors so that we have a diverse group of thought and input

# Composition

The TOC will have five seats with a 2 year term, with a maximum of 2 seats being
held by employees from the same organization (or conglomerate, in the case of companies owning each other).

There will be an annual election to determine the composition of the TOC for the
following year. Three seats will be up for election in one year and two will be
up for election the following year.

# Candidate Eligibility

Current TOC members and
[Approvers](https://github.com/knative/community/blob/main/ROLES.md#approver)
with at least 3 months tenure are eligible to stand for election. Candidates may
self-nominate or be nominated by another eligible member. The approximate time
commitment of a TOC member is around 8 hours per week.

Please note that vendor or organization limits do not restrict candidate eligibility.
We encourage all interested candidates in running as there may be occurrences where a TOC member
steps down and is back-filled by the candidate with the next most votes from the previous election.

# Voter Eligibility

Anyone who has at least 50 contributions in the last 12 months is eligible to
vote in the TOC election. Contributions are defined as opening PRs, reviewing
and commenting on PRs, opening and commenting on issues, writing design docs,
commenting on design docs, helping people on slack, participating in working
groups and etc.

[This dashboard](https://knative.teststats.cncf.io/d/9/developer-activity-counts-by-repository-group-table?orgId=1&var-period_name=Last%20year)
shows only GitHub based contributions and does not capture all the contributions
we value. _We expect this metric not to capture everyone who should be eligible
to vote._ If a community member has had significant contributions over the past
year but is not captured in the stats.knative.dev dashboard, they will be able
to submit an exception form to the steering committee who will then review and
determine whether this member should be marked as an exception.

All eligible voters will be captured at
`knative/community/elections/$YEAR/TOC/voters.md` and the voters’ guide
will be captured at `knative/community/elections/$YEAR/TOC/README.md`
similar to the kubernetes election process.

We are committed to an inclusive process and will adapt future eligibility
requirements based on community feedback.

# Election Process

Elections will be held using a time-limited
[Condorcet](https://en.wikipedia.org/wiki/Condorcet_method) ranking on
[CIVS](http://civs.cs.cornell.edu/) using the
[Schulze](https://en.wikipedia.org/wiki/Schulze_method) method. The top
vote-getters will be elected to the open seats. This is the same process used by
the Kubernetes project.

## Election Officers

The steering committee will be the election officers for the technical oversight
committee elections.

## Vacancies

In the event of a resignation or other loss of an elected TOC member, the
candidate with the next most votes from the previous election will be offered
the seat. This process will continue until the seat is filled.

In case this fails to fill the seat, a special election for that position will
be held as soon as possible. Eligible voters from the most recent election will
vote in the special election (ie: eligibility will not be redetermined at the
time of the special election). Any replacement TOC member will serve out the
remainder of the term for the person they are replacing, regardless of the
length of that remainder.

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
