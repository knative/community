<!-----
NEW: Check the "Suppress top comment" option to remove this info from the output.

Conversion time: 0.849 seconds.


Using this Markdown file:

1. Paste this output into your source file.
2. See the notes and action items below regarding this conversion run.
3. Check the rendered output (headings, lists, code blocks, tables) for proper
   formatting and use a linkchecker before you publish this page.

Conversion notes:

* Docs to Markdown version 1.0β29
* Thu Aug 20 2020 16:30:02 GMT-0700 (PDT)
* Source doc: Evan Proposal: Knative Migration-to-Core Process

ERROR:
undefined internal link to this URL: "#heading=h.xtzluz7y0l31".link text: migration to core
?Did you generate a TOC?


ERROR:
undefined internal link to this URL: "#heading=h.xtzluz7y0l31".link text: migration to core
?Did you generate a TOC?


WARNING:
You have 2 H1 headings. You may want to use the "H1 -> H2" option to demote all headings by one level.

----->

<p style="color: red; font-weight: bold">>>>>>  gd2md-html alert:  ERRORs: 2; WARNINGs: 1; ALERTS: 2.</p>
<ul style="color: red; font-weight: bold"><li>See top comment block for details on ERRORs and WARNINGs. <li>In the converted Markdown or HTML, search for inline alerts that start with >>>>>  gd2md-html alert:  for specific instances that need correction.</ul>

<p style="color: red; font-weight: bold">Links to alert messages:</p><a href="#gdcalert1">alert1</a>
<a href="#gdcalert2">alert2</a>

<p style="color: red; font-weight: bold">>>>>> PLEASE check and correct alert issues and delete this message and the inline alerts.<hr></p>

# Knative Migration-To-Core Process

Status: Migrating to GitHub

Author: Evan Anderson

Proposal: Define a maturity model for projects in the knative-sandbox (broader
community) or elsewhere to graduate into
[the Knative core](https://docs.google.com/document/d/1Jj9Kw5wmsqbal0GHabl0OG9xnOzKqbCxdkYak9QzlSk/edit#),
if appropriate. Note that projects may not map 1:1 with repos, and projects may
even be hierarchical (for example, the "eventing" project includes "eventing
API", "eventing delivery" and "eventing sources", each of which has artifacts in
multiple repositories).

Note that this document intentionally refers to the process of selecting
implemented interfaces from sandbox which should be part of core as
"migration-to-core" rather than "incubation" or "promotion" – while it may be
personally satisfying to have a component be part of the core, components which
are not in the core may be equally valuable and production-worthy.

# Maturity phases

## Initiating

In the initiating phase, a project is starting up, and may have minimal
additional artifacts around the code itself. The minimum bar for an initiating
project in the knative-sandbox org is:

- [Adopted the CLA bot and OWNERS file / Tide merge process](https://github.com/knative/community/blob/master/mechanics/CREATING-A-SANDBOX-REPO.md#technical-requirements),
  to ensure IP ownership
- Adopt the
  [Knative Code of Conduct](https://github.com/knative/community/blob/master/CODE-OF-CONDUCT.md)
- [Sponsored by a WG lead](https://github.com/knative/community/blob/master/mechanics/CREATING-A-SANDBOX-REPO.md#criteria)
- Have a README.md that describes the repo(s)

Projects which do not declare otherwise are assumed to be in the "Initiating"
phase, regardless of release process, users, etc.

This phase is suitable for:

- Experiments
- Proofs-of-Concept
- Examples
- Templates
- etc

Note that Initiating projects do not need to be hosted in knative-sandbox, and
can be hosted under either individual accounts or in other orgs. For projects in
knative-sandbox, review requirements and cadence are at the discretion of the
sponsoring WG lead.

## Usable

A usable project has reached the point where it is releasing executable
artifacts which are in use by at least two end users (preferably different
companies, rather than installations within the same company). Additionally, a
usable project meets certain other criteria (defined below) which allows the
Knative community to have confidence in the project and willingness to invest
further.

Additional requirements beyond "initiating":

- A published release schedule, with at least two successful minor releases.
  - At least two end users (preferably different companies) who are consuming
    the release artifacts. Commercial products with production usage count
    towards this target. This can be documented in something
    [like an ADOPTERS.md file](https://www.google.com/url?q=https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc&sa=D&ust=1597952611892000&usg=AFQjCNFymwghRnNGVqbD0O_01TsfEeto5w).
- Unit test automation meeting the Knative standards with respect to coverage
  and flakiness / test failures.
- Provide at least monthly progress reports at the appropriate WG(s). WGs should
  roll these reports up into their TOC reports.
- User-facing documentation for the following areas: install, usage
- Contributor facing documentation for: development setup
- Ongoing contributions from multiple contributors and management of user
  issues, etc.

Usable projects which meet the criteria for core (i.e. developer-facing
abstractions with wide utility, beta+ APIs) may apply to the TOC and SC for
[migration to core](#migration-to-core), but approval is more likely for
projects in the "Stable" category.

## Stable

A stable project has graduated beyond the bar in "usable" and is widely-adopted
with a stable API and well-understood operational characteristics. This would be
similar to a "GA" level of support for a commercial product in terms of
maturity.

Additional requirements beyond "usable":

- A stable v1 API following documented release and deprecation principles (ref
  [core Knative principles](https://github.com/knative/community/blob/master/mechanics/RELEASE-VERSIONING-PRINCIPLES.md)).
  Depending on the project, the API may include library methods, Kubernetes
  CRDs, network data-plane protocols, container contracts, or some combination.
  The API contract should be explicitly documented in the repo.
- Identified and named set of project leads managing both development roadmap
  and overall project success.
- Release schedule aligned with the core Knative release schedule.
- At least 5 end-users consuming the release artifacts in installations of
  sufficient quality and scope.
  - Commercial products with production usage count towards this target, though
    there must be multiple commercial products or direct upstream users to
    quality (i.e. 5 users of a single-vendor product and no other users may be
    "usable", but would not be accepted for "stable").
- Upgrade and e2e test automation.
- (Reports as for usable projects)
- User-facing documentation for debugging.
- Contributor facing documentation for issue triage and project engagement.
- Members of the development team is varied. Ideally from multiple
  organizations.
- There must be a "stable" set of installation-time dependencies for the
  project. (i.e. for Eventing, the "Kafka" channel implementation could be
  counted as a "stable" dependency)

Projects are generally encouraged to reach stable maturity before considering
[migration to core](#migration-to-core).

## Sunsetting

At any time, a project may move into the "sunsetting" phase. This is a
declaration that the project will no longer managed or maintained. For a project
which is currently in the "stable" phase, there should be a clear commitment to
support the project with bug and security fixes as applicable
[based on the Knative release principles](https://github.com/knative/community/blob/master/mechanics/RELEASE-VERSIONING-PRINCIPLES.md)
(4 versions of Knative).

# Migration to Core

A stable project may apply to the SC and TOC to be admitted to the Knative core.
Not all stable projects are suitable for "core"; this is a judgement as to the
degree to which the project should be considered a part of a "expected developer
experience" across Knative installations. Projects may be graded on the
following criteria (to be updated based on experience):

- Presence of developer-facing or operator-facing capability abstractions
- Independent usability / abstraction interfaces – how much does this require /
  complement existing Knative components (modular at the top)
- Ability to abstract over multiple underlying implementations (pluggable on the
  bottom)
- Number / diversity of installations and adoption as a fraction of total
  Knative usage
- Quality of technical implementation – ability to scale and reach
