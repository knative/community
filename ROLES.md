---
title: "Knative community roles"
linkTitle: "Community roles"
weight: 55
type: "docs"
---

This document describes the set of roles individuals might have within the
Knative community, the requirements of each role, and the privileges that each
role grants.

- [Role Summary](#role-summary)
- [Collaborator](#collaborator)
- [Member](#member)
- [Approver](#approver)
- [Lead](#lead)
- [Administrator](#administrator)

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
        <p>Sponsored by two members</p>
        <p>Has made multiple contributions to the project</p>
    </td>
    <td>
        <p>Member of the GitHub Knative org</p>
        <p>Member of the Knative Slack workspace</p>
        <p>Edit permission on the Knative Team drive</p>
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
        on one or more repos allowing PRs to be merged.
    </td>
    <td>GitHub Directory</td>
  </tr>

  <tr>
    <td>Lead</td>
    <td>
        <p>Set priorities for a functional area and approve proposals</p>
        <p>Triage incoming issues, set milestones, repo labels</p>
        <p>Run their working group</p>
    </td>
    <td>Sponsored by the technical oversight committee as documented
        <a href="./mechanics/WORKING-GROUP-PROCESSES.md">here</a>.
    </td>
    <td>Write permissions on one or more repos allowing issues to be manipulated.</td>
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

- Has made multiple contributions to the project or community. Contributions
  might include, but are not limited to:

  - Authoring and reviewing PRs on GitHub.

  - Filing and commenting on issues on GitHub.

  - Contributing to working group or community discussions.

- Subscribed to
  [knative-dev@googlegroups.com](https://groups.google.com/forum/#!forum/knative-dev).

- Actively contributing to 1 or more areas.

- 1 sponsor must have the approver role.

  - Done by adding GitHub user to Knative organization.

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

- Reviewed or merged at least 30 PRs to the codebase.

- Nominated by an area lead (with no objections from other leads).

  - Done through PR to update an OWNERS file.

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

* Mentor members and contributors.

* Might approve code contributions for acceptance.

## Lead

Working group leads, or just ‘leads’, are approvers of an entire area that have
demonstrated good judgement and responsibility. Leads accept design proposals
and approve design decisions for their area of ownership.

### Requirements

Getting to be a lead of an existing working group:

- Recognized as having expertise in the group’s subject matter.

- Approver for some part of the codebase for at least 3 months.

- Member for at least 1 year or 50% of project lifetime, whichever is shorter.

- Primary reviewer for 20 substantial PRs.

- Reviewed or merged at least 50 PRs.

- Sponsored by the technical oversight committee.

Additional requirements for leads of a new working group:

- Originally authored or contributed major functionality to the group's area.

- An approver in the OWNERS file for the group’s code.

### Responsibilities and privileges

The following apply to the area / component for which one would be an owner.

- Run their working group as explained in the
  [Working Group Processes](./mechanics/WORKING-GROUP-PROCESSES.md).

- Design/proposal approval authority over the area / component, though
  escalation to the technical oversight committee is possible.

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

  - Mentoring and guiding approvers, members, and contributors.

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
