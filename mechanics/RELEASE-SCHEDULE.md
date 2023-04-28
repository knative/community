---
title: "Knative Release Schedule"
linkTitle: "Release schedule"
weight: 50
type: "docs"
---

The Knative project releases quarterly, on Tuesday of the 4th week of January, April, July, October, represented by the following Crontab expression:

```
# minute   hour   day      month        weekday
0          0      22-28    1,4,7,10     2
```

The end of life (EOL) date for given release is calculated, by taking release date of a `n+2` release, and adding one week to it.

As much as possible, releases should be driven by automation, and repos should be ready to release at any point. It should also be possible to produce and consume nightly artifacts.

With that said, it can be useful to have a list of when future releases will happen, so this document provides a schedule for the next 12 months of releases and a list of supported and historical releases.

## Upcoming releases

| Release | Date       | EOL        | Min K8s Version | Notes |
| ------- | ---------- | ---------- | --------------- | ----- |
| 1.10    | 2023-04-25 | 2023-10-31 | 1.24            |       |
| 1.11    | 2023-07-25 | 2024-01-30 | 1.25            |       |
| 1.12    | 2023-10-24 | 2024-04-30 | 1.26            |       |
| 1.13    | 2024-01-23 | 2024-07-30 | -               |       |
| 1.14    | 2024-04-23 | 2024-10-29 | -               |       |
| 1.15    | 2024-07-23 | 2025-01-28 | -               |       |
| 1.16    | 2024-10-22 | 2025-04-29 | -               |       |

## Releases supported by community

| Release | Date       | EOL        | Min K8s Version | Notes |
| ------- | ---------- | ---------- | --------------- | ----- |
| 1.10    | 2023-04-25 | 2023-10-31 | 1.24            |       |
| 1.9     | 2023-01-24 | 2023-08-01 | 1.24            |       |

## No longer supported releases

| Release | Date       | EOL        | Min K8s Version | Notes                                          |
| ------- | ---------- | ---------- | --------------- | ---------------------------------------------- |
| 1.8     | 2022-10-18 | 2023-05-02 | 1.23            |                                                |
| 1.7     | 2022-08-23 | 2023-02-21 | 1.22            |                                                |
| 1.6     | 2022-07-12 | 2023-01-10 | 1.22            |                                                |
| 1.5     | 2022-05-31 | 2022-11-29 | 1.22            |                                                |
| 1.4     | 2022-04-19 | 2022-10-11 | 1.22            |                                                |
| 1.3     | 2022-03-08 | 2022-08-30 | 1.21            |                                                |
| 1.2     | 2022-01-25 | 2022-07-12 | 1.21            |                                                |
| 1.1     | 2021-12-14 | 2022-05-31 | 1.20            |                                                |
| 1.0     | 2021-11-02 | 2022-04-19 | 1.20            |                                                |
| 0.26    | 2021-09-21 | 2022-03-08 | 1.19            |                                                |
| 0.25    | 2021-08-10 | 2022-01-25 | 1.19            |                                                |
| 0.24    | 2021-06-29 | 2021-12-14 | 1.19            |                                                |
| 0.23    | 2021-05-18 | 2021-11-02 | 1.18            | Removes v1beta1 Eventing APIs                  |
| 0.22    | 2021-04-06 | 2021-09-21 | 1.18            |                                                |
| 0.21    | 2021-02-23 | 2021-08-10 | 1.18            |                                                |
| 0.20    | 2021-01-12 | 2021-06-29 | 1.17            | ** Moved by 3 weeks for end of year holidays** |
| 0.19    | 2020-11-10 | 2021-05-18 | 1.17            | Removes v1alpha1 and v1beta1 Serving APIs      |
| 0.18    | 2020-09-29 | 2021-03-16 | 1.17            |                                                |
| 0.17    | 2020-08-18 | 2021-02-02 | 1.16            | Introduces v1 Eventing APIs, removes v1alpha1  |
| 0.16    | 2020-07-07 | 2020-12-22 | 1.16            |                                                |
| 0.15    | 2020-05-26 | 2020-11-10 | 1.16            |                                                |
| 0.14    | 2020-04-14 | 2020-09-29 | 1.15            | Deprecates v1alpha1 Eventing APIs              |
| 0.13    | 2020-03-03 | 2020-08-18 | 1.15            |                                                |
