---
title: "Knative Release Schedule"
linkTitle: "Release schedule"
weight: 50
type: "docs"
---

Knative releases every 6 weeks. As much as possible, releases should be driven by automation, and repos should be ready to release at any point. It should also be possible to produce and consume nightly artifacts.

With that said, it can be useful to have a list of when future releases will happen, so this document provides a schedule for the next 6+ months of releases.

| Release | Date       | EOL        | Min K8s Version | Notes                         |
| ------- | ---------- | ---------- | --------------- | ----------------------------- |
| 0.24    | 2021-06-29 | 2021-12-14 | 1.19            | |
| 0.25    | 2021-08-10 | 2022-01-25 | 1.19            | |
| 0.26    | 2021-09-21 | 2022-03-08 | 1.19            | |
| 1.0     | 2021-11-02 | 2022-04-19 | 1.20            | |
| 1.1     | 2021-12-14 | 2022-05-31 | 1.20            | |
| 1.2     | 2022-01-25 | 2022-06-12 | 1.21            | |
| 1.3     | 2022-03-08 | 2022-07-24 | 1.21            | |


## Historical (no longer supported) releases:

| Release | Date       | EOL        | Min K8s Version | Notes                    |
| ------- | ---------- | ---------- | --------------- | ------------------------ |
| 0.13    | 2020-03-03 | 2020-08-18 | 1.15            | |
| 0.14    | 2020-04-14 | 2020-09-29 | 1.15            | Deprecates v1alpha1 Eventing APIs |
| 0.15    | 2020-05-26 | 2020-11-10 | 1.16            | |
| 0.16    | 2020-07-07 | 2020-12-22 | 1.16            | |
| 0.17    | 2020-08-18 | 2021-02-02 | 1.16            | Introduces v1 Eventing APIs, removes v1alpha1 |
| 0.18    | 2020-09-29 | 2021-03-16 | 1.17            | |
| 0.19    | 2020-11-10 | 2021-05-18 | 1.17            | Removes v1alpha1 and v1beta1 Serving APIs |
| 0.20    | 2021-01-12 | 2021-06-29 | 1.17            | ** Moved by 3 weeks for end of year holidays** |
| 0.21    | 2021-02-23 | 2021-08-10 | 1.18            | |
| 0.22    | 2021-04-06 | 2021-09-21 | 1.18            | |
| 0.23    | 2021-05-18 | 2021-11-02 | 1.18            | Removes v1beta1 Eventing APIs |
