# Knative Threat Model

This document describes the Knative threat model. When vulnerabilities are
reported to the project, we consult this document to determine whether the
report describes a potential exploit, and if so to determine the severity of
the exploit. As we develop new features, we consult this document to consider
their impact on the threat model.

_Status:_ Work in Progress. This document does not yet represent the full
threat model envisioned by the knative community, but we have to start
somewhere.

## Trusted and untrusted workloads

Knative is often used to implement platforms which permit users to submit code
or containers, which the platform runs. In some cases the code is submitted
only by trusted users -- for example in a private cluster -- whereas in others
the code may be entirely untrusted (for example when knative is used to
implement a multi-tenant FaaS-style system).

We differentiate between attacks on the system that require permission to
create a Knative Service or Revision, and attacks that can be performed against
knative clusters or workloads even without permissions to directly deploy a
workload. In a multi-tenant environment it is possible that all users who can
access the environment will also have permission to deploy workloads.

The most significant exploits are those that affect a cluster even without the
ability to submit workloads.

## Availability, confidentiality, and integrity

We consider the confidentiality of user data and integrity of user workloads
the most significant concern. We are also concerned about Denial of Service and
Availability attacks (particularly where exploitable without the ability to
submit workloads), but are aware that these currently exist both within Knative
and the underlying Kubernetes platform. For this reason availability attacks
are currently treated as relatively lower priority.

## Multitenancy in a single cluster

The project differentiates between environments where a single tenant is in
control of the cluster (or trusts all other tenants in the cluster), and
environments where multiple tenants share a single cluster. In the latter case
we assume each tenant owns a namespace, and that appropriate network and access
control policies at the cluster level are defined. In no case do we assume a
security boundary exists between two workloads deployed to the same namespace.

While we are concerned with vulnerabilities that affect multi-tenant clusters,
given the current state of the art with regards to Kubernets multi-tenancy
support, we expect a significant number of users will run knative in
single-tenant clusters.

## Security of the underlying Cluster

Knative relies on the security of the underlying cluster to maintain its security
posture. This currently requires manual work that is not fully documented by
the project (though we would eventually like to fix this).
The project cannot be secure if the underlying platform is not secure,
this requires work on top of the defaults provided by Kubernetes.

## Security of container and network technologies

Similarly to the above paragraph, Knative relies on the security of underlying
container and networking technologies for its security. This includes proper
configuration of these technologies. For example, it is not secure to run
mutually untrusted workloads in a cluster without properly configuring
isolation technologies (for example a VM-based container engine such as gVisor
or firecracker, or enabling - at least - user namespaces and seccomp if using
containerD or crio), and network technologies. The project does not currently
provide guidance on this, although we would like to in future.
