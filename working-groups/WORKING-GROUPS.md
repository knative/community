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

Some working groups (mostly those with a plug-in or extension model) end up
responsible for a set of GitHub repos, one for each extension. This allows for
easier dependency management; in these cases, one or more repo prefix names will
be recorded as canonical "extension names" to allow WGs to be responsible for
their own namespace without needing to get TOC approval for each repo name.

Additionally, all working groups should hold regular meetings, which should be
added to the
[shared knative calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com)
WG leads should have access to be able to create and update events on this
calendar, and should invite knative-dev@googlegroups.com to working group
meetings.

### Calendar import

If you're using Google Calendar, the above should work. If you're using some
other system (Apple Calendar or Outlook, for example),
[here is an iCal export of the community calendar](https://calendar.google.com/calendar/ical/knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com/public/basic.ics).

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
- [Security](#security)
- [User Experience](#user-experience)
  <!-- TODO add charters for each group -->

## API Core

API
[resources](https://github.com/knative/serving/tree/main/pkg/apis/serving),
[validation](https://github.com/knative/pkg/tree/main/webhook), and
[semantics](https://github.com/knative/pkg/tree/main/controller).

| Artifact                   | Link                                                                                                                                                            |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                             |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1HqLJ8uNm5KmKPlMzQmbfDy0haybEUrPydsp7gNBrl20/edit)                                        |
| Community Meeting Calendar | Wednesdays 10:30a-11:00a PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1HqLJ8uNm5KmKPlMzQmbfDy0haybEUrPydsp7gNBrl20/edit)                                                                   |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1eCSmaqJ4LYcuS3TlOqjW0xETnzLmo6Q9)                                                                              |
| Repos                      | [`knative/serving`](https://github.com/knative/serving)                                                                                                         |
| Slack Channel              | [#serving-api](https://slack.knative.dev/messages/serving-api)                                                                                                  |
| Github Team WG Leads       | [@knative/api-core-wg-leads](https://github.com/orgs/knative/teams/api-core-wg-leads/members)                                                                   |

| &nbsp;                                                   | Leads            | Company | Profile                                 |
| -------------------------------------------------------- | ---------------- | ------- | --------------------------------------- |
| <img width="30px" src="https://github.com/dprotaso.png"> | Dave Protasowski | VMware  | [dprotaso](https://github.com/dprotaso) |

## Client

[Client](https://github.com/knative/client), CLI, client libraries, and client
conventions

| Artifact                   | Link                                                                                                                                                                                                                                              |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                                                                                                               |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1cD7NkJJhSBpo2Q6RBHrbrSe6R5zjTZgO_YDGAluQ_oI/edit)                                                                                                                          |
| Community Meeting Calendar | Tuesdays, alternating between 10:30a-11:00a Pacific and 3:30p-4:00p Central European every two weeks<br>[Calendar Invitation](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1cD7NkJJhSBpo2Q6RBHrbrSe6R5zjTZgO_YDGAluQ_oI/edit)                                                                                                                                                     |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1QffYD_XM0vqaXDvFlZVsJHAT6-AahJCO)                                                                                                                                                                |
| Repos                      | [`knative/client`](https://github.com/knative/client), `kn-plugin-*`                                                                                                                                                                              |
| Slack Channel              | [#cli](https://slack.knative.dev)                                                                                                                                                                                                                 |
| Github Team WG leads       | [@knative/client-wg-leads](https://github.com/orgs/knative/teams/client-leads/members)                                                                                                                                                            |

| &nbsp;                                                      | Leads        | Company | Profile                                       |
| ----------------------------------------------------------- | ------------ | ------- | --------------------------------------------- |
| <img width="30px" src="https://github.com/navidshaikh.png"> | Navid Shaikh | VMware | [navidshaikh](https://github.com/navidshaikh) |
| <img width="30px" src="https://github.com/rhuss.png">       | Roland Huß   | Red Hat | [rhuss](https://github.com/rhuss)             |

## Documentation

Knative documentation, especially the [Docs](https://github.com/knative/docs/)
repo.

| Artifact                   | Link                                                                                                                                                                   |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-docs@](https://groups.google.com/forum/#!forum/knative-docs)                                                                                                  |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/13-neXVH-1n1ELWgukSzySnWj-4X6UB2pGGws2z2NH3M/edit)                                               |
| Community Meeting Calendar | Weekly on Tuesdays, 9:30-10:00am PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/13-neXVH-1n1ELWgukSzySnWj-4X6UB2pGGws2z2NH3M/edit)                                                                          |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1cMXOSiGo1kiWl0WXeGz8nV6esLetJS2f)                                                                                     |
| Repos                      | [`knative/docs`](https://github.com/knative/docs), [`knative/website`](https://github.com/knative/website)                                                             |
| Slack Channel              | [#docs](https://slack.knative.dev/messages/docs)                                                                                                                       |
| Github Team WG leads       | [@knative/docs-wg-leads](https://github.com/orgs/knative/teams/docs-wg-leads/members)                                                                                  |


| &nbsp;                                                     | Leads            | Company | Profile                                     |
| ---------------------------------------------------------- | ---------------- | ------- | ------------------------------------------- |

## Eventing

Event sources, bindings, FaaS framework, and orchestration

| Artifact                   | Link                                                                                                                                                         |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Charter / Mission          | [Enable asynchronous application development through event delivery from anywhere.](https://github.com/knative/eventing/blob/main/docs/mission.md)         |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                          |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1Xha-FeunojN49OJN7W0WBnPMcRtp1ycYpbkiir6XsE0/edit)                                     |
| Community Meeting Calendar | Wednesdays 9:00a-9:30a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1Xha-FeunojN49OJN7W0WBnPMcRtp1ycYpbkiir6XsE0/edit)                                                                |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/16A4v6Qv9MNSumpy6geGm0u1izDqAfiNo)                                                                           |
| Repos                      | [`knative/eventing`](https://github.com/knative/eventing), `eventing-*`                                                                                      |
| Slack Channel              | [#eventing](https://slack.knative.dev/messages/eventing)                                                                                                     |
| Github Team WG leads       | [@knative/eventing-wg-leads](https://github.com/orgs/knative/teams/eventing-wg-leads/members)                                                                |

| &nbsp;                                                        | Leads                      | Company | Profile                                           |
| ------------------------------------------------------------- | -------------------------- | ------- | ------------------------------------------------- |
| <img width="30px" src="https://github.com/vaikas.png">        | Ville Aikas (Technical)    | VMware  | [vaikas](https://github.com/vaikas)               |
| <img width="30px" src="https://github.com/lionelvillard.png"> | Lionel Villard (Technical) | IBM     | [lionelvillard](https://github.com/lionelvillard) |
| <img width="30px" src="https://github.com/grantr.png">        | Grant Rodgers (Technical)  | Google  | [grantr](https://github.com/grantr)               |
| <img width="30px" src="https://github.com/devguyio.png">      | Ahmed Abdalla (Execution)  | Red Hat | [devguyio](https://github.com/devguyio)           |

## Event Delivery

Event delivery data plane.

| Artifact                   | Link                                                                                                                                                        |
| -------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                         |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/10pvpEAb6DZ0mplmSeg8fEAmmJtqimt03GEHRieOcMig/edit)                                    |
| Community Meeting Calendar | Tuesdays 8:30a-9:00a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com)  |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/10pvpEAb6DZ0mplmSeg8fEAmmJtqimt03GEHRieOcMig/edit)                                                               |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1YcyeckwBNPApxh8vHAlscYe_w7sIA5TA)                                                                          |
| Repos                      | [`knative/eventing`](https://github.com/knative/eventing), [`knative/eventing-contrib`](https://github.com/knative/eventing-contrib), (shares `eventing-*`) |
| Slack Channel              | [#eventing-delivery](https://slack.knative.dev/messages/eventing-delivery)                                                                                  |
| Github Team WG leads       | [@knative/channel-wg-leads](https://github.com/orgs/knative/teams/channel-wg-leads/members)                                                               |

| &nbsp;                                                   | Leads               | Company | Profile                                 |
| -------------------------------------------------------- | ------------------- | ------- | --------------------------------------- |
| <img width="30px" src="https://github.com/matzew.png">   | Matthias Wessendorf | Red Hat | [matzew](https://github.com/matzew)     |
| <img width="30px" src="https://github.com/slinkydeveloper.png">   | Francesco Guardiani | Red Hat | [slinkydeveloper](https://github.com/slinkydeveloper)     |

## Eventing Sources

Event producers and frameworks.

| Artifact                   | Link                                                                                                                                                       |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                        |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/19txVRqA6_eY6ClGqoLRa0gPB50Ok7PT6_B6zDP1KtKQ/edit)                                   |
| Community Meeting Calendar | Tuesdays 8:30a-9:00a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/19txVRqA6_eY6ClGqoLRa0gPB50Ok7PT6_B6zDP1KtKQ/edit)                                                              |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/13idi0QRqj1vsp5X0vDGfUOULhIyeeLaG)                                                                         |
| Repo prefixes              | (shares `eventing-*` with the eventing WG)                                                                                                                 |
| Slack Channel              | [#eventing-sources](https://slack.knative.dev/messages/eventing-sources)                                                                                   |
| Github Team WG leads       | [@knative/source-wg-leads](https://github.com/orgs/knative/teams/source-wg-leads/members)                                                                  |

| &nbsp;                                                        | Leads          | Company | Profile                                           |
| ------------------------------------------------------------- | -------------- | ------- | ------------------------------------------------- |
| <img width="30px" src="https://github.com/n3wscott.png">      | Scott Nichols  | VMware  | [n3wscott](https://github.com/n3wscott)           |
| <img width="30px" src="https://github.com/vaikas.png">        | Ville Aikas    | VMware  | [vaikas](https://github.com/vaikas)               |
| <img width="30px" src="https://github.com/lionelvillard.png"> | Lionel Villard | IBM     | [lionelvillard](https://github.com/lionelvillard) |

## Networking

Inbound and outbound network connectivity for
[serving](https://github.com/knative/serving) workloads. Specific areas of
interest include: load balancing, routing, DNS configuration and TLS support.

| Artifact                   | Link                                                                                                                                                           |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                            |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/18m2XUDe-5QMFFBpLhIC7ozFnq12niYAqebtS3nMorhk/edit)                                       |
| Community Meeting Calendar | Thursdays at 9:00a-9:30a PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/18m2XUDe-5QMFFBpLhIC7ozFnq12niYAqebtS3nMorhk/edit)                                                                  |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1bx86aDXjXhylDvFmEpjWSMn8o80_2_nR)                                                                             |
| Repos                      | `net-*`                                                                                                                                                        |
| Slack Channel              | [#networking](https://slack.knative.dev/messages/networking)                                                                                                   |
| Github Team WG leads       | [@knative/networking-wg-leads](https://github.com/orgs/knative/teams/networking-wg-leads/members)                                                              |

| &nbsp;                                                      | Leads            | Company   | Profile                                       |
| ----------------------------------------------------------- | ---------------- | --------- | --------------------------------------------- |
| <img width="30px" src="https://github.com/nak3.png">        | Kenjiro Nakayama | Red Hat   | [nak3](https://github.com/nak3)               |
| <img width="30px" src="https://github.com/tcnghia.png">     | Nghia Tran       | Microsoft | [tcnghia](https://github.com/tcnghia)         |
| <img width="30px" src="https://github.com/ZhiminXiang.png"> | Zhimin Xiang     | Google    | [ZhiminXiang](https://github.com/ZhiminXiang) |

## Operations

Managing, assessing system health and maintaining Knative clusters

| Artifact                   | Link                                                                                                                                                       |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                        |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1m9oFlelI292Fzwi_sRUTrf69I-CcCwEVXkGcMAjDYqM/edit)                                   |
| Community Meeting Calendar | Tuesdays at 10:00am PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1m9oFlelI292Fzwi_sRUTrf69I-CcCwEVXkGcMAjDYqM/edit)                                                              |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1GgBCtyWWGx8v8PHGbzJ_tu7DROyBl6x7)                                                                         |
| Repo prefixes              |                                                                                                                                                            |
| Slack Channel              | [#operations](https://slack.knative.dev/messages/operations)                                                                                               |
| Github Team WG leads       | [@knative/operations-wg-leads](https://github.com/orgs/knative/teams/operations-wg-leads/members)                                                          |

| &nbsp;                                                     | Leads       | Company | Profile                                     |
| ---------------------------------------------------------- | ----------- | ------- | ------------------------------------------- |
| <img width="30px" src="https://github.com/houshengbo.png"> | Vincent Hou | IBM     | [houshengbo](https://github.com/houshengbo) |

## Productivity

Project health, test framework, continuous integration & deployment, release,
performance/scale/load testing infrastructure

| Artifact                   | Link                                                                                                                                                  |
| -------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                   |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/16go22yTCiaNBtdjghhrqSxnsWkMbxRs4i_gSIezGzUk/edit)                              |
| Community Meeting Calendar | Thursdays, 10am PST<br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/16go22yTCiaNBtdjghhrqSxnsWkMbxRs4i_gSIezGzUk/edit)                                                         |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1_1oWL7skjVt2211T0aagpwzDEfWmEIQK)                                                                    |
| Repo prefixes              | `actions-*`                                                                                                                                           |
| Slack Channel              | [#productivity](https://slack.knative.dev/messages/productivity)                                                                                      |
| Github Team WG leads       | [@knative/productivity-wg-leads](https://github.com/orgs/knative/teams/productivity-wg-leads/members)                                                 |

| &nbsp;                                                   | Leads         | Company | Profile                                 |
| -------------------------------------------------------- | ------------- | ------- | --------------------------------------- |
| <img width="30px" src="https://github.com/chizhg.png">   | Chi Zhang     | Google  | [chizhg](https://github.com/chizhg)     |
| <img width="30px" src="https://github.com/n3wscott.png"> | Scott Nichols | VMware  | [n3wscott](https://github.com/n3wscott) |

## Scaling

Autoscaling behavior of Knative Serving

| Artifact                   | Link                                                                                                                                                        |
| -------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                         |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1-hezRkkLhnCkkEm6S09SkQR-BpF5xpTyhethYL5Zm6w/edit#)             |
| Community Meeting Calendar | Wednesdays at 9:30am PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1-hezRkkLhnCkkEm6S09SkQR-BpF5xpTyhethYL5Zm6w/edit#)                                        |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1IDDkJ3FD47xFSHY3iA9U2Q8th3Cwdo0K)                                                                          |
| Repo prefixes              |                                                                                                                                                             |
| Slack Channel              | [#autoscaling](https://slack.knative.dev/messages/autoscaling)                                                                                              |
| Github Team WG leads       | [@knative/autoscaling-wg-leads](https://github.com/orgs/knative/teams/autoscaling-wg-leads/members)                                                         |

| &nbsp;                                                         | Leads           | Company | Profile                                             |
| -------------------------------------------------------------- | --------------- | ------- | --------------------------------------------------- |
| <img width="30px" src="https://github.com/markusthoemmes.png"> | Markus Thömmes  | Red Hat | [markusthoemmes](https://github.com/markusthoemmes) |
| <img width="30px" src="https://github.com/vagababov.png">      | Victor Agababov | Google  | [vagababov](https://github.com/vagababov)           |

## Security

Security concerns across Knative components

| Artifact                   | Link                                                                                                                                                   |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Charter / Mission          | [Charter](https://docs.google.com/document/d/194bECSbxPIONzhgGj1g9tccUsP8I-M7WIvbBqcT-tXY/edit)                                                        |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                    |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1b-aCADlUBo2FZEDp5z5XoXlfUT6CvED339fcAdKY66o/edit)                               |
| Community Meeting Calendar | Tuesdays at 9am PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1b-aCADlUBo2FZEDp5z5XoXlfUT6CvED339fcAdKY66o/edit)                                                          |
| Document Folder            | [Folder](https://drive.google.com/drive/u/0/folders/1yOy_yU5vvUTFuxeYvhlY7uzIL4c4SSzX)                                                                 |
| Repo prefixes              |                                                                                                                                                        |
| Slack Channel              | [#security](https://slack.knative.dev/messages/security)                                                                                               |
| Github Team WG leads       | [@knative/security-wg-leads](https://github.com/orgs/knative/teams/security-wg-leads/members)                                                         |

| &nbsp;                                                        | Leads           | Company | Profile                                             |
| ------------------------------------------------------------- | --------------- | ------- | --------------------------------------------------- |
| <img width="30px" src="https://github.com/evankanderson.png"> | Evan Anderson   | VMware  | [evankanderson](https://github.com/evankanderson)   |
| <img width="30px" src="https://github.com/julz.png">          | Julian Friedman | IBM     | [julz](https://github.com/julz)                     |

## User Experience

User Experience concerns across Knative components

| Artifact                   | Link                                                                                                                                                    |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Charter / Mission          | [Charter](https://docs.google.com/document/d/1K_LGymgqedAzRS0ACdVtJH1ktQxrf1OatfxzIb6LjU0/edit?usp=sharing)                                             |
| Forum                      | [knative-dev@](https://groups.google.com/forum/#!forum/knative-dev)                                                                                     |
| Community Meeting VC       | See the top of the [Meeting notes](https://docs.google.com/document/d/1NSlGHen5Dh6c2A0LGavGWibddWrlSLV_PbEbMpNjSTU/edit?usp=sharing)                    |
| Community Meeting Calendar | Thursdays at 8am PST <br>[Calendar](https://calendar.google.com/calendar/embed?src=knative.team_9q83bg07qs5b9rrslp5jor4l6s%40group.calendar.google.com) |
| Meeting Notes              | [Notes](https://docs.google.com/document/d/1NSlGHen5Dh6c2A0LGavGWibddWrlSLV_PbEbMpNjSTU/edit?usp=sharing)                                               |
| Roadmap                    | [Roadmap](https://github.com/orgs/knative/projects/20)                                                                                                  |
| Document Folder            | [Folder](https://drive.google.com/drive/folders/1XzZGqV7yHo38d_l7rH1uSIrbQp3JlbBP?usp=sharing)                                                          |
| Slack Channel              | [#user-experience](https://slack.knative.dev/messages/user-experience)                                                                                  |
| Github Repository          | [#user-experience](https://github.com/knative/ux)                                                                                                       |
| Github Team WG leads       | [@knative/ux-wg-leads](https://github.com/orgs/knative/teams/ux-wg-leads/members)                                                                       |

| &nbsp;                                                        | Leads           | Company | Profile                                             |
| ------------------------------------------------------------- | --------------- | ------- | --------------------------------------------------- |
| <img width="30px" src="https://github.com/omerbensaadon.png"> | Omer Bensaadon  | VMware  | [omerbensaadon](https://github.com/omerbensaadon)   |
| <img width="30px" src="https://github.com/csantanapr.png">    | Carlos Santana  | IBM     | [csantanapr](https://github.com/csantanapr)         |

---

Except as otherwise noted, the content of this page is licensed under the
[Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/),
and code samples are licensed under the
[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).

