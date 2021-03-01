---
title: "Knative community roles"
linkTitle: "Community roles"
weight: 55
type: "docs"
aliases:
  - /contributing/roles/
---

This document describes the set of roles individuals might have within the
Knative community, the requirements of each role, and the privileges that each
role grants.

- [Role Summary](#role-summary)
- [Member](#member)
- [Approver](#approver)
- [Working Group Technical Lead](#working-group-technical-lead)
- [Working Group Execution Lead](#working-group-execution-lead)
- [Scribe](#scribe)

See also [notes](#notes) on roles.

## Role Summary

The following table lists the roles we use within the Knative community. The
table describes:

- General responsibilities expected by individuals in each role
- Requirements necessary to join or stay in a given role
- How the role manifests in terms of permissions and privileges.

<table>
  <thead>
    <tr>
    <th>Role</th>
    <th>Responsibilities</th>
    <th>Requirements</th>
    <th>Privileges</th>
    <th>Scope</th>
    </tr>
  </thead>

  <tr>
    <td>Member</td>
    <td>Regular active contributor in the community</td>
    <td>
        <p>Has made multiple contributions to the project</p>
    </td>
    <td>
        <p>Member of the GitHub Knative org</p>
        <p>Member of the Knative Slack workspace</p>
        <p>Edit access to team drive</p>
    </td>
    <td>GitHub Organization</td>
  </tr>

  <tr>
    <td>Approver</td>
    <td>
        <p>Approve contributions from other members</p>
    </td>
    <td>Highly experienced and active reviewer and contributor to an area</td>
    <td>Entry in one or more OWNERS files in GitHub, and write permissions
        on one or more repos allowing PRs to be merged
    </td>
    <td>GitHub Directory</td>
  </tr>

  <tr>
    <td>Working Group Technical Lead</td>
    <td>
        <p>Set priorities for a functional area and approve proposals</p>
        <p>Triage incoming issues, set milestones, repo labels</p>
        <p>Roadmap alignment with top-level backlog</p>
        <p>Mentor new contributors, project members, and approvers</p>
        <p>Succession - identifying next steps for members of the working group</p>
        <p>Responsible for technical health of their functional area</p>
    </td>
    <td>Sponsored by the technical oversight committee as documented
        <a href="./mechanics/WORKING-GROUP-PROCESSES.md">here</a>
    </td>
    <td>Write permissions on one or more repos allowing issues to be manipulated</td>
    <td>Working Group</td>
  </tr>

  <tr>
    <td>Working Group Execution Lead</td>
    <td>
        <p>Run their working group: Meetings, notes, roadmap, report</p>
        <p>Responsible for the holistic health of the <i>working group</i></p>
        <p>Work organization, planning, high-level execution</p>
        <p>Triage incoming issues, set milestones, repo labels</p>
        <p>Ensure all required skills are present within the working group</p>
        <p>Mentor new contributors, project members, and approvers</p>
    </td>
    <td>Sponsored by the technical oversight committee as documented
        <a href="./mechanics/WORKING-GROUP-PROCESSES.md">here</a>
    </td>
    <td>Write permissions on one or more repos allowing issues to be manipulated</td>
    <td>Working Group</td>
  </tr>

  <tr>
    <td>Scribe</td>
    <td>
        <p>Ensure important information is represented in working group notes</p>
        <p>Post WG recordings to team drive</p>
    </td>
    <td>Sponsored by a working group execution or technical lead.</td>
    <td>Write permissions to team drive and team calendar</td>
    <td>Working Group</td>
  </tr>

</table>

## Member

Established community members are expected to demonstrate their adherence to the
principles in this document, familiarity with project organization, roles,
policies, procedures, conventions, etc., and technical and/or writing ability.

Members are continuously active contributors in the community. They can have
issues and PRs assigned to them, participate in working group meetings, and
pre-submit tests are automatically run for their PRs. Members are expected to
remain active contributors to the community.

All members are encouraged to help with the code review burden, although each PR
must be reviewed by an official [Approver](#approver).

When reviewing, members should focus on code quality and correctness, including
testing and factoring. Members might also review for more holistic issues, but
this is not a requirement.

### Requirements

- Member of the Knative github org. Create a PR adding you to
[knative.yaml](../peribolos/knative.yaml) and/or to
[knative-sandbox.yaml](../peribolos/knative-sandbox.yaml) which after when
merged will send you an invite that you have to accept to become a member
of these organizations.

- Has made multiple contributions to the project or community. Contributions
  might include, but are not limited to:

  - Authoring and reviewing PRs on GitHub.

  - Filing and commenting on issues on GitHub.

  - Contributing to working group or community discussions.

- Subscribed to
  [knative-dev@googlegroups.com](https://groups.google.com/forum/#!forum/knative-dev).

- Actively contributing to 1 or more areas.

### Responsibilities and privileges

- Responsive to issues and PRs assigned to them.

- Active owner of code they have contributed (unless ownership is explicitly
  transferred).

  - Code is well tested.

  - Tests consistently pass.

  - Addresses bugs or issues discovered after code is accepted.

Members who frequently contribute code are expected to proactively perform code
reviews and work towards becoming an approver for the area that they are active
in.

## Approver

Code approvers are able to both review and approve code contributions. While
code review is focused on code quality and correctness, approval is focused on
holistic acceptance of a contribution including: backward / forward
compatibility, adhering to API and flag conventions, subtle performance and
correctness issues, interactions with other parts of the system, etc. Approver
status is scoped to a part of the codebase.

### Requirements

The following apply to the part of the codebase for which one would be an
approver in an OWNERS file:

- Reviewer of the codebase for at least 3 months.

- Primary reviewer for at least 10 substantial PRs to the codebase.

  - One path for getting the necessary reviews is to add yourself to the
    `reviewers` section of the OWNERS file. Note that this does not give you any
    additional privileges. By having yourself listed in this section in OWNERS
    file means that you will get PRs assigned to you for code review. Getting
    added to `reviewers` section is at the discretion of an approver after
    enough evidence of quality contributions.

- Reviewed at least 30 PRs to the codebase.

- Nominated by an a WG lead (with no objections from other leads).

### Responsibilities and privileges

The following apply to the part of the codebase for which one would be an
approver in an OWNERS file:

- Approver status can be a precondition to accepting large code contributions.

- Demonstrate sound technical judgment.

* Responsible for project quality control via [code reviews](./REVIEWING.md).

  - Focus on holistic acceptance of contribution such as dependencies with other
    features, backward / forward compatibility, API and flag definitions, etc.

* Expected to be responsive to review requests as per
  [community expectations](./REVIEWING.md).

* Mentor new contributors and project members.

* Approve code contributions for acceptance.

Approvers are also expected to participate in community contact rotations
([Serving](
https://github.com/knative/serving/blob/master/support/COMMUNITY_CONTACTS.md)
or [Eventing](
https://github.com/knative/eventing/blob/master/support/COMMUNITY_CONTACTS.md))
to support users and keep test quality high, as well as release leads
[rotation](https://github.com/knative/pkg/blob/master/RELEASE-LEADS.md) to
shepherd Knative releases.

# Working Group Leadership Roles

We differentiate here between two key roles, WG Execution Lead and WG Technical
Lead. In small working groups, often these roles will be performed by one
person; in larger working groups it might be best to have different individuals
performing these roles. There is no 'formula', it is up to the working group
leadership to determine what is best for their working group.

## Working Group Technical Lead

Working group technical leads, or just ‘tech leads’, are approvers of an entire
area that have demonstrated good judgement and responsibility. Tech leads accept
design proposals and approve design decisions for their area of ownership, and
are responsible for the overall technical health of their functional area.

### Requirements

Getting to be a tech lead of an existing working group:

- Recognized as having expertise in the group’s subject matter.

- Approver for a relevant part of the codebase for at least 3 months.

- Member for at least 6 months.

- Primary reviewer for 20 substantial PRs.

- Reviewed or merged at least 50 PRs.

- Sponsored by the technical oversight committee.

Additional requirements for leads of a new working group:

- Originally authored or contributed major functionality to the group's area.

### Responsibilities and privileges

The following apply to the area / component for which one would be an owner.

- Design/proposal approval authority over the area / component, though
  escalation to the technical oversight committee is possible.

- Technical review of [feature tracks](./mechanics/FEATURE-TRACKS.md).

- Perform issue triage on GitHub.

- Apply/remove/create/delete GitHub labels and milestones.

- Write access to repo (assign issues/PRs, add/remove labels and milestones,
  edit issues and PRs, edit wiki, create/delete labels and milestones).

- Capable of directly applying lgtm + approve labels for any PR.

  - Expected to respect OWNERS files approvals and use
    [standard procedure for merging code](./REVIEWING.md#merging-prs).

- Expected to work to holistically maintain the health of the project through:

  - Reviewing PRs.

  - Fixing bugs.

  - Identifying needed enhancements / areas for improvement / etc.

  - Execute pay-down of technical debt.

- Mentoring and guiding approvers, members, and new contributors.

## Working Group Execution Lead

Working group execution leads, or just ‘execution leads’, are responsible for
the overall health and execution of the working group itself. Execution leads
work with tech leads to ensure that the working group is making progress toward
its goals, is aligned with the project roadmap, etc. The execution lead may also
be the tech lead in a smaller working group, but they are distinct roles.

### Requirements

- Participant in the working group for at least 3 months, for example as scribe
  or approver.

- Recognized as having expertise in the group’s subject matter.

- Sponsored by the technical oversight committee.

### Responsibilities and privileges

- Run their working group as explained in the
  [Working Group Processes](./mechanics/WORKING-GROUP-PROCESSES.md).

  - Meetings. Prepare the agenda and run the regular working group meetings.

  - Notes. Ensure that meeting notes are kept up to date. Provide a link to the
    recorded meeting in the notes. The lead may delegate note-taking duties to
    the scribe.

  - Roadmap. Establish and maintain a roadmap for the working group outlining
    the areas of focus for the working group over the next 6 months.

  - Report. Report current status to the TOC meeting every 6 weeks.

- Holistic responsibility for their working group's [feature
  tracks](./mechanics/FEATURE-TRACKS.md): tracking, health, and execution.

- Perform issue triage on GitHub.

- Apply/remove/create/delete GitHub labels and milestones.

- Write access to repo (assign issues/PRs, add/remove labels and milestones,
  edit project, issues, and PRs, edit wiki, create/delete labels and milestones).

- Expected to work to holistically maintain the health of the working through:

  - Being a good role model

  - Be an advocate for the working group inside and outside of the community

  - Foster a welcoming and collegial environment

  - Mentoring and guiding approvers, members, and new contributors.

## Scribe

One of the most underrated roles in open source projects is the role of note
taker. The importance and value of this role is frequently overlooked and
underestimated. Since one of the core project values is transparency, we have an
explicit scribe role to recognize these types of contributions. Working group
scribes assist the Working Group leads with the mechanical processes around
Working Group meetings.

### Requirements

- Participant in the working group for at least 1 month.

- Pattern of attendance and note-taking during working group meetings and one-offs.

- Sponsored by a working group execution or technical lead.

### Responsibilities and privileges

- Attend working group meetings and one-offs whenever possible.

- Ensure that important information from meetings makes it into the WG notes.

- Post WG recordings to the team drive.

# Notes

Within this section "manager" refers to a member who is an Execution Lead, Tech
Lead, Approver or Scribe. (this is different from a WG or Organization Member).

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

It is the hope of the steering committee that effective communication will make
the use of these rules something that happens under exceptional circumstances
only. In circumstances where it is unavoidable, these are presented so that the
process is clear.

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
