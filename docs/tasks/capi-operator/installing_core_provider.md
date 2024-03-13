---
sidebar_position: 4
---

# Installing the CoreProvider using CAPIProvider resource

This section describes how to install the `CoreProvider` via `CAPIProvider`, which is responsible for managing the Cluster API CRDs and the Cluster API controller.

:::note
Please refer to installing Core Provider [section](https://cluster-api-operator.sigs.k8s.io/03_topics/03_basic-cluster-api-provider-installation/01_installing-core-provider#installing-the-coreprovider) in CAPI Operator docs for additional details on raw `CoreProvider` resource installation.

Only one CoreProvider can be installed at the same time on a single cluster.
:::

*Example:*

```yaml
apiVersion: turtles-capi.cattle.io/v1alpha1
kind: CAPIProvider
metadata:
  name: cluster-api
  namespace: capi-system
spec:
  version: v1.4.6
  type: core # required
```
