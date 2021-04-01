# Knative Sources Working Group Charter

Author(s): Scott Nichols

Last Updated: Oct 30, 2019

## Mission

The Knative Sources Working Group will reduce the friction of creating, finding, and implementing Knative compliant Sources. We will enable new ways to create, test, install, deploy and observe applications that produce or bridge events onto a cluster for consumption.

## Goals

- As a Developer Persona, enabling the creation of custom sources from new or existing code/containers with little to no changes.
- As a Developer Persona, allow for easy understanding of existing and potential sources installed in a cluster.
- As an Operator, installing new sources is trivial, and not a burden on the overall health of the cluster.
- As a Contributor, enable new methods for allowing new sources to be created allowing for user provided code and containers to adhere to the source specification defined by Knative Eventing.

## Scope

In the existing Eventing repo inside the knative org:

- Continue to evolve the [Source Spec](https://github.com/knative/eventing/blob/master/docs/spec/sources.md).
- Expand on the Pod-Specible Sources work [outlined](https://docs.google.com/document/d/159-xjBwougBWHOigWUk42tHSB_T1As6zdPwb_9zK38s/edit#heading=h.n8a530nnrb), and POC [implemented](https://github.com/n3wscott/sources).
- Explore the concept of sidecars that enable and extend existing pods to emit CloudEvents, implement overrides, and possibly dynamically update the sink target based on configuration.
- Leverage the POC from [Auto-Container-Source](https://github.com/Harwayne/auto-container-source) and expand it to include all implemented Pod-Specable sources.
- Provide examples of custom sources implemented in several languages (Go, Java, Javascript, Python).
- Provide SDK and test suite for event sources
- Provide test conformance suite to verify existing and knative sources are [Source Spec](https://github.com/knative/eventing/blob/master/docs/spec/sources.md) compatible (and versioned). Run automatic tests and publish results for official and nonofficial knative sources.
- Collaborate with CLI for good UX on creating and discovery of Sources installed and/or running.

In working group meetings:

- Provide a forum to discuss and align work items related to the working group, status updates.

Sources WG Overall responsibility:

- Maintaining Sources inside of knative/eventing.
- Advising compatibility with Source contracts contributed to Knative inside of knative/eventing-contrib.
- Ensuring that there are high-quality Sources for the most important (as decided by the WG) event-producing systems.
- Managing and directing the larger Source-authoring ecosystem.
- Tracking long-term efforts around improving the efficiency of Sources when working in coordination with other Knative Eventing components (e.g. upstream filter propagation, batching, or other optimizations).

## Preliminary 3-Month Roadmap

Month 1

- Establish principles of Sources (as both Developer and Operator Personas).
- Define Pod-Specable Run Time Contract (Ref: [work in progress](https://github.com/n3wscott/sources/blob/master/docs/runtime-contract.md)).
- Implement 3 Pod-Specable (TBD: DeploymentSource, JobSource, CronJobSource) sources in the sources.knative.dev API group inside of Eventing.
- Deprecate ContainerSource and CronJobSource in sources.eventing.knative.dev.
- Migrate to new api group: ‘sources.knative.dev’
- Engage with CLI to develop a long term UX plan for creation of sources.

Month 2

- Implement 3 more Pod-Specable Sources (TBD: DaemonSource, StatefulSetSource, ServiceSource)
- Enable Auto\*Source for 2-3 sources (TBD: DeploymentSource, ServiceSource).
- Implement Conformance Test for runtime contract of Pod-Specable Sources.
- Engage with Operator to see path forward for additional source installations.
- Began implementing example sources for 1-2 languages for 1-2 Source types.
- CLI is able to discover and create Sources.
- Start Sidecar and usability research for Advanced Lazy Sources..
- Create

Month 3

- CLI can create and discover Sources from labeled CRDs.
- Implement remaining Pod-Specable resources (TBD: ReplicaSet, Pod)
- Enable Auto\*Source for remaining Sources.
- Implement Examples for remaining languages as needed.
- Operator integration has been determined and work is ongoing.
- TBD: Additional Sidecar work per results of Advanced Lazy Sources research.

## Working Group

We will operate as a partner group to Eventing, giving a brief update in the weekly Eventing working group meeting. We may additionally meet on a weekly basis, cancelled when there is a lack of agenda.

### Nominated Leads:

Nacho Cano, Scott Nichols, Ville Aikas, Lionel Villard
