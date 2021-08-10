---
title: "Knative feature sunsetting"
linkTitle: "Sunsetting features"
weight: 50
type: "docs"
---

# Knative Feature Sunsetting

If a feature (especially in alpha stage) is getting no apparent usage and is
creating mostly busy-work for the community to maintain, the respective working
group should consider to sunset the respective feature. Other cost-benefit-ratio
considerations can lead to sunsetting of features/APIs too, for example if a
beta feature turns out to be a serious scalability issue.

This needs to be in accordance to the
[Knative release principles](RELEASE-VERSIONING-PRINCIPLES.md#api-support-principle)
so the duration of the following process might vary depending on the state of
the respective feature.

## 1. Working Group decision

First of all, there should be a decision at the working group level that there
is a desire to remove the respective feature. This decision should also contain
an explanation, stating why the feature can and should be dropped from Knative.

## 2. Gauge usage

While the working group itself might think there is no usage, that feeling
should be verified by gathering hard data.

### Ask vendors if they explicitly ship the feature

The quickest way of gauging the potential impact of the sunset is to ask vendors
that productize and ship Knative in downstream products if they are shipping the
respective feature explicitly. Explicitly here means that they specifically
chose to enable the respective feature. A good overview of existing downstream
products and projects can been in the
[documentation](https://knative.dev/docs/knative-offerings/).

A post should be sent to
[knative-dev@](https://groups.google.com/forum/#!forum/knative-dev), explaining
which feature is supposed to be removed and why it is supposed to be removed.
Lazy consensus can be applied over a period of two weeks.

### Ask users if they are using the feature

In the same way, it is very important to get to know if any user is actually
using the feature. A post should be sent to
[knative-users@](https://groups.google.com/forum/#!forum/knative-users), with
essentially the same contents as the post to
[knative-dev@](https://groups.google.com/forum/#!forum/knative-dev). Again, lazy
consensus can be applied over a period of two weeks.

## 3. Reevaluate the decision

Once the data has been gathered, the working group should reconsider the
decision with the collected data in mind.

## 4. Sunset

If the decision to remove the feature remains, it can now actually be removed.
This might either be a deprecation notice for Beta and GA feature or a straight
up deletion of the respective feature for alpha features. The deprecation notice
should be visible to both operators and users of the system. For example,
Knative could add a warning notice to the status of entities that use deprecated
features.

Likewise, the documentation should be updated. Deprecation notices should be put
up to the respective features and the deletion of a feature should be
accompanied with a deletion of the relevant documentation.
