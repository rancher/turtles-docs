---
sidebar_position: 0
---

# Create & import a cluster using CAPI provider for AWS

Remember that most Cluster API Providers are upstream projects maintained by the Kubernetes open-source community.

## Prerequisites

- Rancher Manager cluster with Rancher Turtles installed
- Cluster API providers installed for your scenario - we'll be using the AWS infrastructure in these instructions - you can find a guide on how to install a provider using the `CAPIProvider` resource [here](../../tasks/capi-operator/basic_cluster_api_provider_installation.md)
- **clusterctl** CLI - see the [releases](https://github.com/kubernetes-sigs/cluster-api/releases)

## Create Your Cluster Definition

To generate the YAML for the cluster, do the following (assuming the AWS infrastructure provider is being used):

1. Open a terminal and run the following:

```bash
export KUBERNETES_VERSION=v1.28
export AWS_REGION=eu-west-2
export AWS_INSTANCE_TYPE=t3.medium

clusterctl generate cluster cluster1 \
--from https://raw.githubusercontent.com/rancher-sandbox/rancher-turtles-fleet-example/templates/capa.yaml \
> cluster1.yaml
```
2. View **cluster1.yaml** to ensure there are no tokens (i.e. SSH keys or cloud credentials). You can make any changes you want as well.

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
