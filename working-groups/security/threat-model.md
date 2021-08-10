# Knative Threat Model

This document describes the Knative threat model. When vulnerabilities are
reported to the project, we consult this document to determine whether the
report describes a potential exploit, and if so to determine the severity of
the exploit. As we develop new features, we consult this document to consider
their impact on the threat model.

_Status:_ Work in Progress. This document does not yet represent the full
threat model envisioned by the knative community, but we have to start
somewhere.

## Tenancy Model

There are three common tenancy models currently in use in the Kubernetes
Ecosystem (See [Three Tenancy Models for Kubernetes](https://kubernetes.io/blog/2021/04/15/three-tenancy-models-for-kubernetes/)).
Summarising here, in the first, a cluster is fully owned by a single tenant
(Cluster-as-a-service).
In the second, a cluster is partitioned between multiple tenants with the
Namespace forming a security boundary between tenants of the cluster
(Namspace-as-a-service).
In the third, resources such as nodes and schedulers are shared, but the
tenants, are presented with a logical cluster and can perform global operations
such as installing CRDs and managing namespaces (Control-plane-as-a-service).

Currently, Knative assumes a single tenant per cluster (Cluster-as-a-service)
as its security model.

While it is possible to secure Knative for use in non-single-tenant
environments, for example by setting up appropriate NetworkPolicy, using
appropriate container isolation techniques, and filtering requests to shared
components such as the activator, the project does not currently consider
itself to be secure for this use case. We would like eventually to support
secure use in Namespace-as-a-service and Control-plane-as-a-service
environments out of the box, and are very receptive to contributions that
improve support for Namespace-as-a-service clusters, but we do not currently
view exploits in these situations as vulnerabilities for the purpose of our
disclosure policy.

## Trusted and untrusted workloads

Knative is often used to implement platforms which permit developers to submit
code or containers, which the platform runs. In some cases only trusted
developers can submit code (for example in a private cluster), whereas in
others the code may be entirely untrusted (for example when knative is used to
implement a multi-tenant FaaS-style system). In either case the resulting
applications might be available only to trusted users (again for example in a
private cluster), or to untrusted users (a single-tenant cluster providing
internet-accessible applications).

We differentiate between attacks on the system that require permission to
create a Knative Service or Revision, and attacks that can be performed against
knative clusters or workloads even without permissions to directly deploy a
workload. In a multi-tenant environment it is possible that all users who can
access the environment will also have permission to deploy workloads.

The most significant exploits are those that affect a cluster even without the
ability to submit workloads. Another class of attack would involve an
administrator with access to Knative being to perform actions that exceed their
cluster roles; for now, since we primarily consider single-tenant clusters,
while we are keen to fix any bugs in this area, we do not consider cases such
as this as an exploit.

## Availability, confidentiality, and integrity

We consider the confidentiality of user data and integrity of user workloads
the most significant concern. We are also concerned about Denial of Service and
Availability attacks (particularly where exploitable without the ability to
submit workloads), but are aware that these currently exist both within Knative
and the underlying Kubernetes platform. For this reason availability attacks
are currently treated as relatively lower priority.

In general, since many availability risks are inherited from the underlying
platform, we only - in Knative - consider issues which are significantly higher
than the same situation in Kubernetes itself.

## Security of the underlying Cluster

Knative relies on the security of the underlying cluster to maintain its
security posture. This requires manual work that is the responsibility of
upstream SIGs to document.
The project cannot be secure if the underlying platform is not secured.

## Security of container and network technologies

Similarly to the above paragraph, Knative relies on the security of underlying
container and networking technologies for its security. This includes proper
configuration of these technologies. For example, it is not secure to run
mutually untrusted workloads in a cluster without properly configuring
isolation technologies (for example a VM-based container engine such as gVisor
or firecracker, or enabling - at least - user namespaces and seccomp if using
containerD or crio), and network technologies. The project does not currently
provide guidance on this, although we would like to in future.
