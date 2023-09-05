# Roadmap Management for WG Leads

## Overview

All [working groups](../working-groups/WORKING-GROUPS.md) should maintain a current and up-to-date roadmap in a GitHub Project. It's recommended that a _single org-level project_ in the `knative` org should be used; for linking to issues in `knative-extensions`, cards with the URL of the extensions issue can be used.

Working Groups should link to their roadmaps in the Working Group documentation, and should work off the roadmap on a week-to-week or month-to-month basis.

The are multiple benefits of this process:

* Working Groups should already be using GitHub issues to track their upcoming work; adding project boards minimizes the amount of additional bookkeeping needed to produce a roadmap.
* Having a single process for all Working Groups makes it easier for new contributors to approach a WG and contribute.
* Org-level projects are editable via the web GUI for all org members, making the additional permissions required simple and light-weight.
* Having roadmaps which track which items are at which stage of implementation is helpful in [managing the flow of work](https://en.wikipedia.org/wiki/Kanban) for contributors, WG leads, and the [TOC](../TECHNICAL-OVERSIGHT-COMMITTE.md).

## Roadmap Columns and Guidance

Roadmaps should have the following columns. Additional columns can be added temporarily, but it should be possible to map them to these groupings during a TOC review:

### `In Progress` Column
These are issues / feature requests which are currently being implemented (code written) by contributors. The goal should be that contributors should be able to focus their efforts on _one_ of these issues at a time, and complete the issue efficiently without a lot of context switching.

#### Guidance:
- **Work in this column ought to be fleshed out enough that the implementation is straightforward.**
- Each issue in this column has at least one owner (marked as the GitHub "assignee". For large efforts, it's entirely reasonable to have multiple "assignees" if their efforts are focused on the issue.
- Each issue in this column links to something that organizes efforts (more on this below).

### `Ready To Work` Column
These are issues / feature requests which are sufficiently designed that the implementation should be straightforward, which are _staged for implementation_ when contributors finish "In Progress" work. It is the responsibility of Working Group leads to ensure that this column is managed so that contributors have a clear sense of what should be worked on next.

#### Guidance:
- **Work in this column ought to be fleshed out enough to enable implementation, but does not have someone currently working on it.**
- Issues in this column do not currently have an owner assigned. When an owner is assigned, they should be moved to the "In Progress" column.
- Each issue in this column links to something that organizes efforts (issues, other GH Project boards, Milestones, etc).


#### `In Design` Column
These are issues / feature requests where the problem is easily explained, but the solution to the problem is *not well understood*. Typically, for items in this column, a [feature track document](FEATURE-TRACKS.md) will be written to describe the implementation plan and flesh out the details and interactions with other components. Note that for straightforward changes without additional system interactions (for example, "allow field X validation to permit additional value Y"), Working Group leads may move items directly from the "Icebox" to the "Ready To Work" column.

#### Guidance:
- **Work in this column has clear acceptance criteria, but the path to achivement of the acceptance criteria is either unclear or controversial.**
- Items in this column _should_ have an owner (assignee) who is working on designing and planning the implementation. That assignee _should_ link an in-progress document as part of moving this to the "In Design" column.
- May need a [Feature Track](FEATURE-TRACKS.md) or other planning document.
- At the acceptance of the Feature Track or other documentation, there should be:
  * A written document explaining the plan, _documented in the issue_. This could be an issue update for smaller or controversial topics, or a link to a longer separate document.
  * The plan should be generally agreeable to the Working Group -- Working Group Leads are responsible for determining that this agreement has been reached
  Once these have been completed, the issue should be moved to "Ready To Work".
- Each issue in this column links to something that organizes efforts (issues, other GH Project boards, Milestones, etc).

#### `Icebox` Column
This is the intake column for new issues and feature requests which haven't been evaluated yet. This is _also_ the column for work which is known to be large and/or controversial, and which the Working Group is not yet ready to tackle. (Ideally, most Working Groups would have fewer than three large designs under discussion at a time.)

#### Guidance:
- **Work in this column does not need to be fleshed out.**
- These items should proceed to "Ready To Work" or "In Design", depending on their complexity. It is the responsibility of the Working Group leads to manage this process.

## Additional considerations:

See https://github.com/knative/community/issues/746 for the original discussion / impetus of this process.

* These guidelines may be subject to refinement and change in the future. The [TOC](../TECH-OVERSIGHT-COMMITTEE.md) will be responsible for updates and communication of process changes. Working Group leads _are_ encouraged to experiment with the process within the general guidelines laid out above.

* The TOC does not yet have a recommended process for handling issues which have reached "In Progress" but have some type of time-based delay (e.g. feature is released in "alpha", needs to be progressed to "beta" after one release). One option is to create a new "Ready To Work" issue like `Post-1.3: promote $FEATURE to Beta` and close the original issue, but this is up to Working Group leads to manage.

* TOC reviews will generally focus on management and possibly coordination implications of work items, and will generally rely on the judgement of the working group in terms of which items to complete first.
