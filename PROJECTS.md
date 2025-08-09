# Knative Project Overview

## Functions

[func](https://github.com/knative/func) provides a simple programming model for using functions on Knative, without requiring in-depth knowledge of Knative, Kubernetes, containers, or dockerfiles.

## Serving

[Serving](https://github.com/knative/serving) defines a set of resources to manage how your serverless workload behaves on the Kubernetes cluster.

### Related Sub-projects

#### Networking Plugins

[Networking plugins](https://github.com/knative-extensions?q=net-&type=public) provide a shim to converts generic Serving resources to networking specific ones. Example integrations are Istio, Contour, Gateway API and Kourier.

#### Client

The [`kn` cli](https://github.com/knative/client) provides a quick and easy interface for creating Knative resources such as services and event sources, without the need to create or modify YAML files directly. kn also simplifies completion of otherwise complex procedures such as autoscaling and traffic splitting.

#### Operator

The [Operator](https://github.com/knative/operator) helps install, manage and upgrade Knative Serving installations on Kubernetes Clusters.

## Eventing

[Eventing](https://github.com/knative/eventing) is a collection of APIs that enable you to use an event-driven architecture with your applications. You can use these APIs to create components that route events from event producers (known as sources) to event consumers (known as sinks) that receive events.

### Related Sub-projects

#### Eventing Plugins

[Eventing plugins](https://github.com/knative-extensions?q=eventing-&type=public&language=&sort=) allow end-users to integrate Knative with various different eventing buses/queues such as Kafka, Redis, NATS, Rabbit-MQ.

#### Client

The [`kn` cli](https://github.com/knative/client) provides a quick and easy interface for creating Knative resources such as services and event sources, without the need to create or modify YAML files directly. kn also simplifies completion of otherwise complex procedures such as autoscaling and traffic splitting.

#### Operator

The [Operator](https://github.com/knative/operator) helps install, manage and upgrade Knative Eventing installations on Kubernetes Clusters.

## Additional Resources

See the [Knative documentation](https://knative.dev) for in-depth information.
