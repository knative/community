This document describes the set of roles individuals might have within the
Knative community, the requirements of each role, and the privileges that each
role grants. Community members generally start at the first levels of the
"ladder" and advance up it as their involvement in the project grows.  Our
project members are happy to help you advance along the contributor ladder.

- [Role Summary](#role-summary)
- [Contributor](#contributor)
- [Member](#member)
- [Approver](#approver)
- [Working Group Lead](#working-group-lead)

See also [notes](#notes) on roles.

## Role Summary

The following table lists the roles we use within the Knative community. The
table describes general responsibilities expected by individuals in each role.
A detailed description of the responsibilities and requirements for each role
is provided below the table.

<table>
  <thead>
    <tr>
    <th>Role</th>
    <th>Responsibilities</th>
    </tr>
  </thead>

  <tr>
    <td>Contributor</td>
    <td>Follow the CNCF Code of Conduct and contribute to the project</c>
  </tr>

  <tr>
    <td>Member</td>
    <td>Regular active contributor in the community</td>
  </tr>

  <tr>
    <td>Approver</td>
    <td>Approve PRs from Members and Contributors</td>
  </tr>

  <tr>
    <td>Working Group Lead</td>
    <td>Responsible for the health and growth of their functional area or subproject</td>
  </tr>
</table>

Please note that repos in the knative-extensions GitHub org are a reflection of
the WG leads. These leads are free to use their judgement to set the bar for
what is required to become an approver and/or lead for knative-extensions repos
they are responsible for.

## Contributor

This role describes people who have just started contributing, or who
contribute occasionally but don't participate in project governance or
have defined responsibilities.

A Contributor contributes directly to the project and adds value to it.
Contributions need not be code. People at the Contributor level may be new
contributors, or they may only contribute occasionally.
### Requirements

Contributors should be familiar with the project and its processes, and should
contribute to the project in one or more of the following ways.

- Report and sometimes resolve issues
- Occasionally submit PRs
- Contribute to the documentation
- Show up at meetings
- Submit feedback on issues and PRs
- Test releases and patches and submit reviews
- Run or help run events
- Promote the project in public

## Member

A Member is an established Contributor who regularly participates in the
project. Members have privileges in project repositories,
and as such are expected to act in the interests of the whole project.

Members are expected to demonstrate their adherence to the
principles in this document, familiarity with project organization, roles,
policies, procedures, conventions, etc., and technical and/or writing ability.

Members are continuously active contributors in the community. They can have
issues and PRs assigned to them, participate in working group meetings, and
pre-submit tests are automatically run for their PRs. Members are expected to
remain active contributors to the community.

All members are encouraged to help with code reviews, although each PR
must be approved by an official [Approver](#approver).

When reviewing, members should focus on code quality and correctness, including
testing and factoring. Members might also review for more holistic issues, but
this is not a requirement.

### Requirements

- Member of the Knative GitHub org
  - Create a PR adding yourself as a Member to
[knative.yaml](../main/peribolos/knative.yaml) and/or to
[knative-extensions.yaml](../main/peribolos/knative-extensions.yaml)
  - After the PR is merged you will receive an invite that you must accept to become a Member
  - If you would like to work towards becoming an Approver, open a second PR and add yourself as a `reviewer` on the repositories to which you will be contributing.

- Has made multiple contributions to the project or community. Contributions
  might include, but are not limited to:

  - Authoring and reviewing PRs on GitHub
  - Submitting and commenting on issues on GitHub
  - Triaging and resolving issues on GitHub
  - Contributing to working group or community discussions

- Subscribed to
  [knative-dev@googlegroups.com](https://groups.google.com/forum/#!forum/knative-dev)

- Actively contributing to 1 or more areas

- Nominated by two Knative Members, at least one of whom does not work for the same employer

### Responsibilities and privileges

- Responsive to issues and PRs assigned to them

- Active owner of contributions (unless ownership is explicitly
  transferred)

  - Contributions are well tested

  - Tests consistently pass

  - Addresses bugs or issues discovered after a contribution

Members who frequently contribute should proactively perform
reviews and may work towards becoming an Approver for the area that they
are active in.

## Approver

Code approvers are able to both review and approve code contributions. While
code review is focused on code quality and correctness, approval is focused on
holistic acceptance of a contribution including: backward / forward
compatibility, adhering to API and flag conventions, subtle performance and
correctness issues, interactions with other parts of the system, etc. Approver
status is scoped to a part of the codebase.

### Requirements

The following apply to the part of the repository for which one would be an
approver in an OWNERS file:

- Add yourself as a `writer` for the repositories you are responsible for
  in [knative.yaml](../main/peribolos/knative.yaml) and/or
  [knative-extensions.yaml](../main/peribolos/knative-extensions.yaml). After
  this you should be an `approver` in the OWNERS file.

- Reviewer of other people's contributions for about 3 months

- Should have reviewed or contributed to non-trivial PRs

- Nominated by a Working Group Lead with no objections from other leads

### Responsibilities and privileges

The following apply to the part of the repository for which one would be an
approver in an OWNERS file:


- Demonstrate sound judgment

* Responsible for project quality control via [reviews](./REVIEWING.md)

  - Focus on holistic acceptance of contribution such as dependencies with other
    features, backward / forward compatibility, API and flag definitions, etc.

* Expected to be responsive to review requests as per
  [community expectations](./REVIEWING.md)

* Mentor new contributors and project members

* Approve contributions for acceptance

Approvers are also expected to participate in community contact rotations
([Serving](
https://github.com/knative/serving/blob/main/support/COMMUNITY_CONTACTS.md)
or [Eventing](
https://github.com/knative/eventing/blob/main/support/COMMUNITY_CONTACTS.md))
to support users and keep test quality high, as well as release leads
[rotation](https://github.com/knative/pkg/blob/main/RELEASE-LEADS.md) to
shepherd Knative releases.

## Working Group Lead

Working group leads are approvers of an entire
area that have demonstrated good judgement and responsibility. Tech leads accept
design proposals and approve design decisions for their area of ownership, and
are responsible for the overall technical health of their functional area.

### Requirements

For existing working groups:

- Recognized as having expertise in the group’s subject matter

- Approver for a relevant part of the repository for about 3 months

- Should have reviewed or contributed to non-trivial PRs

- Should be an `approver` in the OWNERS file

- Sponsored by the Technical Oversight Committee

- Recognized in the project as an active participant

- Recognized as having expertise in the group’s subject matter


### Responsibilities and privileges

The following apply to the area / component for which one would be a lead.

- Run their working group as explained in the
  [Working Group Processes](./mechanics/WORKING-GROUP-PROCESSES.md).

  - Meetings. Prepare the agenda and run the regular working group meetings.

  - Notes. Ensure that meeting notes are kept up to date. Provide a link to the
    recorded meeting in the notes. The lead may delegate note-taking duties to
    the scribe.

  - Roadmap. Establish and maintain a roadmap for the working group outlining
    the areas of focus for the working group over the next 6 months.

  - Report. Report current status to the TOC meeting every 6 weeks.

- Expected to work to holistically maintain the health of the project through:

  - Reviewing PRs

  - Fixing bugs

  - Identifying needed enhancements / areas for improvement / etc.

  - Be aware of and work to reduce technical debt where it may exist

- Design/proposal approval authority over the area / component, though
  escalation to the technical oversight committee is possible

- Holistic responsibility for their working group's [feature
  tracks](./mechanics/FEATURE-TRACKS.md): review, tracking, health, and execution.

- Perform issue triage on GitHub

- Apply/remove/create/delete GitHub labels and milestones

- Write access to repo (assign issues/PRs, add/remove labels and milestones,
  edit issues and PRs, edit wiki, create/delete labels and milestones)


  - Expected to respect OWNERS files approvals and use
    [standard procedure for merging code](./REVIEWING.md#merging-prs)

- Mentoring and guiding Approvers, Members, and Contributors

- Expected to work to holistically maintain the health of the working group through:

  - Being a good role model

  - Be an advocate for the working group inside and outside of the community

  - Foster a welcoming and collegial environment

  - Mentoring and guiding approvers, members, and new contributors.

# Notes

Within this section "manager" refers to a Member who is a Working Group Lead or
Approver.

- Initial managers are defined at the founding of the WG or Subproject as part
  of the acceptance of that WG or Subproject.
- Managers SHOULD remain active and responsive in their Roles.
- Managers MUST be community members to be eligible to hold a leadership role within a SIG.
- Managers taking an extended leave of 1 or more months SHOULD coordinate with
  other managers to ensure the role is adequately staffed during the leave.
- Managers going on leave for 1-3 months MAY work with other managers to suggest
  a replacement using the normal process
- Managers of a role SHOULD remove any other managers that have not communicated
  a leave of absence and either cannot be reached for more than 1 month or are
  not fulfilling their documented responsibilities for more than 1 month.
  - This MAY be done through a super-majority vote of managers, or if there are
    not enough active managers to get a super-majority of votes cast, then
    removal MAY occur through exception process to the TOC.  The PR removing the
    manager should be open for at least 72 hours.
  - Prior to voting to remove a manager, leads SHOULD reach out to the affected
    manager and see if they need to take a leave.
- Membership disagreements MAY be escalated to the WG leads. WG lead membership
  disagreements MAY be escalated to the TOC.
- Managers MAY decide to step down at anytime and nominate a replacement who
  will be approved through the regular process for that role.

It is the hope of the Steering Committee that effective communication will make
the use of these rules something that happens under exceptional circumstances
only. In circumstances where it is unavoidable, these are presented so that the
process is clear.

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
