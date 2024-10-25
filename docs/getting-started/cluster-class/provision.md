---
sidebar_position: 1
---
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Create a workload cluster

This section will guide you through creating a `ClusterClass` which you will then use to provision a workload `Cluster`. The Cluster API book includes a [ClusterClass section](https://cluster-api.sigs.k8s.io/tasks/experimental-features/cluster-class/) with detailed information on what is supported and how you can use the powerful abstraction this feature provides, including operations such as [patching class values](https://cluster-api.sigs.k8s.io/tasks/experimental-features/cluster-class/write-clusterclass#clusterclass-with-patches) and [operating a managed cluster](https://cluster-api.sigs.k8s.io/tasks/experimental-features/cluster-class/operate-cluster). We recommend you familiarize with the feature to get the most out of it.

:::note
This guide uses the [examples repository](https://github.com/rancher-sandbox/rancher-turtles-fleet-example/tree/templates).
:::

## Providers guide

<Tabs>
  <TabItem value="azure-aks" label="CAPZ AKS" default>
  ## Prerequisites
    - Rancher Manager cluster with Rancher Turtles installed
    - Configure cloud credentials for Azure in Rancher: `Cluster Management` > `Cloud Credentials`.
        - Keep the name you assign to the new set of credentials.
    - Install the [CAPI Infrastructure Provider for Azure](https://github.com/kubernetes-sigs/cluster-api-provider-azure/) using the [`CAPIProvider` resource](../../tasks/capi-operator/basic_cluster_api_provider_installation.md).
        ```yaml
        apiVersion: v1
        kind: Namespace
        metadata:
          name: capz-system
        ---
        apiVersion: turtles-capi.cattle.io/v1alpha1
        kind: CAPIProvider
        metadata:
          name: azure
          namespace: capz-system
        spec:
          type: infrastructure
        ```
  ## Create ClusterClass object

    The `ClusterClass` object represents a template that defines the shape of the control plane and infrastructure of a cluster. This is the base definition of the `Cluster` object/s that will be created from it. If the template is created optimizing flexibility, we could use it to provision workload clusters supporting variants of the same cluster shape, simplifying the configuration applied to each cluster, as the class removes most of the complexity.

    The template we're using in this example will use CAPZ to provision a managed Azure (AKS) cluster. Before applying the yaml file, you will need to export the following environment variables. Remember to adapt the values to your specific scenario as these are just placeholders:
    ```bash
    export CLUSTER_CLASS_NAME="azure-sample"
    export CLUSTER_NAME="azure-aks-cluster"
    export AZURE_LOCATION="northeurope"
    export AZURE_NODE_MACHINE_TYPE="Standard_D2s_v3"
    export AZURE_SUBSCRIPTION_ID=<subs-id> # you can use: az account show --query 'id' --output tsv
    export KUBERNETES_VERSION="v1.30.4"
    export AZURE_CLIENT_ID=<app-id>
    export AZURE_TENANT_ID=<tenant-id>
    export AZURE_CLIENT_SECRET=<password>
    ```

    Using `envsubst` to substitute the exported variables in the original file.

    ```bash
    curl -s https://raw.githubusercontent.com/rancher-sandbox/rancher-turtles-fleet-example/templates/capz/cluster-template-aks-clusterclass.yaml | envsubst > clusterclass1.yaml
    ```

    This will create a new yaml file `clusterclass1.yaml` that contains the class definition formatted with the exported values. You can study the resulting file before applying it to the cluster.

    ```bash
    kubectl apply -f clusterclass1.yaml
    ```

    You can validate that the class has been created successfully and inspect its content via `kubectl`:

    ```bash
    kubectl get clusterclasses.cluster.x-k8s.io
    kubectl describe clusterclasses.cluster.x-k8s.io <class-name>
    ```

  ## Provision workload cluster

    Now that the class resource is available in the cluster, we can go ahead and create a cluster from this topology. Let's first substitute the variables in the template, as we did before:

    ```bash
    curl -s https://raw.githubusercontent.com/rancher-sandbox/rancher-turtles-fleet-example/templates/capz/cluster-template-aks-topology.yaml | envsubst > cluster1.yaml
    ```

    This will create a new yaml file `cluster1.yaml` that contains the cluster definition formatted with the exported values. You can study the resulting file before applying it to the cluster, which will effectively trigger workload cluster creation.

    ```bash
    kubectl apply -f cluster1.yaml
    ```

    Be patient, cluster provisioning will take some time (up to 10min). While you wait for it to become ready, you can go through the `capz-controller-manager` logs, which is responsible for reconciling the cluster resources you just created.

    </TabItem>

</Tabs>

  ## Enable auto-import into Rancher

    As with any other CAPI clusters, you will have to enable auto-import for Turtles to manage importing it into Rancher Manager. Please, refer to [Mark namespace for auto-import](../../getting-started/create-first-cluster/using_fleet#mark-namespace-for-auto-import) notes to enable auto-import.

  ## Post provisioning actions

    The functionality provided by cluster classes makes it possible for you to deploy as many clusters as desired from the topology you created. This template can be written in a way that makes it flexible enough to be used in as many Clusters as possible by supporting variants of the same base Cluster shape.
