# Functions WG Charter

Drafted: Apr 28, 2022

## Mission
The Functions Working Group aims to simplify application development and deployment to Knative by transforming business logic written as a programming language function into OCI container images, and creating or updating the associated on-cluster resources.

By providing higher-level abstractions through functions, the working group aims to make Knative accessible for developers without extensive Kubernetes or container experience.

## Goals
- Provide a developer experience enabling portable, high-level business logic in several programming languages in the form of a “Function” as a Knative Service.
- Provide Function invocation, HTTP serving and CloudEvent marshaling capabilities through language frameworks, capable of wrapping and bootstrapping Knative Function code into a runnable service.
- Provide a developer experience that enables the build and deployment of Knative Functions as OCI container images in a Knative Service.
- Define and document a pluggable ‘language pack’ contract to adapt developer code to the Knative runtime contract.
  - Additionally, curate a set of language plugins which are recommended to new developers.
- Implement a shared CLI which provides commands for creation, build, iteration, and deployment of Knative Services using language plugins.
  - The shared CLI should be as broadly useful as possible, supporting both local and remote build tooling.

## Non-Goals
- Provide a general-purpose client which duplicates the functionality of kn.
- Implement language plugins for all programming languages.

## Scope
- Defining and implementing the interface between developers and the functions CLI.
  - Including command-line options, UX flow, supported commands, source directory requirements, etc.
- Defining and supporting the language plugin interface.
  - This includes developing and supporting third-party language plugin authors as well as first-party language plugins.
  - This also includes any “catalog” or “template” management and shared services provided to language plugins via the CLI or on-cluster build interfaces.
- Providing examples and coordinating with the documentation WG to provide user-facing examples and documentation on using the function tools.

## Initial Leads
- Lance Ball - Red Hat
- Mauricio Salatino - VMWare

## Roadmap
https://github.com/orgs/knative-sandbox/projects/7

