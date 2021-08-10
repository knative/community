# Knative Event Delivery Working Group Charter

Last Updated: Nov 15, 2019

Note: the original document was written to focus on Knative Channels, but the TOC recommended refocusing on consistent data-plane behavior and the group renamed. The charter still contains many references to Channels.

## Mission

The Knative Event Delivery Working Group will reduce the friction of finding, implementing, and using Knative compliant Channels. We will specify the interface for event delivery such that multiple implementations can be used by higher level pieces. As well as create conformance tests to verify a given implementation implements the interface correctly and make that information available for operators and developers.

Event Delivery is the fundamental data-plane building block for most of the Eventing working group's APIs, including: Broker, Parallel, and Sequence. The stability and reliability of event delivery is thus critical for the majority of Eventing API's surface area. Due to the importance of event delivery, this newly proposed Working Group will lead the Spec, API, documentation as well as conformance test in a focused way, helping implementers of eventing components and other APIs depending on event delivery, to guarantee a consistent and efficient data tranfer API.

We expect that other working groups will focus on how channels are used and in particular creating working group for Broker, Parallel, and Sequence (and similar higher level constructs) if needed.

## Goals

- As an [Event Consumer](https://github.com/knative/eventing/blob/master/docs/personas.md#event-consumer-developer), the choice of Channel implementation with identical defined characteristics (e.g. persistence or in-order delivery) has no affect on me.
- As an [Event Producer](https://github.com/knative/eventing/blob/master/docs/personas.md#event-producer), the choice of Channel implementation has no affect on me.
- As a [System Integrator](https://github.com/knative/eventing/blob/master/docs/personas.md#system-integrator), I can create a new Channel implementation and use the conformance tests to verify that it meets all the requirements of a Channel implementation.
- As an Operator, I can find conforming Channel implementations and control which are available within the cluster.
- As an Operator, I can choose Channels based on additional, well defined characteristics, such as persistence or in-order delivery.
- As an Operator, I can understand Channel usage and locate errors.

## Scope

In the existing Eventing repo inside the knative org:

- Continue to evolve the [Channel Spec](https://github.com/knative/eventing/blob/master/docs/spec/channel.md).
- Continue to increase the coverage of the [Channel conformance tests](https://github.com/knative/eventing/tree/master/test/conformance).
- Explore [Channel error handling](https://github.com/knative/eventing/tree/master/docs/delivery).
- Maintain libraries to make Channels more consistent, such as [EventReceiver](https://github.com/knative/eventing/blob/master/pkg/channel/event_receiver.go) and [EventDispatcher](https://github.com/knative/eventing/blob/master/pkg/channel/event_dispatcher.go).
- Maintain OWNERS files within both the eventing and eventing-contrib repo for folders owned by the Channels WG. So that it is easier to track who is on track PRs and who is on track to become an approver or lead.
- Document and evolve the event delivery contract between all Knative pieces dealing with events, and manifest this contract in a “Knative Event Delivery Spec”.

Documentation:

- Explore different styles of Channel implementations, giving nuanced recommendations on the trade-offs between different designs (e.g. Job based vs singleton dispatchers).
- Maintain a list of conforming [Channel implementations](https://knative.dev/docs/eventing/channels/channels-crds/).

In working group meetings:

- Provide a forum to discuss and align work items related to the working group, status updates.

## Preliminary 3-Month Roadmap

Month 1

- Start to define the Event Delivery Spec.
- Verify the Channel Spec and reflect content from Event Delivery Spec.
- InMemoryChannel has error handling (dead letter queue).
- Document the different styles of Channels that currently exist and their trade offs.
- Increase the amount of the specification covered by the conformance tests.

Month 2

- At least one Channel implementation other than InMemoryChannel has error handling.
- Increase the amount of the specification covered by the conformance tests.
- Add observability to the Channel specification.

Month 3

- At least one Channel uses native error handling (e.g. RabbitMQ's dead letter exchange).
- 100% coverage of all MUST and MUST NOT directives in the specification covered by the conformance tests.
- Framework for the conformance tests to test optional Channel attributes, such as guaranteed delivery and in-order delivery.
- At least two Channel implementations that pass all conformance tests.
- Document which Channel implementations are available, whether they pass conformance tests, and any optional properties they guarantee (e.g. persistence, in order delivery, etc.).

## Working Group

We will operate as a partner group to Eventing, giving a brief update in the weekly Eventing working group meeting. We will additionally meet on a weekly basis, cancelled when there is a lack of agenda.

## Proposed Leads

- [Harwayne](https://github.com/Harwayne)
- [matzew](https://github.com/matzew)

## History

The Eventing WG created a [Channel Task Force](https://docs.google.com/document/d/1uxlulaAf2m_yZUqCIeI-inul2gsqP69PElnZdO0FHUo/edit#) in May 2019, lead by [Harwayne](https://github.com/Harwayne) and [matzew](https://github.com/matzew). It has been meeting on a weekly basis and has made progress on many of the issues above. As the scope of the task force has increased, it no longer has a clear exit condition. Therefore it feels more appropriate for it to be a working group, rather than a task force.
