---
title: "Knative working group"
linkTitle: "Join working groups"
weight: 25
type: "docs"
aliases:
  - /contributing/working-groups/
---

Most community activity is organized into _working groups_.

Working groups follow the [contributing](../CONTRIBUTING.md) guidelines although
each of these groups may operate a little differently depending on their needs
and workflow.

When the need arises, a new working group can be created. See the
[working group processes](../mechanics/WORKING-GROUP-PROCESSES.md) for working
group proposal and creation procedures.

The working groups generate design docs which are kept in a
[shared drive](https://drive.google.com/corp/drive/folders/0APnJ_hRs30R2Uk9PVA)
and are available for anyone to read and comment on. The shared drive currently
grants read access to
[knative-users@](https://groups.google.com/forum/#!forum/knative-users) and edit
and comment access to the
[knative-dev@](https://groups.google.com/forum/#!forum/knative-dev) Google
group.

Additionally, all working groups should hold regular meetings, which should be
added to the
[shared knative calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com)
WG leads should have access to be able to create and update events on this
calendar, and should invite knative-dev@googlegroups.com to working group
meetings.

### Calendar import

If you're using Google Calendar, the above should work. If you're using some
other system (Apple Calendar or Outlook, for example),
[here is an iCal export of the community calendar](https://calendar.google.com/calendar/ical/google.com_18un4fuh6rokqf8hmfftm5oqq4@group.calendar.google.com/public/basic.ics).

- [Follow these directions to import into Outlook Web](https://support.office.com/en-us/article/import-or-subscribe-to-a-calendar-in-outlook-on-the-web-503ffaf6-7b86-44fe-8dd6-8099d95f38df)
- [Follow these directions for desktop Outlook](https://support.office.com/en-us/article/See-your-Google-Calendar-in-Outlook-C1DAB514-0AD4-4811-824A-7D02C5E77126)
- [Follow the import directions to import into Apple Calendar](https://support.apple.com/guide/calendar/import-or-export-calendars-icl1023/mac)

# Working Groups

The current working groups are:

- [API Core](#api-core)
- [Client](#client)
- [Documentation](#documentation)
- [Eventing](#eventing)
- [Event Delivery](#event-delivery)
- [Eventing Sources](#eventing-sources)
- [Networking](#networking)
- [Operations](#operations)
- [Productivity](#productivity)
- [Scaling](#scaling)
  <!-- TODO add charters for each group -->

## API Core

API
[resources](https://github.com/knative/serving/tree/master/pkg/apis/serving),
[validation](https://github.com/knative/pkg/tree/master/webhook), and
[semantics](https://github.com/knative/pkg/tree/master/controller).

| Artifact                   | Link                                                                                                                                                          |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                           |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1NC4klOdNaU-N-PsKLyXBqDKgNSHtxCDep29Ta2b5FK0/edit)                                      |
| Community Meeting Calendar | Wednesdays 10:30a-11:00a PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1NC4klOdNaU-N-PsKLyXBqDKgNSHtxCDep29Ta2b5FK0/edit)                                                                 |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1fpBW7VyiBISsKuVdgn1MrgFdtx_JGoC5)                                                                       |
| Slack Channel              | [#serving-api](https://slack.knative.dev/messages/serving-api)                                                                                                |

| &nbsp;                                                   | Leads            | Company | Profile                                 |
| -------------------------------------------------------- | ---------------- | ------- | --------------------------------------- |
| <img width="30px" src="https://github.com/mattmoor.png"> | Matt Moore       | VMware  | [mattmoor](https://github.com/mattmoor) |
| <img width="30px" src="https://github.com/dprotaso.png"> | Dave Protasowski | VMware  | [dprotaso](https://github.com/dprotaso) |

## Client

[Client](https://github.com/knative/client), CLI, client libraries, and client
conventions

| Artifact                   | Link                                                                                                                                                                                                                                            |
| -------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                                                                                                             |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1Uh7jTWQruBBmic-WmTvtc9cMF95kQrKb5lsqWhNuikM/edit)                                                                                                                        |
| Community Meeting Calendar | Tuesdays, alternating between 10:30a-11:00a Pacific and 3:30p-4:00p Central European every two weeks<br>[Calendar Invitation](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1Uh7jTWQruBBmic-WmTvtc9cMF95kQrKb5lsqWhNuikM/edit)                                                                                                                                                   |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1QF-job3rCEqCpJLm8nHkC4mBIi4XANE1)                                                                                                                                                         |
| Slack Channel              | [#cli](https://slack.knative.dev)                                                                                                                                                                                                               |

| &nbsp;                                                      | Leads        | Company | Profile                                       |
| ----------------------------------------------------------- | ------------ | ------- | --------------------------------------------- |
| <img width="30px" src="https://github.com/sixolet.png">     | Naomi Seyfer | Google  | [sixolet](https://github.com/sixolet)         |
| <img width="30px" src="https://github.com/navidshaikh.png"> | Navid Shaikh | Red Hat | [navidshaikh](https://github.com/navidshaikh) |
| <img width="30px" src="https://github.com/rhuss.png">       | Roland Huß   | Red Hat | [rhuss](https://github.com/rhuss)             |

## Documentation

Knative documentation, especially the [Docs](https://github.com/knative/docs/)
repo.

| Artifact                   | Link                                                                                                                                                               |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Forum                      | [knative-docs@](https://groups.google.com/forum/#!forum/knative-docs)                                                                                              |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1Y7rug0XshcQPdKzptdWbQLQjcjgpFdLeEgP1nfkDAe4/edit)                                           |
| Community Meeting Calendar | Bi Weekly on Tuesdays, 9:00-9:30am<br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1Y7rug0XshcQPdKzptdWbQLQjcjgpFdLeEgP1nfkDAe4/edit)                                                                      |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1K5cM9m-b93ySI5WGKalJKbBq_cfjyi-y)                                                                            |
| Slack Channel              | [#docs](https://slack.knative.dev/messages/docs)                                                                                                                   |

| &nbsp;                                                     | Leads            | Company | Profile                                     |
| ---------------------------------------------------------- | ---------------- | ------- | ------------------------------------------- |
| <img width="30px" src="https://github.com/abrennan89.png"> | Ashleigh Brennan | Red Hat | [abrennan89](https://github.com/abrennan89) |

## Eventing

Event sources, bindings, FaaS framework, and orchestration

| Artifact                   | Link                                                                                                                                                       |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Charter / Mission          | [Enable asynchronous application development through event delivery from anywhere.](https://github.com/knative/eventing/blob/master/docs/mission.md)       |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                        |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1uGDehQu493N_XCAT5H4XEw5T9IWlPN1o19ULOWKuPnY/edit)                                   |
| Community Meeting Calendar | Wednesdays 9:00a-9:30a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1uGDehQu493N_XCAT5H4XEw5T9IWlPN1o19ULOWKuPnY/edit)                                                              |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1S22YmGl6B1ppYApwa1j5j9Nc6rEChlPo)                                                                    |
| Slack Channel              | [#eventing](https://slack.knative.dev/messages/eventing)                                                                                                   |

| &nbsp;                                                        | Leads                      | Company | Profile                                           |
| ------------------------------------------------------------- | -------------------------- | ------- | ------------------------------------------------- |
| <img width="30px" src="https://github.com/vaikas.png">        | Ville Aikas (Technical)    | VMware  | [vaikas](https://github.com/vaikas)               |
| <img width="30px" src="https://github.com/lionelvillard.png"> | Lionel Villard (Technical) | IBM     | [lionelvillard](https://github.com/lionelvillard) |
| <img width="30px" src="https://github.com/grantr.png">        | Grant Rodgers (Technical)  | Google  | [grantr](https://github.com/grantr)               |
| <img width="30px" src="https://github.com/devguyio.png">      | Ahmed Abdalla (Execution)  | Red Hat | [devguyio](https://github.com/devguyio)           |

## Event Delivery

Event delivery data plane.

| Artifact                   | Link                                                                                                                                                     |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                      |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1uxlulaAf2m_yZUqCIeI-inul2gsqP69PElnZdO0FHUo/edit#)                                |
| Community Meeting Calendar | Tuesdays 8:30a-9:00a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1uxlulaAf2m_yZUqCIeI-inul2gsqP69PElnZdO0FHUo/edit#)                                                           |
| Document Folder            | [Folder](https://drive.google.com/drive/u/0/folders/1gQymVBlLsQxknScnn1x5wUT5OBBNj5P3)                                                                   |
| Slack Channel              | [#eventing-delivery](https://slack.knative.dev/messages/eventing-delivery)                                                                               |

| &nbsp;                                                   | Leads               | Company | Profile                                 |
| -------------------------------------------------------- | ------------------- | ------- | --------------------------------------- |
| <img width="30px" src="https://github.com/Harwayne.png"> | Adam Harwayne       | Google  | [Harwayne](https://github.com/Harwayne) |
| <img width="30px" src="https://github.com/matzew.png">   | Matthias Wessendorf | Red Hat | [matzew](https://github.com/matzew)     |

## Eventing Sources

Event producers and frameworks.

| Artifact                   | Link                                                                                                                                                     |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                      |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/19txVRqA6_eY6ClGqoLRa0gPB50Ok7PT6_B6zDP1KtKQ/edit#)                                |
| Community Meeting Calendar | Tuesdays 8:30a-9:00a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/19txVRqA6_eY6ClGqoLRa0gPB50Ok7PT6_B6zDP1KtKQ/edit#)                                                           |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1tDlLtLVNXghg8Fn8KlVGr5ZTvkiQjZJs)                                                                       |
| Slack Channel              | [#eventing-sources](https://slack.knative.dev/messages/eventing-sources)                                                                                 |

| &nbsp;                                                        | Leads                | Company | Profile                                           |
| ------------------------------------------------------------- | -------------------- | ------- | ------------------------------------------------- |
| <img width="30px" src="https://github.com/n3wscott.png">      | Scott Nichols        | VMware  | [n3wscott](https://github.com/n3wscott)           |
| <img width="30px" src="https://github.com/vaikas.png">        | Ville Aikas          | VMware  | [vaikas](https://github.com/vaikas)               |
| <img width="30px" src="https://github.com/lionelvillard.png"> | Lionel Villard       | IBM     | [lionelvillard](https://github.com/lionelvillard) |
| <img width="30px" src="https://github.com/nachocano.png">     | Ignacio (Nacho) Cano | Google  | [nachocano](https://github.com/nachocano)         |

## Networking

Inbound and outbound network connectivity for
[serving](https://github.com/knative/serving) workloads. Specific areas of
interest include: load balancing, routing, DNS configuration and TLS support.

| Artifact                   | Link                                                                                                                                                         |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                          |
| Community Meeting VC       | See the top of the [Meeting notes](https://drive.google.com/open?id=1EE1t5mTfnTir2lEasdTMRNtuPEYuPqQCZbU3NC9mHOI)                                            |
| Community Meeting Calendar | Thursdays at 9:00a-9:30a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://drive.google.com/open?id=1EE1t5mTfnTir2lEasdTMRNtuPEYuPqQCZbU3NC9mHOI)                                                                       |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1oVDYbcEDdQ9EpUmkK6gE4C7aZ8u6ujsN)                                                                      |
| Slack Channel              | [#networking](https://slack.knative.dev/messages/networking)                                                                                                 |

| &nbsp;                                                      | Leads            | Company   | Profile                                       |
| ----------------------------------------------------------- | ---------------- | --------- | --------------------------------------------- |
| <img width="30px" src="https://github.com/nak3.png">        | Kenjiro Nakayama | Red Hat   | [nak3](https://github.com/nak3)               |
| <img width="30px" src="https://github.com/tcnghia.png">     | Nghia Tran       | Microsoft | [tcnghia](https://github.com/tcnghia)         |
| <img width="30px" src="https://github.com/ZhiminXiang.png"> | Zhimin Xiang     | Google    | [ZhiminXiang](https://github.com/ZhiminXiang) |

## Operations

Managing, assessing system health and maintaining Knative clusters

| Artifact                   | Link                                                                                                                                                     |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                      |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1DoCG0VGZ0_Nj84Ci443bEefSOnJAWVDXUbcYR-viVfY/edit#heading=h.redlkj86bwwn)          |
| Community Meeting Calendar | Tuesdays at 10:00am PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1DoCG0VGZ0_Nj84Ci443bEefSOnJAWVDXUbcYR-viVfY/edit#heading=h.redlkj86bwwn)                                     |
| Document Folder            | [Folder](https://drive.google.com/drive/u/0/folders/14AI7ClIq2btPZ13WT8BAN4AtNKpK_K6p)                                                                   |
| Slack Channel              | [#operations](https://slack.knative.dev/messages/operations)                                                                                             |

| &nbsp;                                                     | Leads       | Company | Profile                                     |
| ---------------------------------------------------------- | ----------- | ------- | ------------------------------------------- |
| <img width="30px" src="https://github.com/houshengbo.png"> | Vincent Hou | IBM     | [houshengbo](https://github.com/houshengbo) |

## Scaling

Autoscaling

| Artifact                   | Link                                                                                                                                                      |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                       |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1FoLJqbDJM8_tw7CON-CJZsO2mlF8Ia1cWzCjWX8HDAI/edit#heading=h.c0ufqy5rucfa)           |
| Community Meeting Calendar | Wednesdays at 9:30am PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1FoLJqbDJM8_tw7CON-CJZsO2mlF8Ia1cWzCjWX8HDAI/edit#heading=h.c0ufqy5rucfa)                                      |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1qpGIPXVGoMm6IXb74gPrrHkudV_bjIZ9)                                                                   |
| Slack Channel              | [#autoscaling](https://slack.knative.dev/messages/autoscaling)                                                                                            |

| &nbsp;                                                         | Leads           | Company | Profile                                             |
| -------------------------------------------------------------- | --------------- | ------- | --------------------------------------------------- |
| <img width="30px" src="https://github.com/markusthoemmes.png"> | Markus Thömmes  | Red Hat | [markusthoemmes](https://github.com/markusthoemmes) |
| <img width="30px" src="https://github.com/vagababov.png">      | Victor Agababov | Google  | [vagababov](https://github.com/vagababov)           |

## Productivity

Project health, test framework, continuous integration & deployment, release,
performance/scale/load testing infrastructure

| Artifact                   | Link                                                                                                                                                |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                 |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1aPRwYGD4XscRIqlBzbNsSB886PJ0G-vZYUAAUjoydko)                                 |
| Community Meeting Calendar | Thursdays, 10am PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=google.com_18un4fuh6rokqf8hmfftm5oqq4%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1aPRwYGD4XscRIqlBzbNsSB886PJ0G-vZYUAAUjoydko)                                                            |
| Document Folder            | [Folder](https://drive.google.com/corp/drive/folders/1oMYB4LQHjySuMChmcWYCyhH7-CSkz2r_)                                                             |
| Slack Channel              | [#productivity](https://slack.knative.dev/messages/productivity)                                                                                    |

| &nbsp;                                                   | Leads         | Company | Profile                                 |
| -------------------------------------------------------- | ------------- | ------- | --------------------------------------- |
| <img width="30px" src="https://github.com/chizhg.png">   | Chi Zhang     | Google  | [chizhg](https://github.com/chizhg)     |
| <img width="30px" src="https://github.com/n3wscott.png"> | Scott Nichols | VMware  | [n3wscott](https://github.com/n3wscott) |

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
