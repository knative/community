---
title: "Knative Steering Committee Election - 2021 Voters Guide"
linkTitle: "2021 Voters Guide"
weight: 30
type: "docs"
---

# 2021 Voters Guide - Knative Steering Committee Election

## Purpose

The role of this election is to fill out the three (3) seats due for election
this year on the [Knative Steering Committee]. Each elected member will serve a
two (2) year term.

## Background

This election will shape the future of Knative as a community and project. The
Knative Steering Committee (SC) is responsible for the general health of the
Knative community. Seats up for the election are the remaining bootstrap seats
and in addition to contributor seats, we're introducing one (1) end-user seat to
bring more diversity into the SC representation, which will be elected by the
SC after the election has been completed.

## Eligibility

Please refer to the [SC Election Charter] for:

- [Eligibility for candidacy]
- [Eligibility for voting]

The list of eligible voters is stored (TODO: https://github.com/knative/community/issues/772)[in the community repository](./voters.yaml), but it is easier
to check if you are eligible by logging into [Elekto]. See Voting Process below for more details.

## Schedule

| Date         | Event                    |
| ------------ | ------------------------ |
| October 25   | Announcement of Election, call for nominations, exceptions |
| November 8   | All candidate nominations due by 0000 UTC (5pm Pacific) |
| November 10  | Election Begins via Elekto UI |
| November 19  | Voter exception requests due by 0000 UTC (5pm Pacific) |
| November 29  | Election Closes by 0000 UTC (5pm Pacific) |
| December 2   | Announcement of Results |

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
your candidate profile in the `/elections/2021-SC` folder, with the following
filename format:

```
candidate-githubid.md
```

This profile should include:

* Your name
* Your company affiliation (employer or otherwise)
* Your contributions to Knative
* Why you are running

You can find [a sample template in the folder](./nomination-template.md).

Once you have created the PR, you may email knative-dev and/or knative-users
*once* to let people know about your candidacy and encourage endorsements as
comments on the PR. Please use something like this for your email in order to
encourage endorsements in Github and not on the mailing list:

```
Fellow Knative community members,

My name is {your name} and I am running for Steering Committee.  You can read
my profile here: {link to PR}.  If you support my candidacy, please endorse me
by commenting on that PR.  Please do not reply to endorse me on this list.
```

After a candidate has met all election requirements, the Election Officers will
merge the profile PR.

If you want to nominate someone else, you may do so, but PLEASE talk to them
first.

**Campaigning**

Please refer to the [SC Election Charter] and understand
that we care deeply about [limiting corporate campaigning]. The election
officers and members of the steering committee [pledge to recuse] themselves
from any form of electioneering.

You should be running as a "brand free" individual, based on your contribution
to the project as a member of this community, outside of whatever corporate
roles you may hold.

## Voting Process

Elections will be held using [Elekto], an online voting tool created by CNCF 
intern Manish Sahani.  Elekto has one critical advantage over CIVS, which is
that it relies on GitHub Oauth to log you in to vote, instead of relying on 
email. More details on voting are in the [Elekto documentation].

Thus, when you go to [Elekto] you will be prompted to log in your GitHub account.
Please do, so, and then click on "Explore Election" to look at the list of 
elections.  From there you can click on the "2021 Knative SC Election."

The election page will, among other things, tell you if you are eligible to vote,
via a button at the top of the screen. Due to limitations in our contributor
data, many contributors may have been unfairly missed; if you are one of these,
please [file a voting exception] by November 8th 2021 so that we can enfranchise
you.

As candidates file their candidate statements in the community repo, they will
become visible in the [Elekto] UI.  You may click through to any candidate
to see their profile.

Once the vote begins, you will be able to rank the candidates in the order of
your preference, and submit your ballot.  When you submit, you will be offered
a chance to set a password, which is required if you want the ability to return
and re-cast your ballot before November 29th.

Employer diversity is encouraged, and thus maximal representation will be
enforced as spelled out in the [SC Election Charter].

Elekto is being run on an OpenShift instance operated by Red Hat's Open Source 
Practice Office.  Individual ballot data is encrypted, and not retrievable by 
anyone except the voter, or in aggregate form.

### Decision

The newly elected body will be announced in the Knative Steering Meeting on the
scheduled announcement date.

Following the meeting, the raw voting results and winners will be published on the
[Knative Blog].

For more information, definitions, and/or detailed election process, please refer to
the [SC Election Charter].

[Knative Steering Committee]: https://github.com/knative/community/blob/main/STEERING-COMMITTEE.md
[SC Election Charter]: https://github.com/knative/community/blob/main/mechanics/SC.md

[limiting corporate campaigning]: https://github.com/kubernetes/steering/blob/master/elections.md#limiting-corporate-campaigning
[pledge to recuse]: https://github.com/kubernetes/steering/blob/master/elections.md#steering-committee-and-election-officer-recusal

[Elekto]: https://elections.knative.io
[Elekto documentation]: 

[Knative Blog]: https://knative.dev/blog/
[voters.md]: ./voters.yaml

[Eligibility for candidacy]: https://github.com/knative/community/blob/main/mechanics/SC.md#candidate-eligibility
[Eligibility for voting]: https://github.com/knative/community/blob/main/mechanics/SC.md#candidate-eligibility#voter-eligibility
