---
title: "Knative release principles"
linkTitle: "Release principles"
weight: 50
type: "docs"
---

# Knative Release Principles

## Principles

_These principles are forward looking from release 0.13.x onward. There are
plenty of examples where some or all of these principles have not been followed
in the past. We are not looking to change past behavior, but to help define
future behavior for the project._

### Overarching principle:

Wherever possible† follow the principles that are widely adopted by Kubernetes
for
[API changes](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md)
and
[API deprecation](https://kubernetes.io/docs/reference/using-api/deprecation-policy/).

_† - Some things may not be possible due to either the maturity of CRDs or the
mechanism by which they are implemented (e.g. not an API Server, no conversion
webhooks, bugs)._

### K8s minimum version principle:

We (the community) support the range of upstream minor Kubernetes versions that
are Kubernetes community supported at the time of the cut. Only the latest patch
release of the minor version is tested for support. Support for versions prior
to that are best effort, or by vendors.

Adoption of newer K8s minimum versions will also be on a best-effort basis. When a
Kubernetes version is released it's availability in downstream distributions is
often lagging. When these delays affect our planned min K8s version we'll update
the version table below and make it known in the release's notes.

This principle will be revisited when changes to upstream Kubernetes support 
policy occur.

### Knative community support window principle:

We (the community) support the most recent 4 versions of Knative. The term
support as used here means that we will maintain release branches for these
releases. We will consider backporting applicable fixes to these release
branches depending on severity and feasibility.

### Default and optional API versions principle:

There will be at least one common API version available across all community
supported versions (See
[Community support window principle](#knative-community-support-window-principle))
at any given time. At least one of these common API versions will be enabled by
default from the open source installation path. The open source installation
process may allow optional API versions at the discretion of the installing
user.

### API Support Principle:

We (the community) support all API versions (both default and optional)
available from current community supported Knative releases. We will keep Beta
API versions available in releases for at least 9 months and GA API versions in
releases for at least 12 months after deprecation is announced. Earlier API
versions are best effort, or by vendors. See FAQ for v1alpha1.

### Recommended API version principle:

We (the community) recommend that clients developing to a static API version
(including non-versioned tests) develop to the highest common API version
offered across supported community releases at the time of development. Where
feasible a dynamic client should be preferred.

> **Example:** If 0.8.x through 0.11.x are the current community supported
> versions then v1alpha1 would be the recommended client version as 0.8.x does
> not support v1beta1 nor v1 by default. This makes v1alpha1 the common version.
> (Note: This would mean 0.11.x must support v1alpha1 by default otherwise we do
> not have any minimum overlapping version).

### Runtime Conformance Principle:

In order to be runtime conformant for a given release, an installation must pass
the runtime conformance tests for the installed version and 3 prior versions (as
determined by the supported version principle). This means runtime tests for a
given version should follow the recommended API version principle to be forwards
compatible when run against one of the 3 subsequent versions.

### API Conformance Principle:

In order to be API conformant for a given release, an installation must pass
conformance for the installed version and 3 prior versions for common API
versions. There will always be at least one such version due to the principles
above.

### Project, Feature, and Sub-Feature Phase Principle:

There will be a phase identified for each project, feature, and sub-feature for
each release. Each of the phases: Alpha, Beta, and Stable will have clear
definitions related to performance, deprecation and maintenance.

|                      | Alpha                           | Beta                                      | Stable                                                                   |
| -------------------- | ------------------------------- | ----------------------------------------- | ------------------------------------------------------------------------ |
| _Purpose_            | Works with possible limitations | Works end to end                          | Production Ready                                                         |
| _API_                | May not be backward compatible  | Versioned, may not be backward compatible | Versioned / Backward Compatible                                          |
| _Performance_        | No guarantee                    | No guarantee - Baseline                   | Performance is quantified, documented, and guarantees against regression |
| _Deprecation Notice_ | none                            | 9 months                                  | 12 months                                                                |

## FAQ

### What happens during upgrades of Knative?

We support upgrading versions of Knative from any community supported version of
Knative to a minor version one higher. Upgrade steps larger than one minor
version must go through an intermediate version first. As an example, if the
current community versions are 0.9.x through 0.12.x the upgrade matrix will look
like below:

| Knative Version | Upgrades to |
| --------------- | ----------- |
| 0.9.x           | 0.10.x      |
| 0.10.x          | 0.11.x      |
| 0.11.x          | 0.12.x      |
| 0.12.x          | N/A         |

### What happens during downgrades of Knative?

Knative downgrades are a harder story than upgrades as the purpose of the
upgrades are to fix bugs and add features. That said there are many cases where
a downgrade or rollback is appropriate. As downgrades are more difficult to get
right Knative will make a best effort to support downgrade of a single minor
version and only if users have not already taken advantage of new features or
fields.

This means that we need to support the downgrade tests case within our project
to ensure that downgrade is indeed possible before code checked in and releases
are cut.

### What happens during upgrades of Kubernetes?

Strawman: Upgrading Kubernetes may move you from a Kubernetes version for which
a release was intended to a newer version that may have not been originally
qualified. We expect this to work given Kubernetes backwards compatibility
support. We intend for these upgrades to be a non-issue. If particular versions
are found to be incompatible and the Knative version is still community
supported we will either:

1. work with the Kubernetes community on a fix or
2. develop a patch release to workaround the issue. We will publish guidance on
   potential incompatibility on knative.dev

### Which API endpoints should I enable when I vend/host Knative Serving?

Vendors may make their decisions about supporting API versions. They may decide
to support optional API versions by default or may support API versions that are
EOL in the open source community. The recommendation of the community is to
offer all APIs versions that the Open Source project makes available by default
for a given release version. This offers clients the best interoperability and
users the best portability.

### What about Serving v1alpha1 API support?

Due to the length of time we have supported v1alpha1, the ubiquity of
availability across released Knative Serving versions, and the presence of
clients statically developed against this version we will treat the current set
of v1alpha1 Serving APIs as 'Beta' APIs. This means that we will keep v1alpha1
available in releases for at least 9 months after deprecated is announced.
During that period of time of availability, we will evaluate whether the API is
available through the default Open Source installation by following the
[Default and optional API versions principle](#default-and-optional-api-versions-principle).

### What happens when the minimum Kubernetes version for a supported Knative release is no longer supported by Kubernetes?

We will explicitly avoid introducing changes into a Knative Serving release
branch that would raise the minimum version. We will not keep around test
clusters for unsupported Kubernetes versions, but we will attempt to solve user
reported bugs to the best of our ability.

### Can we have the table below on knative.dev?

Yes. We will publish a table similar to below once we have agreement on the
principles so that users can easily determine versions and endpoints available
for a given release.

## Knative Version Tables

Storage Type = The type that is stored in etcd by the Kubernetes API Server

API Endpoints = The rest endpoints that are made available to clients

API Types = The API types that the endpoint understands

Lemonade = v1alpha1 + v1beta1 super-set type (Serving only)

Force Upgrade = Automatically convert v1alpha1 to v1beta1 syntax via Lemonade

Availability =

- Default means installed through typical open source release.
- Optional means available in the open source, but will not be installed by
  default.
- Removed means support no longer exists within the open source release

**Note**: subject to change based on verification of (knative, kube) version
tuple

### Knative Serving Version Table

<table>
  <tr>
   <td><strong>Knative Version</strong>
   </td>
   <td><strong>Release Date</strong>
   </td>
   <td><strong>End of LIfe Date</strong>
   </td>
   <td><strong>Min K8s Version</strong>
   </td>
   <td><strong>Storage Type</strong>
   </td>
   <td><strong>Force Upgrade</strong>
   </td>
   <td><strong>API Endpoints</strong>
   </td>
   <td><strong>API Type Served</strong>
   </td>
   <td><strong>Availability</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.20.x
   </td>
   <td rowspan="3" >2020-12-22
   </td>
   <td rowspan="3" >2021-06-08
   </td>
   <td rowspan="3" >(undetermined)
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td><strong>removed</strong>
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td><strong>removed</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.19.x
   </td>
   <td rowspan="3" >2020-11-10
   </td>
   <td rowspan="3" >2021-04-27
   </td>
   <td rowspan="3" >(undetermined)
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td><strong>removed</strong>
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td><strong>removed</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.18.x
   </td>
   <td rowspan="3" >2020-09-29
   </td>
   <td rowspan="3" >2021-03-16
   </td>
   <td rowspan="3" >(undetermined)
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>optional (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>optional (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.17.x
   </td>
   <td rowspan="3" >2020-08-18
   </td>
   <td rowspan="3" >2021-02-02
   </td>
   <td rowspan="3" >1.16.x
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.16.x
   </td>
   <td rowspan="3" >2020-07-07
   </td>
   <td rowspan="3" >2020-12-22
   </td>
   <td rowspan="3" >1.16.x
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.15.x
   </td>
   <td rowspan="3" >2020-05-26
   </td>
   <td rowspan="3" >2020-11-10
   </td>
   <td rowspan="3" >1.16.x
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.14.x
   </td>
   <td rowspan="3" >2020-04-14
   </td>
   <td rowspan="3" >2020-09-29
   </td>
   <td rowspan="3" >1.15.x
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.13.x
   </td>
   <td rowspan="3" >2020-03-03
   </td>
   <td rowspan="3" >2020-08-18
   </td>
   <td rowspan="3" >1.15.x
   </td>
   <td rowspan="3" >v1alpha1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.12.x
   </td>
   <td rowspan="3" >2020-01-21
   </td>
   <td rowspan="3" >2020-07-07
   </td>
   <td rowspan="3" >1.15.x
   </td>
   <td rowspan="3" >v1alpha1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.11.x
   </td>
   <td rowspan="3" >2019-12-10
   </td>
   <td rowspan="3" >2020-05-26
   </td>
   <td rowspan="3" >1.14.x
   </td>
   <td rowspan="3" >v1alpha1 (Lemonade)
   </td>
   <td rowspan="3" >Yes
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.10.x
   </td>
   <td rowspan="3" >2019-10-29
   </td>
   <td rowspan="3" >2020-04-14
   </td>
   <td rowspan="3" >1.14.x
   </td>
   <td rowspan="3" >v1alpha1 (Lemonade)
   </td>
   <td rowspan="3" >Yes
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default 
<p>
(deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.9.x
   </td>
   <td rowspan="3" >2019-09-17
   </td>
   <td rowspan="3" >2020-03-03
   </td>
   <td>1.13.10
   </td>
   <td rowspan="3" >v1alpha1 (Lemonade)
   </td>
   <td rowspan="3" >Yes
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>optional
   </td>
  </tr>
  <tr>
   <td>1.13.10
   </td>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>optional
   </td>
  </tr>
  <tr>
   <td>1.11
   </td>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.8.x
   </td>
   <td rowspan="2" >2019-08-06
   </td>
   <td rowspan="2" >2020-01-21
   </td>
   <td>1.13.10
   </td>
   <td rowspan="2" >v1alpha1 (Lemonade)
   </td>
   <td rowspan="2" >Yes
   </td>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>optional
   </td>
  </tr>
  <tr>
   <td>1.11
   </td>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.7.x
   </td>
   <td rowspan="2" >2019-06-25
   </td>
   <td rowspan="2" >2019-12-10
   </td>
   <td>1.13.10
   </td>
   <td rowspan="2" >v1alpha1 (Lemonade)
   </td>
   <td rowspan="2" >Yes
   </td>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>optional
   </td>
  </tr>
  <tr>
   <td>1.11
   </td>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>0.6.x
   </td>
   <td>2019-05-14
   </td>
   <td>2019-10-29
   </td>
   <td>1.11
   </td>
   <td>v1alpha1 (Lemonade)
   </td>
   <td>No
   </td>
   <td>v1alpha1
   </td>
   <td>Lemonade
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>0.5.x
   </td>
   <td>2019-04-02
   </td>
   <td>2019-09-17
   </td>
   <td>1.11
   </td>
   <td>v1alpha1
   </td>
   <td>N/A
   </td>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>0.4.x
   </td>
   <td>2019-02-19
   </td>
   <td>2019-08-06
   </td>
   <td>1.11
   </td>
   <td>v1alpha1
   </td>
   <td>N/A
   </td>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>0.3.x
   </td>
   <td>2019-01-08
   </td>
   <td>2019-06-25
   </td>
   <td>1.11
   </td>
   <td>v1alpha1
   </td>
   <td>N/A
   </td>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>0.2.x
   </td>
   <td>2018-10-30
   </td>
   <td>2019-05-14
   </td>
   <td>1.09
   </td>
   <td>v1alpha1
   </td>
   <td>N/A
   </td>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>0.1.x
   </td>
   <td>2018-07-18
   </td>
   <td>2019-04-02
   </td>
   <td>1.09
   </td>
   <td>v1alpha1
   </td>
   <td>N/A
   </td>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>default
   </td>
  </tr>
</table>

### Knative Eventing Version Table

<table>
  <tr>
   <td><strong>Knative Version</strong>
   </td>
   <td><strong>Release Date</strong>
   </td>
   <td><strong>End of Life Date</strong>
   </td>
   <td><strong>Min K8s Version</strong>
   </td>
   <td><strong>Storage Type</strong>
   </td>
   <td><strong>Force Upgrade</strong>
   </td>
   <td><strong>API Endpoints</strong>
   </td>
   <td><strong>API Type Served</strong>
   </td>
   <td><strong>Availability</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.19.x
   </td>
   <td rowspan="2" >2020-12-22
   </td>
   <td rowspan="2" >2021-06-08
   </td>
   <td rowspan="2" >(undetermined)
   </td>
   <td rowspan="2" >v1
   </td>
   <td rowspan="2" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td><strong>removed</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.18.x
   </td>
   <td rowspan="2" >2020-09-29
   </td>
   <td rowspan="2" >2021-03-16
   </td>
   <td rowspan="2" >(undetermined)
   </td>
   <td rowspan="2" >v1
   </td>
   <td rowspan="2" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td><strong>not served</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.17.x
   </td>
   <td rowspan="3" >2020-08-18
   </td>
   <td rowspan="3" >2021-02-02
   </td>
   <td rowspan="3" >1.16.x
   </td>
   <td rowspan="3" >v1
   </td>
   <td rowspan="3" >Yes
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>optional (deprecated)
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td><strong>removed</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="3" >0.16.x
   </td>
   <td rowspan="3" >2020-07-07
   </td>
   <td rowspan="3" >2020-12-22
   </td>
   <td rowspan="3" >1.16.x
   </td>
   <td rowspan="3" >v1beta1
   </td>
   <td rowspan="3" >N/A
   </td>
   <td>v1
   </td>
   <td>v1
   </td>
   <td>optional
   </td>
  </tr>
  <tr>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td><strong>not served</strong>
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.15.x
   </td>
   <td rowspan="2" >2020-05-26
   </td>
   <td rowspan="2" >2020-11-10
   </td>
   <td rowspan="2" >1.16.x
   </td>
   <td rowspan="2" >v1beta1
   </td>
   <td rowspan="2" >Yes
   </td>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>optional (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.14.x
   </td>
   <td rowspan="2" >2020-04-14
   </td>
   <td rowspan="2" >2020-09-29
   </td>
   <td rowspan="2" >1.15.x
   </td>
   <td rowspan="2" >v1alpha1
   </td>
   <td rowspan="2" >N/A
   </td>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>default
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>optional (deprecated)
   </td>
  </tr>
  <tr>
   <td rowspan="2" >0.13.x
   </td>
   <td rowspan="2" >2020-03-03
   </td>
   <td rowspan="2" >2020-08-18
   </td>
   <td rowspan="2" >1.15.x
   </td>
   <td rowspan="2" >v1alpha1
   </td>
   <td rowspan="2" >N/A
   </td>
   <td>v1beta1
   </td>
   <td>v1beta1
   </td>
   <td>optional
   </td>
  </tr>
  <tr>
   <td>v1alpha1
   </td>
   <td>v1alpha1
   </td>
   <td>default
   </td>
  </tr>
</table>
