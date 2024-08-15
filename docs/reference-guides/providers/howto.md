---
sidebar_position: 0
---
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';


# Create & import a cluster using CAPI providers

This guide goes over the process of creating and importing CAPI clusters with a selection of the officially certified providers.

Keep in mind that most Cluster API Providers are upstream projects maintained by the Kubernetes open-source community.

## Prerequisites

<Tabs>
  <TabItem value="aws-rke2" label="AWS RKE2" default>
    - Rancher Manager cluster with Rancher Turtles installed
    - Cluster API Providers: you can find a guide on how to install a provider using the `CAPIProvider` resource [here](../../tasks/capi-operator/basic_cluster_api_provider_installation.md)
        - [Infrastructure provider for AWS](https://github.com/kubernetes-sigs/cluster-api-provider-aws/)
        - [Bootstrap/Control Plane provider for RKE2](https://github.com/rancher/cluster-api-provider-rke2)
    - **clusterctl** CLI - see [install clusterctl from CAPI book](https://cluster-api.sigs.k8s.io/user/quick-start#install-clusterctl)
    </TabItem>
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
  <TabItem value="aws-rke2" label="AWS RKE2" default>


Before creating an AWS+RKE2 workload cluster, it is required to build an AMI for the RKE2 version that is going to be installed on the cluster. You can follow the steps in the [RKE2 image-builder README](https://github.com/rancher/cluster-api-provider-rke2/tree/main/image-builder#aws) to build the AMI.

We recommend you refer to the CAPRKE2 repository where you can find a [samples folder](https://github.com/rancher/cluster-api-provider-rke2/tree/main/samples/aws) with different CAPA+CAPRKE2 cluster configurations that can be used to provision downstream clusters. The [internal folder](https://github.com/rancher/cluster-api-provider-rke2/tree/main/samples/aws/internal) contains cluster templates to deploy an RKE2 cluster on AWS using the internal cloud provider, and the [external folder](https://github.com/rancher/cluster-api-provider-rke2/tree/main/samples/aws/external) contains the cluster templates to deploy a cluster with the external cloud provider.

We will use the `internal` one for this guide, however the same steps apply for `external`.

To generate the YAML for the cluster, do the following:

1. Open a terminal and run the following:

```bash
export CONTROL_PLANE_MACHINE_COUNT=3
export WORKER_MACHINE_COUNT=3
export RKE2_VERSION=v1.30.3+rke2r1
export AWS_NODE_MACHINE_TYPE=t3a.large
export AWS_CONTROL_PLANE_MACHINE_TYPE=t3a.large
export AWS_SSH_KEY_NAME="aws-ssh-key"
export AWS_REGION="aws-region"
export AWS_AMI_ID="ami-id"

clusterctl generate cluster cluster1 \
--from https://github.com/rancher/cluster-api-provider-rke2/blob/main/samples/aws/internal/cluster-template.yaml \
> cluster1.yaml
```
2. View **cluster1.yaml** and examine the resulting yaml file. You can make any changes you want as well.

> The Cluster API quickstart guide contains more detail. Read the steps related to this section [here](https://cluster-api.sigs.k8s.io/user/quick-start.html#required-configuration-for-common-providers).

3. Create the cluster using kubectl

```bash
kubectl create -f cluster1.yaml
```
    </TabItem>
  <TabItem value="aws-kubeadm" label="AWS Kubeadm" default>
To generate the YAML for the cluster, do the following:

1. Open a terminal and run the following:

```bash
export KUBERNETES_VERSION=v1.30
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
export KUBERNETES_VERSION=v1.30.0

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

:::tip
After your cluster is provisioned, you can check functionality of the workload cluster using `clusterctl`:
```bash
clusterctl describe cluster cluster1
```

Remember that clusters are namespaced resources. These examples provision clusters in the `default` namespace, but you will need to provide yours if using a different one.
:::

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
