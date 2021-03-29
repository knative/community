# Contributor Productivity Working Group Charter

This is the proposal of creating contributor productivity working group for Knative.

# Mission

  The mission of Contributor Productivity working group is to develop engineering guidelines, tools and infrastructure for Knative open source community to enable high quality and high velocity deliverables, to drive measurable and predictable releases, and to achieve high community collaboration and overall productivity.

# Goals
- Foster collaboration among contributors for common needs
- Identify top priorities or pain points
- Formulate guidelines and processes around testing, continuous integration, code velocity, and release packaging.
- Develop/Evangelize reusable tools and infrastructure to be utilized by contributors across repos for various purposes such as E2E testing, conformance testing, performance/scalability testing, project health reporting, product troubleshooting/diagnostics and etc. More specifically we’d graduate reusable tools to kubernetes/test-infra.
- Drive up contributor experience with ongoing improvements to automations around code life cycle such as issue creation, triage, code review, test and release. The goal is to reduce contributor overhead, minimize product defects and improve overall efficiency. In turn, it shall help developer and operator experience with better product quality and more efficient releases.

# Scope
- Project health
  - Define metrics (such as listed below) to measure project health, and deliver automation for improving project health.
    - Code coverage
    - API coverage
    - Release churns
    - PR latency (time to merge)
    - Onboarding velocity (time to PR)
    - Issue Triage velocity
    - Issue Resolving Velocity
  - Have automation in place to help drive up project health metrics with opt-in (default off) vs. opt-out (default on) options; and each WG can decide whether they are in or out.
- Test framework
  - Ensure there is common test infrastructure in place to run unit, build, E2E and conformance tests, to output logs and metrics from tests, and to visualize logs/metrics.
- CI/CD workflows
  - Maintain presubmit and postsubmit workflows to gate the quality of code checkin and quality of release
  - Ensure there is sufficient monitoring in place to monitor the health of workflow and infrastructure themselves
- Release criteria/process
  - Define and implement release criteria and process
- [Performance/Scalability/Load testing](https://docs.google.com/document/d/1_LXxZc_dlmFexILC27TnitFSFxP3l9mUkmGf_uTRYHc/edit)
  - Working with individual feature WG (serving, scaling, build, control plane, eventing) to identify common requirements and provide a shared framework for each WG to consume. Initial thinking of common need includes cluster creation/setup, load profiles, test orchestration, result collection, metric visualization and etc. The corresponding framework code will exist at test-infra repo.
  - Individual feature WG will be responsible for designing and implementing the actual testing, the corresponding test code will exist at individual repo.
  - Collaborate with partner teams to reuse existing performance/load testing infrastructure as well as environment whenever possible
  - Link above has more detailed proposals
- Troubleshooting experience
  - Ensure sufficient automation and documentation is in place for self-service Knative troubleshooting
- Will use Knative/test-infra repo for code, maybe a good idea to rename it for broader purpose as outlined in this doc

# Initial set of leads:
- Adriano Cunha (adrcunha on github)
- Jessie Zhu (jessiezcc on github)


# Proposed 3 Month Roadmap
- M8:
  - Set up regular WG meetings and establish forms of communication
- M9:
  - Document existing development/release workflows to form a common understanding within the community
  - Collect feedback on onboarding/development experience from the community from various sources such as meetings, slack, social media
  - Define performance/load testing goals for different scenarios (contributor vs. developer)
  - Finalize methodologies for perf/load testing
- M10:
  - Build tools to measure community satisfaction around onboarding, developer, test, troubleshoot experience.
  - Contributor friendly onboarding & troubleshooting experience based on feedbacks from M8
  - At least one WG adopting perf/load test infra (scaling)
  - Project health dashboard open to community
- M11
  - Shareable performance/load test toolset
  - Build tools to gather and track structured community feedback
  - Drive up community satisfaction based on measurements from M9

