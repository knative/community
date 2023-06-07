---
title: "Knative elections"
linkTitle: "Knative elections"
weight: 30
type: "docs"
---

# Knative Elections

This document outlines how to conduct Knative elections. See [TOC election process](../mechanics/TOC.md) and [Steering Committee election process](../mechanics/SC.md) for more information of how the committee decides when to have elections, eligibility for voting, eligibility for candidacy, maximal representation, etc.

## Process

1. Steering committee selects election officers.

2. Prepare the election repository

    * Make knative/community/elections/$YEAR-TOC or knative/community/elections/$YEAR-SC
    * Create an OWNERS file in the above directory with the election officers as approvers / reviewers.
    * Create the README.md in the above directory, this is the voter's guide
        * Copy over the voter's guide from the previous year. The voter's guide is the single source of truth for the election that year! All annoucements and notices should link to this document.
        * Update with new dates, candidates, and procedures (if necessary)
    * Create voters.yaml in the above directory
        * Initial list is generated via the eligibility for voting requirements.
        * Election officers will update this list as voters submit the voting exception form.
    * Create election.yaml and election_desc.md, which are used by Elekto
        * Copy the one from the previous election and update with new information about the election.
    * Create nomination-template.md
        * Copy from the previous election - no need to update.

3. Annouce voting schedule to the community

    * Should be mostly links to the voter guide, which contains links to other relevant details.
    * On knative-dev list, Knative Slack under #general, and twitter.

4. Executing the Election in Elekto

    * Elections will be held using [Elekto](https://elekto.dev/), an online voting tool created
      by CNCF intern Manish Sahani and administered by Josh Berkus. 
    * It relies on GitHub Oauth for access to ballots
    * More details can be found in the [Elekto documentation](https://elekto.dev/docs/)
    * Remember to send periodic reminders about key deadlines and to encourage people to vote.

5. Announcing the results

    * Reporting
        * Mail results of the election to the Steering Committee to discuss next steps.
        * Contact the candidates in advance of the announcement to let each of them know whether or not
          they were selected and ask them to keep it confidential until the announcement.
        * Steering Committee announces the results to the entire community at the end of the election
    * Push election results into results.md in the above directory _after_ the Steering Commmittee has announced the results

## Roles and Responsibilities:

### Technical Oversight Committee

- [Recuses themselves from public election activities][election-recusal]
- May vote
- May answer questions about general election specifics, ie:
  - Where do I find the schedule?
  - How do I vote?
- Will not answer questions about specific candidates, or anything that could be construed as endorsing, ie:
  - How is $candidate doing so far? (PS - we don't know anyway)
  - Who are your favorite candidates?


### Steering Committee

- Oversees the election, including appointing election officers, approving dates, and monitoring the process
- [Recuse themselves from public election activities][election-recusal] except those required to run the election
- May vote
- May answer questions about general election specifics, ie:
  - Where do I find the schedule?
  - How do I vote?
- Will not answer questions about specific candidates, or anything that could be construed as endorsing, ie:
  - How is $candidate doing so far? (PS - we don't know anyway)
  - Who are your favorite candidates?

### Election Officers
- [Recuse themselves from public election activities][election-recusal] except those required to run the election
- May vote
- May answer questions about general election specifics, ie:
  - Where do I find the schedule?
  - How do I vote?
- Will not answer questions about specific candidates, or anything that could be construed as endorsing, ie:
  - How is $candidate doing so far? (PS - we don't know anyway)
  - Who are your favorite candidates?
- Manage the voting process within Elekto
- Generate the voter guide and list of voters according to the criteria for that year's election
- Manage the exception process - Review applicants and add approved ones to the voter's list 
- Track candidates
- Monitor groups for nominations
- Update the community regularly
- Post deadlines and reminders during the election season
- It is impossible to see the results of the election until the election ends; for purposes of transparency with the community it is encouraged to release some statistics during the election (ie. “65% of the community has voted so far!”)

Diary for 2023 Knative TOC election can be found in this [Gist](https://gist.github.com/aliok/136be152fef14912b9a73eb753b3267b) for reference.

[election-recusal]: https://github.com/kubernetes/steering/blob/main/elections.md#steering-committee-and-election-officer-recusal

