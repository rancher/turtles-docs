---
sidebar_position: 3
---

# Create & Import Your First Cluster Using kubectl

This section will guide you through creating your first cluster and importing it into Rancher Manager using kubectl.

## Prerequisites

- Rancher Manager cluster with Rancher Turtles installed
- Cluster API providers installed for your scenario - we'll be using the [Docker infrastructure](https://github.com/kubernetes-sigs/cluster-api/tree/main/test/infrastructure/docker) and [RKE2 bootstrap/control plane](https://github.com/rancher-sandbox/cluster-api-provider-rke2) providers in these instructions - see [Initialization for common providers using Turtles' `CAPIProvider`](../../tasks/capi-operator/capiprovider_resource.md)
- **clusterctl** CLI - see the [releases](https://github.com/kubernetes-sigs/cluster-api/releases)

## Create Your Cluster Definition

To generate the YAML for the cluster, do the following (assuming the Docker infrastructure provider is being used):

1. Open a terminal and run the following:

```bash
export CLUSTER_NAME=cluster1
export CONTROL_PLANE_MACHINE_COUNT=1
export WORKER_MACHINE_COUNT=1
export KUBERNETES_VERSION=v1.30.0

curl -s https://raw.githubusercontent.com/rancher-sandbox/rancher-turtles-fleet-example/templates/docker-rke2.yaml | envsubst > cluster1.yaml
```

2. View **cluster1.yaml** to ensure there are no tokens. You can make any changes you want as well.

> The Cluster API quickstart guide contains more detail. Read the steps related to this section [here](https://cluster-api.sigs.k8s.io/user/quick-start.html#required-configuration-for-common-providers).

3. Create the cluster using kubectl

```bash
kubectl create -f cluster1.yaml
```

## Mark Namespace or Cluster for Auto-Import

To automatically import a CAPI cluster into Rancher Manager, there are 2 options:

1. Label a namespace so all clusters contained in it are imported.
2. Label an individual cluster definition so that it's imported.

Labeling a namespace:

```bash
kubectl label namespace default cluster-api.cattle.io/rancher-auto-import=true
```

Labeling an individual cluster definition:

```bash
kubectl label cluster.cluster.x-k8s.io cluster1 cluster-api.cattle.io/rancher-auto-import=true
```
