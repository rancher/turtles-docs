---
sidebar_position: 1
---

# Introduction

In this section we cover using **ClusterClass** with Rancher Turtles.

:::caution
ClusterClass is an experimental feature of Cluster API. As with any experimental features it should be used with caution as it may be unreliable. All experimental features are not subject to any compatibility or deprecation promise. **It is planned to be graduated with [CAPI v1.9](https://github.com/kubernetes-sigs/cluster-api/milestone/38)**.
:::

:::tip
Before using ClusterClass, study the provider docs and confirm that the feature is supported. This documentation includes a matrix in the [Certified providers](../../reference-guides/providers/certified.md) section with [cluster topology support status](../../reference-guides/providers/certified.md#clusterclass-support-for-certified-providers). 
:::

## Pre-requisities

To use ClusterClass it needs to be enabled for core Cluster API and any provider that supports it. This is done by setting the `CLUSTER_TOPOLOGY` variable to `true`.

The Turtles Helm chart will automatically enable the feature when installing CAPI. However, when enabling additional providers, ensure `CLUSTER_TOPOLOGY` is set in the provider configuration. Turtles' [CAPIProvider](../../tasks/capi-operator/capiprovider_resource.md) resource supports passing installation parameters to the provider via `variables` as follows:

```yaml
apiVersion: turtles-capi.cattle.io/v1alpha1
kind: CAPIProvider
metadata:
  name: azure
  namespace: capz-system
spec:
  type: infrastructure
  name: azure
  configSecret:
    name: azure-variables
  variables:
    CLUSTER_TOPOLOGY: "true"
    EXP_CLUSTER_RESOURCE_SET: "true"
    EXP_MACHINE_POOL: "true"
    EXP_AKS_RESOURCE_HEALTH: "true"
```

The resource defined in this yaml file will install CAPZ with support for a number of features, including `CLUSTER_TOPOLOGY`.
