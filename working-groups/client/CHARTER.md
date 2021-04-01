# Knative Client Working Group Charter

## Mission

The Knative Client Working Group maintains a delightful baseline operator-neutral Knative client developer experience, including documentation about what normative clients do, and a reference implementation of a normative client.

## Goals

- Provide a portable reference command-line client for Knative with a normative workflow.
- Provide guidance about normative uses of the Knative APIs to client developers.
- Provide reference implementations of Knative developer workflows as a library to be incorporated into other clients.

## Scope

In a new repo in the knative org:

- Maintain documentation around how the API can be used to create a delightful user experience for core Developer journeys, and client conventions for Knative as a whole and each Knative building block.
- Create and maintain a golang library for implementing the normative client flows for the Developer persona, for general Knative-wide conventions as well as workflows specific to each building block and their interactions. This library should support workflows for both graphical and command-line client tooling, though the group will only maintain an implementation of a command-line client.
- Create and maintain a standalone command-line client, `kn`, using the aforementioned library, that implements a developer experience for Knative that is portable across operators.

In working group meetings:

- Provide a forum to discuss non-CLI user tooling as well (web, etc.), and how the client libraries can support those too.

### Non-scope:

- Implementation of other experimental interfaces for Knative.

## Preliminary 3-Month Roadmap

**Month 1 (December. December barely exists.)**

- Get repo structure set up
- Start library for Serving
- Pick tech stack for CLI (set of libs to use for terminal manipulation & display etc)

**Month 2**

- Start CLI, using library for Serving
- One major user journey complete for Serving, ex "deploy a new image to a Service"
- With Build: Settle on supported path for deploy-from-source-dir
- With Eventing: Discuss user support for discovery story

**Month 3**

- Have at least two major user journeys completed e2e for Serving

## Details

### Documentation details

New WG will own <code>[normative-examples.md](https://github.com/knative/serving/blob/master/docs/spec/normative_examples.md)</code> and <code>[client-conventions.md](https://github.com/knative/serving/blob/master/docs/client-conventions.md)</code> from Serving; similar docs can start here, concerning any of Serving, Eventing, Build.

### Library details

- A Go library implementing core Knative functionality for the Developer persona, including:
  - General Knative-wide functionality
    - Library for reading and dealing with ready conditions and terminal conditions.
      - Watching or polling them for completeness
      - Reporting on partial completeness
  - Support for the core user journeys across Knative projects, including cross-project integrations. Examples:
    - Serving
      - Creating Services
      - Deploying software to Services
        - Deploying from source: integrated with Build
        - Deploying an image
      - Managing the configurations of services
        - Resource limits
        - Concurrency
        - Environment
        - (etc)
      - Rolling out traffic to revisions
      - Rolling back
      - Fetching and watching all Knative Serving objects
    - Build/Pipeline
      - Fetching build templates
      - Creating builds using build templates
        - From a source repository version
        - From a directory of source (possibly using a technique similar to [Matt's proof-of-concept](https://github.com/mattmoor/kontext#kontext))
      - Creating builds not using build templates (NB: Probably not part of most mainline developer experiences, but included for completeness)
      - Creating namespace-scoped build templates
    - Eventing
      - Creating bindings from events to endpoints in services
      - Listing available event sources
      - Describing source capabilities and configuration options
      - Creating event source instances (diff between CRD defining the source and an actual instance of a source) in a generic way, these vary wildly
      - CRUD on Channels
      - A way to find unused channels
  - Any future Knative building blocks

### CLI details

A basic command-line client (or plugin to kubectl, chosen based on overall developer experience), built on the Go library, that provides a command-line interface for all the tasks.

- We like short command names. `kn`?
- Authentication using kubeconfig
- Output formatting and filtering similar to kubectl (tabular, json, yaml formats, templating)

### Responsibility also includes

- Tests for all of the above.
- Code coverage for all the above
- Any extra github automation for the repo

### Possible future scope:

- Client libs in other languages. (interest in py from google, at least)

### Out of scope:

- Operator concerns (Focusing the interface on the role of developer will lead to a better developer experience, and operator concerns are irrelevant to most developers. Also, operator concerns are more likely to involve non-Knative APIs, like RBAC)
  - Installing cluster-scoped build templates
  - Installing Knative
  - Making cluster-scoped configuration choices
- Any functionality that relies on non-Knative APIs (Knative clusters are built using a variety of features, due to the "pluggable backends" and "duck typing" work we've been doing, and so if an API hasn't been standardized into Knative we can't assume a Knative cluster has it)
  - (Currently) exposing a service outside the cluster
  - RBAC management
  - Binding a service to a domain name
  - Installing Knative
- Any functionality that relies on a particular implementation of Knative. (Client libs should work the same against any implementation of Knative)
  - Vendor extensions (aside from codifying where in the API they can go)
  - Any particular happenstance/subtleties that aren't codified, for example in exact ordering of conditions becoming True for a happy-path operation.
- Conformance tests for Knative. These can and should **use** the client libraries we're working on, but deserve their own home.

## Working Group

To be a peer of the other Knative working groups, run weekly (to start), cancelled when we don't have enough to talk about.
