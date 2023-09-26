---
sidebar_position: 4
---

# Install Rancher Turtles Operator

This section walks through different installation options for the Rancher Turtles Operator.

### Install Rancher Turtles Operator with `Cluster API Operator` as a Helm dependency

*Note: this section will be extended with additional details later*

A `rancher-turtles` chart repository should be added first:

```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```

To install `Cluster API Operator` as a dependency to the `Rancher Turtles`, a set of additional helm flags should be specified:

```bash
helm install rancher-turtles turtles/rancher-turtles
    -n rancher-turtles-system
    --dependency-update
    # Passing secret name and namespace for additional environment variables to be used when deploying CAPI provider
    --set cluster-api-operator.cluster-api.configSecret.name=<secret_name>
    --set cluster-api-operator.cluster-api.configSecret.namespace=<secret_namespace>
    --create-namespace --wait
    --timeout 180s
```

Any values passed to `helm` with the `cluster-api-operator` key will be passed along to the `Cluster API Operator` project. A full set of available values for the `Cluster API Operator` can be found in the operator [values.yaml](https://github.com/kubernetes-sigs/cluster-api-operator/blob/main/hack/charts/cluster-api-operator/values.yaml).

Currently the available set of values for the `cluster-api-operator` setup in the `rancher-turtles`:

```yaml
cluster-api-operator:
  enabled: true # indicates if CAPI operator should be installed (default: true)
  cert-manager:
    enabled: true # indicates if cert-manager should be installed (default: true)
  cluster-api:
    enabled: true # indicates if core CAPI controllers should be installed (default: true)
    version: v1.4.6 # version of CAPI to install (default: v1.4.6)
    configSecret:
      name: "" # name of the config secret to use for core CAPI controllers, used by the CAPI operator. See [CAPI operator](https://github.com/kubernetes-sigs/cluster-api-operator/tree/main/docs#installing-azure-infrastructure-provider) docs for more details.
      namespace: "" # namespace of the config secret to use for core CAPI controllers, used by the CAPI operator.
    core:
      namespace: capi-system
      fetchConfig: # (only required for airgapped environments)
        url: ""  # url to fetch config from, used by the CAPI operator. See [CAPI operator](https://github.com/kubernetes-sigs/cluster-api-operator/tree/main/docs#provider-spec) docs for more details.
        selector: ""  # selector to use for fetching config, used by the CAPI operator.
    kubeadmBootstrap:
      namespace: capi-kubeadm-bootstrap-system
      fetchConfig:
        url: ""
        selector: ""
    kubeadmControlPlane:
      namespace: capi-kubeadm-control-plane-system
      fetchConfig:
        url: ""
        selector: ""
```

A `secret` with a set of environment variables should be passed to the `Cluster API Operator` installation.

Example `variables/default` `secret` configuration with `CLUSTER_TOPOLOGY` and `EXP_CLUSTER_RESOURCE_SET` feature flags set:

```sh title="helm install flags"
helm install ...
--set cluster-api-operator.cluster-api.configSecret.name=variables
--set cluster-api-operator.cluster-api.configSecret.namespace=default
```

```yaml title="secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: variables
  namespace: default
type: Opaque
stringData:
  CLUSTER_TOPOLOGY: "true"
  EXP_CLUSTER_RESOURCE_SET: "true"
```

### Install Rancher Turtles Operator without `Cluster API Operator` as a Helm dependency

A `Rancher Turtles` requires a connection to the `Rancher Manager` cluster. This can be achieved by:

1. Installing it in the same cluster as the `Rancher Manager`.

*Note: In the future, we will support different deployment topologies*

The recommended path of installation for the operator is by using `Helm`. To install it in the cluster, a chart repository should be added first:

```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```
and then it can be installed into the `rancher-turtles-system` namespace with:

```bash
helm install rancher-turtles turtles/rancher-turtles
    -n rancher-turtles-system
    --set cluster-api-operator.enabled=false
    --set cluster-api-operator.cluster-api.enabled=false
    --create-namespace --wait
    --dependency-update
```