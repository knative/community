---
title: "Knative Golang policy"
linkTitle: "Golang policy"
weight: 50
type: "docs"
---

Knative's components are mostly written in Golang. As such the project has to deal with
major and minor upgrades of the Go language itself. The following principles are applied
regarding the Golang version Knative tests and releases with.

## Recency

Knative generally strives to always use the most recent Golang version available, to
ensure future compatibility and to benefit from the ongoing optimization to the Golang
runtime.

CI systems (Prow and Github Actions) should try to always be on the most recent patch
level.

## Caution

Major Golang version bumps (i.e. 1.14 to 1.15) should be made deliberately and should not
be taken too lightly. Such a change should not happen too closely to a release of
Knative. If the bump would be within two weeks of the Knative release cycle, it should be
postponed until the next release to allow for thorough testing and hardening. Exceptions
to this should be discussed and agreed upon through the TOC.

Since such a version bump can potentially be breaking to some repositories, the change
should be announced and discussed and agreed upon between the impacted working groups.
