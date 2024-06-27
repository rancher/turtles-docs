---
sidebar_position: 0
---
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';


# Create & import a cluster using CAPI providers

This guide goes over the process of creating and importing CAPI clusters with a selection of the officially supported providers.

Remember that most Cluster API Providers are upstream projects maintained by the Kubernetes open-source community.

## Prerequisites

<Tabs>
  <TabItem value="aws-kubeadm" label="AWS Kubeadm" default>
    - Rancher Manager cluster with Rancher Turtles installed
    - Cluster API Providers: you can find a guide on how to install a provider using the `CAPIProvider` resource [here](../../tasks/capi-operator/basic_cluster_api_provider_installation.md)
        - [Infrastructure provider for AWS](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)
        - [Bootstrap/Control Plane provider for Kubeadm](https://github.com/kubernetes-sigs/cluster-api)
    - **clusterctl** CLI - see [install clusterctl from CAPI book](https://cluster-api.sigs.k8s.io/user/quick-start#install-clusterctl)
    </TabItem>
  <TabItem value="docker-kubeadm" label="Docker Kubeadm">
    - Rancher Manager cluster with Rancher Turtles installed
    - Cluster API Providers: you can find a guide on how to install a provider using the `CAPIProvider` resource [here](../../tasks/capi-operator/basic_cluster_api_provider_installation.md)
        - [Infrastructure provider for Docker](https://github.com/kubernetes-sigs/cluster-api)
        - [Bootstrap/Control Plane provider for Kubeadm](https://github.com/kubernetes-sigs/cluster-api)
    - **clusterctl** CLI - see [install clusterctl from CAPI book](https://cluster-api.sigs.k8s.io/user/quick-start#install-clusterctl)
  </TabItem>
</Tabs>

## Create Your Cluster Definition

<Tabs>
  <TabItem value="aws-kubeadm" label="AWS Kubeadm" default>
To generate the YAML for the cluster, do the following:

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
    </TabItem>
  <TabItem value="docker-kubeadm" label="Docker Kubeadm">
To generate the YAML for the cluster, do the following:

1. Open a terminal and run the following:

```bash
export CONTROL_PLANE_MACHINE_COUNT=1
export WORKER_MACHINE_COUNT=1
export KUBERNETES_VERSION=v1.26.4

clusterctl generate cluster cluster1 \
--from https://raw.githubusercontent.com/rancher-sandbox/rancher-turtles-fleet-example/templates/docker-kubeadm.yaml \
> cluster1.yaml
```

2. View **cluster1.yaml** to ensure there are no tokens. You can make any changes you want as well.

> The Cluster API quickstart guide contains more detail. Read the steps related to this section [here](https://cluster-api.sigs.k8s.io/user/quick-start.html#required-configuration-for-common-providers).

3. Create the cluster using kubectl

```bash
kubectl create -f cluster1.yaml
```
    </TabItem>
</Tabs>

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
kubectl label cluster.cluster.x-k8s.io -n default cluster1 cluster-api.cattle.io/rancher-auto-import=true
```
