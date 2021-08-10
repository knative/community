# Security Working Group Charter

## Intro

The “aaS” in FaaS stands for “as-a-service”. Many users would like to use Knative as the basis of a multi-tenant FaaS or PaaS-style platform. Even users who do not need hard multi-tenancy will wish to be confident of Knative’s security in single-tenant and soft multi-tenant scenarios. Currently the project provides limited guidance for how to configure and deploy itself successfully in these environments. The Security Working Group’s mission will be to fix this, by improving documentation, code, and process as needed to harden Knative for these use cases.

## Mission

The Security Working Group is responsible for helping users successfully run Knative in security sensitive environments. This mission will involve documentation, proactive exploratory testing, fixing of gaps, and production of necessary manifests and example security policies for various use cases.

The Working Group will evaluate new and existing features for their security impact, and where necessary will work with other Working Groups to add features and close gaps in order to support security-sensitive and multi-tenant use cases. The overarching goal is to, as well as to promote a security-conscious mindset throughout the project.

Example features and tasks which could fit within this mission (see also “Goals”, below):

- Documented recommendations and example manifests for setting up NetworkPolicy to prevent access between namespaces and for enabling mTLS.
- Documented recommendations and example manifests for securing traffic between Eventing and Serving workloads.
- Filling gaps in the security story for integrating eventing and serving. .
- Document how to expose Knative as a multi-tenant service. For example configuring a secure runtime (gVisor, Kata containers, firecracker etc), adding the AlwaysPullImages admission controller and appropriate security policies, setting up appropriate NetworkPolicy.
- Documenting and improving the current security posture of Knative for various specific use cases, such as multi-tenant environments, air-gapped environments, and security-sensitive single tenant environments.
- A suggested PodSecurityPolicy (or OPA policy etc) to be used with Knative.
- Tests and configuration to run Knative with gVisor, ClearContainers, and other isolation technologies.
- Integration with open source security tools, eg. Falco.
- Setting up and maintaining Security Scanners, Linters and other automated tools (and initially fixing any found problems).
- Defining processes for reporting, handling and pro-actively auditing the system for vulnerabilities; for example organising exploratory testing sessions.
- Defining and owning the process for handling reported security vulnerabilities (but see “non-goals” section w.r.t.not directly handling vulnerability reports, since this likely would mandate a restricted membership group).

## Working Group

Since this work cross-cuts Serving, Eventing, and Documentation, and since Security and Multi-tenancy are on-going considerations for the project, this will be a top-level working group which will work closely with other working groups.

## Goals

- Ensure that documentation and sample configs exist which permit Knative to be deployed securely (to the extent possible) in security-sensitive and multi-tenant environments. Close gaps (working with other WGs) and add features as needed to make this possible.
- Share and document best practices between those using Knative for security-sensitive use cases.
- Provide example configurations that lock down the upstream cluster to maximize security (i.e. to configure the system so that only the Knative APIs are available, and security-sensitive fields such as securityContext and serviceAccountTokens are blocked).
- Validate new feature proposals for their ability to be implemented in secure and multi-tenant environments.
- Consider the API surface of Knative to identify and fix potential DOS attacks and other security vulnerabilities, and to ensure Knative can be successfully used in multi-tenant environments, and for security-sensitive use cases.
- Publish design docs on Knative Threat Model, Security Analysis docs etc.
- Set up processes to ensure ongoing security of the platform. For example security linters, scanners, fuzzers, patch managers, and at a higher level through ensuring security is considered holistically and pro-actively by the project as a whole.
- Run a (ideally regular) brainstorming/exploratory session where we proactively brainstorm, investigate and mitigate potential attack vectors against the platform.

## Non-Goals

- Hard multi-tenancy in the unrestricted upstream platform. We will seek to upstream features as needed, but we assume achieving the goal for Knative in the short term may imply blocking access to some features in the underlying platform.
- Guaranteeing the security of the platform - this is ultimately up to vendors, administrators and operators; the WG mission is to provide guidance and share best practices to enable these people to be successful.
- Responsibility for specific vulnerabilities (though the WG may be a good place to define and steward a general process for handling vulnerabilities, security scanning etc, actively being responsible for managing specific reported vulnerabilities would likely require a closed WG).

## Scope

### In the existing repos:

- Provide cross-cutting input on security and multi-tenancy concerns.
- Provide documentation and example configurations for deploying Knative in multi-tenant environments, for example PodSecurityPolicies, NetworkPolicies, Secure Runtimes (some of this might end up being in a different repo).
- Fill gaps needed to use securely, as appropriate. For example rate limiting for DOS attack mitigation, network policy, podsecuritypolicy (or related new hotness as it’s deprecated :)).
- Set up & own automation and security scanning tools, linters.
- Set up & own test infrastructure for deploying and testing against an environment with limited RBAC rules.

### In working group meetings:

- Provide a forum to discuss and document security-sensitive use cases, experiences and requirements, and to surface gaps and needed features.
- Provide a forum to discuss potential new feature proposals WRT security and multi-tenant/security-sensitive-single-tenant use cases.
- Perform pro-active exploratory testing for vulnerabilities, DOS vectors etc.
- Discuss, prioritise and document the current and desired security posture of the project with regard to various scenarios (e.g. security-sensitive single tenant, air-gapped, multi-tenant etc).

## Preliminary 3-month Roadmap

- Document, with appropriate warnings about limitations, the current security recommendations for deploying Knative in various environments, for example hosted multi-tenant, private/trusted multi-tenant, secure single tenant.
- Provide and document an optional set of RBAC rules, PodSecurityPolicies etc which give access only to public Knative APIs, and block access to insecure features (securityContext, service tokens etc).
- Define and document a process for reporting and handling vulnerabilities (see discussion in [https://github.com/knative/community/pull/258](https://github.com/knative/community/pull/258))
- Set up infrastructure and CI automation to test changes against a restricted-RBAC, restricted PodSecurityPolicy environment. Close any found gaps.
- Define and add a security section to the Project Proposal Template to ensure we explicitly consider security implications of new features.
- Working with Networking WG, create NetworkPolicies to secure access between namespaces and between the knative-serving namespace and container namespaces.

## Potential Future (6-12 month) Items

- Work with Eventing and Serving WG to provide secure communication between Events and Services in as networking-provider agnostic a way as possible.
- Working with Documentation WG, provide deployment instructions for configuring Knative with a secure container runtime (e.g., gVisor, Kata Containers)...

## Proposed Leads

- Julian Friedman (IBM)
- Evan Anderson (VMware)
