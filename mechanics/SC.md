---
title: "Knative Steering Committee election process"
linkTitle: "SC election process"
weight: 30
type: "docs"
aliases:
  - /contributing/mechanics/sc/
---

# Guiding Principles

- The SC seat should be earned by individuals and not company participation
- Members of the SC should be exemplary members of the community who have
  demonstrated our community values
- The SC should be made up of individuals from different working groups and
  vendors so that we have a diverse group of thought and input

# Composition

The SC will have five seats with a 2 year term, with a maximum of 2 seats being
held by employees from the same vendor.

There will be an annual election to determine the composition of the SC for the
following year. Two seats will be up for election in one year and three will be
up for election the following year.

# Candidate Eligibility

Community members must be eligible to vote in order to stand for election (this
includes voters who qualify for an exception). Candidates may self-nominate or
be nominated by another eligible member.  There are no term limits for KSC
members. Nothing prevents a qualified member from serving on both the TOC and SC
simultaneously.

To run for a Contributing seat, a candidate must additionally be at least a
project Member as defined in
[ROLES.md](https://github.com/knative/community/blob/main/ROLES.md).

End-user seats are planned to enable employees of organizations which use but do
not sell Knative-derived products. The mechanics of eligibility for end-user
seats are still in progress; it is expected that end-user seats will add to the
current Steering options to specific include end-user perspectives in the
decision-making process.

# Voter Eligibility

Anyone who has at least 50 contributions in the last 12 months is eligible to
vote in the SC election. Contributions are defined as opening PRs, reviewing
and commenting on PRs, opening and commenting on issues, writing documentation,
commenting on docs, helping people on slack, participating in working
groups, actively promoting the Knative project, running project events, and
anything else that helps the Knative project succeed and grow.

[This dashboard](https://knative.teststats.cncf.io/d/9/developer-activity-counts-by-repository-group-table?orgId=1&var-period_name=Last%20year)
shows only GitHub based contributions and does not capture all the contributions
we value. _We expect this metric not to capture everyone who should be eligible
to vote._ If a community member has had significant contributions over the past
year but is not captured in the stats.knative.dev dashboard, they will be able
to submit an exception form to the steering committee who will then review and
determine whether this member should be marked as an exception.

All eligible voters will be captured at
`knative/community/elections/$YEAR/SC/voters.md` and the voters’ guide
will be captured at `knative/community/elections/$YEAR/SC/README.md`
similar to the kubernetes election process.

We are committed to an inclusive process and will adapt future eligibility
requirements based on community feedback.

# Election Process

Elections will be held using a time-limited
[Condorcet](https://en.wikipedia.org/wiki/Condorcet_method) ranking on
[CIVS](http://civs.cs.cornell.edu/) using the
[Schulze](https://en.wikipedia.org/wiki/Schulze_method) method. The top
vote-getters will be elected to the open seats. This is the same process used
for TOC elections.

## Election Officers

The Steering Committee will appoint
between one and three Election Officer(s) to administer the election. Elections
Officers will be Knative community members in good standing who are eligible to
vote, are not running for Steering in that election, who are not currently part
of the Steering Committee and can make a public promise of impartiality. They
will be responsible for:

- Making all announcements associated with the election
- Preparing and distributing electronic ballots
- Judging exception requests
- Assisting candidates in preparing and sharing statements
- Tallying voting results according to the rules

## Limitations on Company Representation

No more than two seats may be held by employees of the same organization (or
conglomerate, in the case of companies owning each other). Additionally, each
End-User seat (when available) must belong to an employee from a different
organization. If the results of an election result in greater than two employees
of the same organization, the lowest vote getters from any particular employer
will be removed until representation on the committee is down to two.

If employers change because of job changes, acquisitions, or other events, in a
way that would yield more than 2 seats being held by employees of the same
organization, sufficient members of the committee must resign until only two
employees of the same employer are left. If it is impossible to find sufficient
members to resign, all employees of that organization will be removed and new
special elections held. In the event of a question of company membership (for
example evaluating independence of corporate subsidiaries) a majority of all
non-involved Steering Committee members will decide.

## Vacancies

In the event of a resignation or other loss of an elected SC member, the
candidate with the next most votes from the previous election will be offered
the seat, provided that person otherwise qualifies to join the SC. This process
will continue until the seat is filled.

In case this fails to fill the seat, a special election for that position will
be held as soon as possible, unless the regular SC election is less than 7 weeks
away. Eligible voters from the most recent election will vote in the special
election. Eligibility will not be redetermined at the time of the special
election. Any replacement SC member will serve out the remainder of the term for
the person they are replacing, regardless of the length of that remainder.

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
