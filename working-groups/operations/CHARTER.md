# Knative Operations Working Group Charter

Created: 2019-04-26 / Last Updated: 2019-05-06

## Charter

### Mission

The Knative Operation Working Group provides best practices for managing, assessing system health and maintaining Knative clusters. The initial focus will be on how Knative is installed, removed, and upgraded.

### Goals

- Provide a means to install Knative [Serving/Build/Eventing] with operator(s) specified options
  - Operator specified defaults such as timeouts and concurrency
  - Network preferences, Istio / Gloo / others
  - Build Template selection
  - Eventing component selection
- Provide a means to install Knative with limited/no internet access
- Provide a means to upgrade/downgrade Knative with limited/no user visible impact to their service
- Provide guidelines on determining the overall health of Knative
- Provide cluster sizing and lifecycle recommendations
- An entrypoint and forum for dogfooders to give feedback

### Non-Goals

- The knative operator does not manage or install Istio, or other dependents.
  - But we might want to have the option to recommend other operators for this task. Like `knative.spec.istioOperator = yes` please.

### Scope

- Tooling for installation/removal/upgrades/downgrades
- Management mechanisms for:
  - namespace setup (if needed)
  - Quota and RBAC management
  - Build and registry configuration
  - integration with existing cluster o11y
- Debugging tools for operator personas.
- Define more specific "operator" sub-personas within the Knative "operator" persona (e.g. auditor, cloud provider, IT provider, etc)
  - Collect feedback from operators on the usability of Knative
- Collaborate w/ Documentation Working Group on documentation for installation/removal/upgrades/downgrades
- Collaborate w/ Observability Working Group on metrics for health of Knative
- Collaborate w/ Productivity Working Group on using best-practice installation for testing

## Lead Nomination

Ben Browning (bbrowning), Gregory Haynes (greghaynes), Kenny Leung (k4leung4)

## Roadmap\*

- Month 1 ~ Knative 0.7
  - Implement all tasks under [Setting Up A Working Group](https://github.com/knative/community/blob/master/WORKING-GROUP-PROCESSES.md#setting-up-a-working-group)
  - Set up separate repo for each of Knative serving/eventing
  - Enable operator installs for current version of Knative Serving
- Month 2 ~ Knative 0.8
  - Adopt operator installation usage for automated Knative Serving
  - Decide on parameters for customization for Knative Serving
  - Enable operator installs for Knative Build/Eventing
- Month 3 ~ Knative 0.9

  - Investigate using operator for test execution
  - Explore upgrade scenarios for Knative Serving using operator

- Assumes adoption of operators

## Previous Efforts

- [Knative Operator Working Group Proposal](https://drive.google.com/open?id=11cenPfZrEU04OFQYTE_6e_LrOw07HPOkeT1_Gigmy7U)
