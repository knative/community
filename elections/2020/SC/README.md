---
title: "Knative Steering Committee Election - 2020 Voters Guide"
linkTitle: "2020 Voters Guide"
weight: 30
type: "docs"
---

* [Purpose](#purpose)
* [Background](#background)
* [Eligibiliy](#eligibility)
* [Exceptions](#exceptions-and-email-addresses)
* [Schedule](#schedule)
* [Election Officers](#election-officers)
* [Candidacy Process](#candidacy-process)
* [Voting Process](#voting-process)
* [Nominees](#nominees)

# 2020 Voters Guide - Knative Steering Committee Election

## Purpose

The role of this election is to fill out the two (2) seats due for election
this year on the [Knative Steering Committee]. Each elected member will serve
a two (2) year term.

## Background

This election will shape the future of Knative as a community and project.
The Knative Steering Committee (KSC) is
responsible for the general health of the Knative community, and this is the
first general election for that body.

## Eligibility

Please refer to the [SC Election Charter] for:

- [Eligibility for candidacy]
- [Eligibility for voting] and
  [Eligible voters](voters.md) list (not in charter)

## Exceptions and Email Addresses

As we do not have a reliable source of contributor email information, all
potential voters are strongly encouraged to provide us with an up-to-date email
address by filling out just the email portion of the [voter exception form].  This will
help prevent any delays in receiving your ballot.

Additionally, if your contributions to Knative are not completely represented
by devstats data, and thus do not appear in the [Eligible voters](voters.md) list,
 please file an exception request via the [voter exception form].  We believe that
the data only shows us half of the real Knative contributors, so do not hesitate
to request an exception if you are in the other half.

If you are eligible, but have not received your ballot on time, plese fill out
the [Ballot Replacement Form](https://bit.ly/knsc2020-ballot).

## Schedule

| Date         | Event                    |
| ------------ | ------------------------ |
| October 27     | Announcement of Election, call for nominations |
| November 9     | All candidate nominations due by 0000 UTC (5pm Pacific) |
| November 12     | Election Begins via email ballots |
| November 22    | Last call for exception & replacement ballots |
| November 29       | Election Closes by 0000 UTC (5pm Pacific) |
| December 3       | Announcement of Results at Knative Steering Meeting |

## Election Officers

In an effort to run a fair and transparent election, the following people
have been asked by the Steering Committee to run this election:

- [Josh Berkus](https://github.com/jberkus)
- [Dawn Foster](https://github.com/geekygirldawn)
- [Maria Cruz](https://github.com/macruzbar-zz)

You can reach us by emailing elections@knative.team

## Candidacy Process

**Nomination**

If you want to stand for election, open a PR against the
[knative/community repository](https://github.com/knative/community) to include
your candidate profile in the `/elections/2020/SC` folder, with your name as the
filename.  This profile should include:

* Your name
* Your company affiliation (employer or otherwise)
* Your contributions to Knative
* Why you are running

You can find [a sample template in the folder](./candidate-template.md).

Once you have created the PR, you may email knative-dev and/or knative-users *once* to let
people know about your candidacy and encourage endorsements as comments on the PR.
Please use something like for your email in order to encourage endorsements
in Github and not on the mailing list:

```
Fellow Knative community members,

My name is {your name} and I am running for Steering Committee.  You can read
my profile here: {link to PR}.  If you support my candidacy, please endorse me
by commenting on that PR.  Please don't reply to endorse me on this list.
```

After a candidate has met all election requirements, the Election Officers will
merge the profile PR.

If you want to nominate someone else, you may do so, but PLEASE talk to them
first.

**Campaigning**

Please refer to the [SC Election Charter] and understand
that we care deeply about [limiting campaigning]. The election
officers and members of the steering committee [pledge to recuse] themselves
from any form of electioneering.

You should be running as a "brand free" individual, based on your contribution
to the project as a member of this community, outside of whatever corporate
roles you may hold.

## Voting Process

Knative members in [voters.md] will receive a ballot via email. If you are
not on that list and feel you have worked on Knative in a way that is NOT
reflected in GitHub contributions, you can use the [voter exception form] to ask
to participate in the election.  Please also use that form just to update your
email address with us, as the one we have for you may be out of date.

Elections will be held using time-limited [Condorcet] ranking on [CIVS]
using the [Schulze](https://en.wikipedia.org/wiki/Schulze_method) method. The top vote getters will be elected to the open
seats.

You will be ranking your choices of the candidates with an option for
"no opinion". In the event of a tie, a coin will be flipped.

The election will open for voting according to the calendar above.
You will receive an email
to the address on file at the start of the election from
"Knative Election Officers (CIVS Poll Supervisor)" `<civs@cs.cornell.edu>`,
make sure it passes spam filters if necessary. Detailed
voting instructions will be addressed in email and the CIVS polling page.

If you do not receive your ballot by email within 48 hours of the election start,
please file a [ballot replacement request].

### Decision

The newly elected body will be announced in the Knative Steering Meeting on the
scheduled announcement date.

Following the meeting, the raw voting results and winners will be published on the
[Knative Blog].

For more information, definitions, and/or detailed election process, please refer to
the [SC Election Charter].

## Nominees

|    Name    | Organization/Company |  GitHub  |
|:----------:|:--------------------:|:--------:|
| [Arghya Sadhu](arghya-sadhu.md) | JP Morgan Chase |  [arghya88](https://github.com/arghya88) |
| [Carlos Santana](csantanapr.md) | IBM | [csantanapr](https://github.com/csantanapr) |
| [Sebastien Goasguen](goasguen.md) | TriggerMesh | [sebgoa](https://github.com/sebgoa) |
| [Jacques Chester](jacques-chester.md)| VMware | [jchesterpivotal](https://github.com/jchesterpivotal) |
| [Paul Morie](pmorie.md) | Red Hat | [pmorie](https://github.com/pmorie) |
| [Ville Aikas](vaikas.md) | VMware | [vaikas](https://github.com/vaikas) |

[Knative Steering Committee]: https://github.com/knative/community/blob/master/STEERING-COMMITTEE.md
[SC Election Charter]: https://github.com/knative/community/blob/master/mechanics/SC.md

[limiting corporate campaigning]: https://github.com/kubernetes/steering/blob/master/elections.md#limiting-corporate-campaigning

[Condorcet]: https://en.wikipedia.org/wiki/Condorcet_method
[CIVS]: http://civs.cs.cornell.edu/
[IRV method]: https://www.daneckam.com/?p=374

[Knative Blog]: https://knative.dev/blog/
[voter exception form]: https://bit.ly/knative-sc20-exception
[voters.md]: ./voters.md
[ballot replacement request]: https://bit.ly/knative-sc20-ballot

[Eligibility for candidacy]: https://github.com/knative/community/blob/master/mechanics/SC.md#candidate-eligibility
[Eligibility for voting]: https://github.com/knative/community/blob/master/mechanics/SC.md#candidate-eligibility#voter-eligibility
