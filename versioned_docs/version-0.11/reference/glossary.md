---
sidebar_position: 2
---

# Glossary

## Table of Contents

[A](glossary.md#a) | [B](glossary.md#b) | [C](glossary.md#c) | [F](glossary.md#f) | [I](glossary.md#i) | [K](glossary.md#k) | [M](glossary.md#m) | [N](glossary.md#n) | [O](glossary.md#o) | [P](glossary.md#p) | [R](glossary.md#r) | [S](glossary.md#s) | [T](glossary.md#t) | [W](glossary.md#w)

A
---

### Add-ons

Services beyond the fundamental components required to deploy a Kubernetes-conformant cluster and categorized into two types:

* __Core Add-ons__: Addons that are required to deploy a Kubernetes-conformant cluster: DNS, kube-proxy, CNI.
* __Additional Add-ons__: Addons that are not required for a Kubernetes-conformant cluster (e.g. metrics/Heapster, Dashboard).

### Air-gapped environment

Setting up and running Kubernetes clusters without direct access to the internet.

B
---

### Bootstrap

The process of turning a server into a Kubernetes node. This may involve assembling data to provide when creating the server that backs the Machine, as well as runtime configuration of the software running on that server.

### Bootstrap cluster

A temporary cluster that is used to provision a Target Management cluster.

### Bootstrap provider

Refers to a [provider](#provider) that implements a solution for the [bootstrap](#bootstrap) process.

C
---

### CAPI
Core Cluster API

### CAPA
Cluster API Provider AWS

### CAPD
Cluster API Provider Docker

### CAPG
Cluster API Google Cloud Provider

### CAPIO
Cluster API Operator

### CAPRKE2

Cluster API Provider RKE2

### CAPV
Cluster API Provider vSphere

### CAPZ
Cluster API Provider Azure

### CAPI Provider

A public API that facilitates provisioning and operations over the [CAPI Operator](#cluster-api-operator) and [Cluster API](#cluster-api) resources.


### Child Cluster

Term commonly used interchangeably with the [workload cluster](#workload-cluster).

### Cluster

A full Kubernetes deployment. See Management Cluster and Workload Cluster.

### ClusterClass

A collection of templates that define a topology (control plane and workers) to be used to continuously reconcile one or more Clusters.
See [ClusterClass](../getting-started/cluster-class/create_cluster.md)

### Cluster API

Or __Cluster API project__

The Cluster API is a sub-project of the SIG-cluster-lifecycle. It is also used to refer to the software components, APIs, and community that produce them.

See [core provider](#core-provider)

### Cluster API Operator

Or __Cluster API Operator project__

The Cluster API Operator is a sub-project of the SIG-cluster-lifecycle. It is designed to empower cluster administrators to handle the lifecycle of Cluster API providers within a management cluster using a declarative approach.

### Cluster API Provider RKE2

[Cluster API Provider RKE2](#caprke2) is a combination of two provider types: a Cluster API Control Plane Provider for provisioning Kubernetes control plane nodes and a Cluster API Bootstrap Provider for bootstrapping Kubernetes on a machine where [RKE2](#rke2) is used as the Kubernetes distro.

### Control plane

The set of Kubernetes services that form the basis of a cluster. See also [https://kubernetes.io/docs/concepts/#kubernetes-control-plane](https://kubernetes.io/docs/concepts/#kubernetes-control-plane) There are two variants:

* __Self-provisioned__: A Kubernetes control plane consisting of pods or machines wholly managed by a single Cluster API deployment.
* __External__ or __Managed__: A control plane offered and controlled by some system other than Cluster API (e.g., GKE, AKS, EKS, IKS).

### Control plane provider

Refers to a [provider](#provider) that implements a solution for the management of a Kubernetes [control plane](#control-plane).

See [CAPRRKE2](#caprke2), [KCP](#kcp).

### Core provider

Refers to a [provider](#provider) that implements Cluster API core controllers; if you
consider that the first project that must be deployed in a management Cluster is Cluster API itself, it should be clear why
the Cluster API project is also referred to as the core provider.

See [CAPI](#cluster-api).

F
---

### Fleet

A container management and deployment engine designed to offer users more control on the local cluster and constant monitoring through GitOps. Take a look at [fleet documentation](https://fleet.rancher.io/) to know more about Fleet.

I
---

### Infrastructure provider

Refers to a [provider](#provider) that implements provisioning of infrastructure/computational resources required by
the Cluster or by Machines (e.g. VMs, networking, etc.).
Clouds infrastructure providers include AWS, Azure, or Google; while VMware, MAAS, or metal3.io can be defined as bare metal providers.

### IPAM provider

Refers to a [provider](#provider) that allows Cluster API to interact with IPAM solutions.
IPAM provider's interaction with Cluster API is based on the `IPAddressClaim` and `IPAddress` API types.

K
---

### Kubernetes-conformant

Or __Kubernetes-compliant__

A cluster that passes the Kubernetes conformance tests.

### Kubernetes Operator

A Kubernetes Operator is a method of packaging, deploying, and managing a Kubernetes application. See also [https://kubernetes.io/docs/concepts/extend-kubernetes/operator/](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) for more information.

### k/k

Refers to the [main Kubernetes git repository](https://github.com/kubernetes/kubernetes) or the main Kubernetes project.

### KCP

Kubeadm Control plane Provider

M
---

### Machine

Or __Machine Resource__

The Custom Resource for Kubernetes that represents an infrastructure component that hosts a Kubernetes node.


### Manage a cluster

Perform create, scale, upgrade, or destroy operations on the cluster.

### Managed Kubernetes

Managed Kubernetes refers to any Kubernetes cluster provisioning and maintenance abstraction, usually exposed as an API, that is natively available in a Cloud provider. For example: [EKS](https://aws.amazon.com/eks/), [OKE](https://www.oracle.com/cloud/cloud-native/container-engine-kubernetes/), [AKS](https://azure.microsoft.com/en-us/products/kubernetes-service), [GKE](https://cloud.google.com/kubernetes-engine), [IBM Cloud Kubernetes Service](https://www.ibm.com/cloud/kubernetes-service), [DOKS](https://www.digitalocean.com/products/kubernetes), and many more throughout the Kubernetes Cloud Native ecosystem.

### Managed Topology

See [Topology](#topology)

### Management cluster

The cluster where one or more Infrastructure Providers run, and where resources (e.g. Machines) are stored. Typically referred to when you are provisioning multiple workload clusters.

N
---

### Node pools

A node pool is a group of nodes within a cluster that all have the same configuration.

O
---

### Operating system

Or __OS__

A generically understood combination of a kernel and system-level userspace interface, such as Linux or Windows, as opposed to a particular distribution.

P
---

### Pivot

Pivot is a process for moving the provider components and declared cluster-api resources from a Source Management cluster to a Target Management cluster.

The pivot process is also used for deleting a management cluster and could also be used during an upgrade of the management cluster.

### Provider

Or __Cluster API provider__

This term was originally used as abbreviation for [Infrastructure provider](#infrastructure-provider), but currently it is used
to refer to any project that can be deployed and provides functionality to the Cluster API management Cluster.

See [Bootstrap provider](#bootstrap-provider), [Control plane provider](#control-plane-provider), [Core provider](#core-provider),
[Infrastructure provider](#infrastructure-provider), [IPAM provider](#ipam-provider), [Runtime extension provider](#runtime-extension-provider).

### Provider components

Refers to the YAML artifact published as part of the release process for [providers](#provider);
it usually includes Custom Resource Definitions (CRDs), Deployments (to run the controller manager), RBAC, etc.

In some cases, the same expression is used to refer to the instances of above components deployed in a management cluster.

See [Provider repository](#provider-repository)

### Provider repository

Refers to the location where the YAML for [provider components](#provider-components) are hosted; usually a provider repository hosts
many version of provider components, one for each released version.

R
---

### Rancher

An open-source [platform](https://www.rancher.com/) designed to simplify the deployment and management of Kubernetes clusters.

### Rancher Cluster Agent

A component deployed by Rancher in each Kubernetes cluster it manages. Its primary role is to establish a secure communication channel between the Rancher server and the Kubernetes cluster, enabling Rancher to manage and interact with the cluster.

### Rancher Manager

The Rancher Manager (or Rancher Server) is where the Rancher UI and API are hosted, and it communicates with managed clusters through components like the [Rancher Cluster Agent](#rancher-cluster-agent). It allows users to manage their Kubernetes clusters, applications, and Rancher-specific resources such as Catalogs, Users, Global Roles, and more. 

### Rancher Turtles

A [Kubernetes operator](#kubernetes-operator) that provides integration between Rancher Manager and Cluster API (CAPI) with the aim of bringing full CAPI support to Rancher.

### RKE2

Rancher's next-generation, fully conformant Kubernetes distribution that focuses on security and compliance within the U.S. Federal Government sector. See [documentation](https://docs.rke2.io/) for more details.

### Runtime Extension

An external component which is part of a system built on top of Cluster API that can handle requests for a specific Runtime Hook.

### Runtime Extension provider

Refers to a [provider](#provider) that implements one or more [runtime extensions](#runtime-extension).

S
---

### Scaling

Unless otherwise specified, this refers to horizontal scaling.

### Stacked control plane

A control plane node where etcd is colocated with the Kubernetes API server, and
is running as a static pod.

### Server

The infrastructure that backs a [Machine Resource](#machine), typically either a cloud instance, virtual machine, or physical host.

T
---

### Topology

A field in the Cluster object spec that allows defining and managing the shape of the Cluster's control plane and worker machines from a single point of control. The Cluster's topology is based on a [ClusterClass](#clusterclass).
Sometimes it is also referred as a managed topology.

See [ClusterClass](#clusterclass)

### Turtles

Refers to [Rancher Turtles](#rancher-turtles)

W
---

### Workload Cluster

A cluster created by a ClusterAPI controller, which is *not* a bootstrap cluster, and is meant to be used by end-users, as opposed to by CAPI tooling.

### WorkerClass

A collection of templates that define a set of worker nodes in the cluster. A ClusterClass contains zero or more WorkerClass definitions.

See [ClusterClass](#clusterclass)