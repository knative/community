---
title: "Knative elections"
linkTitle: "Knative elections"
weight: 30
type: "docs"
---

# Knative Elections

This document outlines how to conduct a Knative Technical Oversight Committee election. See [TOC election process](../mechanics/TOC.md) for more information of how the committee decides when to have elections, eligibility for voting, eligibility for candidacy, maximal representation, etc.

## Process

1. Steering committee prepare the election repository

    * Make knative/community/elections/$YEAR
    * Make knative/community/elections/$YEAR/README.md, this is the voter's guide
        * Copy over the voter's guide from the previous year. The voter's guide is the single source of truth for the election that year! All annoucements and notices should link to this document.
        * Update with new dates, candidates, and procedures (if necessary)
    * Make knative/community/elections/$YEAR/voters.md
        * Initial list is generated via the eligibility for voting requirements.
        * Steering will update this list as voters submit the voting exemption form.

1. Annouce voting schedule to the community

    * Should be mostly links to the voter guide and the TOC charter.
    * On knative-dev list, Knative Slack under #general, and twitter.

3. Executing the Election in CIVS

    * Use [CIVS](http://civs.cs.cornell.edu/civs_create.html) to create the election, which CIVS calls a poll. Once you send out the ballots you cannot UNSEND the emails, ensure everything in the form is correct!
    * Name of the poll - Knative TOC Election for $YEAR”
    * Name of supervisor - “Knative Steering Committee”
    * Email - elections@knative.team
    * Date and Time: Write in the date and time the election will stop. This field is not programmatic, the election is stopped by hand, so you can write this in plain text
    * Description: Use the following text, modify it for either 3 or 2 positions, depending on the amount of open seats:

        This election is to nominate the Technical Oversight Committee for the Knative project. Order the candidates by preference, the top $NUMBER candidates will be selected. Please see the voter's guide for more information.  PLEASE NOTE: "No opinion" is also a voting option if you do not feel comfortable ranking every single candidate

    * Add the candidate list to the form
    * How many choices will win: This number needs to be set to the amount of open seats of a given election (and updated in the description)
    * More options, check the boxes for:
        * Do not release results to all voters
        * Enable detailed ballot reporting
        * Allow voters to select “no opinion” for some choices
        * Enforce proportional representation
    * Click create poll, this will send elections@knative.team an email with instructions
    * It will send you a link to “Poll Control”, bookmark this generated page as this is where you will add voters and also resend ballots to people if their ballot gets lost or filtered
    * This page is where the “Start Poll” and “Stop Poll” buttons are, start the poll
    * WARNING: This is the point of no return:
    * Paste in the registered voters and click add voters
        * It will mail the ballots to the participants
        * It does duplicate detection so multiple entries are fine
        * This might take a while and the web page will not update, this has taken up to 10m in the past as it's sending each ballot. Don't panic or refresh the page
    * Leave the poll open for the duration of voting
        * Remember to send a 24 hour reminder before closing the poll
        * Click "Stop poll" at the end of the election, check the previously generated URL that CIVS mailed you when you started the poll
        * Select "Condorcet IRV" on the right hand side of the page to select the results method
    * Reporting
        * Mail results of the election to the Steering Committee members and incumbent TOC members who are not running for election
        * Steering Committee announces the results to the entire community at once at the end of the election
    * Push election results into knative/community/elections/$YEAR/results.md _after_ the Steering Commmittee has announced the results

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

- Runs the election, including setting dates and executing the voting
- [Recuse themselves from public election activities][election-recusal] except those required to run the election
- May vote
- May answer questions about general election specifics, ie:
  - Where do I find the schedule?
  - How do I vote?
- Will not answer questions about specific candidates, or anything that could be construed as endorsing, ie:
  - How is $candidate doing so far? (PS - we don't know anyway)
  - Who are your favorite candidates?
- Generate the voter guide and list of voters according to the criteria for that year's election
- Generate exemption form for non-code contributors to apply for voting
  - Review and commit applicants to approved voter's list at least once a week until the election begins
- Track candidates
- Monitor groups for nominations
  - Keep track of nominees in a spreadsheet
  - All nominations are conducted in the public, so sharing this sheet during the nomination process is encouraged
- Update the community regularly
- Post deadlines and reminders during the election season
- Reissue ballots from CIVS to voters who might have not received their ballot.
- Guard the privacy of the email addresses of voters
- It is impossible for the steering committee to see the results of the election until the election ends; for purposes of transparency with the community it is encouraged to release some statistics during the election (ie. “65% of the community has voted so far!”)


[election-recusal]: https://github.com/kubernetes/steering/blob/master/elections.md#steering-committee-and-election-officer-recusal
