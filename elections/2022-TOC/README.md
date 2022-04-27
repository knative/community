---
title: "Knative TOC Election - 2022 Voters Guide"
linkTitle: "2022 Voters Guide"
weight: 30
type: "docs"
---

# 2022 Voters Guide - Knative TOC Election

## Purpose

The role of this election is to fill the four (4) seats due for election this year on the [Knative Technical Oversight Committee]. The elected
members will serve two (2) year terms.

## Background

This election will shape the future of Knative as a community and project.
While WGs help shape the individual direction of the project, the
[TOC Charter] covers the technical direction, health of the project
and community as a whole.

## Eligibility

Please refer to the [TOC Election Charter] for:

- [Eligibility for candidacy]
- [Eligibility for voting] 

The list of eligible voters is stored [in the community repository](https://github.com/knative/community/elections/2022-TOC/), but it is easier
to check if you are eligible by logging into [Elekto]. See Voting Process below for more details.

### Schedule

| Date         | Event                    |
| ------------ | ------------------------ |
| April 26     | Announcement of Election |
| May 10       | All candidate nominations are due by 23:59 UTC (4:59pm Pacific) |
| May 13       | Election Begins via Elekto UI |
| May 24       | Voting exception forms are due by 23:59 UTC (4:59pm Pacific)
| May 27       | Election Closes by 23:59 UTC (4:59pm Pacific) |
| June 1       | Announcement of Results |

## Election Officers

In an effort to run a fair and transparent election, the following people
have been asked by the Steering Committee to run this election:

- [Josh Berkus](https://github.com/jberkus)
- [Dawn Foster](https://github.com/geekygirldawn)

You can reach us by emailing elections@knative.team

## Candidacy Process

**Nomination**

If you want to stand for election, open a PR against the
[knative/community repository](https://github.com/knative/community) to include
your candidate profile in the `/elections/2022-TOC` folder, with the following
filename format:

```
candidate-yourname.md
```

`yourname` in the template file should match your Github handle.

We have included a template file as an example. This profile should include:

* Your name
* Your company affiliation (employer or otherwise)
* Your contributions to Knative
* Why you are running

You can find [examples from the last TOC election](https://github.com/knative/community/tree/main/elections/2021-TOC).

Once you have created the PR, you may email knative-dev and/or knative-users *once* to let
people know about your candidacy and encourage endorsements as comments on the PR.
Please use something like for your email in order to encourage endorsements
in Github and not on the mailing list:

```
Fellow Knative community members,

My name is {your name} and I am running for TOC.  You can read
my profile here: {link to PR}.  If you support my candidacy, please endorse me
by commenting on that PR.  Please don't reply to endorse me on this list.
```

After a candidate has met all election requirements, the Election Officers will
merge the profile PR.

If you want to nominate someone else, you may do so, but PLEASE talk to them
first.

**Campaigning**

Please refer to the [TOC Election Charter] and understand
that we care deeply about [limiting corporate campaigning]. The election
officers and members of the steering committee [pledge to recuse] themselves
from any form of electioneering.

You should be running as a "brand free" individual, based on your contribution
to the project as a member of this community, outside of whatever corporate
roles you may hold.

## Voting Process

Elections will be held once again using [Elekto], an online voting tool created
by CNCF intern Manish Sahani. As a reminder it relies on GitHub Oauth to log you
in to vote, instead of relying on email.

Thus, when you go to [Elekto] you will be prompted to log in your GitHub account.
Please do so, and then click on "Explore Election" to look at the list of 
elections.  From there you can click on the "2022 Knative TOC Election."

The election page will, among other things, tell you if you are eligible to vote,
via a button at the top of the screen. Due to limitations in our contributor
data, many contributors may have been unfairly missed; if you are one of these,
please [file a voting exception] by the deadline above so that we can enfranchise you.

As candidates file their candidate statements in the community repo, they will
become visible in the [Elekto] UI.  You may click through to any candidate
to see their profile.

Once the vote begins, you will be able to rank the candidates in the order of
your preference, and submit your ballot.  When you submit, you will be offered
a chance to set a password, which is required if you want the ability to return
and re-cast your ballot before the voting deadline above.

All data for Elekto is stored on a pair of GCP instances operated by Josh Berkus
and Manish Sahani, and is not shared with third parties.  Individual ballot data
is encrypted, and not retrievable by anyone except in aggregate form.

### Decision

Ballots are compiled by Elekto and all candidates are ranked using the Condorcet
method, Schultze variant.

For more information, definitions, and/or detailed election process, please refer to
the [TOC Election Charter].

[Knative Technical Oversight Committee]: https://github.com/knative/community/blob/master/TECH-OVERSIGHT-COMMITTEE.md
[TOC Charter]: https://github.com/knative/community/blob/main/TECH-OVERSIGHT-COMMITTEE.md#charter
[TOC Election Charter]: https://github.com/knative/community/blob/master/mechanics/TOC.md

[limiting corporate campaigning]: https://github.com/kubernetes/steering/blob/master/elections.md#limiting-corporate-campaigning
[pledge to recuse]: https://github.com/kubernetes/steering/blob/master/elections.md#steering-committee-and-election-officer-recusal

[Eligibility for candidacy]: https://github.com/knative/community/blob/master/mechanics/TOC.md#candidate-eligibility
[Eligibility for voting]: https://github.com/knative/community/blob/master/mechanics/TOC.md#candidate-eligibility#voter-eligibility
[Elekto]: https://elections.knative.dev
[file a voting exception]: https://elections.knative.dev/app/elections/2022-TOC/exception
